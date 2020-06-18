package main

import (
	"net/http"

	"github.com/arunvm/recro/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (server *server) userSearch(c *gin.Context) {
	var args models.UserSearchArgs

	err := c.ShouldBindJSON(&args)
	if err != nil {
		log.WithFields(log.Fields{
			"func": "userSearch",
			"info": "error decoding request body",
		}).Error(err)
		c.JSON(http.StatusBadRequest, "Request body not properly formatted")
		return
	}

	results, err := models.UserSearch(server.db, &args)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "userSearch",
			"subFunc": "models.UserSearch",
			"args":    args,
		}).Error(err)
		c.JSON(http.StatusInternalServerError, "Error when retrieving users")
		return
	}

	c.JSON(http.StatusOK, results)
	return
}
