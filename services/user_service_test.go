package services

import (
	"backend/models"
	repository "backend/repository/mocks"
	"backend/requests"
	"backend/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUserService_RegisterSuccess(t *testing.T) {
	repo := repository.NewMockUserRepository(t)

	repo.On("CheckEmailExists", "test@example.com").
		Return(false, nil).
		Once()

	repo.On("Create", mock.AnythingOfType("*models.User")).
		Return(nil).
		Once()

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	username := "testuser"
	email := "test@example.com"
	role := requests.UserRoleBasic
	ReminderEnabled := false
	ReminderDays := 0
	timezone := "UTC"
	input := requests.UserRequest{
		Username:        &username,
		Email:           email,
		Password:        "password",
		Role:            &role,
		ReminderEnabled: &ReminderEnabled,
		ReminderDays:    &ReminderDays,
		Timezone:        &timezone,
	}

	_, err := service.RegisterService(&input)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	assert.NoError(t, err)
}

func TestUserService_RegisterEmailExists(t *testing.T) {
	repo := repository.NewMockUserRepository(t)

	repo.On("CheckEmailExists", "test@example.com").
		Return(true, nil).
		Once()

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	username := "testuser"
	email := "test@example.com"
	role := requests.UserRoleBasic
	ReminderEnabled := false
	ReminderDays := 0
	timezone := "UTC"
	input := requests.UserRequest{
		Username:        &username,
		Email:           email,
		Password:        "password",
		Role:            &role,
		ReminderEnabled: &ReminderEnabled,
		ReminderDays:    &ReminderDays,
		Timezone:        &timezone,
	}

	_, err := service.RegisterService(&input)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	require.ErrorIs(t, err, utils.ErrEmailExists)
}

func TestUserService_RegisterRepoError(t *testing.T) {
	repo := repository.NewMockUserRepository(t)

	repo.On("CheckEmailExists", "test@example.com").
		Return(false, nil).
		Once()

	repo.On("Create", mock.AnythingOfType("*models.User")).
		Return(utils.ErrDatabase).
		Once()

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	username := "testuser"
	email := "test@example.com"
	role := requests.UserRoleBasic
	ReminderEnabled := false
	ReminderDays := 0
	timezone := "UTC"
	input := requests.UserRequest{
		Username:        &username,
		Email:           email,
		Password:        "password",
		Role:            &role,
		ReminderEnabled: &ReminderEnabled,
		ReminderDays:    &ReminderDays,
		Timezone:        &timezone,
	}

	_, err := service.RegisterService(&input)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	require.ErrorIs(t, err, utils.ErrDatabase)
}

func TestUserService_LoginSuccess(t *testing.T) {

	repo := repository.NewMockUserRepository(t)

	hashedPassword, err := utils.HashPassword("password")
	require.NoError(t, err)

	repo.On("GetByEmail", "test@example.com").
		Return(&models.User{
			Password: hashedPassword,
			Email:    "test@example.com",
		}, nil).
		Once()
	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	input := requests.UserLoginRequest{
		Email:    "test@example.com",
		Password: "password",
	}

	_, err = service.LoginService(&input)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	assert.NoError(t, err)
}

func TestUserService_LoginEmailNotFound(t *testing.T) {
	repo := repository.NewMockUserRepository(t)
	email := "test@example.com"

	repo.On("GetByEmail", email).
		Return(nil, nil).
		Once()

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	input := requests.UserLoginRequest{
		Email:    email,
		Password: "password",
	}

	_, err := service.LoginService(&input)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	require.ErrorIs(t, err, utils.ErrInvalidCredentials)
}

func TestUserService_LoginInvalidPassword(t *testing.T) {
	repo := repository.NewMockUserRepository(t)

	hashedPassword, err := utils.HashPassword("password")
	require.NoError(t, err)

	email := "test@example.com"
	repo.On("GetByEmail", email).
		Return(&models.User{
			Password: hashedPassword,
			Email:    email,
		}, nil).
		Once()

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	input := requests.UserLoginRequest{
		Email:    email,
		Password: "wrongpassword",
	}

	_, err = service.LoginService(&input)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	require.ErrorIs(t, err, utils.ErrInvalidCredentials)
}

func TestUserService_LoginRepoError(t *testing.T) {
	repo := repository.NewMockUserRepository(t)

	email := "test@example.com"

	repo.On("GetByEmail", email).
		Return(nil, utils.ErrDatabase).
		Once()

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	input := requests.UserLoginRequest{
		Email:    email,
		Password: "password",
	}

	_, err := service.LoginService(&input)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	require.ErrorIs(t, err, utils.ErrDatabase)
}

func TestUserService_GetDetailSuccess(t *testing.T) {
	repo := repository.NewMockUserRepository(t)

	repo.On("GetByID", uint(1)).
		Return(&models.User{}, nil).
		Once()

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	_, err := service.GetDetail(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	assert.NoError(t, err)
}

func TestUserService_GetDetailNotFound(t *testing.T) {
	repo := repository.NewMockUserRepository(t)

	repo.On("GetByID", uint(1)).
		Return(nil, utils.ErrNotFound).
		Once()

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	_, err := service.GetDetail(1)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	require.ErrorIs(t, err, utils.ErrNotFound)
}

func TestUserService_GetDetailRepoError(t *testing.T) {
	repo := repository.NewMockUserRepository(t)

	repo.On("GetByID", uint(1)).
		Return(nil, utils.ErrDatabase).
		Once()

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	_, err := service.GetDetail(1)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	require.ErrorIs(t, err, utils.ErrDatabase)
}

func TestUserService_UpdateSuccess(t *testing.T) {
	repo := repository.NewMockUserRepository(t)

	email := "test@example.com"
	hashedPassword, err := utils.HashPassword("password")
	require.NoError(t, err)
	repo.On("GetByID", uint(1)).Return(&models.User{
		Email:    email,
		Password: hashedPassword,
		Id:       1,
	}, nil).Once()

	repo.On("Update", uint(1), mock.AnythingOfType("*models.User")).
		Return(nil).
		Once()

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	input := requests.UserRequest{
		Email:    email,
		Password: "password",
	}

	_, err = service.Update(1, &input)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	assert.NoError(t, err)
}

func TestUserService_UpdateOtherUser(t *testing.T) {
	repo := repository.NewMockUserRepository(t)

	email := "test@example.com"

	repo.On("GetByID", uint(1)).Return(&models.User{
		Email: email,
		Id:    2,
	}, nil).Once()

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	input := requests.UserRequest{
		Email:    email,
		Password: "password",
	}

	_, err := service.Update(1, &input)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	require.ErrorIs(t, err, utils.ErrNotFound)
}

func TestUserService_UpdateEmailExists(t *testing.T) {
	repo := repository.NewMockUserRepository(t)

	email := "test@example.com"
	hashedPassword, err := utils.HashPassword("password")
	require.NoError(t, err)

	repo.On("GetByID", uint(1)).Return(&models.User{
		Email:    "different@example.com",
		Password: hashedPassword,
		Id:       1,
	}, nil).Once()

	repo.On("CheckEmailExists", email).
		Return(true, nil).
		Once()

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	input := requests.UserRequest{
		Email:    email,
		Password: "password",
	}

	_, err = service.Update(1, &input)
	require.Error(t, err)
	require.ErrorIs(t, err, utils.ErrEmailExists)
}

func TestUserService_UpdateRepoError(t *testing.T) {
	repo := repository.NewMockUserRepository(t)

	email := "test@example.com"

	repo.On("GetByID", uint(1)).
		Return(&models.User{
			Email: email,
			Id:    1,
		}, nil).Once()

	repo.On("Update", uint(1), mock.AnythingOfType("*models.User")).
		Return(utils.ErrDatabase).
		Once()

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	input := requests.UserRequest{
		Email:    email,
		Password: "password",
	}

	_, err := service.Update(1, &input)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	require.ErrorIs(t, err, utils.ErrDatabase)
}

func TestUserService_DeleteSuccess(t *testing.T) {
	repo := repository.NewMockUserRepository(t)

	repo.On("Delete", uint(1)).
		Return(nil).
		Once()

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	err := service.Delete(1, 1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	assert.NoError(t, err)
}

func TestUserService_DeleteNotFound(t *testing.T) {
	repo := repository.NewMockUserRepository(t)

	repo.On("Delete", uint(1)).
		Return(utils.ErrNotFound).
		Once()

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	err := service.Delete(1, 1)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	require.ErrorIs(t, err, utils.ErrNotFound)
}

func TestUserService_DeleteOtherUser(t *testing.T) {

	repo := repository.NewMockUserRepository(t)

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	err := service.Delete(1, 2)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	require.ErrorIs(t, err, utils.ErrNotFound)
}

func TestUserService_DeleteRepoError(t *testing.T) {
	repo := repository.NewMockUserRepository(t)

	repo.On("Delete", uint(1)).
		Return(utils.ErrDatabase).
		Once()

	service := NewUserService(repo, []byte("secret"), 24*time.Hour)

	err := service.Delete(1, 1)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	require.ErrorIs(t, err, utils.ErrDatabase)
}
