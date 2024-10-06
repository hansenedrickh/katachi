package user

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/matthewhartstonge/argon2"
)

type repository interface {
	Insert(ctx context.Context, user User) error
	Fetch(ctx context.Context, username string) (User, error)
}

type Usecase struct {
	jwtSecretToken string
	repository     repository
}

func NewUsecase(jwtSecretToken string, repo repository) *Usecase {
	return &Usecase{
		jwtSecretToken: jwtSecretToken,
		repository:     repo,
	}
}

func (u *Usecase) Register(ctx context.Context, username, password string) error {
	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return err
	}

	user := User{
		Username: username,
		Password: string(encoded),
	}

	return u.repository.Insert(ctx, user)
}

func (u *Usecase) Login(ctx context.Context, username, password string) (string, error) {
	user, err := u.repository.Fetch(ctx, username)
	if err != nil {
		return "", err
	}

	isVerified, err := argon2.VerifyEncoded([]byte(password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	if !isVerified {
		return "", errors.New("wrong password")
	}

	token, err := createToken(user.Username, u.jwtSecretToken)
	if err != nil {
		return "", err
	}

	return token, nil
}

func createToken(username, secret string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": "katachi",
		"aud": "user",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	token, err := claims.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
