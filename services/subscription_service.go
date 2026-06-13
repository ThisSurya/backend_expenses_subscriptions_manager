package services

import (
	"backend/models"
	"backend/repository"
	"backend/requests"
	"backend/utils"
	"errors"

	"github.com/shopspring/decimal"
)

type SubscriptionService struct {
	subRepo repository.SubscriptionRepository
}

func NewSubscriptionService(subRepo repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{
		subRepo: subRepo,
	}
}

func checkAmount(amount decimal.Decimal) bool {
	return amount.GreaterThan(decimal.Zero)
}

func (s *SubscriptionService) GetByUserId(userId uint) ([]models.Subscription, error) {
	subscriptions, err := s.subRepo.GetAllByUserID(userId)

	if err != nil {
		return nil, utils.ErrDatabase
	}

	return subscriptions, nil
}

func (s *SubscriptionService) GetDetail(id uint, userId uint) (*models.Subscription, error) {
	subscription, err := s.subRepo.GetByIDAndUserId(id, userId)

	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func (s *SubscriptionService) Create(input requests.SubscriptionRequest, userId int) (*models.Subscription, error) {

	if !checkAmount(input.Amount) {
		return nil, utils.ErrInvalidAmount
	}

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

	err := s.subRepo.Create(&subscription)

	if err != nil {
		return nil, utils.ErrDatabase
	}

	return &subscription, nil
}

func (s *SubscriptionService) Update(id uint, input *requests.SubscriptionRequest, userId int) (*models.Subscription, error) {
	if !checkAmount(input.Amount) {
		return nil, utils.ErrInvalidAmount
	}

	subscription := models.Subscription{
		Name:            input.Name,
		UserId:          userId,
		CategoryId:      input.CategoryId,
		ExpenseId:       input.ExpenseId,
		Amount:          input.Amount,
		BillingCycle:    input.BillingCycle,
		NextBillingDate: input.NextBillingDate,
		IsActive:        input.IsActive,
	}

	_, err := s.subRepo.GetByIDAndUserId(id, uint(userId))
	if err != nil {
		return nil, err
	}

	err = s.subRepo.Update(id, &subscription)

	if err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (s *SubscriptionService) Delete(id uint, userId int) error {
	_, err := s.subRepo.GetByIDAndUserId(id, uint(userId))

	if err != nil {
		return err
	}

	err = s.subRepo.Delete(id)

	if err != nil {
		// return err

		if errors.Is(err, utils.ErrNotFound) {
			return utils.ErrNotFound
		}
		return utils.ErrDatabase
	}

	return err
}
