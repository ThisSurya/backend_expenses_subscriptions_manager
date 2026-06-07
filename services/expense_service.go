package services

import (
	"backend/models"
	"backend/repository"
	"backend/requests"
)

type ExpenseService struct {
	expoRepo repository.ExpenseRepository
}

func NewExpenseService(expo repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{
		expoRepo: expo,
	}
}

func (r *ExpenseService) GetByUserId(UserId uint) ([]models.Expense, error) {
	expenses, err := r.expoRepo.GetAllByUserID(UserId)

	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *ExpenseService) Create(input *requests.ExpenseRequest, userId int) (*models.Expense, error) {
	expense := models.Expense{
		UserId:        userId,
		CategoryId:    input.CategoryId,
		Amount:        input.Amount,
		Description:   input.Description,
		ExpenseDate:   input.ExpenseDate,
		PaymentMethod: input.PaymentMethod,
		IsRecurring:   input.IsRecurring,
	}

	err := r.expoRepo.Create(&expense)

	if err != nil {
		return nil, err
	}

	return &expense, nil
}

func (r *ExpenseService) GetDetail(id uint, userId uint) (*models.Expense, error) {
	expense, err := r.expoRepo.GetByIdAndUserId(id, userId)

	if err != nil {
		return nil, err
	}

	return expense, nil
}

func (r *ExpenseService) Update(id uint, input *requests.ExpenseRequest, userId uint) (*models.Expense, error) {
	expense, err := r.expoRepo.GetByIdAndUserId(id, userId)

	if err != nil {
		return nil, err
	}

	expense.CategoryId = input.CategoryId
	expense.Amount = input.Amount
	expense.Description = input.Description
	expense.ExpenseDate = input.ExpenseDate
	expense.PaymentMethod = input.PaymentMethod
	expense.IsRecurring = input.IsRecurring

	err = r.expoRepo.Update(expense)
	if err != nil {
		return nil, err
	}

	return expense, nil
}

func (r *ExpenseService) Delete(id uint, userId uint) (bool, error) {
	_, err := r.expoRepo.GetByIdAndUserId(id, userId)
	if err != nil {
		return false, err
	}

	err = r.expoRepo.Delete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}
