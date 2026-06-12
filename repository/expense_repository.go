package repository

import (
	"backend/models"
	"backend/utils"
	"errors"

	"gorm.io/gorm"
)

type ExpenseRepository interface {
	Create(expense *models.Expense) error
	GetByIdAndUserId(ID uint, userID uint) (*models.Expense, error)
	Update(expense *models.Expense) error
	Delete(ID uint) error
	GetAllByUserID(userID uint) ([]models.Expense, error)
}

type ExpenseRepositoryImpl struct {
	DB *gorm.DB
}

func NewExpenseRepository(
	db *gorm.DB,
) *ExpenseRepositoryImpl {
	return &ExpenseRepositoryImpl{
		DB: db,
	}
}

func (r *ExpenseRepositoryImpl) Create(
	expense *models.Expense,
) error {
	return r.DB.Create(expense).Error
}

func (r *ExpenseRepositoryImpl) GetAllByUserID(
	userID uint,
) ([]models.Expense, error) {

	var expenses []models.Expense

	err := r.DB.
		Where("user_id = ?", userID).
		Find(&expenses).
		Error

	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *ExpenseRepositoryImpl) GetByIdAndUserId(ID uint, userID uint) (*models.Expense, error) {
	var expense models.Expense

	err := r.DB.Where("id = ? and user_id = ?", ID, userID).First(&expense).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}

		return nil, utils.ErrDatabase
	}

	return &expense, nil
}

func (r *ExpenseRepositoryImpl) Update(model *models.Expense) error {
	var expense models.Expense

	return r.DB.Model(&expense).Updates(model).Error
}

func (r *ExpenseRepositoryImpl) Delete(ID uint) error {
	result := r.DB.Where("id = ?", ID).Delete(&models.Expense{})

	if result.RowsAffected == 0 {
		return utils.ErrNotFound
	}

	return result.Error
}
