package controllers

import (
	"backend/requests"
	"backend/services"
	"backend/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
		utils.ErrorResponse(ctx, "Unauthorized", err, http.StatusUnauthorized)
		return
	}

	subscription, err := s.Service.GetByUserId(userId)

	if err != nil {
		utils.ErrorResponse(ctx, "Error getting subscriptions", err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Success get subscriptions", subscription, http.StatusOK)
}

func (s *SubscriptionController) CreateSubscription(ctx *gin.Context) {
	var input requests.SubscriptionRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		errs := utils.FormatValidationError(err)
		utils.ErrorResponse(ctx, "An Error occured!", errs, http.StatusBadRequest)
		return
	}

	userId, err := utils.GetUserIdFromSession(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, "Unauthorized", nil, http.StatusUnauthorized)
		return
	}

	subscription, err := s.Service.Create(input, int(userId))

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

	subscription, err := s.Service.GetDetail(uint(idsubs), userId)

	if err != nil {
		if errors.Is(err, utils.ErrForbidden) {
			utils.ErrorResponse(ctx, "Forbidden", nil, http.StatusForbidden)
			return
		}
		if errors.Is(err, utils.ErrNotFound) {
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
	if err != nil {
		utils.ErrorResponse(ctx, "Unauthorized", nil, http.StatusUnauthorized)
		return
	}

	var input requests.SubscriptionRequest
	if err = ctx.ShouldBindJSON(&input); err != nil {
		errs := utils.FormatValidationError(err)
		utils.ErrorResponse(ctx, "An error occured!", errs, http.StatusBadRequest)
		return
	}

	subscription, err := s.Service.Update(uint(idSubs), &input, int(userId))

	if err != nil {
		if errors.Is(err, utils.ErrForbidden) {
			utils.ErrorResponse(ctx, "Forbidden", nil, http.StatusForbidden)
			return
		}
		if errors.Is(err, utils.ErrNotFound) {
			utils.ErrorResponse(ctx, "Subscription not found!", nil, http.StatusNotFound)
			return
		}
		utils.ErrorResponse(ctx, "Error updating subscription", nil, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Success updating subscriptions", subscription, http.StatusAccepted)
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

	err = s.Service.Delete(uint(idSubs), int(userId))

	if err != nil {
		if errors.Is(err, utils.ErrForbidden) {
			utils.ErrorResponse(ctx, "Forbidden", nil, http.StatusForbidden)
			return
		}
		if errors.Is(err, utils.ErrNotFound) {
			utils.ErrorResponse(ctx, "Subscription not found!", nil, http.StatusNotFound)
			return
		}
		utils.ErrorResponse(ctx, "Error while deleting subscription", nil, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(ctx, "Success deleting subscription", nil, http.StatusAccepted)
}
