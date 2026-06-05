package controllers

import (
	"backend/requests"
	"backend/services"
	"backend/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserService *services.UserService
	jwtSecret   []byte
	tokenExp    time.Duration
}

func NewAuthController(s *services.UserService, jwt []byte) *AuthController {
	return &AuthController{
		UserService: s,
		jwtSecret:   jwt,
		tokenExp:    1 * time.Hour,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var input requests.UserRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		fmt.Println("Validation error: ", err)
		errors := utils.FormatValidationError(err)
		utils.ErrorResponse(ctx, "An error occured!", errors, http.StatusBadRequest)
		return
	}

	exists, err := c.UserService.CheckEmailExists(input.Email)
	if err != nil {
		utils.ErrorResponse(ctx, "An error occured!", err.Error(), http.StatusInternalServerError)
		return
	}

	if exists {
		utils.ErrorResponse(ctx, "Email already exists!", nil, http.StatusBadRequest)
		return
	}

	user, err := c.UserService.Create(&input)

	if err != nil {
		utils.ErrorResponse(ctx, "An error occured!", err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "User created successfully!", user, http.StatusCreated)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var input requests.UserLoginRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		errors := utils.FormatValidationError(err)
		utils.ErrorResponse(ctx, "An error occured!", errors, http.StatusBadRequest)
		return
	}

	user, err := c.UserService.GetByEmail(input.Email)

	if err != nil {
		utils.ErrorResponse(ctx, "An error occured!", err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		utils.ErrorResponse(ctx, "Invalid Credentials email", nil, http.StatusUnauthorized)
		return
	}

	check := utils.CheckPasswordHash(input.Password, user.Password)

	if !check {
		utils.ErrorResponse(ctx, "Invalid Credentials password", nil, http.StatusUnauthorized)
		return
	}

	now := time.Now()

	claims := jwt.MapClaims{
		"user_id": user.Id,
		"email":   user.Email,
		"iat":     now.Unix(),
		"exp":     now.Add(c.tokenExp).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(c.jwtSecret)
	if err != nil {
		utils.ErrorResponse(ctx, "An error occured", nil, http.StatusInternalServerError)
		return
	}

	resp := map[string]any{
		"token":      tokenString,
		"expires_in": int64(c.tokenExp.Seconds()),
		"type":       "Bearer",
	}

	utils.SuccessResponse(ctx, "Login success!", resp, http.StatusOK)
}

func (c *AuthController) Logout(ctx *gin.Context) {
	resp := map[string]string{
		"instructions": "Delete the token on client side",
	}

	utils.SuccessResponse(ctx, "Logout Success!", resp, http.StatusOK)
}
