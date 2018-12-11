package utility

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Render HTML, JSON or XML based on 'Accept' header of the request
// The Default is HTML
func Render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
