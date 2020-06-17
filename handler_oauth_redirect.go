package main

import (
	"net/http"

	"github.com/arunvm/recro/config"
	"github.com/gin-gonic/gin"
)

func (server *server) oauthRedirect(c *gin.Context) {
	// Retrieve provider from route
	provider := c.Param("provider")

	config := config.GetConfig()

	providerData := config.Providers[provider]
	scopes := config.Scopes[provider]
	authURL, err := server.gocial.New().
		Driver(provider).
		Scopes(scopes).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)

	// Check for errors (usually driver not valid)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Redirect with authURL
	c.Redirect(http.StatusFound, authURL)
}
