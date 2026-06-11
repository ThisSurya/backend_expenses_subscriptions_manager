package repository

import (
	"backend/models"
	"backend/utils"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	GetByIDAndUserId(ID uint, userId uint) (*models.Category, error)
	Update(category *models.Category) error
	Delete(ID uint) error
	GetAllByUserID(userID uint) ([]models.Category, error)
}

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{
		DB: db,
	}
}

func (r *CategoryRepositoryImpl) Create(model *models.Category) error {
	return r.DB.Create(model).Error
}

func (r *CategoryRepositoryImpl) GetByIDAndUserId(ID uint, userId uint) (*models.Category, error) {
	var category models.Category
	err := r.DB.Where("id = ? AND user_id = ?", ID, userId).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepositoryImpl) Update(model *models.Category) error {
	var category models.Category

	return r.DB.Model(&category).Updates(model).Error
}

func (r *CategoryRepositoryImpl) Delete(ID uint) error {
	var category models.Category

	result := r.DB.Where("id = ?", ID).Delete(&category)

	if result.RowsAffected == 0 {
		return utils.ErrNotFound
	}

	return result.Error
}

func (r *CategoryRepositoryImpl) GetAllByUserID(userID uint) ([]models.Category, error) {
	var categories []models.Category

	err := r.DB.Where("user_id = ?", userID).Find(&categories).Error

	if err != nil {
		return nil, err
	}
	return categories, nil
}
