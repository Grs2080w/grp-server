package addMessage

// "github.com/Grs2080w/grp_server/core/api/chat/addMessage"

import (
	"context"
	"log"
	"time"

	"github.com/Grs2080w/grp_server/core/domains/chat/addMessage"
	"github.com/gin-gonic/gin"

	"github.com/google/uuid"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/chat/.model"
	add "github.com/Grs2080w/grp_server/core/db/dynamo/chat/addMessage"
	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
)

// MessageHandler godoc
// @Summary Add a message
// @Description Add a message with the request body
// @Tags chat
// @Accept json
// @Produce json
// @Param user body addMessage.Request true "Request body"
// @Success 200 {object} addMessage.SuccessResponse
// @Failure 400 {object} addMessage.ErrorResponse
// @Router /chat [post]
func AddMessageHandler(c *gin.Context) {

	username := c.GetString("username")

	var message addMessage.Request
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	err := message.ParseReq()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id_message := uuid.New().String()

	newMessage := model.Message{
		Pk: username + "#CHAT",
		Sk: "CHAT#" + id_message,
		Id: id_message,
		Message: message.Message,
		Date: time.Now().String(),
		Hour: time.Now().Format("15:04:05"),
		Username: username,
		Size: message.Size,
	}

	errChan := make(chan error, 1)
		
	go func ()  {

		defer close(errChan)
		
		err = add.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.AddMessage(context.TODO(), newMessage, clientDb.CDB.Cfg)
		
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


	c.JSON(200, gin.H{"message": "message senting schedule"})

}