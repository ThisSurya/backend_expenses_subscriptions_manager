package services

import (
	"backend/models"
	"backend/requests"
	"backend/utils"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (s *UserService) GetAll() ([]models.User, error) {
	var users []models.User

	err := s.DB.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) CheckEmailExists(email string) (bool, error) {
	var exists bool

	err := s.DB.Raw("SELECT EXISTS(SELECT 1 FROM users where email = $1)", email).Scan(&exists).Error

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (s *UserService) Create(input *requests.UserRequest) (*models.User, error) {
	password, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username:        input.Username,
		Email:           input.Email,
		Password:        password,
		Role:            "basic",
		ReminderEnabled: false,
		ReminderDays:    0,
		Timezone:        "UTC",
	}

	err = s.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetDetail(id int) (*models.User, error) {
	var user models.User
	err := s.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) Update(id int, input *requests.UserRequest) (*models.User, error) {
	var user models.User
	err := s.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	user.Role = models.UserRole(*input.Role)
	user.ReminderEnabled = *input.ReminderEnabled
	user.ReminderDays = *input.ReminderDays
	user.Timezone = *input.Timezone

	s.DB.Save(&user)

	return &user, nil
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {
	var user models.User

	err := s.DB.Where("email", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) Delete(id int) error {
	var user models.User
	err := s.DB.First(&user, id).Error
	if err != nil {
		return err
	}
	s.DB.Delete(&user)
	return nil
}
