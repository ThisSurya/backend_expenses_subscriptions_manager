package controllers

import (
	"backend/requests"
	"backend/services"
	"backend/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserService *services.UserService
}

func NewAuthController(s *services.UserService) *AuthController {
	return &AuthController{
		UserService: s,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var input requests.UserRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		errs := utils.FormatValidationError(err)
		utils.ErrorResponse(ctx, "An error occured!", errs, http.StatusBadRequest)
		return
	}

	user, err := c.UserService.RegisterService(&input)

	if err != nil {
		if errors.Is(err, utils.ErrEmailExists) {
			utils.ErrorResponse(ctx, "Email already exists!", nil, http.StatusBadRequest)
			return
		}
		utils.ErrorResponse(ctx, "An error occured!", err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "User created successfully!", user, http.StatusCreated)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var input requests.UserLoginRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		errs := utils.FormatValidationError(err)
		utils.ErrorResponse(ctx, "An error occured!", errs, http.StatusBadRequest)
		return
	}

	data, err := c.UserService.LoginService(&input)

	if err != nil {
		if errors.Is(err, utils.ErrInvalidCredentials) {
			utils.ErrorResponse(ctx, "Invalid credentials", nil, http.StatusUnauthorized)
			return
		}
		utils.ErrorResponse(ctx, "An error occured!", err.Error(), http.StatusInternalServerError)
		return
	}

	if data == nil {
		utils.ErrorResponse(ctx, "Invalid Credentials email", nil, http.StatusUnauthorized)
		return
	}

	utils.SuccessResponse(ctx, "Login success!", data, http.StatusOK)
}

func (c *AuthController) Logout(ctx *gin.Context) {
	resp := map[string]string{
		"instructions": "Delete the token on client side",
	}

	utils.SuccessResponse(ctx, "Logout Success!", resp, http.StatusOK)
}
