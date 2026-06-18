package middleware

import (
	"backend/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/time/rate"
)

func AuthMiddleware(jwtSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			utils.ErrorResponse(c, "Authorization header is required", nil, http.StatusUnauthorized)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.ErrorResponse(c, "Invalid authorization header format", nil, http.StatusUnauthorized)
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				utils.ErrorResponse(c, "Unexpected signing method", nil, http.StatusUnauthorized)
				return nil, jwt.ErrSignatureInvalid
			}

			return jwtSecret, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				utils.ErrorResponse(c, "Invalid token signature", nil, http.StatusUnauthorized)
			} else {
				utils.ErrorResponse(c, "Invalid token", nil, http.StatusUnauthorized)
			}
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			utils.ErrorResponse(c, "Invalid token claims", nil, http.StatusUnauthorized)
			c.Abort()
			return
		}

		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				utils.ErrorResponse(c, "Token has expired", nil, http.StatusUnauthorized)
				c.Abort()
				return
			}
		}

		var userId int
		switch v := claims["user_id"].(type) {
		case float64:
			userId = int(v)
		case int:
			userId = v
		case string:
			parsed, err := strconv.Atoi(v)
			if err != nil {
				utils.ErrorResponse(c, "Invalid token claims", nil, http.StatusUnauthorized)
				c.Abort()
				return
			}
			userId = parsed
		default:
			utils.ErrorResponse(c, "Invalid token claims", nil, http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("user_id", userId)
		c.Set("email", claims["email"])

		c.Next()
	}
}

func RateLimitMiddleware() gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(time.Second), 10)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			utils.ErrorResponse(c, "Too many requests", nil, http.StatusTooManyRequests)
			c.Abort()
			return
		}

		c.Next()
	}
}
