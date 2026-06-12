package controllers

import (
	"backend/requests"
	"backend/services"
	"backend/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExpenseController struct {
	Service *services.ExpenseService
}

func NewExpenseController(s *services.ExpenseService) *ExpenseController {
	return &ExpenseController{
		Service: s,
	}
}

func (c *ExpenseController) GetAllExpenses(ctx *gin.Context) {
	userId, err := utils.GetUserIdFromSession(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, "Error occured when fetching expenses", err, http.StatusUnauthorized)
		return
	}

	expenses, err := c.Service.GetByUserId(userId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve expenses", "success": false})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Expenses retrieved successfully",
		"data":    expenses,
		"success": true,
	})
}

func (c *ExpenseController) GetExpenseByUserId(ctx *gin.Context) {
	userId, err := utils.GetUserIdFromSession(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, "Error occured when fetching expenses", err, http.StatusUnauthorized)
		return
	}

	expenses, err := c.Service.GetByUserId(userId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve expenses", "success": false})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Expenses retrieved successfully",
		"data":    expenses,
		"success": true,
	})
}

func (c *ExpenseController) CreateExpense(ctx *gin.Context) {
	var input requests.ExpenseRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		errors := utils.FormatValidationError(err)
		utils.ErrorResponse(ctx, "An error occured!", errors, http.StatusBadRequest)
		return
	}

	userId, err := utils.GetUserIdFromSession(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, "Unauthorized", err, http.StatusUnauthorized)
		return
	}

	expense, err := c.Service.Create(&input, int(userId))
	if err != nil {
		utils.ErrorResponse(ctx, "Error while insert expenses!", err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Expense created successfully", expense, http.StatusCreated)
}

func (c *ExpenseController) GetExpenseDetail(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		utils.ErrorResponse(ctx, "Invalid id expense format", nil, http.StatusBadRequest)
		return
	}

	userId, err := utils.GetUserIdFromSession(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, "Unauthorized", err, http.StatusUnauthorized)
		return
	}

	expenses, err := c.Service.GetDetail(uint(id), userId)
	if err != nil {
		if errors.Is(err, utils.ErrForbidden) {
			utils.ErrorResponse(ctx, "Forbidden", nil, http.StatusForbidden)
			return
		}
		utils.ErrorResponse(ctx, "Error while fetching the detail expenses", err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Found The expenses!", expenses, http.StatusOK)
}

func (c *ExpenseController) UpdateExpenses(ctx *gin.Context) {
	idExpenseParams := ctx.Param("id")
	idExpense, err := strconv.Atoi(idExpenseParams)

	if err != nil {
		utils.ErrorResponse(ctx, "Missing id Format", nil, http.StatusBadRequest)
		return
	}

	var input requests.ExpenseRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		errors := utils.FormatValidationError(err)
		utils.ErrorResponse(ctx, "Please check the form", errors, http.StatusBadRequest)
		return
	}

	userId, err := utils.GetUserIdFromSession(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, "Unauthorized", err, http.StatusUnauthorized)
		return
	}

	result, err := c.Service.Update(uint(idExpense), &input, userId)
	if err != nil {
		if errors.Is(err, utils.ErrForbidden) {
			utils.ErrorResponse(ctx, "Forbidden", nil, http.StatusForbidden)
			return
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ErrorResponse(ctx, "Not Found", nil, http.StatusNotFound)
			return
		}
		utils.ErrorResponse(ctx, "Error while update the data", err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Success update the data", result, http.StatusOK)
}

func (c *ExpenseController) DeleteExpense(ctx *gin.Context) {
	idExpenseParam := ctx.Param("id")
	id, err := strconv.Atoi(idExpenseParam)

	if err != nil {
		utils.ErrorResponse(ctx, "Invalid id format", nil, http.StatusBadRequest)
		return
	}

	userId, err := utils.GetUserIdFromSession(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, "Unauthorized", err, http.StatusUnauthorized)
		return
	}

	err = c.Service.Delete(uint(id), userId)
	if err != nil {
		if errors.Is(err, utils.ErrForbidden) {
			utils.ErrorResponse(ctx, "Forbidden", nil, http.StatusForbidden)
			return
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ErrorResponse(ctx, "Not Found", nil, http.StatusNotFound)
			return
		}
		utils.ErrorResponse(ctx, "Error while deleting the data", err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Success delete the data", nil, http.StatusOK)
}
