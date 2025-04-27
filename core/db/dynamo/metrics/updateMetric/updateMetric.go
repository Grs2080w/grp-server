package updateMetric

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/.model"
	get "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/getMetric"
)


type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func ParseJson(obj interface{}) string {
	items, _ := json.Marshal(obj)
	return string(items)
}

func ParseUnmarshal(jsonString string) model.Metrics {
	var result model.Metrics
	json.Unmarshal([]byte(jsonString), &result)
	return result
}

func (basics TableBasics) UpdateMetric(ctx context.Context, metrics model.Metrics) (string, error) {
	

	metricReturned, errr := get.TableBasics{DynamoDbClient: basics.DynamoDbClient, TableName: basics.TableName}.GetMetric(ctx, metrics.Pk, metrics.Sk)
	
	if errr != nil {
		log.Printf("Couldn't get info about %v. Here's why: %v\n", metrics.Sk, errr)
	} 

	metricsToUpdate := ParseUnmarshal(metricReturned)
	valueToUpdated := metricsToUpdate.Value + metrics.Value
	
	var err error
	var response *dynamodb.UpdateItemOutput
	update := expression.Set(expression.Name("value"), expression.Value(valueToUpdated))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		log.Printf("Couldn't build expression for update. Here's why: %v\n", err)
	} else {
		response, err = basics.DynamoDbClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
			TableName:                 aws.String(basics.TableName),
			Key:                       metrics.GetKey(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
			ReturnValues:              types.ReturnValueUpdatedNew,
		})
		if err != nil {
			log.Printf("Couldn't update metric %v. Here's why: %v\n", metrics.Sk, err)
		} else {
			err = attributevalue.UnmarshalMap(response.Attributes, &metrics)
			if err != nil {
				log.Printf("Couldn't unmarshall update response. Here's why: %v\n", err)
			}
		}
	}

	attributeJson := ParseJson(metrics)
	return attributeJson, err
}


