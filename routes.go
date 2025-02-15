package main

import (
	"github.com/gin-gonic/gin"
)

func initialiseRoutes(server *server) *gin.Engine {
	r := gin.Default()

	r.POST("/signup", server.signup)
	r.POST("/login", server.login)

	r.GET("/", server.indexHandler)
	r.GET("/auth/:provider", server.oauthRedirect)
	r.GET("/callback/:provider", server.oauthCallback)

	r.GET("/user/all", server.getAllUsers)
	r.GET("/user", server.getUserDetails)

	private := r.Group("/")
	private.Use(server.tokenAuthorisationMiddleware())
	private.PUT("/user/set_password", server.setPassword)
	private.POST("/user/search", server.userSearch)

	return r
}
