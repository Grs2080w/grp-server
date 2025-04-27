package passwords

import (
	create "github.com/Grs2080w/grp_server/core/api/passwords/createPwd"
	del "github.com/Grs2080w/grp_server/core/api/passwords/deletePwd"
	get "github.com/Grs2080w/grp_server/core/api/passwords/getPwd"
	q "github.com/Grs2080w/grp_server/core/api/passwords/getPwds"

	// middle
	"github.com/Grs2080w/grp_server/core/middleware/auth"
	"github.com/Grs2080w/grp_server/core/middleware/cache"
	"github.com/Grs2080w/grp_server/core/middleware/log"

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(rg *gin.RouterGroup) {
	passwords := rg.Group("/passwords")
	passwords.Use(auth.AuthMiddle())
	passwords.Use(log.LogMiddleware())
	passwords.POST("/", create.AddPwdHandler)
	passwords.DELETE("/:id", del.DeletePwdHandler)
	passwords.GET("/", cache.CacheMiddleware(120), q.GetPwdsHandler)
	passwords.GET("/:id", cache.CacheMiddleware(120), get.GetPwdHandler)
}
