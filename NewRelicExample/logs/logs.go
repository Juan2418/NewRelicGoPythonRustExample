package logs

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		incomingRequestId := c.Request.Header.Get("X-Request-Id")

		if incomingRequestId != "" {
			c.Set("requestId", incomingRequestId)
		} else {
			c.Set("requestId", uuid.NewString())
		}

		requestId := c.MustGet("requestId").(string)
		c.Writer.Header().Set("X-Request-Id", requestId)
		log.Default().SetPrefix("[" + requestId + "] ")

		c.Next()

	}
}
