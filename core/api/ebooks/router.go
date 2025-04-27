package ebooks

import (
	add "github.com/Grs2080w/grp_server/core/api/ebooks/addEbook"
	del "github.com/Grs2080w/grp_server/core/api/ebooks/deleteEbook"
	get "github.com/Grs2080w/grp_server/core/api/ebooks/getEbook"
	q "github.com/Grs2080w/grp_server/core/api/ebooks/getEbooks"

	// middle
	"github.com/Grs2080w/grp_server/core/middleware/auth"
	"github.com/Grs2080w/grp_server/core/middleware/cache"
	"github.com/Grs2080w/grp_server/core/middleware/log"

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(rg *gin.RouterGroup) {
	ebooks := rg.Group("/ebooks")
	ebooks.Use(auth.AuthMiddle())
	ebooks.Use(log.LogMiddleware())
	ebooks.POST("/", add.AddEbookHandler)
	ebooks.DELETE("/:id", del.DeleteEbookHandler)
	ebooks.GET("/", cache.CacheMiddleware(120), q.QueryEbooksHandler)
	ebooks.GET("/:id", cache.CacheMiddleware(120), get.GetEbookHandler)
}
