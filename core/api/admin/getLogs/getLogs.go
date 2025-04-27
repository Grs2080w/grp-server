package getLogs

import (
	"context"

	"github.com/gin-gonic/gin"

	config "github.com/Grs2080w/grp_server/core/config"
	client "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	g "github.com/Grs2080w/grp_server/core/db/dynamo/logs/query"
	getlogs "github.com/Grs2080w/grp_server/core/domains/logs/getLogs"
)

// GetLogsHandler godoc
// @Summary Get all logs
// @Description Get all logs with no params
// @Tags admin
// @Accept json
// @Produce json
// @Param token header string true "admin token"
// @Success 200 {object} []getlogs.Logs
// @Failure 400 {object} getlogs.ErrorResponse
// @Router /logs [get]
func GetLogsHandler(c *gin.Context) {
	
	username := c.GetString("username")

	if username != config.GetValueEnv("ADMIN") {
		c.JSON(400, gin.H{"error": "Just admin can access this endpoint"})
		return
	}
	
	logsGet, err := g.TableBasics{DynamoDbClient: client.CDB.DynamoClient, TableName: client.CDB.LogTableName}.Query(context.TODO(), "LOGS")

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve logs"})
		return
	}

	logs, err := getlogs.ParseLogs(logsGet)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to parse logs"})
		return
	}

	c.JSON(200, logs)
}

