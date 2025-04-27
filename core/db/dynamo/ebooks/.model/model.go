package model

// "github.com/Grs2080w/grp_server/core/db/dynamo/ebooks/.model/model.go"

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)


type Ebook struct {
	Pk string  `dynamodbav:"pk"`
	Sk  string  `dynamodbav:"sk"`
	Username string  `dynamodbav:"username"`
	Id  string  `dynamodbav:"id"`
    Name string `dynamodbav:"name"` 
	Date string `dynamodbav:"date"`
	Type string `dynamodbav:"type"`
	Tags []string `dynamodbav:"tags"`
	Size float32 `dynamodbav:"size"`
	Extension string `dynamodbav:"extension"`	
}


func (i Ebook) GetKey() map[string]types.AttributeValue {
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



