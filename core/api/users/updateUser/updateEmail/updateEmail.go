package updateEmail

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	clientDb "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	model "github.com/Grs2080w/grp_server/core/db/dynamo/users/.model"
	getUser "github.com/Grs2080w/grp_server/core/db/dynamo/users/getUser"
	upUser "github.com/Grs2080w/grp_server/core/db/dynamo/users/updateUser"
	"github.com/Grs2080w/grp_server/core/domains/users/updateUser/updateEmail"
)

// UserHandler godoc
// @Summary Update a user email
// @Description Update a user email with the request body
// @Tags user
// @Accept json
// @Produce json
// @Param user body updateEmail.UserRequest true "email in body and token in headers"
// @Success 200 {object} updateEmail.SuccessResponse
// @Failure 400 {object} updateEmail.ErrorResponse
// @Router /users/email [patch]
func UpdateupdateEmailHandler(c *gin.Context) {

	username := c.GetString("username")

	// get request body

	var user updateEmail.UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "error getting user"})
		return
	}

	userParsed, err := updateEmail.ParseUnmarshal(userGet)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update user

	_, err = upUser.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.UpdateUser(context.TODO(), model.Users{
		Pk: userParsed.Pk,
		Sk: userParsed.Sk,
	}, "email", user.Email)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user email updated"})
}
