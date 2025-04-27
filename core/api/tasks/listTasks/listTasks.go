package listTasks

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	list "github.com/Grs2080w/grp_server/core/db/dynamo/tasks/query"
	"github.com/Grs2080w/grp_server/core/domains/tasks/listTasks"
)

// ListTasksHandler godoc
// @Summary Get all tasks from a user
// @Description Get all tasks from a user with token
// @Tags task
// @Accept json
// @Produce json
// @Param user body listTasks.UserRequest true "Token in headers"
// @Success 200 {object} listTasks.TasksList
// @Failure 400 {object} listTasks.ErrorResponse
// @Router /tasks [get]
func ListTaskHandler(c *gin.Context) {

	username := c.GetString("username")

	tasks, err := (&list.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).Query(context.TODO(), username + "#TASKS")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error getting tasks"})
		return
	}

	taskResponse, err := listTasks.ParseUnmarshal(tasks)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing task"})
		return
	}

	if taskResponse[0].Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tasks not found"})
		return
	}
	
	c.JSON(http.StatusOK, taskResponse)
}