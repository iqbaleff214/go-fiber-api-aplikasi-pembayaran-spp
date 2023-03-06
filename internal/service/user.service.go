package service

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/iqbaleff214/go-fiber-api-aplikasi-pembayaran-spp/config"
	"github.com/iqbaleff214/go-fiber-api-aplikasi-pembayaran-spp/internal/entity"
	"github.com/iqbaleff214/go-fiber-api-aplikasi-pembayaran-spp/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService interface {
	Authenticate(username, password string) (entity.User, string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return userService{repo: repo}
}

func (u userService) Authenticate(username, password string) (entity.User, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if username == "" {
		return entity.User{}, "", ErrUsernameMandatory
	}

	if password == "" {
		return entity.User{}, "", ErrPasswordMandatory
	}

	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil || user == (entity.User{}) {
		return entity.User{}, "", ErrUserNotFound
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return entity.User{}, "", ErrPasswordIncorrect
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["name"] = user.Name
	claims["username"] = user.Username
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.NewConfig().SecretJWT()))
	if err != nil {
		return entity.User{}, "", ErrSignedToken
	}

	return user, t, nil
}
