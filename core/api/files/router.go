package files

import (
	add "github.com/Grs2080w/grp_server/core/api/files/addFile"
	del "github.com/Grs2080w/grp_server/core/api/files/deleteFile"
	ex "github.com/Grs2080w/grp_server/core/api/files/exists"
	get "github.com/Grs2080w/grp_server/core/api/files/getFile"
	query "github.com/Grs2080w/grp_server/core/api/files/getFiles"

	// middle
	"github.com/Grs2080w/grp_server/core/middleware/auth"
	"github.com/Grs2080w/grp_server/core/middleware/cache"
	"github.com/Grs2080w/grp_server/core/middleware/log"

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(rg *gin.RouterGroup) {
	files := rg.Group("/files")
	files.Use(auth.AuthMiddle())
	files.Use(log.LogMiddleware())
	files.POST("/", add.AddFileHandler)
	files.GET("/exists", ex.ExistsHandler)
	files.DELETE("/", del.DelFileHandler)
	files.GET("/", cache.CacheMiddleware(120), query.GetFilesHandler)
	files.GET("/one", cache.CacheMiddleware(120), get.GetFileHandler)
}
