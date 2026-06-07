package services

import (
	"backend/models"
	"backend/repository"
	"backend/requests"
	"errors"

	"gorm.io/gorm"
)

type CategoryService struct {
	CR repository.CategoryRepository
}

func NewCategoryService(cr repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		CR: cr,
	}
}

func (s *CategoryService) GetByUserId(userId uint) ([]models.Category, error) {
	categories, err := s.CR.GetAllByUserID(userId)
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

	err := s.CR.Create(&category)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (s *CategoryService) GetDetail(id uint, userId uint) (*models.Category, error) {
	category, err := s.CR.GetByIDAndUserId(id, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Resources not found")
		}
		return nil, err
	}
	return category, nil
}

func (s *CategoryService) Update(id uint, input *requests.CategoryRequest, userId uint) (*models.Category, error) {
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

func (s *CategoryService) Delete(id uint, userId uint) (bool, error) {
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
