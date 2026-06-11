package services

import (
	"backend/models"
	"backend/repository"
	"backend/requests"
	"backend/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	userRepo  repository.UserRepository
	jwtSecret []byte
	tokenExp  time.Duration
}

func NewUserService(r repository.UserRepository, jwtSecret []byte, tokenExp time.Duration) *UserService {
	return &UserService{
		userRepo:  r,
		jwtSecret: jwtSecret,
		tokenExp:  tokenExp,
	}
}

// func (s *UserService) GetAll() ([]models.User, error) {
// 	users := s.userRepo.
// }

func (s *UserService) CheckEmailExists(email string) (bool, error) {

	exists, err := s.userRepo.CheckEmailExists(email)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (s *UserService) RegisterService(input *requests.UserRequest) (*models.User, error) {
	password, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	exists, err := s.userRepo.CheckEmailExists(input.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, utils.ErrEmailExists
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

	err = s.userRepo.Create(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) LoginService(input *requests.UserLoginRequest) (interface{}, error) {
	user, err := s.userRepo.GetByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	check := utils.CheckPasswordHash(input.Password, user.Password)

	if !check {
		return nil, utils.ErrInvalidCredentials
	}

	now := time.Now()

	claims := jwt.MapClaims{
		"user_id": user.Id,
		"email":   user.Email,
		"iat":     now.Unix(),
		"exp":     now.Add(s.tokenExp).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"token":      tokenString,
		"user":       user,
		"type":       "Bearer",
		"expires_in": int64(s.tokenExp.Seconds()),
	}

	return data, nil
}

func (s *UserService) GetDetail(id uint) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Update(id uint, input *requests.UserRequest) (*models.User, error) {
	check, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if check.Id != int(id) {
		return nil, utils.ErrNotFound
	}

	user := models.User{
		Username:        input.Username,
		Email:           input.Email,
		Password:        input.Password,
		AvatarUrl:       input.AvatarUrl,
		ReminderEnabled: false,
		ReminderDays:    3,
		Timezone:        "UTC",
	}

	err = s.userRepo.Update(id, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) Delete(id int) error {
	err := s.userRepo.Delete(uint(id))

	if err != nil {
		return err
	}

	return nil
}
