package main

import (
	"time"

	"github.com/arunvm/recro/config"
	"github.com/dgrijalva/jwt-go"
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
