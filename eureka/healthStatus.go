package pkg

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
func Status(startTime time.Time) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"uptime": time.Since(startTime).String(),
			"status": "running",
		})
	}
}
