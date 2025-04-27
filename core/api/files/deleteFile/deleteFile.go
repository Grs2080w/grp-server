package deleteFile

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Grs2080w/grp_server/core/domains/files/deleteFile"

	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	model "github.com/Grs2080w/grp_server/core/db/dynamo/files/.model"
	up "github.com/Grs2080w/grp_server/core/db/dynamo/files/UpdateVersion"
	delDy "github.com/Grs2080w/grp_server/core/db/dynamo/files/deleteFile"
	get "github.com/Grs2080w/grp_server/core/db/dynamo/files/getFile"
	delS3 "github.com/Grs2080w/grp_server/core/db/s3/delFile"
)

// DeleteFileHandler godoc
// @Summary Delete a file
// @Description Delete a file with the request body
// @Tags file
// @Accept json
// @Produce json
// @Param user body deleteFile.UserRequest true "Request body"
// @Success 200 {object} deleteFile.SuccessResponse
// @Failure 400 {object} deleteFile.ErrorResponse
// @Router /files [delete]
func DelFileHandler(c *gin.Context) {
	
	username := c.GetString("username")

	var userRequest deleteFile.UserRequest
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

	fileParsed, err := deleteFile.ParseFile(file)
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

	errChan := make(chan error)

	if len(fileParsed.Versions) == 1 && fileParsed.Versions[0].Id == userRequest.Id {
		

		go func() {

			defer close(errChan)

			err = (&delDy.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).DeleteFile(context.TODO(), fileParsed, clientDb.CDB.Cfg)
			
			if err != nil {
				log.Printf("Failed to delete file dynamo: %v", err)
				errChan <- err
				return
			}
			_, err := (&delS3.S3Actions{S3Client: clientDb.CDB.S3Client}).DeleteObject(context.TODO(), clientDb.CDB.BucketName, userRequest.Id + "." + fileParsed.Extension, "", false)
	
			if err != nil {
				log.Printf("Failed to delete file s3: %v", err)
				errChan <- err
				return
			}

			errChan <- nil

		}()


	} else {
		
		fileVersions := fileParsed.Versions
		var newFileVersions []model.Version
	
		for _, version := range fileVersions {
	
			if version.Id != userRequest.Id {
				newFileVersions = append(newFileVersions, version)
			}
	
		}
	
		fileParsed.Versions = newFileVersions
		fileParsed.Versions[len(fileParsed.Versions)-1].Is_latest = true
	
		
		go func(){

			defer close(errChan)
	
			_, err = up.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.UpdateVersion(context.TODO(), fileParsed, "versions", fileParsed.Versions)
		
			if err != nil {
				log.Printf("Failed to update file: %v", err)
				errChan <- err
				return
			}
		
			_, err = (&delS3.S3Actions{S3Client: clientDb.CDB.S3Client}).DeleteObject(context.TODO(), clientDb.CDB.BucketName, userRequest.Id + "." + fileParsed.Extension, "", false)
			
			if err != nil {
				log.Printf("Failed to delete file s3: %v", err)
				errChan <- err
				return
			}

			errChan <- nil
			
		}()

	}

	go func() {
		err := <- errChan

		if err != nil {
			log.Println("Background delete file error:", err)
		}

	}()

		
	c.JSON(202, gin.H{"message": "file deletion scheduled"})
}