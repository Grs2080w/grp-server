package model

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)


type Users struct {
	Pk string  `dynamodbav:"pk"`
	Sk  string  `dynamodbav:"sk"`
	Username string  `dynamodbav:"username"`
	Password string  `dynamodbav:"password"`
	Email string  `dynamodbav:"email"`
	Storage_used string  `dynamodbav:"storage_used"`
	Total_files string  `dynamodbav:"total_files"`
	Avatar_url string  `dynamodbav:"avatar_url"`
	Theme_preferences string  `dynamodbav:"theme_preferences"`
	Language string  `dynamodbav:"language"`
	Failed_logins string  `dynamodbav:"failed_logins"`
	Extra_verification string  `dynamodbav:"extra_verification"`
	Master_password_hash string  `dynamodbav:"master_password_hash"`
	Secret_deterministic string  `dynamodbav:"secret_deterministic"`
}
 

func (i Users) GetKey() map[string]types.AttributeValue {
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

