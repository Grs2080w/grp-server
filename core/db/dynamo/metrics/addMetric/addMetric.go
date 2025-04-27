package addMetric

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/.model"
)


type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}


func (basics TableBasics) AddMetric(ctx context.Context, metric model.Metrics) error {
	item, err := attributevalue.MarshalMap(metric)
	if err != nil {
		panic(err)
	}
	_, err = basics.DynamoDbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(basics.TableName), Item: item,
	})
	if err != nil {
		log.Printf("Couldn't add metric to table. Here's why: %v\n", err)
	}
	return err
}


