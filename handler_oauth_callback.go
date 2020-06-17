package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (server *server) oauthCallback(c *gin.Context) {
	// Retrieve query params for state and code
	state := c.Query("state")
	code := c.Query("code")
	// provider := c.Param("provider")

	// Handle callback and check for errors
	user, token, err := server.gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Print in terminal user information
	fmt.Printf("%#v", token)
	fmt.Printf("%#v", user)

}
