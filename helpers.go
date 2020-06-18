package main

import (
	"errors"
	"time"

	"github.com/arunvm/recro/config"
	"github.com/arunvm/recro/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func getJWTToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	config := config.GetConfig()

	signedToken, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "getJWTToken",
			"subFunc": "token.SignedString",
			"userID":  userID,
		}).Error(err)
		return "", err
	}

	return signedToken, nil
}

func getUserFromContext(c *gin.Context) (*models.User, error) {
	user, ok := c.Keys["user"].(*models.User)
	if !ok {
		log.WithFields(log.Fields{
			"func": "getUserFromContext",
			"info": "retrieving user info from context",
		}).Error(errors.New("Error while retrieving user info from context"))
		return nil, errors.New("Error fetching user")
	}

	return user, nil
}
