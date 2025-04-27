package queryUsername

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/logs/.model"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func ParseJson (obj interface{}) string {
	items, _ := json.Marshal(obj)
	return string(items)
}


func (basics TableBasics) QueryUsername(ctx context.Context, username string) (string, error) {
	var err error
	var response *dynamodb.QueryOutput
	var logs []model.Logs
	keyEx := expression.Key("username").Equal(expression.Value(username))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		log.Printf("Couldn't build expression for query. Here's why: %v\n", err)
	} else {
		queryPaginator := dynamodb.NewQueryPaginator(basics.DynamoDbClient, &dynamodb.QueryInput{
			TableName:                 aws.String(basics.TableName),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			KeyConditionExpression:    expr.KeyCondition(),
			IndexName: aws.String("LogsByUsernameIndex"),
		})
		for queryPaginator.HasMorePages() {
			response, err = queryPaginator.NextPage(ctx)
			if err != nil {
				log.Printf("Couldn't query for logs released in %v. Here's why: %v\n", username, err)
				break
			} else {
				var logsPage []model.Logs
				err = attributevalue.UnmarshalListOfMaps(response.Items, &logsPage)
				if err != nil {
					log.Printf("Couldn't unmarshal query response. Here's why: %v\n", err)
					break
				} else {
					logs = append(logs, logsPage...)
				}
			}
		}
	}

	logJson := ParseJson(logs)
	return logJson, err
}


