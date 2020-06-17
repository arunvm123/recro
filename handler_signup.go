package main

import (
	"net/http"

	"github.com/arunvm/recro/models"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func (server *server) signup(c *gin.Context) {
	var args models.SignUpArgs
	err := c.ShouldBindJSON(&args)
	if err != nil {
		log.WithFields(log.Fields{
			"func": "signup",
			"info": "decoding request body",
		}).Error(err)
		c.JSON(http.StatusBadRequest, "Request body not properly formatted")
		return
	}

	if models.CheckIfUserExists(server.db, args.Email) == true {
		c.JSON(http.StatusConflict, "Email already exists")
		return
	}

	err = models.UserSignup(server.db, &args)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "signup",
			"subFunc": "models.UserSignup",
			"email":   args.Email,
		})
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
	return
}
