package addTask

import (
	"context"
	"log"
	"net/http"
	"time"

	u "github.com/google/uuid"

	a "github.com/Grs2080w/grp_server/core/db/dynamo/tasks/addTask"

	cDb "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	model "github.com/Grs2080w/grp_server/core/db/dynamo/tasks/.model"
	"github.com/Grs2080w/grp_server/core/domains/tasks/addTask"

	"github.com/gin-gonic/gin"
)

// TaskHandler godoc
// @Summary Add a task
// @Description Add a task with the request body
// @Tags task
// @Accept json
// @Produce json
// @Param task body addTask.UserRequest true "Request body"
// @Success 200 {object} addTask.Tasks
// @Failure 400 {object} addTask.ErrorResponse
// @Router /tasks [post]
func AddTaskHandler(c *gin.Context) {

	username := c.GetString("username")
	
	var task addTask.UserRequest
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if ok, err := task.Verify(); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id_task := u.New().String()

	new_task := model.Tasks{
		Pk: username + "#TASKS",
		Sk: "TASKS#" + id_task,
		Id: id_task,
		Status: "open",
		Description: task.Description,
		Title: task.Title,
		Date: time.Now().String(),
		Tags: task.Tags,
		Username: username,
		Size: task.Size,
	}



	errChan := make(chan error, 1)


	go func ()  {

		defer close(errChan)
		
		err := a.TableBasics{
			TableName: cDb.CDB.TableName,
			DynamoDbClient :    cDb.CDB.DynamoClient,
		}.AddTask(context.TODO(), new_task, cDb.CDB.Cfg)
	
		if err != nil {
			errChan <- err
			return
		}

		errChan <- nil
	}()

	go func ()  {
		
		err := <-errChan

		if err != nil {
			log.Println("Error adding task:", err)
			return
		}

	}()


	c.JSON(http.StatusOK, gin.H{
		"message": "Task adding scheduled",
	})
}