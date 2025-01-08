package middlewares

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	logCounter int
	counterMu  sync.Mutex
)

func UserHistoryLogger() gin.HandlerFunc {
	// Tạo và mở file log
	file, _ := os.OpenFile("./log/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	logger := log.New(file, "", log.LstdFlags)

	return func(c *gin.Context) {
		startTime := time.Now()

		// Lưu thời gian bắt đầu --> Xử lý yêu cầu request --> lưu các thông tin của khách hàng : reqBody,....
		// Đọc và sao chép request body
		reqBody := readRequestBody(c)

		// Xử lý yêu cầu
		c.Next()

		// Ghi log
		logRequest(c, startTime, reqBody, logger)
	}
}

// Hàm đọc request body
func readRequestBody(c *gin.Context) string {
	if c.Request.Body == nil {
		return ""
	}
	bodyBytes, _ := io.ReadAll(c.Request.Body) // --> bytes array
	// Gán lại body để không ảnh hưởng đến các handler khác
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return string(bodyBytes)
}

// Hàm ghi log request
func logRequest(c *gin.Context, startTime time.Time, reqBody string, logger *log.Logger) {
	// con trỏ counterMu giúp đồng bộ logCounter
	counterMu.Lock()
	logCounter++
	currentLogID := logCounter
	counterMu.Unlock()

	// Thêm đường dẫn đầy đủ vào log
	fullURL := c.Request.URL.String()

	logEntry := map[string]interface{}{
		"sequence":     currentLogID,                    // số thứ tự người dùng
		"timestamp":    time.Now().Format(time.RFC3339), // thời điểm log được tạo ra theo iso 8601
		"client_ip":    c.ClientIP(),
		"method":       c.Request.Method,
		"path":         c.Request.URL.Path,
		"full_url":     fullURL,
		"status":       c.Writer.Status(),
		"latency":      time.Since(startTime).String(), // thời gian xử lý req
		"user_agent":   c.Request.UserAgent(),          // thông tin trình duyệt, hệ điều hành
		"request_body": reqBody,
	}

	logger.Printf("Log #%d: %+v\n", currentLogID, logEntry)
}
