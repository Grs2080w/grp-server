package users

import (
	c "github.com/Grs2080w/grp_server/core/api/users/createUser"
	uau "github.com/Grs2080w/grp_server/core/api/users/updateUser/updateAvatarUrl"
	ue "github.com/Grs2080w/grp_server/core/api/users/updateUser/updateEmail"
	um "github.com/Grs2080w/grp_server/core/api/users/updateUser/updateMasterPassword"
	up "github.com/Grs2080w/grp_server/core/api/users/updateUser/updatePassword"
	ut "github.com/Grs2080w/grp_server/core/api/users/updateUser/updateTheme"
	utv "github.com/Grs2080w/grp_server/core/api/users/updateUser/updateTypeVerification"

	// middle
	"github.com/Grs2080w/grp_server/core/middleware/auth"
	"github.com/Grs2080w/grp_server/core/middleware/log"

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.Use(log.LogMiddleware())
	users.POST("/", c.CreateUserHandler)
	users.Use(auth.AuthMiddle())
	users.PATCH("/avatar_url", uau.UpdateAvatarUrlrHandler)
	users.PATCH("/email", ue.UpdateupdateEmailHandler)
	users.PATCH("/master", um.UpdateMasterHandler)
	users.PATCH("/theme", ut.UpdateThemeHandler)
	users.PATCH("/type", utv.UpdateTypeHandler)
	users.PATCH("/password", up.UpdatePasswordHandler)

}
