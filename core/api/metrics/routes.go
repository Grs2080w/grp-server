package metrics

import (
	get "github.com/Grs2080w/grp_server/core/api/metrics/getMetrics"

	// middle
	"github.com/Grs2080w/grp_server/core/middleware/auth"
	"github.com/Grs2080w/grp_server/core/middleware/cache"
	"github.com/Grs2080w/grp_server/core/middleware/log"


	"github.com/gin-gonic/gin"
)


func RegisterRoutes(rg *gin.RouterGroup) {
	metrics := rg.Group("/metrics")
	metrics.Use(auth.AuthMiddle())
	metrics.Use(log.LogMiddleware())
	metrics.Use(cache.CacheMiddleware(120))
	metrics.GET("/", get.GetMetricsHandler)

}
