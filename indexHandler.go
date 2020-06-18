package main

import "github.com/gin-gonic/gin"

// Show homepage with login URL.
// TESTING PURPOSES ONLY
func (server *server) indexHandler(c *gin.Context) {
	c.Writer.Write([]byte("<html><head><title>Gocialite example</title></head><body>" +
		"<a href='/auth/github'><button>Login with GitHub</button></a><br>" +
		"<a href='/auth/linkedin'><button>Login with LinkedIn</button></a><br>" +
		"<a href='/auth/facebook'><button>Login with Facebook</button></a><br>" +
		"<a href='/auth/google'><button>Login with Google</button></a><br>" +
		"<a href='/auth/bitbucket'><button>Login with Bitbucket</button></a><br>" +
		"<a href='/auth/amazon'><button>Login with Amazon</button></a><br>" +
		"<a href='/auth/amazon'><button>Login with Slack</button></a><br>" +
		"</body></html>"))
}
