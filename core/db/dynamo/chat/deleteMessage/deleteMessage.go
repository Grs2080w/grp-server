package deleteMessage

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/chat/.model"
	modelM "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/.model"
	aim "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/addIfMissing"
)


type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}


func (basics TableBasics) DeleteMessage(ctx context.Context, message model.Message, cfg aws.Config) error {

	// metrics

	err := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: message.Username +"#METRICS", Sk: "METRICS#RECORDS_PER_DOMAIN#MESSAGE", Value: -1, Username: message.Username, Domain: "MESSAGE", Type: "RECORDS_PER_DOMAIN"})

	if err != nil {
		log.Fatalf("err adding metric: %v", err)
	}

	err2 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: message.Username +"#METRICS", Sk: "METRICS#STORAGE_PER_DOMAIN#MESSAGE", Value: -message.Size, Username: message.Username, Domain: "MESSAGE", Type: "STORAGE_PER_DOMAIN"})

	if err2 != nil {
		log.Fatalf("err adding metric: %v", err)
	}

	// message

	_, err = basics.DynamoDbClient.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(basics.TableName), Key: message.GetKey(),
	})
	if err != nil {
		log.Printf("Couldn't delete %v from the table. Here's why: %v\n", message.Sk, err)
	}
	return err
}


