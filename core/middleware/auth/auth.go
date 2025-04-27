package auth

// "github.com/Grs2080w/grp_server/core/middleware/auth"

import (
	"github.com/Grs2080w/grp_server/core/utils/authToken"
	"github.com/gin-gonic/gin"
)



func AuthMiddle() gin.HandlerFunc {

	return func(c *gin.Context) {

		username, err := authToken.VerifyToken(c)

		if err != nil {
			c.JSON(401, gin.H{
				"error":  err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("username", username)
		c.Next()

	}

}