package updatePassword

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	r "github.com/Grs2080w/grp_server/core/db/redis"

	hash "github.com/Grs2080w/grp_server/core/crypto/hash_password"
	clientDb "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	model "github.com/Grs2080w/grp_server/core/db/dynamo/users/.model"
	getUser "github.com/Grs2080w/grp_server/core/db/dynamo/users/getUser"
	upUser "github.com/Grs2080w/grp_server/core/db/dynamo/users/updateUser"
	"github.com/Grs2080w/grp_server/core/domains/users/updateUser/updatePassword"
)

// UserHandler godoc
// @Summary Update a user password
// @Description Update a user password with the request body
// @Tags user
// @Accept json
// @Produce json
// @Param user body updatePassword.UserRequest true "password in body and token in headers"
// @Success 200 {object} updatePassword.SuccessResponse
// @Failure 400 {object} updatePassword.ErrorResponse
// @Router /users/password [patch]
func UpdatePasswordHandler(c *gin.Context) {

	username := c.GetString("username")

	// get request body
	var user updatePassword.UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if ok, err := user.Verify(); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	// verify the code

	code, err := r.R_get(username + "#OTP")

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "code not found, try send the code again"})
		return
	}

	if code != user.Code {
		attempts, err := r.R_get(username + "#ATTEMPTS#pass")
		attemptsInt, _ := strconv.Atoi(attempts)
		
		if err != nil {
			r.R_set(username + "#ATTEMPTS#pass", "1", 0)
		} else {
			r.R_set(username + "#ATTEMPTS#pass", fmt.Sprint(attemptsInt + 1), 0)
		}

		if attemptsInt == 2 {
			r.R_del(username + "#OTP")
			r.R_del(username + "#ATTEMPTS#pass")
		}


		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid code"})
		return
	}

	// conection client dynamo

	userGet, err := (&getUser.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).GetUser(context.TODO(), username + "#" + "PROFILE", "USERS" + "#" + username)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error getting user"})
		return
	}

	userParsed, err := updatePassword.ParseUnmarshal(userGet)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update user

	hsh_pass, err := hash.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}

	_, err = upUser.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.UpdateUser(context.TODO(), model.Users{
		Pk: userParsed.Pk,
		Sk: userParsed.Sk,
	}, "password", hsh_pass)
	
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user password updated"})
}
