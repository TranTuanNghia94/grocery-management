package middleware

import (
	"bytes"
	"encoding/json"
	"grocery-management/pkg/logger"
	"io"
	"strings"
	"time"

	"slices"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware(skipPaths []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip logging for certain paths (like health checks)
		path := c.Request.URL.Path
		if slices.Contains(skipPaths, path) {
			c.Next()
			return
		}

		// Start timer
		start := time.Now()

		// Create a custom response writer
		w := &responseBodyLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = w

		// Log request body if it exists
		var requestBody string
		if c.Request.Body != nil {
			bodyBytes, _ := c.GetRawData()
			requestBody = string(bodyBytes)
			// Re-create the body for the actual handler to read
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		if strings.Contains(requestBody, "password") {
			// Mask the "password" field in the JSON request body
			if strings.Contains(c.Request.Header.Get("Content-Type"), "application/json") {
				var jsonBody map[string]interface{}
				if err := json.Unmarshal([]byte(requestBody), &jsonBody); err == nil {
					if _, exists := jsonBody["password"]; exists {
						jsonBody["password"] = "******"
					}
					maskedBody, _ := json.Marshal(jsonBody)
					requestBody = string(maskedBody)
				}
			}
		}

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Log basic request information
		fields := []zap.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.Duration("latency", latency),
			zap.Int("size", c.Writer.Size()),
		}

		// Add request body if not empty and not too large
		if len(requestBody) > 0 && len(requestBody) < 1024 {
			fields = append(fields, zap.String("request", requestBody))
		}

		// Add response body if not too large
		responseBody := w.body.String()
		if len(responseBody) > 0 && len(responseBody) < 1024 {
			if strings.Contains(responseBody, "jwt") {
				fields = append(fields, zap.String("response", "*****"))
			} else {
				fields = append(fields, zap.String("response", responseBody))
			}
		}

		// Add any errors from the context
		if len(c.Errors) > 0 {
			fields = append(fields, zap.String("errors", c.Errors.String()))
		}

		// Log at appropriate level based on status code
		msg := "HTTP Request"
		if c.Writer.Status() >= 500 {
			logger.Log.Error(msg, fields...)
		} else if c.Writer.Status() >= 400 {
			logger.Log.Warn(msg, fields...)
		} else {
			logger.Log.Info(msg, fields...)
		}
	}
}
