package getPassword

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/passwords/.model"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func ParseJson(obj interface{}) string {
	items, _ := json.Marshal(obj)
	return string(items)
}

func (basics *TableBasics) GetPassword(ctx context.Context, pk string, sk string) (string, error) {
	password := model.Passwords{
		Pk: pk,
		Sk: sk,
	}
	
	response, err := basics.DynamoDbClient.GetItem(ctx, &dynamodb.GetItemInput{
		Key: password.GetKey(), 
		TableName: aws.String(basics.TableName),
	})
	
	if err != nil {
		log.Printf("Couldn't get info about %v. Here's why: %v\n", pk, err)
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &password)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}

	passwordJson := ParseJson(password)
	return passwordJson, err
}