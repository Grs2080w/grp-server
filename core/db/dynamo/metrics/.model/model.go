package model

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)


type Metrics struct {
	Pk string `json:"Pk" dynamodbav:"pk"`
	Sk  string  `json:"Sk" dynamodbav:"sk"`
	Value float32  `json:"Value" dynamodbav:"value"`
	Username string  `json:"Username" dynamodbav:"username"`
	Domain string  `json:"Domain" dynamodbav:"domain"`
	Type string  `json:"Type" dynamodbav:"type"`
}
 

func (i Metrics) GetKey() map[string]types.AttributeValue {
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



