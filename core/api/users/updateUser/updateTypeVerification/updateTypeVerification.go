package updateTypeVerification

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	clientDb "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	model "github.com/Grs2080w/grp_server/core/db/dynamo/users/.model"
	getUser "github.com/Grs2080w/grp_server/core/db/dynamo/users/getUser"
	upUser "github.com/Grs2080w/grp_server/core/db/dynamo/users/updateUser"
	"github.com/Grs2080w/grp_server/core/domains/users/updateUser/updateTypeVerification"
)

// UserHandler godoc
// @Summary Update a user Type verification
// @Description Update a user Type verification with the request body
// @Tags user
// @Accept json
// @Produce json
// @Param user body updateTypeVerification.UserRequest true "Type verification in body and token in headers"
// @Success 200 {object} updateTypeVerification.SuccessResponse
// @Failure 400 {object} updateTypeVerification.ErrorResponse
// @Router /users/type [patch]
func UpdateTypeHandler(c *gin.Context) {
	
	username := c.GetString("username")
	
	// get request body

	var user updateTypeVerification.UserRequest
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

	userParsed, err := updateTypeVerification.ParseUnmarshal(userGet)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update user

	_, err = upUser.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.UpdateUser(context.TODO(), model.Users{
		Pk: userParsed.Pk,
		Sk: userParsed.Sk,
	}, "extra_verification", user.Type)
	
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user type verification updated"})
}
