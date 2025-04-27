package getEbooks

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	q "github.com/Grs2080w/grp_server/core/db/dynamo/ebooks/query"
	"github.com/Grs2080w/grp_server/core/domains/ebooks/getEbooks"
)

// GetEbooksHandler godoc
// @Summary Get all ebooks
// @Description Get all ebooks
// @Tags ebook
// @Accept json
// @Produce json
// @Success 200 {object} []getEbooks.Ebook
// @Failure 400 {object} getEbooks.ErrorResponse
// @Router /ebooks [get]
func QueryEbooksHandler(c *gin.Context) {
	
	username := c.GetString("username")
	
	filesGet, err := q.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.Query(context.TODO(), username + "#EBOOKS")

	if err != nil {
		c.JSON(500, gin.H{"error": "internal error server"})
		return
	}

	files := getEbooks.ParseEbooks(filesGet)
	
	c.JSON(200, files)

}