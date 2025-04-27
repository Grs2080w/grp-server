package authSimple

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Grs2080w/grp_server/core/domains/auth/authSimple"

	hashV "github.com/Grs2080w/grp_server/core/crypto/hash_password"
	jwt "github.com/Grs2080w/grp_server/core/crypto/jwt_encode_verify"
	clientDb "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	getU "github.com/Grs2080w/grp_server/core/db/dynamo/users/getUser"
)

// AuthHandler godoc
// @Summary Auth the user
// @Description Validate the user and the password and return a temp jwt
// @Tags auth
// @Accept json
// @Produce json
// @Param user body authSimple.UserAuth true "User Credentials"
// @Success 200 {object} authSimple.SuccessResponse
// @Failure 400 {object} authSimple.ErrorResponse
// @Router /auth [post]
func AuthHandler(c *gin.Context) {

	var user authSimple.UserAuth
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := user.Verify() ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := (&getU.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).GetUser(context.TODO(), user.Username + "#" + "PROFILE", "USERS" + "#" + user.Username)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	userParsed := authSimple.ParseUnmarshal(userResponse)
	
	if userParsed.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if !hashV.VerifyPassword(user.Password, userParsed.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	type_verif := userParsed.Extra_verification

	if type_verif != "master_password_hash" && type_verif != "secret_deterministic" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid type of verification"})
		return
	}

	token := jwt.Token{Username: user.Username, Type_verification: type_verif}.CreateVerifyToken()

	c.JSON(http.StatusCreated, gin.H{
		"status": 200,
		"type": type_verif,
		"data": token,
	})
}
