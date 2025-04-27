package queryMetric

import (
	"context"
	"encoding/json"
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


func ParseJson(obj interface{}) string {
	items, _ := json.Marshal(obj)
	return string(items)
}


func (basics TableBasics) GetMetric(ctx context.Context, pk string, sk string) (string, error) {
	metric := model.Metrics{Pk: pk, Sk: sk}
	response, err := basics.DynamoDbClient.GetItem(ctx, &dynamodb.GetItemInput{
		Key: metric.GetKey(), TableName: aws.String(basics.TableName),
	})
	if err != nil {
		log.Printf("Couldn't get info about %v. Here's why: %v\n", sk, err)
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &metric)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}
	movieJson := ParseJson(metric)
	return movieJson, err
}


