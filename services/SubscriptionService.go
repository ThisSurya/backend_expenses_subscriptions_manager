package services

import (
	"backend/models"
	"backend/requests"

	"gorm.io/gorm"
)

type SubscriptionService struct {
	DB *gorm.DB
}

func NewSubscriptionService(db *gorm.DB) *SubscriptionService {
	return &SubscriptionService{
		DB: db,
	}
}

func (s *SubscriptionService) GetAll() ([]models.Subscription, error) {
	var subscriptions []models.Subscription

	err := s.DB.Find(&subscriptions).Error

	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (s *SubscriptionService) GetByUserId(userId int) ([]models.Subscription, error) {
	var subscriptions []models.Subscription

	err := s.DB.Where("user_id = ?", userId).Find(&subscriptions).Error

	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (s *SubscriptionService) GetDetail(id int, userId int) (*models.Subscription, error) {
	var subscription models.Subscription

	err := s.DB.Where("id = ? AND user_id = ?", id, userId).First(&subscription).Error

	if err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (s *SubscriptionService) Create(input requests.SubscriptionRequest, userId int) (*models.Subscription, error) {

	subscription := models.Subscription{
		UserId:          userId,
		Name:            input.Name,
		CategoryId:      input.CategoryId,
		ExpenseId:       input.ExpenseId,
		Amount:          input.Amount,
		BillingCycle:    input.BillingCycle,
		NextBillingDate: input.NextBillingDate,
		IsActive:        input.IsActive,
	}

	err := s.DB.Create(&subscription).Error
	if err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (s *SubscriptionService) Update(id int, input *requests.SubscriptionRequest, userId int) (*models.Subscription, error) {
	var subscription models.Subscription

	err := s.DB.First(&subscription, id).Error

	if err != nil {
		return nil, err
	}

	if userId != subscription.UserId {
		return nil, ErrForbidden
	}

	subscription.Name = input.Name
	subscription.IsActive = input.IsActive
	subscription.Amount = input.Amount
	subscription.BillingCycle = input.BillingCycle
	subscription.CategoryId = input.CategoryId
	subscription.ExpenseId = input.ExpenseId

	err = s.DB.Save(&subscription).Error

	if err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (s *SubscriptionService) Delete(id int, userId int) error {
	var subscription models.Subscription

	err := s.DB.First(&subscription, id).Error

	if err != nil {
		return err
	}

	if userId != subscription.UserId {
		return ErrForbidden
	}

	err = s.DB.Delete(&subscription).Error

	if err != nil {
		return err
	}

	return nil
}
