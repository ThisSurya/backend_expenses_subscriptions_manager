package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRateLimitMiddleware(t *testing.T) {
	var got429 bool
	r := gin.New()

	r.Use(RateLimitMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	for i := 0; i < 15; i++ {
		req, _ := http.NewRequest("GET", "/", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code == http.StatusTooManyRequests {
			got429 = true
			break
		}
	}

	if !got429 {
		t.Errorf("Expected to receive 429 Too Many Requests, but did not")
	}

	assert.True(t, got429)
}
