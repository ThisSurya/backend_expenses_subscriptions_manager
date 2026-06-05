package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromSession(ctx *gin.Context) (int, error) {
	userId, exists := ctx.Get("user_id")
	if !exists {
		return 0, errors.New("user id not found")
	}

	uid, ok := userId.(int)
	if !ok {
		return 0, errors.New("invalid user id")
	}

	return uid, nil
}
