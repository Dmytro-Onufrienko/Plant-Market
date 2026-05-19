package middleware

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(isProd bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		if c.Request.URL.RawQuery != "" {
			path += "?" + c.Request.URL.RawQuery
		}

		if isProd {
			entry := map[string]any{
				"time":     start.UTC().Format(time.RFC3339),
				"method":   method,
				"path":     path,
				"status":   status,
				"duration": duration.Milliseconds(),
			}
			b, _ := json.Marshal(entry)
			fmt.Println(string(b))
		} else {
			color := statusColor(status)
			fmt.Printf("%s | %3d | %8s | %-7s %s\n",
				start.Format("15:04:05"),
				status,
				duration.Round(time.Microsecond),
				method,
				path,
			)
			_ = color
		}
	}
}

func statusColor(status int) string {
	switch {
	case status >= 500:
		return "\033[31m" // red
	case status >= 400:
		return "\033[33m" // yellow
	case status >= 300:
		return "\033[36m" // cyan
	default:
		return "\033[32m" // green
	}
}
