package getTask

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	get "github.com/Grs2080w/grp_server/core/db/dynamo/tasks/getTask"
	"github.com/Grs2080w/grp_server/core/domains/tasks/getTask"
)

// TaskHandler godoc
// @Summary Get a task
// @Description Get a task with the url id and token
// @Tags task
// @Accept json
// @Produce json
// @Param task body getTask.UserRequest true "url id and token"
// @Success 200 {object} getTask.Tasks
// @Failure 400 {object} getTask.ErrorResponse
// @Router /tasks/:id [get]
func GetTaskHandler(c *gin.Context) {

	username := c.GetString("username")
	id_task := c.Param("id")

	task, err := (&get.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).GetTask(context.TODO(), username + "#TASKS", "TASKS#" + id_task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error getting task"})
		return
	}

	taskResponse, err := getTask.ParseUnmarshal(task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing task"})
		return
	}

	if taskResponse.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task not found"})
		return
	}
	
	c.JSON(http.StatusOK, taskResponse)
}