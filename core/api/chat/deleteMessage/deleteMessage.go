package deleteMessage

// "github.com/Grs2080w/grp_server/core/api/chat/deleteMessages"

import (
	"context"
	"log"

	"github.com/Grs2080w/grp_server/core/domains/chat/deleteMessage"
	"github.com/gin-gonic/gin"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/chat/.model"
	del "github.com/Grs2080w/grp_server/core/db/dynamo/chat/deleteMessage"
	get "github.com/Grs2080w/grp_server/core/db/dynamo/chat/getMessage"
	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
)

// DeleteMessageHandler godoc
// @Summary Delete a message
// @Description Delete a message with the header token and id in path
// @Tags chat
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Id of message"
// @Success 200 {object} deleteMessage.SucessResponse
// @Failure 400 {object} deleteMessage.ErrorResponse
// @Router /chat/{id} [delete]
func DeleteMessageHandler(c *gin.Context) {

	username := c.GetString("username")
	id_message := c.Param("id")

	if id_message == "" {
		c.JSON(400, gin.H{"error": "id message is required"})
		return
	}

	message, err := (&get.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).GetMessage(context.TODO(), username + "#CHAT", "CHAT#" + id_message)

	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"error": "error internal server"})
		return
	}

	messParsed, err := deleteMessage.ParseMessages(message)

	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"error": "error internal server"})
		return
	}


	errChan := make(chan error, 1)
	
	go func() {
		defer close(errChan)
		
		err = del.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.DeleteMessage(context.TODO(), model.Message{
			Pk: messParsed.Pk,
			Sk: messParsed.Sk,
			Size: messParsed.Size,
		}, clientDb.CDB.Cfg)
	
		if err != nil {
			log.Print(err)
			errChan <- err
			return
		}

		errChan <- nil

	}()

	go func ()  {
		
		err := <-errChan

		if err != nil {
			log.Print(err)
			return
		}

	}()


	c.JSON(200, gin.H{"status": "message deletion schuedule"})

}