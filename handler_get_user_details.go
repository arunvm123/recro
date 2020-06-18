package main

import (
	"net/http"
	"strconv"

	"github.com/arunvm/recro/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (server *server) getUserDetails(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Please provide a valid id")
		return
	}

	user, err := models.GetUserDetails(server.db, id)
	if err != nil {
		log.WithFields(log.Fields{
			"func":    "getUserDetails",
			"subFunc": "models.GetUserDetails",
		}).Error(err)
		c.JSON(http.StatusInternalServerError, "Error retrieving user details")
		return
	}

	c.JSON(http.StatusOK, user)
	return
}
