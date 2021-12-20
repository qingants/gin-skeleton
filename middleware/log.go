package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

//if param.Latency > time.Minute {
//// Truncate in a golang < 1.8 safe way
//param.Latency = param.Latency - param.Latency%time.Second
//}
//return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
//param.TimeStamp.Format("2006/01/02 - 15:04:05"),
//statusColor, param.StatusCode, resetColor,
//param.Latency,
//param.ClientIP,
//methodColor, param.Method, resetColor,
//param.Path,
//param.ErrorMessage,
//)
func LoggerWithZap(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		timestamp := time.Now()
		latency := timestamp.Sub(start)

		if latency > time.Minute {
			// Truncate in a golang < 1.8 safe way
			latency = latency - latency%time.Second
		}

		if len(c.Errors) > 0 {
			logger.Error(path,
				zap.String("ip", c.ClientIP()),
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("protocol", c.Request.Proto),
				zap.String("content-type", c.GetHeader("Content-Type")),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("query", query),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.String("latency", latency.String()),
			)
		} else {
			logger.Info(path,
				zap.String("ip", c.ClientIP()),
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("protocol", c.Request.Proto),
				zap.String("content-type", c.GetHeader("Content-Type")),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("query", query),
				zap.String("latency", latency.String()),
			)
		}

	}
}
