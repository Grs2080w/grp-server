package updateUser

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/users/.model"
)


type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func ParseJson(obj interface{}) string {
	items, _ := json.Marshal(obj)
	return string(items)
}

func (basics TableBasics) UpdateUser(ctx context.Context, user model.Users, toBeUpdated string, updatedValue string) (string, error) {
	var err error
	var response *dynamodb.UpdateItemOutput
	update := expression.Set(expression.Name(toBeUpdated), expression.Value(updatedValue))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		log.Printf("Couldn't build expression for update. Here's why: %v\n", err)
	} else {
		response, err = basics.DynamoDbClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
			TableName:                 aws.String(basics.TableName),
			Key:                       user.GetKey(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
			ReturnValues:              types.ReturnValueUpdatedNew,
		})
		if err != nil {
			log.Printf("Couldn't update user %v. Here's why: %v\n", user.Pk, err)
		} else {
			err = attributevalue.UnmarshalMap(response.Attributes, &user)
			if err != nil {
				log.Printf("Couldn't unmarshall update response. Here's why: %v\n", err)
			}
		}
	}
	userJson := ParseJson(user)
	return userJson, err
}


