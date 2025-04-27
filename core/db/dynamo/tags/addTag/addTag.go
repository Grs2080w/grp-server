package addTag

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	modelM "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/.model"
	aim "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/addIfMissing"
	model "github.com/Grs2080w/grp_server/core/db/dynamo/tags/.model"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}


func (basics TableBasics) AddTag(ctx context.Context, tag model.Tags, cfg aws.Config) error {

	err := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: tag.Username + "#METRICS", Sk: "METRICS#RECORDS_PER_DOMAIN#TAGS", Value: 1, Username: tag.Username, Domain: "TAGS", Type: "RECORDS_PER_DOMAIN"})

	if err != nil {
		log.Fatalf("err adding metric: %v", err)
	}

	item, err := attributevalue.MarshalMap(tag)
	if err != nil {
		panic(err)
	}
	_, err = basics.DynamoDbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(basics.TableName), Item: item,
	})
	if err != nil {
		log.Printf("Couldn't add tag to table. Here's why: %v\n", err)
	}
	return err
}


