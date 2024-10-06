package user

import (
	"context"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/matthewhartstonge/argon2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Insert(ctx context.Context, user User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockRepository) Fetch(ctx context.Context, username string) (User, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(User), args.Error(1)
}

func TestRegister_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUsecase("secret", mockRepo)
	ctx := context.TODO()

	mockRepo.On("Insert", ctx, mock.AnythingOfType("User")).Return(nil)

	err := usecase.Register(ctx, "username", "password")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUsecase("secret", mockRepo)
	ctx := context.TODO()

	argon := argon2.DefaultConfig()
	hashedPassword, _ := argon.HashEncoded([]byte("password"))
	mockRepo.On("Fetch", ctx, "username").Return(User{Username: "username", Password: string(hashedPassword)}, nil)

	token, err := usecase.Login(ctx, "username", "password")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestLogin_WrongPassword(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUsecase("secret", mockRepo)
	ctx := context.TODO()

	argon := argon2.DefaultConfig()
	hashedPassword, _ := argon.HashEncoded([]byte("password"))
	mockRepo.On("Fetch", ctx, "username").Return(User{Username: "username", Password: string(hashedPassword)}, nil)

	token, err := usecase.Login(ctx, "username", "wrongpassword")
	assert.Error(t, err)
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestLogin_UserNotFound(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUsecase("secret", mockRepo)
	ctx := context.TODO()

	mockRepo.On("Fetch", ctx, "username").Return(User{}, errors.New("user not found"))

	token, err := usecase.Login(ctx, "username", "password")
	assert.Error(t, err)
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestCreateToken_ValidToken(t *testing.T) {
	token, err := createToken("username", "secret")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)
}
