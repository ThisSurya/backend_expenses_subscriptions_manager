package controllers

import (
	"backend/models"
	"backend/requests"
	"backend/services"
	"backend/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	Service *services.CategoryService
}

func NewCategoryController(s *services.CategoryService) *CategoryController {
	return &CategoryController{
		Service: s,
	}
}

func (c *CategoryController) GetByUserId(ctx *gin.Context) {
	var category []models.Category

	userId, err := utils.GetUserIdFromSession(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, "Unauthorized", err, http.StatusUnauthorized)
		return
	}

	category, err = c.Service.GetByUserId(userId)

	if err != nil {
		utils.ErrorResponse(ctx, "Error while fetch the data", err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Success fetch data", category, http.StatusOK)
}

func (c *CategoryController) GetById(ctx *gin.Context) {
	var category *models.Category

	idParams := ctx.Param("id_category")
	id, err := strconv.Atoi(idParams)

	if err != nil {
		utils.ErrorResponse(ctx, "Invalid id format", err, http.StatusBadRequest)
		return
	}

	userId, err := utils.GetUserIdFromSession(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, "Unauthorized", err, http.StatusUnauthorized)
		return
	}

	category, err = c.Service.GetDetail(uint(id), userId)
	if err != nil {
		if errors.Is(err, utils.ErrForbidden) {
			utils.ErrorResponse(ctx, "Forbidden", nil, http.StatusForbidden)
			return
		}
		if errors.Is(err, utils.ErrNotFound) {
			utils.ErrorResponse(ctx, "Not found", nil, http.StatusNotFound)
			return
		}
		utils.ErrorResponse(ctx, "Error while fetching detail", err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Success get detail category", category, http.StatusOK)
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var input requests.CategoryRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		errs := utils.FormatValidationError(err)
		utils.ErrorResponse(ctx, "Validation failed", errs, http.StatusBadRequest)
		return
	}

	userId, err := utils.GetUserIdFromSession(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, "Unauthorized", err, http.StatusUnauthorized)
		return
	}

	category, err := c.Service.Create(&input, int(userId))
	if err != nil {
		utils.ErrorResponse(ctx, "Error while creating category", err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Success creating category", category, http.StatusAccepted)
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	var input requests.CategoryRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		errs := utils.FormatValidationError(err)
		utils.ErrorResponse(ctx, "Validation error", errs, http.StatusBadRequest)
		return
	}

	userId, err := utils.GetUserIdFromSession(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, "Unauthorized", err, http.StatusUnauthorized)
		return
	}

	idParams := ctx.Param("id_category")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		utils.ErrorResponse(ctx, "Invalid id format", err, http.StatusBadRequest)
		return
	}

	category, err := c.Service.Update(uint(id), &input, userId)

	if err != nil {
		if errors.Is(err, utils.ErrForbidden) {
			utils.ErrorResponse(ctx, "Forbidden", nil, http.StatusForbidden)
			return
		}
		if errors.Is(err, utils.ErrNotFound) {
			utils.ErrorResponse(ctx, "Not found", nil, http.StatusNotFound)
			return
		}
		utils.ErrorResponse(ctx, "Error while update data", err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Success update data", category, http.StatusAccepted)
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	idParams := ctx.Param("id_category")

	id, err := strconv.Atoi(idParams)

	if err != nil {
		utils.ErrorResponse(ctx, "Invalid id format", err, http.StatusBadRequest)
		return
	}

	userId, err := utils.GetUserIdFromSession(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, "Unauthorized", err, http.StatusUnauthorized)
		return
	}

	_, err = c.Service.Delete(uint(id), userId)
	if err != nil {
		if errors.Is(err, utils.ErrForbidden) {
			utils.ErrorResponse(ctx, "Forbidden", nil, http.StatusForbidden)
			return
		}
		if errors.Is(err, utils.ErrNotFound) {
			utils.ErrorResponse(ctx, "Not found", nil, http.StatusNotFound)
			return
		}
		utils.ErrorResponse(ctx, "Error while deleting data", err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Success deleting data", nil, http.StatusOK)
}
