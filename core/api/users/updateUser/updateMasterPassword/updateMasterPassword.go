package updateMasterPassword

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	hash "github.com/Grs2080w/grp_server/core/crypto/hash_password"
	clientDb "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	model "github.com/Grs2080w/grp_server/core/db/dynamo/users/.model"
	getUser "github.com/Grs2080w/grp_server/core/db/dynamo/users/getUser"
	upUser "github.com/Grs2080w/grp_server/core/db/dynamo/users/updateUser"
	"github.com/Grs2080w/grp_server/core/domains/users/updateUser/updateMasterPassword"
)

// UserHandler godoc
// @Summary Update a user master password
// @Description Update a user master password with the request body
// @Tags user
// @Accept json
// @Produce json
// @Param user body updateMasterPassword.UserRequest true "master password in body and token in headers"
// @Success 200 {object} updateMasterPassword.SuccessResponse
// @Failure 400 {object} updateMasterPassword.ErrorResponse
// @Router /users/master [patch]
func UpdateMasterHandler(c *gin.Context) {

	username := c.GetString("username")
	
	// get request body

	var user updateMasterPassword.UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if ok, err := user.Verify(); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	// conection client dynamo

	userGet, err := (&getUser.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).GetUser(context.TODO(), username + "#" + "PROFILE", "USERS" + "#" + username)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error getting user"})
		return
	}

	userParsed, err := updateMasterPassword.ParseUnmarshal(userGet)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update user

	hash_master, err := hash.HashPassword(user.Master_password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error unexpected ocurred"})
		return
	}

	_, err = upUser.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.UpdateUser(context.TODO(), model.Users{
		Pk: userParsed.Pk,
		Sk: userParsed.Sk,
	}, "master_password_hash", hash_master)
	
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user master password updated"})
}
