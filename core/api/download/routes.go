package download

import (
	down "github.com/Grs2080w/grp_server/core/api/download/downloadFile"

	"github.com/Grs2080w/grp_server/core/middleware/auth"
	"github.com/Grs2080w/grp_server/core/middleware/log"
	"github.com/gin-gonic/gin"
)


func RegisterRoutes(rg *gin.RouterGroup) {
	download := rg.Group("/download")
	download.Use(auth.AuthMiddle())
	download.Use(log.LogMiddleware())
	download.GET("/", down.DownloadFileHandler)
}
