package deletePwd

import (
	"context"
	"log"

	"github.com/Grs2080w/grp_server/core/domains/passwords/deletePwd"
	"github.com/gin-gonic/gin"

	client "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	del "github.com/Grs2080w/grp_server/core/db/dynamo/passwords/deletePassword"
	get "github.com/Grs2080w/grp_server/core/db/dynamo/passwords/getPassword"
)

// @Summary Delete a password
// @Description Delete a password with the id in the path, this endpoint require a username in the header
// @Tags passwords
// @Accept  json
// @Produce  json
// @Param id path string true "Password id"
// @Success 200 {object} deletePwd.SuccessResponse
// @Failure 400 {object} deletePwd.ErrorResponse
// @Router /passwords/{id} [delete]
func DeletePwdHandler(c *gin.Context) {

	username := c.GetString("username")

	id_pwd := c.Param("id")

	if id_pwd == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}

	errChan := make(chan error, 1)
	
	go func ()  {
		defer close(errChan)
		
		pwdGet, err := (&get.TableBasics{DynamoDbClient: client.CDB.DynamoClient,TableName: client.CDB.TableName}).GetPassword(context.TODO(), username + "#PASSWORDS", "PASSWORD#" + id_pwd)
		
		if err != nil {
			log.Println("Failed to query password:", err)
			errChan <- err
			return
		}
	
		pwd := deletePwd.ParsePassword(pwdGet)
	
		err = (&del.TableBasics{DynamoDbClient: client.CDB.DynamoClient,TableName: client.CDB.TableName}).DeletePassword(context.TODO(), pwd, client.CDB.Cfg)
	
		if err != nil {
			log.Println("Failed to delete password:", err)
			errChan <- err
			return
		}

		errChan <- nil
		
	}()

	go func() {
		err := <-errChan

		if err != nil {
			log.Println("Error occurred:", err)
			return
		}

	}()


	c.JSON(200, gin.H{"message": "Password deletion scheduled"})
}