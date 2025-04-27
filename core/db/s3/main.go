package main

/*
import (
	u "example/uploadFile"
	d "example/downloadFile"
	df "example/deleteFile"

	"context"
	"fmt"
	"log"

	"os"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)



func main() {

	//
	//
	//
	// CREATE THE CLIENT TAHT WILL USED IN THE UPLOAD, DOWNLOAD AND DELETE FILES
	//
	//
	//

	ctx := context.TODO()

	 Load the shared AWS configuration (from ~/.aws/config)
	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile("default"))
	if err != nil {
		log.Fatalf("err loading aws config: %v", err)
	}

	client := s3.NewFromConfig(cfg)


	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Err getting current directory", err)
		return
	}


	//
	//
	//
	// VARS THAT WILL BE USED IN THE UPLOAD, DOWNLOAD AND DELETE FILES
	//
	//
	//
	//

	bucketName := "gab2080-server-ps"
	key := "CURRICULUM_ATUALIZADO_2Sergio.docx"
	file := "CURRICULUM_ATUALIZADO_2Sergio.docx"
	fileName := dir + "/files/" + file


	//
	//
	//
	//
	//UPLOAD FILE TO A BUCKET ____________________________________
	//
	//
	//
	//
	//


	err = u.BucketBasics{S3Client: client}.UploadFile(ctx, bucketName, objectKey, fileName)

	if err != nil {
		log.Fatalf("Couldn't upload file %v to %v:%v. Here's why: %v\n", fileName, bucketName, objectKey, err)
	}



	//
	//
	//
	//
	//DOWNLOAD FILE FROM A BUCKET ____________________________________
	//
	//
	//
	//

	err = d.BucketBasics{S3Client: client}.DownloadFile(ctx, bucketName, key, fileName)

	if err != nil {
		log.Fatalf("Couldn't download file %v from %v:%v. Here's why: %v\n", fileName, bucketName, key, err)
	}



	//
	//
	//
	//
	// DELETE FILE FROM A BUCKET ____________________________________
	//
	//
	//
	//

	deleted, err := df.S3Actions{S3Client: client}.DeleteObject(ctx, bucketName, key, "", false)

	if err != nil {
		log.Fatalf("Couldn't delete file %v from %v:%v. Here's why: %v\n", fileName, bucketName, key, err)
	}

	if deleted {
		fmt.Printf("Deleted file %v from %v:%v.\n", fileName, bucketName, key)
	}


}
*/