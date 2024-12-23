package usecase

import (
	"errors"
	"user/internal/domain"
	"user/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	userRepo *repository.UserRepository
	jwtUtils *domain.JWTUtils
}

func NewAuthUsecase(userRepo *repository.UserRepository, jwtSecret string) *AuthUsecase {
	return &AuthUsecase{
		userRepo: userRepo,
		jwtUtils: domain.NewJWTUtils(jwtSecret),
	}
}

func (u *AuthUsecase) Register(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := domain.User{
		Username: username,
		Password: string(hashedPassword),
	}
	return u.userRepo.CreateUser(user)
}

func (u *AuthUsecase) Login(username, password string) (domain.Token, error) {
	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return domain.Token{}, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return domain.Token{}, errors.New("invalid credentials")
	}

	return u.jwtUtils.GenerateToken(user.ID)
}
