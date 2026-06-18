package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func setupRouter(jwtSecret []byte) *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(AuthMiddleware(jwtSecret))

	r.GET("/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Protected route accessed"})
	})
	return r
}

func TestAuthMiddleware_Success(t *testing.T) {
	now := time.Now()
	jwtSecret := []byte("test_secret")
	claims := jwt.MapClaims{
		"user_id": 1,
		"email":   "test@example.com",
		"iat":     now.Unix(),
		"exp":     now.Add(1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		t.Fatalf("Error generating token: %v", err)
	}

	w := httptest.NewRecorder()
	r := setupRouter(jwtSecret)

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestAuthMiddleware_Unauthorized(t *testing.T) {
	jwtSecret := []byte("test_secret")
	w := httptest.NewRecorder()
	r := setupRouter(jwtSecret)

	req, _ := http.NewRequest("GET", "/protected", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	jwtSecret := []byte("test_secret")
	w := httptest.NewRecorder()
	r := setupRouter(jwtSecret)

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer beareawikowk")

	r.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
}

func TestAuthMiddleware_ExpiredToken(t *testing.T) {
	now := time.Now()
	jwtSecret := []byte("test_secret")
	claims := jwt.MapClaims{
		"user_id": 1,
		"email":   "test@example.com",
		"iat":     now.Unix(),
		"exp":     now.Add(-1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString(jwtSecret)

	w := httptest.NewRecorder()
	r := setupRouter(jwtSecret)

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	r.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
}

func TestAuthMiddleware_InvalidFormatHeader(t *testing.T) {
	jwtSecret := []byte("test_secret")

	w := httptest.NewRecorder()
	r := setupRouter(jwtSecret)

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "InvalidFormatToken")

	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}

func TestAuthMiddleware_WrongSecretKey(t *testing.T) {
	now := time.Now()
	jwtSecret := []byte("test_secret1")
	claims := jwt.MapClaims{
		"user_id": 1,
		"email":   "test@example.com",
		"iat":     now.Unix(),
		"exp":     now.Add(1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		t.Fatalf("Error generating token: %v", err)
	}

	w := httptest.NewRecorder()
	r := setupRouter([]byte("test_secret2"))

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}

func TestAuthMiddleware_InvalidUserIDType(t *testing.T) {
	now := time.Now()
	jwtSecret := []byte("test_secret")

	claims := jwt.MapClaims{
		"user_id": "invalid_user_id",
		"email":   "test@example.com",
		"iat":     now.Unix(),
		"exp":     now.Add(1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtSecret)

	w := httptest.NewRecorder()
	r := setupRouter(jwtSecret)

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}
