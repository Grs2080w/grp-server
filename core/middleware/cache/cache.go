package cache

// "github.com/Grs2080w/grp_server/core/middleware/cache"

import (
	"bytes"

	r "github.com/Grs2080w/grp_server/core/db/redis"
	"github.com/gin-gonic/gin"
)


func CacheMiddleware(tll int) gin.HandlerFunc {
	return func(c *gin.Context) {

		username := c.GetString("username")

		key := username + c.Request.RequestURI
        if data, err := r.R_get(key); err == nil {
            c.Data(200, "application/json", []byte(data))
            c.Abort()
            return
        }

        writer := &responseWriter{body: bytes.NewBuffer(nil), ResponseWriter: c.Writer}
        c.Writer = writer
        c.Next()  

        r.R_set(key, writer.body.Bytes(), tll)
	}
}

type responseWriter struct {
    gin.ResponseWriter
    body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
    w.body.Write(b)
    return w.ResponseWriter.Write(b)
}

func InvalidateCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		username := c.GetString("username")

		key := username + c.Request.RequestURI
        r.R_del(key)

        c.Next()
	}
}

