package main

import (
	"net/http"

	"github.com/arunvm/recro/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (server *server) setPassword(c *gin.Context) {

	user, err := getUserFromContext(c)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "setPassword",
			"subFunc": "getUserFromContext",
		}).Error(err)
		c.JSON(http.StatusBadRequest, "Error fetching user")
		return
	}

	var args models.SetPasswordArgs

	err = c.ShouldBindJSON(&args)
	if err != nil {
		log.WithFields(log.Fields{
			"func": "setPassword",
			"info": "error decoding request body",
		}).Error(err)
		c.JSON(http.StatusBadRequest, "Request body not properly formatted")
		return
	}

	err = models.SetPassword(server.db, user, &args)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "setPassword",
			"subFunc": "models.SetPassword",
			"userId":  user.ID,
		}).Error(err)
		c.JSON(http.StatusInternalServerError, "Error when updateing password")
		return
	}

	c.Status(http.StatusOK)
	return
}
