package addEbook

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Grs2080w/grp_server/core/domains/ebooks/addEbook"
	uuid "github.com/google/uuid"

	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	model "github.com/Grs2080w/grp_server/core/db/dynamo/ebooks/.model"
	add "github.com/Grs2080w/grp_server/core/db/dynamo/ebooks/addEbook"
	up "github.com/Grs2080w/grp_server/core/db/s3/upFile"
)

// EBookHandler godoc
// @Summary Add a ebook
// @Description Add a ebook with the request body
// @Tags ebook
// @Accept multipart/form-data
// @Produce json
// @Param ebook formData file true "Request body"
// @Param size formData string true "Request body"
// @Success 200 {object} addEbook.Response
// @Failure 400 {object} addEbook.ErrorResponse
// @Router /ebooks [post]
func AddEbookHandler(c *gin.Context) {

	username := c.GetString("username")

	fileReq, err := c.FormFile("ebook")
    
    if err != nil {
      log.Println(err)
      c.JSON(400, gin.H{"error": "error on uploading ebook"})
      return
    }

	size, err := addEbook.ParseSize(c.PostForm("size"))
	if err != nil {
	  log.Println(err)
	  c.JSON(400, gin.H{"error": "error on converting size"})
	  return
	}

	tags, err := addEbook.ParseTags(c.PostForm("tags"))
	if err != nil {
	  log.Println(err)
	  c.JSON(400, gin.H{"error": "error on converting tags"})
	  return
	}

    metaforms := addEbook.Request{
      Name: c.PostForm("name"),
      Extension: c.PostForm("extension"),
      Size: size,
      Tags: tags,
    }

	new_id_version := uuid.New().String()

	path := "./storage/" + username + "#" + new_id_version

    c.SaveUploadedFile(fileReq, path)

	newEbook := model.Ebook{
		Pk: username + "#EBOOKS",
		Sk: "EBOOK#" + new_id_version,
		Username: username,
		Id: new_id_version,
		Name: metaforms.Name,
		Date: time.Now().Format(time.RFC3339),
		Type: "ebooks",
		Tags: metaforms.Tags,
		Size: metaforms.Size,
		Extension: metaforms.Extension,
	}

	errChan := make(chan error, 1)

	go func() {

		defer close(errChan)

		defer os.Remove(path)
		
		err = add.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.AddEbook(context.TODO(), newEbook, clientDb.CDB.Cfg)
	
		if err != nil {
			log.Println(err)
			errChan <- err
			return
		}
	
		err = up.BucketBasics{S3Client: clientDb.CDB.S3Client}.UploadFile(context.TODO(), clientDb.CDB.BucketName, new_id_version + "." + metaforms.Extension, path)
	
		if err != nil {
			log.Println(err)
			errChan <- err
			return
		}

		errChan <- nil

	}()

	go func() {
		err = <-errChan

		if err != nil {
			log.Println(err)
			return
		}

	}()	

	c.JSON(200, gin.H{
		"message": "Ebook adition scheduled",
		"ebook": newEbook,
	})

}