package getTag

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/tags/.model"
)


type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}


func ParseJson(obj interface{}) string {
	items, _ := json.Marshal(obj)
	return string(items)
	
}


func (basics TableBasics) GetTag(ctx context.Context, pk string, sk string) (string, error) {
	tag := model.Tags{Pk: pk, Sk: sk}
	log.Print(tag)
	response, err := basics.DynamoDbClient.GetItem(ctx, &dynamodb.GetItemInput{
		Key: tag.GetKey(), TableName: aws.String(basics.TableName),
	})
	if err != nil {
		log.Printf("Couldn't get info about %v. Here's why: %v\n", sk, err)
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &tag)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}
	tagJson := ParseJson(tag)
	return tagJson, err
}


