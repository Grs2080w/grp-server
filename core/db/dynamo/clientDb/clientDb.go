package clientDb

import (
	"context"
	"log"

	c "github.com/Grs2080w/grp_server/core/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// clientDb is a package that contains the logic for creating aws clients
// that can be used to interact with dynamo db and s3.
//
// The package exposes a single function, NewClient, which takes a
// configuration object and returns a ClientType object. The ClientType
// object contains references to the created clients as well as the
// configuration object.
//
// The ClientType object has methods to create a dynamo db table and
// an s3 bucket. These methods will create the resources if they do
// not already exist.


type ClientType struct {
	DynamoClient *dynamodb.Client
	S3Client *s3.Client
	Cfg aws.Config
	TableName string
	BucketName string
	LogTableName string
	PresignedClient *s3.PresignClient
}

var CDB ClientType

func InitDynamoClient() error {
	//Load the shared AWS configuration (from ~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))
	if err != nil {
		log.Fatalf("err loading aws config: %v", err)
		return err
	}

	CDB.DynamoClient = dynamodb.NewFromConfig(cfg)
	CDB.S3Client = s3.NewFromConfig(cfg)
	CDB.Cfg = cfg
	CDB.TableName = c.GetValueEnv("TABLE_NAME_DB")
	CDB.BucketName = c.GetValueEnv("BUCKET_NAME")
	CDB.PresignedClient = s3.NewPresignClient(s3.NewFromConfig(cfg))
	CDB.LogTableName = c.GetValueEnv("LOG_TABLE_NAME_DB")
	return nil
}