package model

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)


type Version struct {
	Version string `dynamodbav:"version"`
	Date string `dynamodbav:"date"`
	Id string `dynamodbav:"id"`
	Is_latest bool `dynamodbav:"is_latest"`
	Size float32 `dynamodbav:"size"`
}


type Files struct {
	Pk string  `dynamodbav:"pk"`
	Sk  string  `dynamodbav:"sk"`
	Filename string `dynamodbav:"filename"`
    Type string `dynamodbav:"type"`
	Versions []Version `dynamodbav:"versions"`
	Username string  `dynamodbav:"username"`
	Tags []string `dynamodbav:"tags"`
	Extension string `dynamodbav:"extension"`
}
 

func (i Files) GetKey() map[string]types.AttributeValue {
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



