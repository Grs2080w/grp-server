package getMessage

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/chat/.model"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func ParseJson(obj interface{}) string {
	items, _ := json.Marshal(obj)
	return string(items)
}

func (basics *TableBasics) GetMessage(ctx context.Context, pk string, sk string) (string, error) {
	item := model.Message{
		Pk: pk,
		Sk: sk,
	}
	
	response, err := basics.DynamoDbClient.GetItem(ctx, &dynamodb.GetItemInput{
		Key: item.GetKey(), 
		TableName: aws.String(basics.TableName),
	})
	
	if err != nil {
		log.Printf("Couldn't get info about %v. Here's why: %v\n", pk, err)
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &item)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}
	messageJson := ParseJson(item)
	return messageJson, err
}
