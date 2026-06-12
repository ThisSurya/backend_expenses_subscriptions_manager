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

func TestExpenseService_CreateSuccess(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)

	repo.On("Create", mock.AnythingOfType("*models.Expense")).
		Return(nil).
		Once()

	service := NewExpenseService(repo)
	desc := "Test Expense"
	paymentMethod := "Credit Card"
	expenseDate, _ := time.Parse(
		"2006-01-02",
		"2024-06-01",
	)

	input := requests.ExpenseRequest{
		UserId:        1,
		CategoryId:    new(int),
		Amount:        decimal.NewFromFloat(100.0),
		Description:   &desc,
		ExpenseDate:   expenseDate,
		PaymentMethod: &paymentMethod,
		IsRecurring:   false,
	}

	_, err := service.Create(&input, 1)
	assert.NoError(t, err)
}

func TestExpenseService_CreateInvalidAmount(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)

	service := NewExpenseService(repo)
	desc := "Test Expense"
	paymentMethod := "Credit Card"
	expenseDate, _ := time.Parse(
		"2006-01-02",
		"2024-06-01",
	)

	input := requests.ExpenseRequest{
		UserId:        1,
		CategoryId:    new(int),
		Amount:        decimal.NewFromFloat(-50.0),
		Description:   &desc,
		ExpenseDate:   expenseDate,
		PaymentMethod: &paymentMethod,
		IsRecurring:   false,
	}

	_, err := service.Create(&input, 1)
	require.ErrorIs(t, err, utils.ErrInvalidAmount)
}

func TestExpenseService_CreateDateInFuture(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)

	service := NewExpenseService(repo)
	desc := "Test Expense"
	paymentMethod := "Credit Card"
	expenseDate := time.Now().Add(24 * time.Hour) // Future date

	input := requests.ExpenseRequest{
		UserId:        1,
		CategoryId:    new(int),
		Amount:        decimal.NewFromFloat(100.0),
		Description:   &desc,
		ExpenseDate:   expenseDate,
		PaymentMethod: &paymentMethod,
		IsRecurring:   false,
	}
	_, err := service.Create(&input, 1)
	require.ErrorIs(t, err, utils.ErrDateInFuture)
}

func TestExpenseService_CreateRepoError(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)
	repo.On("Create", mock.AnythingOfType("*models.Expense")).
		Return(utils.ErrDatabase).
		Once()

	service := NewExpenseService(repo)
	desc := "Test Expense"
	paymentMethod := "Credit Card"
	expenseDate, _ := time.Parse(
		"2006-01-02",
		"2024-06-01",
	)

	input := requests.ExpenseRequest{
		UserId:        1,
		CategoryId:    new(int),
		Amount:        decimal.NewFromFloat(100.0),
		Description:   &desc,
		ExpenseDate:   expenseDate,
		PaymentMethod: &paymentMethod,
		IsRecurring:   false,
	}
	_, err := service.Create(&input, 1)
	require.ErrorIs(t, err, utils.ErrDatabase)
}

func TestExpenseService_GetByUserIdSuccess(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)
	repo.On("GetAllByUserID", uint(1)).
		Return([]models.Expense{}, nil).
		Once()

	service := NewExpenseService(repo)
	expenses, err := service.GetByUserId(1)
	assert.NoError(t, err)
	assert.Empty(t, expenses)
}

func TestExpenseService_GetByUserIdRepoError(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)
	repo.On("GetAllByUserID", uint(1)).
		Return(nil, utils.ErrDatabase).
		Once()

	service := NewExpenseService(repo)
	expenses, err := service.GetByUserId(1)
	assert.Error(t, err)
	assert.Nil(t, expenses)
}

func TestExpenseService_GetDetailSuccess(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)

	repo.On("GetByIdAndUserId", uint(1), uint(1)).
		Return(&models.Expense{}, nil).
		Once()

	service := NewExpenseService(repo)
	expense, err := service.GetDetail(1, 1)
	assert.NoError(t, err)
	assert.NotNil(t, expense)
}

func TestExpenseService_GetDetailNotFound(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)

	repo.On("GetByIdAndUserId", uint(1), uint(1)).Return(nil, utils.ErrNotFound).Once()

	service := NewExpenseService(repo)
	expense, err := service.GetDetail(1, 1)
	assert.ErrorIs(t, err, utils.ErrNotFound)
	assert.Nil(t, expense)
}

func TestExpenseService_GetDetailRepoError(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)

	repo.On("GetByIdAndUserId", uint(1), uint(1)).Return(nil, utils.ErrDatabase).Once()

	service := NewExpenseService(repo)
	expense, err := service.GetDetail(1, 1)
	assert.ErrorIs(t, err, utils.ErrDatabase)
	assert.Nil(t, expense)
}

func TestExpenseService_UpdateSuccess(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)

	repo.On("GetByIdAndUserId", uint(1), uint(1)).
		Return(&models.Expense{}, nil).
		Once()

	repo.On("Update", mock.AnythingOfType("*models.Expense")).
		Return(nil).
		Once()

	service := NewExpenseService(repo)

	desc := "Updated Expense"
	paymentMethod := "Debit Card"
	expenseDate, _ := time.Parse(
		"2006-01-02",
		"2024-06-01",
	)

	input := requests.ExpenseRequest{
		UserId:        1,
		CategoryId:    new(int),
		Amount:        decimal.NewFromFloat(150.0),
		Description:   &desc,
		ExpenseDate:   expenseDate,
		PaymentMethod: &paymentMethod,
		IsRecurring:   false,
	}

	expense, err := service.Update(1, &input, 1)

	require.Nil(t, err)
	require.NotNil(t, expense)
}

func TestExpenseService_UpdateNotFound(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)

	repo.On("GetByIdAndUserId", uint(1), uint(1)).
		Return(nil, utils.ErrNotFound).
		Once()

	service := NewExpenseService(repo)
	desc := "Updated Expense"
	paymentMethod := "Debit Card"
	expenseDate, _ := time.Parse(
		"2006-01-02",
		"2024-06-01",
	)

	input := requests.ExpenseRequest{
		UserId:        1,
		CategoryId:    new(int),
		Amount:        decimal.NewFromFloat(150.0),
		Description:   &desc,
		ExpenseDate:   expenseDate,
		PaymentMethod: &paymentMethod,
		IsRecurring:   false,
	}

	expense, err := service.Update(1, &input, 1)

	require.ErrorIs(t, err, utils.ErrNotFound)
	require.Nil(t, expense)
}

func TestExpenseService_UpdateInvalidAmount(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)

	service := NewExpenseService(repo)

	desc := "Updated Expense"
	paymentMethod := "Debit Card"
	expenseDate, _ := time.Parse(
		"2006-01-02",
		"2024-06-01",
	)

	input := requests.ExpenseRequest{
		UserId:        1,
		CategoryId:    new(int),
		Amount:        decimal.NewFromFloat(-150.0),
		Description:   &desc,
		ExpenseDate:   expenseDate,
		PaymentMethod: &paymentMethod,
		IsRecurring:   false,
	}

	expense, err := service.Update(1, &input, 1)
	require.ErrorIs(t, err, utils.ErrInvalidAmount)
	require.Nil(t, expense)
}

func TestExpenseService_UpdateDateInFuture(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)

	service := NewExpenseService(repo)
	desc := "Updated Expense"
	paymentMethod := "Debit Card"
	expenseDate := time.Now().Add(24 * time.Hour) // Future date

	input := requests.ExpenseRequest{
		UserId:        1,
		CategoryId:    new(int),
		Amount:        decimal.NewFromFloat(150.0),
		Description:   &desc,
		ExpenseDate:   expenseDate,
		PaymentMethod: &paymentMethod,
		IsRecurring:   false,
	}

	expense, err := service.Update(1, &input, 1)
	require.ErrorIs(t, err, utils.ErrDateInFuture)
	require.Nil(t, expense)
}

func TestExpenseService_UpdateRepoError(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)

	repo.On("GetByIdAndUserId", uint(1), uint(1)).
		Return(nil, utils.ErrDatabase).
		Once()

	service := NewExpenseService(repo)
	desc := "Updated Expense"
	paymentMethod := "Debit Card"
	expenseDate, _ := time.Parse(
		"2006-01-02",
		"2024-06-01",
	)
	input := requests.ExpenseRequest{
		UserId:        1,
		CategoryId:    new(int),
		Amount:        decimal.NewFromFloat(150.0),
		Description:   &desc,
		ExpenseDate:   expenseDate,
		PaymentMethod: &paymentMethod,
		IsRecurring:   false,
	}
	expense, err := service.Update(1, &input, 1)
	require.ErrorIs(t, err, utils.ErrDatabase)
	require.Nil(t, expense)
}

func TestExpenseService_DeleteSuccess(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)

	repo.On("GetByIdAndUserId", uint(1), uint(1)).
		Return(&models.Expense{}, nil).
		Once()

	repo.On("Delete", uint(1)).
		Return(nil).
		Once()

	service := NewExpenseService(repo)
	err := service.Delete(1, 1)
	require.Nil(t, err)
}

func TestExpenseService_DeleteNotFound(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)

	repo.On("GetByIdAndUserId", uint(1), uint(1)).
		Return(nil, utils.ErrNotFound).
		Once()

	service := NewExpenseService(repo)

	err := service.Delete(1, 1)
	require.ErrorIs(t, err, utils.ErrNotFound)
}

func TestExpenseService_DeleteRepoError(t *testing.T) {
	repo := repository.NewMockExpenseRepository(t)
	repo.On("GetByIdAndUserId", uint(1), uint(1)).
		Return(&models.Expense{}, nil).
		Once()
	repo.On("Delete", uint(1)).
		Return(utils.ErrDatabase).
		Once()

	service := NewExpenseService(repo)

	err := service.Delete(1, 1)
	require.ErrorIs(t, err, utils.ErrDatabase)
}
