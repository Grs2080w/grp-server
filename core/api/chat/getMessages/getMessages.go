package getMessages

// "github.com/Grs2080w/grp_server/core/api/chat/getMessages"

import (
	"context"
	"log"
	"sort"

	"github.com/Grs2080w/grp_server/core/domains/chat/getMessages"
	"github.com/gin-gonic/gin"

	q "github.com/Grs2080w/grp_server/core/db/dynamo/chat/query"
	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
)

// QueryMessagesHandler godoc
// @Summary Get all messages
// @Description Get all messages with the header token
// @Tags chat
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} []getMessages.Message
// @Failure 400 {object} getMessages.ErrorResponse
// @Router /chat [get]
func QueryMessagesHandler(c *gin.Context) {

	username := c.GetString("username")

	mess, err := q.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.Query(context.TODO(), username + "#CHAT")

	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"error": "error internal server"})
		return
	}

	messagesParsed, err := getMessages.ParseMessages(mess)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"error": "error internal server"})
	}

	sort.Slice(messagesParsed, func(a, b int) bool {
		return messagesParsed[a].Date < messagesParsed[b].Date
	})

	c.JSON(200, messagesParsed)

}