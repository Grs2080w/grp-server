package tasks

import (
	"github.com/gin-gonic/gin"

	add "github.com/Grs2080w/grp_server/core/api/tasks/addTask"
	up "github.com/Grs2080w/grp_server/core/api/tasks/changeStatus"
	del "github.com/Grs2080w/grp_server/core/api/tasks/deleteTask"
	get "github.com/Grs2080w/grp_server/core/api/tasks/getTask"
	list "github.com/Grs2080w/grp_server/core/api/tasks/listTasks"

	// middle
	"github.com/Grs2080w/grp_server/core/middleware/auth"
	"github.com/Grs2080w/grp_server/core/middleware/cache"
	"github.com/Grs2080w/grp_server/core/middleware/log"
)



func RegisterRoutes(rg *gin.RouterGroup) {
	tasks := rg.Group("/tasks")
	tasks.Use(auth.AuthMiddle())
	tasks.Use(log.LogMiddleware())
	tasks.POST("/", add.AddTaskHandler)
	tasks.PATCH("/:id", up.ChangeTaskHandler)
	tasks.DELETE("/:id", del.DeleteTaskHandler)
	tasks.GET("/", cache.CacheMiddleware(150), list.ListTaskHandler)
	tasks.GET("/:id", cache.CacheMiddleware(150), get.GetTaskHandler)
}
