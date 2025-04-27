package deleteTask

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	modelTag "github.com/Grs2080w/grp_server/core/db/dynamo/tags/.model"
	deleteT "github.com/Grs2080w/grp_server/core/db/dynamo/tags/deleteTag"
	model "github.com/Grs2080w/grp_server/core/db/dynamo/tasks/.model"

	modelM "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/.model"
	aim "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/addIfMissing"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func (basics TableBasics) DeleteTask(ctx context.Context, task model.Tasks, cfg aws.Config) error {

	// metrics

	err := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: task.Username + "#METRICS", Sk: "METRICS#RECORDS_PER_DOMAIN#TASKS", Value: -1, Username: task.Username, Domain: "TASKS", Type: "RECORDS_PER_DOMAIN"})

	if err != nil {
		log.Fatalf("err adding metric: %v", err)
	}

	err2 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: task.Username + "#METRICS", Sk: "METRICS#STORAGE_PER_DOMAIN#TASKS", Value: -task.Size, Username: task.Username, Domain: "TASKS", Type: "STORAGE_PER_DOMAIN"})

	if err2 != nil {
		log.Fatalf("err adding metric: %v", err)
	}


	// tags

	for _, tag := range task.Tags {
		err := deleteT.TableBasics{DynamoDbClient: basics.DynamoDbClient, TableName: basics.TableName}.DeleteTag(ctx, modelTag.Tags{
			Pk: task.Username + "#TAGS",
			Sk: "TAGS#" + task.Id + "#" + tag,
			Domain: "tasks",
			Item_id: task.Id,
			Tag: task.Username + "#" + tag,
			Date: task.Date,
			Username: task.Username,
		}, cfg)
		if err != nil {
			log.Printf("Couldn't add tag to table. Here's why: %v\n", err)
		}
	}

	// tasks

	_, err = basics.DynamoDbClient.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(basics.TableName), Key: task.GetKey(),
	})
	if err != nil {
		log.Printf("Couldn't delete %v from the table. Here's why: %v\n", task.Sk, err)
	}
	return err
}


