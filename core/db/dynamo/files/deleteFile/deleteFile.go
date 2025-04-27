package deleteFile

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/files/.model"
	modelTag "github.com/Grs2080w/grp_server/core/db/dynamo/tags/.model"
	deleteT "github.com/Grs2080w/grp_server/core/db/dynamo/tags/deleteTag"

	modelM "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/.model"
	aim "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/addIfMissing"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}


func (basics TableBasics) DeleteFile(ctx context.Context, file model.Files, cfg aws.Config) error {


	// Metrics

	err := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: file.Username + "#METRICS", Sk: "METRICS#RECORDS_PER_DOMAIN#FILES", Value: -1, Username: file.Username, Domain: "FILES", Type: "RECORDS_PER_DOMAIN"})

	if err != nil {
		log.Fatalf("err deleting metric: %v", err)
	}

	err2 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: file.Username + "#METRICS", Sk: "METRICS#FILES_PER_EXTENSION#" + file.Extension, Value: -1, Username: file.Username, Domain: "FILES", Type: "FILES_PER_EXTENSION"})  

	if err2 != nil {
		log.Fatalf("err deleting metric: %v", err)
	}

	err3 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: file.Username + "#METRICS", Sk: "METRICS#STORAGE_PER_TYPE#" + file.Extension, Value: -file.Versions[len(file.Versions)-1].Size, Username: file.Username, Domain: "FILES", Type: "STORAGE_PER_TYPE"})  

	if err3 != nil {
		log.Fatalf("err deleting metric: %v", err)
	}


	err4 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: file.Username + "#METRICS", Sk: "METRICS#STORAGE_PER_DOMAIN#FILES", Value: -file.Versions[len(file.Versions)-1].Size, Username: file.Username, Domain: "FILES", Type: "STORAGE_PER_DOMAIN"})  

	if err4 != nil {
		log.Fatalf("err adding metric: %v", err)
	}


	// Tags


	for _, tag := range file.Tags {
		err := deleteT.TableBasics{DynamoDbClient: basics.DynamoDbClient, TableName: basics.TableName}.DeleteTag(ctx, modelTag.Tags{
			Pk: file.Username + "#TAGS",
			Sk: "TAGS#" + file.Versions[len(file.Versions)-1].Id + "#" + tag,
			Domain: "files",
			Item_id: file.Versions[len(file.Versions)-1].Id,
			Tag: file.Username + "#" + tag,
			Date: file.Versions[len(file.Versions)-1].Date,
			Username: file.Username,
		}, cfg)
		if err != nil {
			log.Printf("Couldn't add tag to table. Here's why: %v\n", err)
		}
	}


	// File

	_, err = basics.DynamoDbClient.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(basics.TableName), Key: file.GetKey(),
	})
	if err != nil {
		log.Printf("Couldn't delete %v from the table. Here's why: %v\n", file.Sk, err)
	}
	return err
}


