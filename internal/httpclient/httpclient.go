// internal/httpclient/httpclient.go

package httpclient

import (
	"log"
	"net/http"
	"sync/atomic"

	"github.com/geekxflood/orion/internal/localtypes"
	"github.com/gin-gonic/gin"
)

// Client is a struct that holds an atomic.Value for the configuration.
type Client struct {
	Conf *atomic.Value
}

// RunClient runs the HTTP client with the provided configuration and endpoints.
// It sets up routes for handling HTTP requests and starts the HTTP server.
func (client *Client) RunClient(endpoints []localtypes.Endpoints) {

	// Disable Gin debug messages
	gin.SetMode(gin.ReleaseMode)

	// Create a new router with default middleware
	r := gin.Default()

	// Load the current configuration
	currentConf := client.Conf.Load().(*localtypes.Config)

	// Set up insecure client if needed
	if currentConf.Insecure {
		SetInsecure()
	}

	// Set up routes
	// Default route returns 403 Forbidden
	r.GET("", func(c *gin.Context) {
		c.String(http.StatusForbidden, "Forbidden")
	})

	// /targets route returns the list of targets
	r.GET("/targets", func(c *gin.Context) {
		// Load the current configuration for each request
		conf := client.Conf.Load().(*localtypes.Config)
		c.JSON(http.StatusOK, conf.Endpoints)
	})

	// /ready endpoint returns 200 OK
	r.GET("/ready", func(c *gin.Context) {
		c.String(http.StatusOK, "Ready")
	})

	// /health endpoint returns 200 OK
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// /config route returns config file contents
	r.GET("/config", func(c *gin.Context) {
		// Load the current configuration for each request
		conf := client.Conf.Load().(*localtypes.Config)
		c.JSON(http.StatusOK, conf)
	})

	// Start the HTTP server
	err := r.Run(":" + currentConf.Port)
	if err != nil {
		return
	}

	log.Println("Application available at http://localhost:" + currentConf.Port + "/")

}
