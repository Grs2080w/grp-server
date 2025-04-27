package createPwd

import (
	"context"
	"log"

	pwd "github.com/Grs2080w/grp_server/core/crypto/pwd_manager"
	client "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	"github.com/Grs2080w/grp_server/core/domains/passwords/createPwd"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	rp "github.com/Grs2080w/grp_server/core/utils/ramdomPaswword"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/passwords/.model"
	add "github.com/Grs2080w/grp_server/core/db/dynamo/passwords/addPassword"
)

// @Summary Create a new password
// @Description Create a new password with a random password, this endpoint require a username in the header
// @Tags passwords
// @Accept  json
// @Produce  json
// @Param request body createPwd.Request true "Request body"
// @Success 200 {object} createPwd.SuccessResponse
// @Failure 400 {object} createPwd.ErrorResponse
// @Router /passwords [post]
func AddPwdHandler(c *gin.Context) {

	username := c.GetString("username")

	var req createPwd.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	err := req.Validate()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var password string = req.Password

	if req.Mode == "auto" {
		
		password, err = rp.RandomPassword(req.Params.Length, req.Params.Uppercase_letters, req.Params.Lowercase_letters, req.Params.Digits, req.Params.Special_characters)

		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to generate random password"})
			return
		}

	}

	key := req.Master + createPwd.ReverseString(req.Master)

	log.Print("Key: ", key)
	
	encryptedPassword, err := pwd.EncryptAES_GCM([]byte(password), []byte(key))
	
	if err != nil {
		log.Print("Failed to encrypt password:", err)
		return
	}


	errChan := make(chan error, 1)

	go func ()  {

		defer close(errChan)
		
		id_pwd := uuid.New().String()

		newPassword := model.Passwords{
			Pk: username + "#PASSWORDS",
			Sk: "PASSWORD#" + id_pwd,
			Id: id_pwd,
			Hash: encryptedPassword,
			Identifier: req.Identifier,
			Tags: req.Tags,
			Username: username,
			Size: req.Size,
		}
	
		err = add.TableBasics{DynamoDbClient: client.CDB.DynamoClient, TableName: client.CDB.TableName}.AddPassword(context.TODO(), newPassword, client.CDB.Cfg)
	
		if err != nil {
			errChan <- err
			return
		}

		errChan <- nil

	}()

	go func ()  {
		
		err := <- errChan

		if err != nil {
			log.Print("Failed to add password to DynamoDB:", err)
			return
		}


	}()


	c.JSON(200, gin.H{"message": "Password ading scheduled"})
}