package model

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)


type Message struct {
	Pk string  `dynamodbav:"pk"`
	Sk  string  `dynamodbav:"sk"`
	Id string `dynamodbav:"id"`
	Message string `dynamodbav:"message"`
	Date string `dynamodbav:"date"`
	Hour string `dynamodbav:"hour"`
	Username string  `dynamodbav:"username"`
	Size float32 `dynamodbav:"size"`
}


func (i Message) GetKey() map[string]types.AttributeValue {
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
