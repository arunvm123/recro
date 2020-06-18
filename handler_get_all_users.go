package main

import (
	"net/http"

	"github.com/arunvm/recro/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (server *server) getAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers(server.db)
	if err != nil {
		log.WithFields(log.Fields{
			"func":   "getAllUsers",
			"subunc": "models.GettAllUsers",
		}).Error(err)
		c.JSON(http.StatusInternalServerError, "Error retrieving users")
		return
	}

	c.JSON(http.StatusOK, users)
	return
}
