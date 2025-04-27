package deleteEbook

import (
	"context"
	"log"

	del "github.com/Grs2080w/grp_server/core/db/dynamo/ebooks/deleteEbook"
	get "github.com/Grs2080w/grp_server/core/db/dynamo/ebooks/getEbook"
	delS3 "github.com/Grs2080w/grp_server/core/db/s3/delFile"

	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	"github.com/Grs2080w/grp_server/core/domains/ebooks/deleteEbook"
	"github.com/gin-gonic/gin"
)

// DeleteEbookHandler godoc
// @Summary Delete a ebook
// @Description Delete a ebook with the request body
// @Tags ebook
// @Accept json
// @Produce json
// @Param id path string true "Ebook id"
// @Success 200 {object} getEbooks.Ebook
// @Failure 400 {object} getEbooks.ErrorResponse
// @Router /ebooks/{id} [delete]
func DeleteEbookHandler(c *gin.Context) {

	username := c.GetString("username")

	id_ebook := c.Param("id")

	if id_ebook == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}

	ebookGet, err := (&get.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).GetEbook(context.TODO(), username + "#EBOOKS", "EBOOK#" + id_ebook)
	
	if err != nil {
		log.Println("Error getting ebook:", err)
		c.JSON(500, gin.H{"error": "internal error server"})
		return
	}

	ebook, err := deleteEbook.ParseEbook(ebookGet)

	if err != nil {
		log.Println("Error parsing ebook:", err)
		c.JSON(500, gin.H{"error": "internal error server"})
		return
	}

	c.JSON(200, gin.H{"message": "ebook deletion scheduled"})

	errChan := make(chan error, 1)

	go func() {
		
		defer close(errChan)

		err = (&del.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).DeleteEbook(context.TODO(), ebook, clientDb.CDB.Cfg)
		
		if err != nil {
			log.Println("Error deleting ebook from dynamo:", err)
			c.JSON(500, gin.H{"error": "internal error server"})
			errChan <- err
			return
		}
	
		_, err = delS3.S3Actions{S3Client: clientDb.CDB.S3Client}.DeleteObject(context.TODO(), clientDb.CDB.BucketName, ebook.Id + "." + ebook.Extension, "", false)
	
		if err != nil {
			log.Println("Error deleting ebook from S3:", err)
			c.JSON(500, gin.H{"error": "internal error server"})
			errChan <- err
			return
		}

		errChan <- nil
	}()

	go func() {
		err := <-errChan
	
		if err != nil {
			log.Println("Background delete ebook error:", err)
		}
	}()

}