package model

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Tags struct {
	Pk    string                 `dynamodbav:"pk"`
	Sk    string                 `dynamodbav:"sk"`
	Domain string                 `dynamodbav:"domain"`
	Item_id string                 `dynamodbav:"item_id"`
	Tag string                 `dynamodbav:"tag"`
	Date string                 `dynamodbav:"date"`
	Username string  `dynamodbav:"username"`
}


func (tags Tags) GetKey() map[string]types.AttributeValue {
	pk, err := attributevalue.Marshal(tags.Pk)
	if err != nil {
		panic(err)
	}
	sk, err := attributevalue.Marshal(tags.Sk)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"pk": pk, "sk": sk}
}


