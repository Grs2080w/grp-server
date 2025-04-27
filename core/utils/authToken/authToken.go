package authToken

// "github.com/Grs2080w/grp_server/core/utils/authToken"

import (
	"errors"
	"strings"

	jwt "github.com/Grs2080w/grp_server/core/crypto/jwt_decode"
	"github.com/gin-gonic/gin"
)

func VerifyToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("missing token in headers")
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")

	user_username, err := jwt.DecodeToken(token); 
	if err != nil {
		return "", errors.New("invalid token")
	}

	return user_username, nil
}