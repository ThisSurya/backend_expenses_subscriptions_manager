package repository

import (
	"backend/models"
	"backend/utils"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(expense *models.User) error
	GetByID(ID uint) (*models.User, error)
	Update(ID uint, expense *models.User) error
	Delete(ID uint) error
	CheckEmailExists(email string) (bool, error)
	GetByEmail(email string) (*models.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (r *UserRepositoryImpl) Create(model *models.User) error {
	return r.DB.Create(model).Error
}

func (r *UserRepositoryImpl) GetByID(ID uint) (*models.User, error) {
	var user *models.User

	err := r.DB.First(&user, ID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) Update(ID uint, model *models.User) error {
	var user models.User

	return r.DB.Model(&user).Updates(model).Error
}

func (r *UserRepositoryImpl) Delete(ID uint) error {
	var user models.User

	result := r.DB.Where("id = ? ", ID).Delete(&user)

	if result.RowsAffected == 0 {
		return utils.ErrNotFound
	}

	return result.Error
}

func (r *UserRepositoryImpl) CheckEmailExists(email string) (bool, error) {
	var exists bool

	err := r.DB.Raw("SELECT 1 FROM users WHERE email = ?", email).Scan(&exists).Error

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.DB.Where("email = ?", email).Find(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}
