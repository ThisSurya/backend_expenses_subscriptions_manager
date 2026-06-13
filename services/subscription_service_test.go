package services

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"backend/models"
	repository "backend/repository/mocks"
	"backend/requests"
	"backend/utils"
)

func TestSubscriptionService_CreateSuccess(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	repo.On("Create", mock.AnythingOfType("*models.Subscription")).Return(nil).Once()

	service := NewSubscriptionService(repo)

	billingCycle := "Monthly"
	nextBillingDate, _ := time.Parse(
		"2006-01-02",
		"2024-07-01",
	)

	input := requests.SubscriptionRequest{
		CategoryId:      new(int),
		BillingCycle:    billingCycle,
		NextBillingDate: nextBillingDate,
		Amount:          decimal.NewFromFloat(50.0),
		Name:            "Test Subscription",
		ExpenseId:       nil,
		IsActive:        nil,
	}

	_, err := service.Create(input, 1)
	assert.NoError(t, err)
}

func TestSubscriptionService_CreateInvalidAmount(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	service := NewSubscriptionService(repo)
	billingCycle := "Monthly"
	nextBillingDate, _ := time.Parse(
		"2006-01-02",
		"2024-07-01",
	)

	input := requests.SubscriptionRequest{
		CategoryId:      new(int),
		BillingCycle:    billingCycle,
		NextBillingDate: nextBillingDate,
		Amount:          decimal.NewFromFloat(-10.0),
		Name:            "Test Subscription",
		ExpenseId:       nil,
		IsActive:        nil,
	}
	_, err := service.Create(input, 1)
	require.ErrorIs(t, err, utils.ErrInvalidAmount)
}

func TestSubscriptionService_CreateInvalidAmountZero(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	service := NewSubscriptionService(repo)
	billingCycle := "Monthly"
	nextBillingDate, _ := time.Parse(
		"2006-01-02",
		"2024-07-01",
	)

	input := requests.SubscriptionRequest{
		CategoryId:      new(int),
		BillingCycle:    billingCycle,
		NextBillingDate: nextBillingDate,
		Amount:          decimal.NewFromFloat(0.0),
		Name:            "Test Subscription",
		ExpenseId:       nil,
		IsActive:        nil,
	}
	_, err := service.Create(input, 1)
	require.ErrorIs(t, err, utils.ErrInvalidAmount)
}

func TestSubscriptionService_CreateRepoError(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)
	repo.On("Create", mock.AnythingOfType("*models.Subscription")).Return(utils.ErrDatabase).Once()

	service := NewSubscriptionService(repo)
	billingCycle := "Monthly"
	nextBillingDate, _ := time.Parse(
		"2006-01-02",
		"2024-07-01",
	)
	input := requests.SubscriptionRequest{
		CategoryId:      new(int),
		BillingCycle:    billingCycle,
		NextBillingDate: nextBillingDate,
		Amount:          decimal.NewFromFloat(50.0),
		Name:            "Test Subscription",
		ExpenseId:       nil,
		IsActive:        nil,
	}
	_, err := service.Create(input, 1)
	require.ErrorIs(t, err, utils.ErrDatabase)
}

func TestSubscriptionService_GetByUserIdSuccess(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	repo.On("GetAllByUserID", uint(1)).Return([]models.Subscription{}, nil).Once()

	service := NewSubscriptionService(repo)

	_, err := service.GetByUserId(1)
	assert.NoError(t, err)
}

func TestSubscriptionService_GetByUserIdRepoError(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	repo.On("GetAllByUserID", uint(1)).Return(nil, utils.ErrDatabase).Once()

	service := NewSubscriptionService(repo)
	_, err := service.GetByUserId(1)
	require.ErrorIs(t, err, utils.ErrDatabase)
}

func TestSubscriptionService_GetDetailSuccess(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	repo.On("GetByIDAndUserId", uint(1), uint(1)).
		Return(&models.Subscription{}, nil).
		Once()

	service := NewSubscriptionService(repo)

	_, err := service.GetDetail(1, 1)
	assert.NoError(t, err)
}

func TestSubscriptionService_GetDetailNotFound(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	repo.On("GetByIDAndUserId", uint(1), uint(1)).
		Return(nil, utils.ErrNotFound).
		Once()

	service := NewSubscriptionService(repo)

	_, err := service.GetDetail(1, 1)
	require.ErrorIs(t, err, utils.ErrNotFound)
}

func TestSubscriptionService_GetDetailRepoError(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	repo.On("GetByIDAndUserId", uint(1), uint(1)).
		Return(nil, utils.ErrDatabase).
		Once()

	service := NewSubscriptionService(repo)

	_, err := service.GetDetail(1, 1)
	require.ErrorIs(t, err, utils.ErrDatabase)
}

func TestSubscriptionService_UpdateSuccess(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	repo.On("GetByIDAndUserId", uint(1), uint(1)).
		Return(&models.Subscription{}, nil).
		Once()

	repo.On("Update", uint(1), mock.AnythingOfType("*models.Subscription")).
		Return(nil).
		Once()

	service := NewSubscriptionService(repo)

	input := requests.SubscriptionRequest{
		CategoryId:      new(int),
		BillingCycle:    "Monthly",
		NextBillingDate: time.Now().AddDate(0, 1, 0),
		Amount:          decimal.NewFromFloat(50.0),
		Name:            "Updated Subscription",
		ExpenseId:       nil,
		IsActive:        nil,
	}
	_, err := service.Update(1, &input, 1)
	assert.NoError(t, err)
}

func TestSubscriptionService_UpdateInvalidAmount(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	service := NewSubscriptionService(repo)

	input := requests.SubscriptionRequest{
		CategoryId:      new(int),
		BillingCycle:    "Monthly",
		NextBillingDate: time.Now().AddDate(0, 1, 0),
		Amount:          decimal.NewFromFloat(-10.0),
		Name:            "Updated Subscription",
		ExpenseId:       nil,
		IsActive:        nil,
	}
	_, err := service.Update(1, &input, 1)
	require.ErrorIs(t, err, utils.ErrInvalidAmount)
}

func TestSubscriptionService_UpdateInvalidAmountZero(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	service := NewSubscriptionService(repo)

	input := requests.SubscriptionRequest{
		CategoryId:      new(int),
		BillingCycle:    "Monthly",
		NextBillingDate: time.Now().AddDate(0, 1, 0),
		Amount:          decimal.NewFromFloat(0.0),
		Name:            "Updated Subscription",
		ExpenseId:       nil,
		IsActive:        nil,
	}
	_, err := service.Update(1, &input, 1)
	require.ErrorIs(t, err, utils.ErrInvalidAmount)
}

func TestSubscriptionService_UpdateNotFound(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	repo.On("GetByIDAndUserId", uint(1), uint(1)).
		Return(nil, utils.ErrNotFound).
		Once()

	service := NewSubscriptionService(repo)

	input := requests.SubscriptionRequest{
		CategoryId:      new(int),
		BillingCycle:    "Monthly",
		NextBillingDate: time.Now().AddDate(0, 1, 0),
		Amount:          decimal.NewFromFloat(50.0),
		Name:            "Updated Subscription",
		ExpenseId:       nil,
		IsActive:        nil,
	}

	_, err := service.Update(1, &input, 1)
	require.ErrorIs(t, err, utils.ErrNotFound)
}

func TestSubscriptionService_UpdateRepoError(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	repo.On("GetByIDAndUserId", uint(1), uint(1)).
		Return(&models.Subscription{}, nil).
		Once()

	repo.On("Update", uint(1), mock.AnythingOfType("*models.Subscription")).
		Return(utils.ErrDatabase).
		Once()

	service := NewSubscriptionService(repo)
	input := requests.SubscriptionRequest{
		CategoryId:      new(int),
		BillingCycle:    "Monthly",
		NextBillingDate: time.Now().AddDate(0, 1, 0),
		Amount:          decimal.NewFromFloat(50.0),
		Name:            "Updated Subscription",
		ExpenseId:       nil,
		IsActive:        nil,
	}
	_, err := service.Update(1, &input, 1)
	require.ErrorIs(t, err, utils.ErrDatabase)
}

func TestSubscriptionService_DeleteSuccess(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	repo.On("GetByIDAndUserId", uint(1), uint(1)).
		Return(&models.Subscription{}, nil).
		Once()

	repo.On("Delete", uint(1)).Return(nil).Once()

	service := NewSubscriptionService(repo)

	err := service.Delete(1, 1)
	assert.NoError(t, err)
}

func TestSubscriptionService_DeleteNotFound(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	repo.On("GetByIDAndUserId", uint(1), uint(1)).
		Return(nil, utils.ErrNotFound).
		Once()

	service := NewSubscriptionService(repo)

	err := service.Delete(1, 1)
	require.ErrorIs(t, err, utils.ErrNotFound)
}

func TestSubscriptionService_DeleteNoRowDeleted(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	repo.On("GetByIDAndUserId", uint(1), uint(1)).
		Return(&models.Subscription{}, nil).
		Once()

	repo.On("Delete", uint(1)).
		Return(utils.ErrNotFound).
		Once()

	service := NewSubscriptionService(repo)

	err := service.Delete(1, 1)
	require.ErrorIs(t, err, utils.ErrNotFound)
}

func TestSubscriptionService_DeleteRepoError(t *testing.T) {
	repo := repository.NewMockSubscriptionRepository(t)

	repo.On("GetByIDAndUserId", uint(1), uint(1)).
		Return(&models.Subscription{}, nil).
		Once()

	repo.On("Delete", uint(1)).
		Return(utils.ErrDatabase).
		Once()

	service := NewSubscriptionService(repo)

	err := service.Delete(1, 1)
	require.ErrorIs(t, err, utils.ErrDatabase)
}
