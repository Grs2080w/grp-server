package getMetrics

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	getM "github.com/Grs2080w/grp_server/core/db/dynamo/metrics/query"
	"github.com/Grs2080w/grp_server/core/domains/metrics/getMetrics"
)

// GetMetricsHandler godoc
// @Summary Retrieve metrics
// @Description Retrieve metrics for a user using their authentication token.
// @Tags metrics
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} getMetrics.Response
// @Failure 401 {object} getMetrics.ErrorResponse
// @Failure 500 {object} getMetrics.ErrorResponse
// @Router /metrics [get]
func GetMetricsHandler(c *gin.Context) {

	username := c.GetString("username")

	metrics, err := getM.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.Query(context.TODO(), username + "#METRICS")
	if err != nil {
		log.Printf("Failed to get metrics: %v", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}


	metricsParsed, err := getMetrics.ParseMetrics(metrics)
	if err != nil {
		log.Printf("Failed to parse metrics: %v", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	data_RECORDS_PER_DOMAIN := make(map[string]int)
	data_FILES_PER_EXTENSION := make(map[string]int)
	data_STORAGE_PER_TYPE := make(map[string]int)
	data_STORAGE_PER_DOMAIN := make(map[string]int)

	for _, metric := range metricsParsed {

		switch metric.Type {
		case "RECORDS_PER_DOMAIN":
			separated := getMetrics.ExtractKey(metric.Sk)
			data_RECORDS_PER_DOMAIN[separated] = int(metric.Value)
		case "FILES_PER_EXTENSION":
			separated := getMetrics.ExtractKey(metric.Sk)
			data_FILES_PER_EXTENSION[separated] = int(metric.Value)
		case "STORAGE_PER_TYPE":
			separated := getMetrics.ExtractKey(metric.Sk)
			data_STORAGE_PER_TYPE[separated] = int(metric.Value)
		case "STORAGE_PER_DOMAIN":
			separated := getMetrics.ExtractKey(metric.Sk)
			data_STORAGE_PER_DOMAIN[separated] = int(metric.Value)
		}

	}
	
	response := getMetrics.Response{
		Files_per_extension: data_FILES_PER_EXTENSION,
		Records_per_domain: data_RECORDS_PER_DOMAIN,
		Storage_per_type: data_STORAGE_PER_DOMAIN,
		Storage_per_domain: data_STORAGE_PER_DOMAIN,
	}
	
	c.JSON(200, response)
}