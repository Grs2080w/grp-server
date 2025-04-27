package queryTag

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
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

func (basics TableBasics) QueryMetric(ctx context.Context, sk string) (string, error) {
	var err error
	var response *dynamodb.QueryOutput
	var metric []model.Metrics
	keyEx := expression.Key("sk").Equal(expression.Value(sk))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		log.Printf("Couldn't build expression for query. Here's why: %v\n", err)
	} else {
		queryPaginator := dynamodb.NewQueryPaginator(basics.DynamoDbClient, &dynamodb.QueryInput{
			TableName:                 aws.String(basics.TableName),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			KeyConditionExpression:    expr.KeyCondition(),
			IndexName: aws.String("MetricsBySkIndex"),
		})
		for queryPaginator.HasMorePages() {
			response, err = queryPaginator.NextPage(ctx)
			if err != nil {
				log.Printf("Couldn't query for tags released in %v. Here's why: %v\n", sk, err)
				break
			} else {
				var skPage []model.Metrics
				err = attributevalue.UnmarshalListOfMaps(response.Items, &skPage)
				if err != nil {
					log.Printf("Couldn't unmarshal query response. Here's why: %v\n", err)
					break
				} else {
					metric = append(metric, skPage...)
				}
			}
		}
	}
	metricJson := ParseJson(metric)
	return metricJson, err
}


