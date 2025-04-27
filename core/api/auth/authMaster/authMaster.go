package authMaster

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Grs2080w/grp_server/core/domains/auth/authMaster"
	
	hashV "github.com/Grs2080w/grp_server/core/crypto/hash_password"
	jwt "github.com/Grs2080w/grp_server/core/crypto/jwt_decode"
	jwt_encode "github.com/Grs2080w/grp_server/core/crypto/jwt_encode_acess"
	clientDb "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	getU "github.com/Grs2080w/grp_server/core/db/dynamo/users/getUser"
)

// AuthHandler godoc
// @Summary Auth and verify the master_password of the user
// @Description Validate the master_password sent by the user and return a temp jwt with more duration
// @Tags auth
// @Accept json
// @Produce json
// @Param user body authMaster.UserAuthMaster true "User token and master_password"
// @Success 200 {object} authMaster.SuccessResponse
// @Failure 400 {object} authMaster.ErrorResponse
// @Router /auth/master [post]
func AuthHandler(c *gin.Context) {

	var user authMaster.UserAuthMaster
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := user.Verify() ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_username, err := jwt.DecodeToken(user.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	} 

	userResponse, err := (&getU.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).GetUser(context.TODO(), user_username + "#" + "PROFILE", "USERS" + "#" + user_username)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	userParsed := authMaster.ParseUnmarshal(userResponse)

	if !hashV.VerifyPassword(user.Master_password, userParsed.Master_password_hash) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid master_password"})
		return
	}

	token := jwt_encode.Token{Username: user_username}.CreateAcessToken()

	c.JSON(http.StatusCreated, gin.H{
		"status": 200,
		"data": token,
	})
}
