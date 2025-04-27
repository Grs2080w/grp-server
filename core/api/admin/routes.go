package admin

import (
	get "github.com/Grs2080w/grp_server/core/api/admin/getLogs"

	"github.com/Grs2080w/grp_server/core/middleware/auth"
	"github.com/Grs2080w/grp_server/core/middleware/cache"
	"github.com/Grs2080w/grp_server/core/middleware/log"
	"github.com/gin-gonic/gin"
)


func RegisterRoutes(rg *gin.RouterGroup) {
	admin := rg.Group("/admin")
	admin.Use(auth.AuthMiddle())
	admin.Use(log.LogMiddleware())
	admin.Use(cache.CacheMiddleware(60))
	admin.GET("/logs", get.GetLogsHandler)
}
