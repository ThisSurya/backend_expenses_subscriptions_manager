package services

import (
	"backend/models"
	"backend/requests"

	"gorm.io/gorm"
)

type ExpenseService struct {
	DB *gorm.DB
}

func NewExpenseService(db *gorm.DB) *ExpenseService {
	return &ExpenseService{
		DB: db,
	}
}

func (s *ExpenseService) GetAll() ([]models.Expense, error) {
	var expenses []models.Expense

	err := s.DB.Find(&expenses).Error
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (s *ExpenseService) GetByUserId(UserId int) ([]models.Expense, error) {
	var expenses []models.Expense

	err := s.DB.Where("user_id = ?", UserId).Find(&expenses).Error

	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func (s *ExpenseService) Create(input *requests.ExpenseRequest, userId int) (*models.Expense, error) {
	expense := models.Expense{
		UserId:        userId,
		CategoryId:    input.CategoryId,
		Amount:        input.Amount,
		Description:   input.Description,
		ExpenseDate:   input.ExpenseDate,
		PaymentMethod: input.PaymentMethod,
		IsRecurring:   input.IsRecurring,
	}

	err := s.DB.Create(&expense).Error
	if err != nil {
		return nil, err
	}

	return &expense, nil
}

func (s *ExpenseService) GetDetail(id int, userId int) (*models.Expense, error) {
	var expense models.Expense
	err := s.DB.First(&expense, id).Error
	if err != nil {
		return nil, err
	}

	if expense.UserId != userId {
		return nil, ErrForbidden
	}

	return &expense, nil
}

func (s *ExpenseService) Update(id int, input *requests.ExpenseRequest, userId int) (*models.Expense, error) {
	var expense models.Expense
	err := s.DB.First(&expense, id).Error
	if err != nil {
		return nil, err
	}

	if expense.UserId != userId {
		return nil, ErrForbidden
	}

	expense.CategoryId = input.CategoryId
	expense.Amount = input.Amount
	expense.Description = input.Description
	expense.ExpenseDate = input.ExpenseDate
	expense.PaymentMethod = input.PaymentMethod
	expense.IsRecurring = input.IsRecurring

	err = s.DB.Save((&expense)).Error
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

func (s *ExpenseService) Delete(id int, userId int) (bool, error) {
	var expense models.Expense
	err := s.DB.First(&expense, id).Error
	if err != nil {
		return false, err
	}

	if expense.UserId != userId {
		return false, ErrForbidden
	}

	err = s.DB.Delete(&expense).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
