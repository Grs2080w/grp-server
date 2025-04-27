package health

import (
	"github.com/gin-gonic/gin"
)

// HealthCheckHandler godoc
// @Summary Check server health
// @Description Check server health
// @Tags health
// @Accept json
// @Produce json
// @Router /health [get]
func HealthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}