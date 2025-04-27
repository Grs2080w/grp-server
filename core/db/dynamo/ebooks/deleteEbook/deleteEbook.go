package deleteEbook

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	model "github.com/Grs2080w/grp_server/core/db/dynamo/ebooks/.model"
	modelTag "github.com/Grs2080w/grp_server/core/db/dynamo/tags/.model"
	deleteT "github.com/Grs2080w/grp_server/core/db/dynamo/tags/deleteTag"

	modelM "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/.model"
	aim "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/addIfMissing"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}


func (basics TableBasics) DeleteEbook(ctx context.Context, ebook model.Ebook, cfg aws.Config) error {

	// Metrics

	err := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: ebook.Username + "#METRICS", Sk: "METRICS#RECORDS_PER_DOMAIN#EBOOKS", Value: -1, Username: ebook.Username, Domain: "EBOOKS", Type: "RECORDS_PER_DOMAIN"})

	if err != nil {
		log.Fatalf("err deleting metric: %v", err)
	}

	err2 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: ebook.Username + "#METRICS", Sk: "METRICS#FILES_PER_EXTENSION#" + ebook.Extension, Value: -1, Username: ebook.Username, Domain: "EBOOKS", Type: "FILES_PER_EXTENSION"})

	if err2 != nil {
		log.Fatalf("err deleting metric: %v", err)
	}

	err3 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: ebook.Username + "#METRICS", Sk: "METRICS#STORAGE_PER_TYPE#" + ebook.Extension, Value: -ebook.Size, Username: ebook.Username, Domain: "EBOOKS", Type: "STORAGE_PER_TYPE"})  

	if err3 != nil {
		log.Fatalf("err deleting metric: %v", err)
	}

	err4 := aim.AddIfMissing(cfg, basics.TableName, modelM.Metrics{Pk: ebook.Username + "#METRICS", Sk: "METRICS#STORAGE_PER_DOMAIN#EBOOKS", Value: -ebook.Size, Username: ebook.Username, Domain: "EBOOKS", Type: "STORAGE_PER_DOMAIN"})  

	if err4 != nil {
		log.Fatalf("err adding metric: %v", err)
	}


	// Tags

	for _, tag := range ebook.Tags {
		err := deleteT.TableBasics{DynamoDbClient: basics.DynamoDbClient, TableName: basics.TableName}.DeleteTag(ctx, modelTag.Tags{
			Pk: ebook.Username + "#TAGS",
			Sk: "TAGS#" + ebook.Id + "#" + tag,
			Domain: "ebooks",
			Item_id: ebook.Id,
			Tag: ebook.Username + "#" + tag,
			Date: ebook.Date,
			Username: ebook.Username,
		}, cfg)
		if err != nil {
			log.Printf("Couldn't add tag to table. Here's why: %v\n", err)
		}
	}


	// Ebooks
	

	_, err = basics.DynamoDbClient.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(basics.TableName), Key: ebook.GetKey(),
	})
	if err != nil {
		log.Printf("Couldn't delete %v from the table. Here's why: %v\n", ebook.Sk, err)
	}
	return err
}


