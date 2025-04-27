package addTask

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	modelTag "github.com/Grs2080w/grp_server/core/db/dynamo/tags/.model"
	addTag "github.com/Grs2080w/grp_server/core/db/dynamo/tags/addTag"
	model "github.com/Grs2080w/grp_server/core/db/dynamo/tasks/.model"

	modelM "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/.model"
	aim "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/addIfMissing"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func (basics TableBasics) AddTask(ctx context.Context, Task model.Tasks, cfg aws.Config) error {

	// metrics

	err := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: Task.Username + "#METRICS", Sk: "METRICS#RECORDS_PER_DOMAIN#TASKS", Value: 1, Username: Task.Username, Domain: "TASKS", Type: "RECORDS_PER_DOMAIN"})

	if err != nil {
		log.Fatalf("err adding metric: %v", err)
	}

	err2 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: Task.Username + "#METRICS", Sk: "METRICS#STORAGE_PER_DOMAIN#TASKS", Value: Task.Size, Username: Task.Username, Domain: "TASKS", Type: "STORAGE_PER_DOMAIN"})

	if err2 != nil {
		log.Fatalf("err adding metric: %v", err)
	}

	// tags

	for _, tag := range Task.Tags {
		err := addTag.TableBasics{DynamoDbClient: basics.DynamoDbClient, TableName: basics.TableName}.AddTag(ctx, modelTag.Tags{
			Pk: Task.Username + "#TAGS",
			Sk: "TAGS#" + Task.Id + "#" + tag,
			Domain: "tasks",
			Item_id: Task.Id,
			Tag: Task.Username + "#" + tag,
			Date: Task.Date,
			Username: Task.Username,
		}, cfg)
		if err != nil {
			log.Printf("Couldn't add tag to table. Here's why: %v\n", err)
		}
	}

	// tasks

	task, err := attributevalue.MarshalMap(Task)
	if err != nil {
		panic(err)
	}
	_, err = basics.DynamoDbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(basics.TableName), Item: task,
	})
	if err != nil {
		log.Printf("Couldn't add task to table. Here's why: %v\n", err)
	}
	return err
}


