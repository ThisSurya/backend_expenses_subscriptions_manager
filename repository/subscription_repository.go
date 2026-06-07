package repository

import (
	"backend/models"
	"errors"

	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	Create(subscription *models.Subscription) error
	GetByIDAndUserId(ID uint, userID uint) (*models.Subscription, error)
	Update(ID uint, subscription *models.Subscription) error
	Delete(ID uint) error
	GetAllByUserID(userID uint) ([]models.Subscription, error)
}

type SubscriptionRepositoryImpl struct {
	DB *gorm.DB
}

func NewSubscriptionController(db *gorm.DB) *SubscriptionRepositoryImpl {
	return &SubscriptionRepositoryImpl{
		DB: db,
	}
}

func (r *SubscriptionRepositoryImpl) Create(model *models.Expense) error {
	return r.DB.Create(model).Error
}

func (r *SubscriptionRepositoryImpl) GetByIDAndUserId(ID uint, userID uint) (*models.Subscription, error) {
	var subscription models.Subscription

	err := r.DB.Where("id = ? AND user_id = ?", ID, userID).
		First(&subscription).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &subscription, nil
}

func (r *SubscriptionRepositoryImpl) Update(ID uint, model *models.Subscription) error {
	var subscription models.Subscription

	return r.DB.Model(&subscription).Updates(model).Error
}

func (r *SubscriptionRepositoryImpl) Delete(ID uint) error {
	var subscription models.Subscription

	result := r.DB.Where("id = ?", ID).Delete(&subscription)

	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return result.Error
}

func (r *SubscriptionRepositoryImpl) GetAllByUserId(userID uint) ([]models.Subscription, error) {
	var subscriptions []models.Subscription

	err := r.DB.Where("user_id = ?", userID).Find(&subscriptions).Error

	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}
