package changeStatus

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	model "github.com/Grs2080w/grp_server/core/db/dynamo/tasks/.model"
	get "github.com/Grs2080w/grp_server/core/db/dynamo/tasks/getTask"
	up "github.com/Grs2080w/grp_server/core/db/dynamo/tasks/updateTask"
	"github.com/Grs2080w/grp_server/core/domains/tasks/changeStatus"
)

// ChangeTaskHandler godoc
// @Summary Change status task
// @Description Change status task with the request body and url id and token
// @Tags task
// @Accept json
// @Produce json
// @Param task body changeStatus.UserRequest true "url id and token"
// @Success 200 {object} changeStatus.SuccessResponse
// @Failure 400 {object} changeStatus.ErrorResponse
// @Router /tasks/:id [patch]
func ChangeTaskHandler(c *gin.Context) {

	username := c.GetString("username")

	id_task := c.Param("id")

	task, err := (&get.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).GetTask(context.TODO(), username + "#TASKS", "TASKS#" + id_task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error getting task"})
		return
	}

	taskToUpdate, err := changeStatus.ParseUnmarshal(task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing task"})
		return
	}

	if taskToUpdate.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task not found"})
		return
	}

	if taskToUpdate.Status == "open" {
		taskToUpdate.Status = "closed"
	} else {
		taskToUpdate.Status = "open"
	}


	errChan := make(chan error, 1)

	go func () {

		defer close(errChan)
		
		_, err = up.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}.UpdateTask(context.TODO(), model.Tasks{
			Pk: username + "#TASKS",
			Sk: "TASKS#" + id_task,
			Status: taskToUpdate.Status,
		})
	
		if err != nil {
			errChan <- err
			return
		}

		errChan <- nil

	}()


	go func ()  {
		
		err := <- errChan

		if err != nil {
			log.Println("error updating task: ", err)
			return
		}

	}()

	
	c.JSON(http.StatusOK, changeStatus.SuccessResponse{
		Message: "task status change schedule",
	})
}