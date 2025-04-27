package deletePassword

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/passwords/.model"
	modelTag "github.com/Grs2080w/grp_server/core/db/dynamo/tags/.model"
	deleteT "github.com/Grs2080w/grp_server/core/db/dynamo/tags/deleteTag"

	modelM "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/.model"
	aim "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/addIfMissing"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func (basics TableBasics) DeletePassword(ctx context.Context, password model.Passwords, cfg aws.Config) error {


	// metrics

	err := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: password.Username + "#METRICS", Sk: "METRICS#RECORDS_PER_DOMAIN#PASSWORDS", Value: -1, Username: password.Username, Domain: "PASSWORDS", Type: "RECORDS_PER_DOMAIN"})

	if err != nil {
		log.Fatalf("err adding metric: %v", err)
	}

	err2 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: password.Username + "#METRICS", Sk: "METRICS#STORAGE_PER_DOMAIN#PASSWORDS", Value: -password.Size, Username: password.Username, Domain: "PASSWORDS", Type: "STORAGE_PER_DOMAIN"})

	if err2 != nil {
		log.Fatalf("err adding metric: %v", err)
	}


	// tags

	for _, tag := range password.Tags {
		err := deleteT.TableBasics{DynamoDbClient: basics.DynamoDbClient, TableName: basics.TableName}.DeleteTag(ctx, modelTag.Tags{
			Pk: password.Username + "#TAGS",
			Sk: "TAGS#" + password.Id + "#" + tag,
			Domain: "passwords",
			Item_id: password.Id,
			Tag: password.Username + "#" + tag,
			Date: time.Now().String(),
			Username: password.Username,
		}, cfg)
		if err != nil {
			log.Printf("Couldn't add tag to table. Here's why: %v\n", err)
		}
	}



	// passwords

	_, err = basics.DynamoDbClient.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(basics.TableName), Key: password.GetKey(),
	})
	if err != nil {
		log.Printf("Couldn't delete %v from the table. Here's why: %v\n", password.Sk, err)
	}
	return err
}


