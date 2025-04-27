package addPassword

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/passwords/.model"
	modelTag "github.com/Grs2080w/grp_server/core/db/dynamo/tags/.model"
	addTag "github.com/Grs2080w/grp_server/core/db/dynamo/tags/addTag"

	modelM "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/.model"
	aim "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/addIfMissing"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func (basics TableBasics) AddPassword(ctx context.Context, Password model.Passwords, cfg aws.Config) error {

	// metrics

	err := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: Password.Username + "#METRICS", Sk: "METRICS#RECORDS_PER_DOMAIN#PASSWORDS", Value: 1, Username: Password.Username, Domain: "PASSWORDS", Type: "RECORDS_PER_DOMAIN"})

	if err != nil {
		log.Fatalf("err adding metric: %v", err)
	}

	err2 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: Password.Username + "#METRICS", Sk: "METRICS#STORAGE_PER_DOMAIN#PASSWORDS", Value: Password.Size, Username: Password.Username, Domain: "PASSWORDS", Type: "STORAGE_PER_DOMAIN"})

	if err2 != nil {
		log.Fatalf("err adding metric: %v", err)
	}


	// tags

	for _, tag := range Password.Tags {
		err := addTag.TableBasics{DynamoDbClient: basics.DynamoDbClient, TableName: basics.TableName}.AddTag(ctx, modelTag.Tags{
			Pk: Password.Username + "#TAGS",
			Sk: "TAGS#" + Password.Id + "#" + tag,
			Domain: "passwords",
			Item_id: Password.Id,
			Tag: Password.Username + "#" + tag,
			Date: time.Now().String(),
			Username: Password.Username,
		}, cfg)
		if err != nil {
			log.Printf("Couldn't add tag to table. Here's why: %v\n", err)
		}
	}


	// passwords


	password, err := attributevalue.MarshalMap(Password)
	if err != nil {
		panic(err)
	}
	_, err = basics.DynamoDbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(basics.TableName), Item: password,
	})
	if err != nil {
		log.Printf("Couldn't add password to table. Here's why: %v\n", err)
	}
	return err
}


