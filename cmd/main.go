package main

import (
	"log"
	"time"

	admin "github.com/Grs2080w/grp_server/core/api/admin"
	auth "github.com/Grs2080w/grp_server/core/api/auth"
	chat "github.com/Grs2080w/grp_server/core/api/chat"
	down "github.com/Grs2080w/grp_server/core/api/download"
	ebooks "github.com/Grs2080w/grp_server/core/api/ebooks"
	files "github.com/Grs2080w/grp_server/core/api/files"
	health "github.com/Grs2080w/grp_server/core/api/health"
	metrics "github.com/Grs2080w/grp_server/core/api/metrics"
	passwords "github.com/Grs2080w/grp_server/core/api/passwords"
	tags "github.com/Grs2080w/grp_server/core/api/tags"
	tasks "github.com/Grs2080w/grp_server/core/api/tasks"
	users "github.com/Grs2080w/grp_server/core/api/users"

	_ "github.com/Grs2080w/grp_server/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	"github.com/Grs2080w/grp_server/core/db/redis"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title grp@server
// @version 1.0
// @description API for personal server.
// @host localhost:8080
// @BasePath /
func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*",},
        AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Request-ID"},
        ExposeHeaders:    []string{"Content-Length", "X-Request-ID"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

	err := clientDb.InitDynamoClient()
	if err != nil {
		log.Print(err)
		log.Print("failed to init conection with db aws")
	}

	redis.Init_Client_Redis()

	// Serve na rota /swagger/index.html
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	api := router.Group("/api")
	users.RegisterRoutes(api)
	auth.RegisterRoutes(api)
	tasks.RegisterRoutes(api)
	files.RegisterRoutes(api)
	metrics.RegisterRoutes(api)
	chat.RegisterRoutes(api)
	tags.RegisterRoutes(api)
	ebooks.RegisterRoutes(api)
	down.RegisterRoutes(api)
	passwords.RegisterRoutes(api)
	admin.RegisterRoutes(api)
	health.RegisterRoutes(api)

	router.Run(":8080")
}