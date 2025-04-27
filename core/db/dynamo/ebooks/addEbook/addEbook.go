package addEbook


import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/ebooks/.model"
	modelTag "github.com/Grs2080w/grp_server/core/db/dynamo/tags/.model"
	addTag "github.com/Grs2080w/grp_server/core/db/dynamo/tags/addTag"

	modelM "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/.model"
	aim "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/addIfMissing"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func (basics TableBasics) AddEbook(ctx context.Context, Item model.Ebook, cfg aws.Config) error {

	// metrics

	err := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: Item.Username + "#METRICS", Sk: "METRICS#RECORDS_PER_DOMAIN#EBOOKS", Value: 1, Username: Item.Username, Domain: "EBOOKS", Type: "RECORDS_PER_DOMAIN"})

	if err != nil {
		log.Fatalf("err adding metric: %v", err)
	}

	err2 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: Item.Username + "#METRICS", Sk: "METRICS#FILES_PER_EXTENSION#" + Item.Extension, Value: 1, Username: Item.Username, Domain: "EBOOKS", Type: "FILES_PER_EXTENSION"})

	if err2 != nil {
		log.Fatalf("err adding metric: %v", err)
	}

	err3 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: Item.Username + "#METRICS", Sk: "METRICS#STORAGE_PER_TYPE#" + Item.Extension, Value: Item.Size, Username: Item.Username, Domain: "EBOOKS", Type: "STORAGE_PER_TYPE"})  

	if err3 != nil {
		log.Fatalf("err adding metric: %v", err)
	}

	err4 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: Item.Username + "#METRICS", Sk: "METRICS#STORAGE_PER_DOMAIN#EBOOKS", Value: Item.Size, Username: Item.Username, Domain: "EBOOKS", Type: "STORAGE_PER_DOMAIN"})  

	if err4 != nil {
		log.Fatalf("err adding metric: %v", err)
	}


	// tags

	for _, tag := range Item.Tags {
		err := addTag.TableBasics{DynamoDbClient: basics.DynamoDbClient, TableName: basics.TableName}.AddTag(ctx, modelTag.Tags{
			Pk: Item.Username + "#TAGS",
			Sk: "TAGS#" + Item.Id + "#" + tag,
			Domain: "ebooks",
			Item_id: Item.Id,
			Tag: Item.Username + "#" + tag,
			Date: Item.Date,
			Username: Item.Username,
		}, cfg)
		if err != nil {
			log.Printf("Couldn't add tag to table. Here's why: %v\n", err)
		}
	}


	// ebook

	item, err := attributevalue.MarshalMap(Item)
	if err != nil {
		panic(err)
	}
	_, err = basics.DynamoDbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(basics.TableName), Item: item,
	})
	if err != nil {
		log.Printf("Couldn't add Ebook to table. Here's why: %v\n", err)
	}
	return err
}


