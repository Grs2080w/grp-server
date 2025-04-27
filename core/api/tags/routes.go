package tags

import (
	get "github.com/Grs2080w/grp_server/core/api/tags/getTag"
	q "github.com/Grs2080w/grp_server/core/api/tags/getTags"

	// middle
	"github.com/Grs2080w/grp_server/core/middleware/auth"
	"github.com/Grs2080w/grp_server/core/middleware/cache"
	"github.com/Grs2080w/grp_server/core/middleware/log"

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(rg *gin.RouterGroup) {
	tags := rg.Group("/tags")
	tags.Use(auth.AuthMiddle())
	tags.Use(log.LogMiddleware())
	tags.GET("/", cache.CacheMiddleware(120), q.GetTagsHandler)
	tags.GET("/:tag", cache.CacheMiddleware(120), get.GetTagHandler)
}
