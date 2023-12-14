package middleware

import (
	"time"

	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	// InfoLogger for informational logs
	InfoLogger *log.Logger

	// ErrorLogger for error logs
	ErrorLogger *log.Logger
)

func init() {
	// Create loggers. You may customize the output and behavior as needed.
	InfoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LoggerMiddleware is a middleware for logging HTTP requests.
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Record the start time
		startTime := time.Now()

		// Process the request
		c.Next()

		// Record the end time
		endTime := time.Now()

		// Calculate the latency
		latency := endTime.Sub(startTime)

		// Log the request details
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		InfoLogger.Printf("%v | %3d | %13v | %15s | %-7s %s\n",
			endTime.Format("2006/01/02 - 15:04:05"),
			status,
			latency,
			clientIP,
			method,
			path+"?"+query,
		)

		// Log errors (if any)
		if status >= 500 {
			ErrorLogger.Printf("%v | %3d | %13v | %15s | %-7s %s\n",
				endTime.Format("2006/01/02 - 15:04:05"),
				status,
				latency,
				clientIP,
				method,
				path+"?"+query,
			)
		}
	}
}
