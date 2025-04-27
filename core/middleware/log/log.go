package log

// "github.com/Grs2080w/grp_server/core/middleware/log"

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"bytes"
	"io"

	client "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	model "github.com/Grs2080w/grp_server/core/db/dynamo/logs/.model"
	addLog "github.com/Grs2080w/grp_server/core/db/dynamo/logs/addLog"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		raw, err := c.GetRawData()  
		if err != nil {
			c.Next()  
			return
		}

    	c.Request.Body = io.NopCloser(bytes.NewReader(raw))

		writer := &responseWriter{body: bytes.NewBuffer(nil), ResponseWriter: c.Writer}
        c.Writer = writer
        c.Next()  

		duration := time.Since(start)
		status := c.Writer.Status()
		ip := c.ClientIP()
		method := c.Request.Method
		path := c.FullPath()
		
		username, _ := c.Get("username")
		if username == nil {
			username = "anonymous"
		}

		var LogLevelInfo map[int]string = map[int]string{
			200: "INFO",
			201: "INFO",
			204: "INFO",
			301: "WARN",
			302: "WARN",
			400: "ERROR",
			401: "ERROR",
			403: "ERROR",
			404: "ERROR",
			500: "ERROR",
			502: "ERROR",
			503: "ERROR",
			504: "ERROR",
		}

		var message string
		
		if LogLevelInfo[status] == "ERROR" {
			message = "Error ocurred during request"
		} 
		
		if LogLevelInfo[status] == "WARN" {
			message = "Something went wrong"
		}

		if LogLevelInfo[status] == "INFO" {
			message = "Request completed successfully"
		}

		id_log := uuid.New().String()

		log := model.Logs{
			Pk:          "LOGS",
			Sk:          "LOG#" + id_log,
			Log_id:      id_log,
			Timestamp:   int(start.Unix()),
			Level:       LogLevelInfo[status],
			Domain:      getDomainByPath(path),
			Message:     message,
			Metadata:    "",
			Username:    username.(string),
			Ip_address:  ip,
			Stack_trace: "none",
			Operation:   method + " " + path,
			Status_code: strconv.Itoa(status),
		}

		var body string
		if LogLevelInfo[status] == "ERROR" || LogLevelInfo[status] == "WARN" {
			body = string(raw)
		} else {
			body = "none"
		}

		var responseContent string
		if LogLevelInfo[status] == "ERROR" || LogLevelInfo[status] == "WARN" {
			responseContent = writer.body.String()
		} else {
			responseContent = "none"
		}

		metadata := map[string]interface{}{
			"method":     method,
			"path":       path,
			"duration_ms": duration.Milliseconds(),
			"status":     status,
			"body":       body,
			"headers":    c.Request.Header,
			"query":      c.Request.URL.Query(),
			"params":     c.Params,
			"response":   c.Writer.Header(),
			"response_body": c.Writer.Size(),
			"response_content": responseContent,
			"date":      time.Now().Format(time.RFC3339),
		}
		
		log.Metadata = fmt.Sprintf("%v", metadata)

		go func ()  {
			
			err := addLog.TableBasics{DynamoDbClient: client.CDB.DynamoClient, TableName: client.CDB.LogTableName}.AddLogs(context.TODO(), log)

			if err != nil {
				fmt.Println("Error adding log to DynamoDB:", err)
			}

		}()
	}
}

func getDomainByPath(path string) string {
	const apiPrefix = "/api/"
    if !strings.HasPrefix(path, apiPrefix) {
        return ""
    }

    subPath := strings.TrimPrefix(path, apiPrefix)
    parts := strings.SplitN(subPath, "/", 2)

    return parts[0]
}

type responseWriter struct {
    gin.ResponseWriter
    body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
    w.body.Write(b)
    return w.ResponseWriter.Write(b)
}