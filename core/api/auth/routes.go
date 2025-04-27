package auth

import (
	authM "github.com/Grs2080w/grp_server/core/api/auth/authMaster"
	authO "github.com/Grs2080w/grp_server/core/api/auth/authOtp"
	authSe "github.com/Grs2080w/grp_server/core/api/auth/authSecret"
	authS "github.com/Grs2080w/grp_server/core/api/auth/authSimple"
	"github.com/Grs2080w/grp_server/core/middleware/log"
	"github.com/gin-gonic/gin"
)


func RegisterRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	auth.Use(log.LogMiddleware())
	auth.POST("/", authS.AuthHandler)
	auth.POST("/master", authM.AuthHandler)
	auth.POST("/secret", authSe.AuthHandler)
	auth.POST("/otp", authO.AuthHandler)
}
