package addIfMissing

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	model "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/.model"
	up "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/updateMetric"

	add "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/addMetric"
	get "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/getMetric"
)

func ParseUnmarshal(jsonString string) model.Metrics {
	var result model.Metrics
	json.Unmarshal([]byte(jsonString), &result)
	return result
}

func AddIfMissing(cfg aws.Config,tableName string, metric model.Metrics) error {
	metricGet, err := get.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.GetMetric(context.TODO(), metric.Pk, metric.Sk)
	if err != nil {
		return err
	}

	metricParsed := ParseUnmarshal(metricGet)

	if metricParsed.Value == 0.0 && metricParsed.Username == "" {
		
		err = add.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.AddMetric(context.TODO(), metric)
		if err != nil {
			return err
		}

	} else {

		_, err := up.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.UpdateMetric(context.TODO(), model.Metrics{
			Pk: metric.Pk,
			Sk: metric.Sk,
			Value: metric.Value,
			Username: metric.Username,
		})
		if err != nil {
			return err
		}

	}

	return nil

}