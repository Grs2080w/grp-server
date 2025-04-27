package query

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/tasks/.model"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func ParseJson(obj interface{}) string {
	items, _ := json.Marshal(obj)
	return string(items)
}

func (basics TableBasics) Query(ctx context.Context, pk string) (string, error) {
	var err error
	var response *dynamodb.QueryOutput
	var tasks []model.Tasks
	keyEx := expression.Key("pk").Equal(expression.Value(pk))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		log.Printf("Couldn't build expression for query. Here's why: %v\n", err)
	} else {
		queryPaginator := dynamodb.NewQueryPaginator(basics.DynamoDbClient, &dynamodb.QueryInput{
			TableName:                 aws.String(basics.TableName),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			KeyConditionExpression:    expr.KeyCondition(),
		})
		for queryPaginator.HasMorePages() {
			response, err = queryPaginator.NextPage(ctx)
			if err != nil {
				log.Printf("Couldn't query for items released in %v. Here's why: %v\n", pk, err)
				break
			} else {
				var itemPage []model.Tasks
				err = attributevalue.UnmarshalListOfMaps(response.Items, &itemPage)
				if err != nil {
					log.Printf("Couldn't unmarshal query response. Here's why: %v\n", err)
					break
				} else {
					tasks = append(tasks, itemPage...)
				}
			}
		}
	
	}
	tasksJson := ParseJson(tasks)
	return tasksJson, err
}


