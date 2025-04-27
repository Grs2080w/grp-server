package model

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)


type Logs struct {
	Pk string  `dynamodbav:"pk"`
	Sk  string  `dynamodbav:"sk"`
	Log_id string `dynamodbav:"log_id"`
	Timestamp int `dynamodbav:"timestamp"`
	Level string `dynamodbav:"level"`
	Domain string `dynamodbav:"domain"`
	Message string `dynamodbav:"message"`
	Metadata string `dynamodbav:"metadata"`
	Username string `dynamodbav:"username"`
	Ip_address string `dynamodbav:"ip_address"`
	Stack_trace string `dynamodbav:"stack_trace"`
	Operation string `dynamodbav:"operation"`
	Status_code string `dynamodbav:"status_code"`
}
 

func (i Logs) GetKey() map[string]types.AttributeValue {
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



