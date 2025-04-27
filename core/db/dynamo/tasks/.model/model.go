package model

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)


type Tasks struct {
	Pk string  `dynamodbav:"pk"`
	Sk  string  `dynamodbav:"sk"`
	Id string  `dynamodbav:"id"`
	Status string  `dynamodbav:"status"`
	Description string  `dynamodbav:"description"`
	Title string  `dynamodbav:"title"`
	Date string  `dynamodbav:"date"`
	Tags []string `dynamodbav:"tags"`
	Username string  `dynamodbav:"username"`
	Size float32 `dynamodbav:"size"`
}


func (i Tasks) GetKey() map[string]types.AttributeValue {
	pk, err := attributevalue.Marshal(i.Pk)
	if err != nil {
		panic(err)
	}
	sk, err := attributevalue.Marshal(i.Sk)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"pk": pk, "sk": sk}
}

