package addFile

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Grs2080w/grp_server/core/domains/files/addFile"
	uuid "github.com/google/uuid"

	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	model "github.com/Grs2080w/grp_server/core/db/dynamo/files/.model"
	update "github.com/Grs2080w/grp_server/core/db/dynamo/files/UpdateVersion"
	add "github.com/Grs2080w/grp_server/core/db/dynamo/files/addFile"
	get "github.com/Grs2080w/grp_server/core/db/dynamo/files/getFile"
	up "github.com/Grs2080w/grp_server/core/db/s3/upFile"
)

// FileHandler godoc
// @Summary Add a file
// @Description Add a file with the request body
// @Tags file
// @Accept json
// @Produce json
// @Param user body addFile.Request true "Request body"
// @Success 200 {object} addFile.Response
// @Failure 400 {object} addFile.ErrorResponse
// @Router /files [post]
func AddFileHandler(c *gin.Context) {

	username := c.GetString("username")

	fileReq, err := c.FormFile("file")
    
    if err != nil {
      log.Println(err)
      c.JSON(400, gin.H{"error": "error on uploading file"})
      return
    }

	size, err := addFile.ParseSize(c.PostForm("size"))
	if err != nil {
	  log.Println(err)
	  c.JSON(400, gin.H{"error": "error on converting size"})
	  return
	}

	tags, err := addFile.ParseTags(c.PostForm("tags"))
	if err != nil {
	  log.Println(err)
	  c.JSON(400, gin.H{"error": "error on converting tags"})
	  return
	}

    metaforms := addFile.UserRequest{
      Filename: c.PostForm("filename"),
      Type: c.PostForm("type"),
      Size: size,
      Tags: tags,
    }

	new_id_version := uuid.New().String()

	path := "./storage/" + username + "#" + metaforms.Filename

	// save the file
    c.SaveUploadedFile(fileReq, path)

	file, err := (&get.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).GetFile(context.TODO(), username + "#FILES", "FILES#" + metaforms.Filename)

	if err != nil {
		log.Printf("Failed to get file: %v", err)
		c.JSON(500, gin.H{"error": "error on searching file"})
		return
	}

	fileParsed, err := addFile.ParseFile(file)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "an unexpected error occurred"})
		return
	}


	errChan := make(chan error, 1)


	// if file not exists
	if fileParsed.Extension != metaforms.Type {

		go func ()  {

			defer os.Remove(path)
			defer close(errChan)
			
			newFile := model.Files{
				Pk: username + "#FILES",
				Sk: "FILES#" + metaforms.Filename,
				Username: username,
				Filename: metaforms.Filename,
				Type: "file",
				Tags: metaforms.Tags,
				Extension: metaforms.Type,
				Versions: []model.Version{
					{
					Version: "1", 
					Date: time.Now().UTC().Format(time.RFC3339),
					Id: new_id_version,
					Is_latest: true,
					Size: metaforms.Size,
					},
				},
			}
			
			err := add.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.AddFile(context.TODO(), newFile, clientDb.CDB.Cfg)
	
			if err != nil {
				log.Println(err)
				errChan <- err
				return
			}
	
			err = up.BucketBasics{S3Client: clientDb.CDB.S3Client}.UploadFile(context.TODO(), clientDb.CDB.BucketName, new_id_version + "." + metaforms.Type, path)
	
			if err != nil {
				log.Println(err)
				errChan <- err
				return
			}

			errChan <- nil

		}()


	} else {
		
		go func ()  {

			defer os.Remove(path)
			defer close(errChan)
			
			oldVersions := fileParsed.Versions
			lengthOldVersion := len(oldVersions)
			oldVersions[len(oldVersions)-1].Is_latest = false
		
			newVersion := model.Version{
				Version: strconv.Itoa(lengthOldVersion + 1), 
				Date: time.Now().UTC().Format(time.RFC3339),
				Id: new_id_version,
				Is_latest: true,
				Size: metaforms.Size,
			}
		
			newVersions := append(oldVersions, newVersion)
		
			fileParsed.Versions = newVersions
		
			_, err = update.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.UpdateVersion(context.TODO(), fileParsed, "versions", newVersions)
		
			if err != nil {
				log.Println(err)
				errChan <- err
				return
			}
		
			err = up.BucketBasics{S3Client: clientDb.CDB.S3Client}.UploadFile(context.TODO(), clientDb.CDB.BucketName, new_id_version + "." + metaforms.Type, path)
		
			if err != nil {
				log.Println(err)
				errChan <- err
				return
			}

			errChan <- nil

		}()
	}

	go func ()  {
		err = <-errChan
		
		if err != nil {
			log.Println("background error on add file: ", err)
		}

	}()

	c.JSON(200, gin.H{
		"message": "File adition scheduled",
		"file": fileParsed,
	})

}