package main

import (
	"net/http"

	"github.com/arunvm/recro/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (server *server) oauthCallback(c *gin.Context) {
	provider := c.Param("provider")
	// Retrieve query params for state and code
	state := c.Query("state")
	code := c.Query("code")

	// Handle callback and check for errors
	userData, _, err := server.gocial.Handle(state, code)
	if err != nil {
		log.WithFields(log.Fields{
			"func":     "oauthCallback",
			"subFunc":  "server.gocial.Handle",
			"provider": provider,
		})
		c.JSON(http.StatusInternalServerError, "Error retrieving user details from "+provider)
		return
	}

	var user *models.User
	if models.CheckIfUserExists(server.db, userData.Email) == true {
		user, err = models.UpdateProviderDetails(server.db, userData.Email, provider, userData.Raw)
		if err != nil {
			log.WithFields(log.Fields{
				"func":     "oauthCallback",
				"subFunc":  "models.UpdateProviderDetails",
				"provider": provider,
			}).Error(err)
			c.JSON(http.StatusInternalServerError, "Error when updating provider details")
			return
		}
	} else {
		args := models.OauthSignupArgs{
			Email: userData.Email,
			Name:  userData.FirstName + " " + userData.LastName,
		}

		switch provider {
		case "github":
			args.ProviderData = models.ProviderData{
				Github: userData.Raw,
			}
		}

		user, err = models.UserOauthSignup(server.db, &args, provider)
		if err != nil {
			log.WithFields(log.Fields{
				"func":     "oauthCallback",
				"subFunc":  "models.UserOauthSignup",
				"provider": provider,
			}).Error(err)
			c.JSON(http.StatusInternalServerError, "Error when signing up")
			return
		}
	}

	signedToken, err := getJWTToken(user.ID)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "getJWTToken",
			"subFunc": "token.SignedString",
			"email":   user.Email,
		}).Error(err)
		c.JSON(http.StatusInternalServerError, "Error while retrieving token")
		return
	}

	c.JSON(http.StatusOK, struct {
		Token       string  `json:"token"`
		Name        string  `json:"name"`
		ID          int     `json:"id"`
		Email       string  `json:"email"`
		PhoneNumber *string `json:"phone_number"`
	}{
		Token:       signedToken,
		Email:       user.Email,
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	})
	return
}
