package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(p gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s \n",
			p.ClientIP,
			p.TimeStamp.Format(time.RFC822),
			p.Method,
			p.Path,
			p.StatusCode,
			p.Latency,
		)
	})
}
