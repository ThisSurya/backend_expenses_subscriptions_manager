package controllers

import (
	"backend/requests"
	"backend/services"
	"backend/utils"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SubscriptionController struct {
	Service *services.SubscriptionService
}

func NewSubscriptionController(s *services.SubscriptionService) *SubscriptionController {
	return &SubscriptionController{
		Service: s,
	}
}

func (s *SubscriptionController) GetSubscriptionByUserId(ctx *gin.Context) {
	userId, err := utils.GetUserIdFromSession(ctx)

	if err != nil {
		// fmt.Println("Error somethinc")
		utils.ErrorResponse(ctx, "Error getting user Id at session", err, http.StatusInternalServerError)
		return
	}

	subscription, err := s.Service.GetByUserId(userId)

	if err != nil {
		utils.ErrorResponse(ctx, "Error getting subscriptions", err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Success get subcsriptions", subscription, http.StatusOK)
}

func (s *SubscriptionController) CreateSubscription(ctx *gin.Context) {
	var input requests.SubscriptionRequest
	err := ctx.ShouldBindJSON(&input)

	userId, err := utils.GetUserIdFromSession(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, "Unauthorized", nil, http.StatusUnauthorized)
		return
	}

	if err != nil {
		errors := utils.FormatValidationError(err)
		utils.ErrorResponse(ctx, "An Error occured!", errors, http.StatusBadRequest)
		return
	}

	subscription, err := s.Service.Create(input, userId)

	if err != nil {
		utils.ErrorResponse(ctx, "Error creating subscription", err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Success create subscription", subscription, http.StatusAccepted)
}

func (s *SubscriptionController) GetSubscriptionDetail(ctx *gin.Context) {
	idParam := ctx.Param("id")

	idsubs, err := strconv.Atoi(idParam)

	if err != nil {
		utils.ErrorResponse(ctx, "Invalid id subscription format", nil, http.StatusBadRequest)
		return
	}

	userId, err := utils.GetUserIdFromSession(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, "Unauthorized", nil, http.StatusUnauthorized)
		return
	}

	subscription, err := s.Service.GetDetail(idsubs, userId)

	if err != nil {
		if errors.Is(err, services.ErrForbidden) {
			utils.ErrorResponse(ctx, "Forbidden", nil, http.StatusForbidden)
			return
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ErrorResponse(ctx, "Subscription not found!", nil, http.StatusNotFound)
			return
		}

		utils.ErrorResponse(ctx, "Error while getting detail subscriptions", nil, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Found subscription detail", subscription, http.StatusOK)
}

func (s *SubscriptionController) UpdateSubscription(ctx *gin.Context) {
	idParam := ctx.Param("id")

	idSubs, err := strconv.Atoi(idParam)

	if err != nil {
		utils.ErrorResponse(ctx, "Invalid id subscription format", nil, http.StatusBadRequest)
		return
	}

	userId, err := utils.GetUserIdFromSession(ctx)

	var input requests.SubscriptionRequest
	err = ctx.ShouldBindJSON(&input)

	if err != nil {
		errors := utils.FormatValidationError(err)
		utils.ErrorResponse(ctx, "An error occured!", errors, http.StatusBadRequest)
		return
	}

	subscription, err := s.Service.Update(idSubs, &input, userId)

	if err != nil {
		if errors.Is(err, services.ErrForbidden) {
			utils.ErrorResponse(ctx, "Unauthorized", nil, http.StatusUnauthorized)
			return
		}

		utils.ErrorResponse(ctx, "Error updating subscription", nil, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Sucess updating subcriptions", subscription, http.StatusAccepted)
}

func (s *SubscriptionController) DeleteSubscription(ctx *gin.Context) {
	idParams := ctx.Param("id")
	idSubs, err := strconv.Atoi(idParams)

	if err != nil {
		utils.ErrorResponse(ctx, "Invalid id subscription format", nil, http.StatusBadRequest)
		return
	}

	userId, err := utils.GetUserIdFromSession(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, "Unauthorized", nil, http.StatusUnauthorized)
		return
	}

	err = s.Service.Delete(idSubs, userId)

	if err != nil {
		fmt.Println("[ERROR OCCURED!]: ", err)
		if errors.Is(err, services.ErrForbidden) {
			utils.ErrorResponse(ctx, "Unauthorized", nil, http.StatusUnauthorized)
			return
		}

		utils.ErrorResponse(ctx, "Error while deleting subscription", nil, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Sucess deleting subscription", nil, http.StatusAccepted)
}
