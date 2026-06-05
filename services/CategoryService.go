package services

import (
	"backend/models"
	"backend/requests"

	"gorm.io/gorm"
)

type CategoryService struct {
	DB *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{
		DB: db,
	}
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	var categories []models.Category

	err := s.DB.Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoryService) GetByUserId(userId int) ([]models.Category, error) {
	var categories []models.Category

	err := s.DB.Where("user_id = ?", userId).Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoryService) Create(input *requests.CategoryRequest, userId int) (*models.Category, error) {
	category := models.Category{
		UserId:  userId,
		Name:    input.Name,
		IconUrl: input.IconUrl,
		Color:   input.Color,
	}

	err := s.DB.Create(&category).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (s *CategoryService) GetDetail(id int, userId int) (*models.Category, error) {
	var category models.Category
	err := s.DB.First(&category, id).Error

	if err != nil {
		return nil, err
	}

	if category.UserId != userId {
		return nil, ErrForbidden
	}

	return &category, nil
}

func (s *CategoryService) Update(id int, input *requests.CategoryRequest, userId int) (*models.Category, error) {
	var category models.Category

	err := s.DB.First(&category, id).Error

	if err != nil {
		return nil, err
	}

	if category.UserId != userId {
		return nil, ErrForbidden
	}

	category.UserId = userId
	category.Name = input.Name
	category.IconUrl = input.IconUrl
	category.Color = input.Color

	err = s.DB.Save(&category).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (s *CategoryService) Delete(id int, userId int) (bool, error) {
	var category models.Category

	err := s.DB.First(&category, id).Error

	if err != nil {
		return false, err
	}

	if category.UserId != userId {
		return false, ErrForbidden
	}

	err = s.DB.Delete(&category).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
