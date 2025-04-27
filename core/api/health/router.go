package health

import (
	g "github.com/Grs2080w/grp_server/core/api/health/health"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	health := rg.Group("/health")
	health.GET("/", g.HealthCheckHandler)
	
}