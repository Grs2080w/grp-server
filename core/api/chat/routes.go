package chat

// github.com/Grs2080w/grp_server/core/api/chat

import (
	add "github.com/Grs2080w/grp_server/core/api/chat/addMessage"
	del "github.com/Grs2080w/grp_server/core/api/chat/deleteMessage"
	q "github.com/Grs2080w/grp_server/core/api/chat/getMessages"

	// middle
	"github.com/Grs2080w/grp_server/core/middleware/auth"
	"github.com/Grs2080w/grp_server/core/middleware/cache"
	"github.com/Grs2080w/grp_server/core/middleware/log"

	"github.com/gin-gonic/gin"
)



func RegisterRoutes(rg *gin.RouterGroup) {
	chat := rg.Group("/chat")
	chat.Use(auth.AuthMiddle())
	chat.Use(log.LogMiddleware())
	chat.POST("/", add.AddMessageHandler)
	chat.DELETE("/:id", del.DeleteMessageHandler)
	chat.GET("/", cache.CacheMiddleware(150), q.QueryMessagesHandler)
}
