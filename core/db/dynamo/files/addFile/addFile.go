package addFile

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/files/.model"
	modelTag "github.com/Grs2080w/grp_server/core/db/dynamo/tags/.model"
	addTag "github.com/Grs2080w/grp_server/core/db/dynamo/tags/addTag"

	modelM "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/.model"
	aim "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/addIfMissing"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}


func (basics TableBasics) AddFile(ctx context.Context, File model.Files, cfg aws.Config) error {

	// Metrics

	err := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: File.Username + "#METRICS", Sk: "METRICS#RECORDS_PER_DOMAIN#FILES", Value: 1, Username: File.Username, Domain: "FILES", Type: "RECORDS_PER_DOMAIN"})

	if err != nil {
		log.Fatalf("err adding metric: %v", err)
	}

	err2 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: File.Username + "#METRICS", Sk: "METRICS#FILES_PER_EXTENSION#" + File.Extension, Value: 1, Username: File.Username, Domain: "FILES", Type: "FILES_PER_EXTENSION"})

	if err2 != nil {
		log.Fatalf("err adding metric: %v", err)
	}

	err3 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: File.Username + "#METRICS", Sk: "METRICS#STORAGE_PER_TYPE#" + File.Extension, Value: File.Versions[len(File.Versions)-1].Size, Username: File.Username, Domain: "FILES", Type: "STORAGE_PER_TYPE"})  

	if err3 != nil {
		log.Fatalf("err adding metric: %v", err)
	}

	err4 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: File.Username + "#METRICS", Sk: "METRICS#STORAGE_PER_DOMAIN#FILES", Value: File.Versions[len(File.Versions)-1].Size, Username: File.Username, Domain: "FILES", Type: "STORAGE_PER_DOMAIN"})  

	if err4 != nil {
		log.Fatalf("err adding metric: %v", err)
	}


	// Tags

	for _, tag := range File.Tags {
		err := addTag.TableBasics{DynamoDbClient: basics.DynamoDbClient, TableName: basics.TableName}.AddTag(ctx, modelTag.Tags{
			Pk: File.Username + "#TAGS",
			Sk: "TAGS#" + File.Versions[len(File.Versions)-1].Id + "#" + tag,
			Domain: "files",
			Item_id: File.Versions[len(File.Versions)-1].Id,
			Tag: File.Username + "#" + tag,
			Date: File.Versions[len(File.Versions)-1].Date,
			Username: File.Username,
		}, cfg)
		if err != nil {
			log.Printf("Couldn't add tag to table. Here's why: %v\n", err)
		}
	}



	// File

	file, err := attributevalue.MarshalMap(File)
	if err != nil {
		panic(err)
	}
	_, err = basics.DynamoDbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(basics.TableName), Item: file,
	})
	if err != nil {
		log.Printf("Couldn't add file to table. Here's why: %v\n", err)
	}
	return err
}


