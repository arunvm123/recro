package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initialiseRoutes(server *server) *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	return r
}
