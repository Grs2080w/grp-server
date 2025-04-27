package createUser

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	hash "github.com/Grs2080w/grp_server/core/crypto/hash_password"
	clientDb "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	model "github.com/Grs2080w/grp_server/core/db/dynamo/users/.model"
	createU "github.com/Grs2080w/grp_server/core/db/dynamo/users/addUser"
	"github.com/Grs2080w/grp_server/core/domains/users/createUser"
)

// UserHandler godoc
// @Summary Create a user
// @Description Create a user with the request body
// @Tags user
// @Accept json
// @Produce json
// @Param user body createUser.UserRequest true "Request body"
// @Success 200 {object} createUser.SuccessResponse
// @Failure 400 {object} createUser.ErrorResponse
// @Router /users [post]
func CreateUserHandler(c *gin.Context) {

	var user createUser.UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if ok, err := user.Verify(); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_password, err := hash.HashPassword(user.Password)
	if err != nil {
		log.Fatal("err on create user hash password",err)
	}

	var master_password_hash, secret_deterministic string;

	if user.Extra_verification == "master_password_hash" {
		master_password_hash, _ = hash.HashPassword(user.Code)
	} else {
		secret_deterministic = user.Code
	}
	
	newUser := model.Users{
		Pk: user.Username + "#" + "PROFILE",
		Sk: "USERS#" + user.Username,
		Username: user.Username,
		Password: user_password,
		Email: user.Email,
		Storage_used: "0",
		Total_files: "0",
		Avatar_url: "",
		Theme_preferences: user.Theme_preferences,
		Language: user.Language,
		Failed_logins: "0",
		Extra_verification: user.Extra_verification,
		Master_password_hash: master_password_hash,
		Secret_deterministic: secret_deterministic,
	}


	errChan := make(chan error, 1)

	go func () {

		defer close(errChan)

		err = createU.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.AddUser(c, newUser)
	
		if err != nil {
			errChan <- err
			return
		}

		errChan <- nil

	}()

	go func ()  {
		
		err := <-errChan

		if err != nil {
			log.Println("Error creating user:", err)
			return
		}

	}()


	c.JSON(http.StatusCreated, gin.H{
		"message": "User creating scheduled",
	})
}
