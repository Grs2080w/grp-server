package addMessage

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/chat/.model"
	modelM "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/.model"
	aim "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/addIfMissing"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func (basics TableBasics) AddMessage(ctx context.Context, Message model.Message, cfg aws.Config) error {

	// metrics

	err := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: Message.Username + "#METRICS", Sk: "METRICS#RECORDS_PER_DOMAIN#MESSAGE", Value: 1, Username: Message.Username, Domain: "MESSAGE", Type: "RECORDS_PER_DOMAIN"})

	if err != nil {
		log.Fatalf("err adding metric: %v", err)
	}

	err2 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: Message.Username +"#METRICS", Sk: "METRICS#STORAGE_PER_DOMAIN#MESSAGE", Value: Message.Size, Username: Message.Username, Domain: "MESSAGE", Type: "STORAGE_PER_DOMAIN"})

	if err2 != nil {
		log.Fatalf("err adding metric: %v", err)
	}


	// message

	message, err := attributevalue.MarshalMap(Message)
	if err != nil {
		panic(err)
	}
	_, err = basics.DynamoDbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(basics.TableName), Item: message,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
	}
	return err
}


