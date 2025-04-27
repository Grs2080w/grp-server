package addLogs

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/logs/.model"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func (basics TableBasics) AddLogs(ctx context.Context, logs model.Logs) error {
	item, err := attributevalue.MarshalMap(logs)
	if err != nil {
		panic(err)
	}
	_, err = basics.DynamoDbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(basics.TableName), Item: item,
	})
	if err != nil {
		log.Printf("Couldn't add log to table. Here's why: %v\n", err)
	}
	return err
}


