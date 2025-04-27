package getFiles

import (
	"context"
	"log"
	"github.com/gin-gonic/gin"

	"github.com/Grs2080w/grp_server/core/domains/files/getFiles"

	cDb "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	q "github.com/Grs2080w/grp_server/core/db/dynamo/files/query"

)

// GetFilesHandler godoc
// @Summary Get a files
// @Description Get a files with the request body
// @Tags file
// @Accept json
// @Produce json
// @Param user body getFiles.UserRequest true "Request body"
// @Success 200 {object} []getFiles.Files
// @Failure 400 {object} getFiles.ErrorResponse
// @Router /files [get]
func GetFilesHandler(c *gin.Context) {
	
	username := c.GetString("username")

	files, err := q.TableBasics{TableName: cDb.CDB.TableName, DynamoDbClient: cDb.CDB.DynamoClient}.Query(context.TODO(), username + "#FILES")

	if err != nil {
		log.Printf("Failed to get file: %v", err)
		c.JSON(500, gin.H{"error": "error on searching file"})
		return
	}

	filesParsed, err := getFiles.ParseFiles(files)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "an unexpected error occurred"})
		return
	}

	c.JSON(200, filesParsed)

}