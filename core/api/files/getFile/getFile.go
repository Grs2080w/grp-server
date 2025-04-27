package getFile

import (
	"context"
	"log"
	"github.com/gin-gonic/gin"

	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	get "github.com/Grs2080w/grp_server/core/db/dynamo/files/getFile"
	
	"github.com/Grs2080w/grp_server/core/domains/files/getFile"
)

// GetFileHandler godoc
// @Summary Get a file
// @Description Get a file with the request body
// @Tags file
// @Accept json
// @Produce json
// @Param user body getFile.UserRequest true "Request body"
// @Success 200 {object} getFile.Files
// @Failure 400 {object} getFile.ErrorResponse
// @Router /files/one [get]
func GetFileHandler(c *gin.Context) {

	username := c.GetString("username")

	var userRequest getFile.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := userRequest.Validate()
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	file, err := (&get.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).GetFile(context.TODO(), username + "#FILES", "FILES#" + userRequest.Filename)

	if err != nil {
		log.Printf("Failed to get file: %v", err)
		c.JSON(500, gin.H{"error": "error on searching file"})
		return
	}

	fileParsed, err := getFile.ParseFile(file)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "an unexpected error occurred"})
		return
	}

	// if file not exists
	if fileParsed.Extension != userRequest.Type {
		c.JSON(404, gin.H{"error": "file not found"})
		return
	}

	c.JSON(200, fileParsed)

}