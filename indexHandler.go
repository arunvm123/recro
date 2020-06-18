package main

import "github.com/gin-gonic/gin"

// Show homepage with login URL.
// TESTING PURPOSES ONLY
func (server *server) indexHandler(c *gin.Context) {
	c.Writer.Write([]byte("<html><head><title>Gocialite example</title></head><body>" +
		"<a href='/auth/github'><button>Login with GitHub</button></a><br>" +
		"</body></html>"))
}
