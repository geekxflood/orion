// internal/httpclient/httpclient.go

package httpclient

import (
	"net/http"

	"github.com/geekxflood/orion/internal/localtypes"
	"github.com/labstack/echo/v4"
)

func RunClient(conf localtypes.Config) {
	e := echo.New()

	// Disable TLS verification if insecure flag is set
	if conf.Insecure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig.InsecureSkipVerify = true
	}

	// Set up routes

	// Default route returns 403 Forbidden
	e.GET("/*", func(c echo.Context) error {
		return c.String(http.StatusForbidden, "Forbidden")
	})

	// /targets route returns list of targets from config file
	e.GET("/targets", func(c echo.Context) error {
		return c.JSON(http.StatusOK, conf.Targets)
	})

	e.Logger.Fatal(e.Start(":" + conf.Port))
}
