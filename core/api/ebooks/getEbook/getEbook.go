package getEbook

import (
	"context"

	get "github.com/Grs2080w/grp_server/core/db/dynamo/ebooks/getEbook"

	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	"github.com/Grs2080w/grp_server/core/domains/ebooks/getEbook"
	"github.com/gin-gonic/gin"
)

// GetEbookHandler godoc
// @Summary Get a ebook
// @Description Get a ebook with the request body
// @Tags ebook
// @Accept json
// @Produce json
// @Param id path string true "Ebook id"
// @Success 200 {object} getEbooks.Ebook
// @Failure 400 {object} getEbooks.ErrorResponse
// @Router /ebooks/{id} [get]
func GetEbookHandler(c *gin.Context) {

	username := c.GetString("username")

	id_ebook := c.Param("id")

	if id_ebook == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}

	ebookGet, err := (&get.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).GetEbook(context.TODO(), username + "#EBOOKS", "EBOOK#" + id_ebook)
	
	if err != nil {
		c.JSON(500, gin.H{"error": "internal error server"})
		return
	}

	ebook, err := getEbook.ParseEbook(ebookGet)

	if err != nil {
		c.JSON(500, gin.H{"error": "internal error server"})
		return
	}

	c.JSON(200, ebook)
}