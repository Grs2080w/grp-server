package deleteTask

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	model "github.com/Grs2080w/grp_server/core/db/dynamo/tasks/.model"
	del "github.com/Grs2080w/grp_server/core/db/dynamo/tasks/deleteTask"
	get "github.com/Grs2080w/grp_server/core/db/dynamo/tasks/getTask"
	"github.com/Grs2080w/grp_server/core/domains/tasks/deleteTask"
)

// TaskHandler godoc
// @Summary Delete a task
// @Description Delete a task with the url id and token
// @Tags task
// @Accept json
// @Produce json
// @Param task body deleteTask.UserRequest true "url id and token"
// @Success 200 {object} deleteTask.SuccessResponse
// @Failure 400 {object} deleteTask.ErrorResponse
// @Router /tasks/:id [delete]
func DeleteTaskHandler(c *gin.Context) {

	username := c.GetString("username")

	id_task := c.Param("id")

	task, err := (&get.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).GetTask(context.TODO(), username + "#TASKS", "TASKS#" + id_task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error getting task"})
		return
	}

	taskParsed, err := deleteTask.ParseUnmarshal(task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing task"})
		return
	}


	errChan := make(chan error, 1)

	go func () {

		defer close(errChan)

		err = del.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.DeleteTask(context.TODO(), model.Tasks{
			Pk: username + "#TASKS",
			Sk: "TASKS#" + id_task,
			Username: username,
			Id: id_task,
			Date: taskParsed.Date,
			Tags: taskParsed.Tags,
			Size: taskParsed.Size,
		}, clientDb.CDB.Cfg)
	
		if err != nil {
			errChan <- err
			return
		}

		errChan <- nil
		
	}()


	go func ()  {
		
		err := <-errChan

		if err != nil {
			log.Println("error deleting task: ", err)
			return
		}

	}()


	c.JSON(http.StatusOK, deleteTask.SuccessResponse{
		Message: "task deletion scheduled",
	})
}