package main

import (
	"net/http"

	"github.com/arunvm/recro/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (server *server) login(c *gin.Context) {
	var args models.LoginArgs

	err := c.ShouldBindJSON(&args)
	if err != nil {
		log.WithFields(log.Fields{
			"func": "login",
			"info": "decoding request body",
		}).Error(err)
		c.JSON(http.StatusBadRequest, "Request body not properly formatted")
		return
	}

	user, err := models.GetUserFromEmail(server.db, args.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, "User does not exist, Please sign up")
			return
		}
		log.WithFields(log.Fields{
			"func":    "login",
			"subFunc": "models.GetUserFromEmail",
			"email":   args.Email,
		}).Error(err)
		c.JSON(http.StatusInternalServerError, "Server error")
		return
	}

	if len(user.Password) == 0 {
		c.JSON(http.StatusUnauthorized, "Please use social login")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(args.Password))
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "login",
			"subFunc": "bcrypt.CompareHashAndPassword",
			"email":   user.Email,
		}).Error(err)
		c.JSON(http.StatusUnauthorized, "Wrong password")
		return
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
