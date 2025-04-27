package authSecret

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Grs2080w/grp_server/core/domains/auth/authSecret"
	jwt "github.com/Grs2080w/grp_server/core/crypto/jwt_decode"
	jwt_encode "github.com/Grs2080w/grp_server/core/crypto/jwt_encode_acess"
	clientDb "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	getU "github.com/Grs2080w/grp_server/core/db/dynamo/users/getUser"
	scode "github.com/Grs2080w/grp_server/core/utils/secret_deterministic"
)

// AuthHandler godoc
// @Summary Auth and verify the secret_code of the user
// @Description Validate the secret_code sent by the user and return a temp jwt with more duration
// @Tags auth
// @Accept json
// @Produce json
// @Param user body authSecret.UserAuthSecret true "User token and secret_code"
// @Success 200 {object} authSecret.SuccessResponse
// @Failure 400 {object} authSecret.ErrorResponse
// @Router /auth/secret [post]
func AuthHandler(c *gin.Context) {

	var user authSecret.UserAuthSecret
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

	userParsed := authSecret.ParseUnmarshal(userResponse)

	if !scode.SecretDeterministic(user.Secret_code, userParsed.Secret_deterministic) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid secret code"})
		return
	}

	token := jwt_encode.Token{Username: user_username}.CreateAcessToken()

	c.JSON(http.StatusCreated, gin.H{
		"status": 200,
		"data": token,
	})
}
