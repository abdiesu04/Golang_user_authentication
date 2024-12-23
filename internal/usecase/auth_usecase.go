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

	token, err := u.jwtUtils.GenerateToken(user.ID)
	if err != nil {
		return domain.Token{}, err
	}

	// Save refresh token in the database
	if err := u.userRepo.SaveRefreshToken(user.ID, token.RefreshToken); err != nil {
		return domain.Token{}, err
	}

	return token, nil
}

func (u *AuthUsecase) RefreshToken(refreshToken string) (string, error) {
	claims, err := u.jwtUtils.ValidateToken(refreshToken)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return "", errors.New("invalid refresh token claims")
	}

	// Verify refresh token from database
	storedToken, err := u.userRepo.GetRefreshToken(int64(userID))
	if err != nil || storedToken != refreshToken {
		return "", errors.New("refresh token mismatch or not found")
	}

	// Generate a new access token
	newToken, err := u.jwtUtils.GenerateToken(int64(userID))
	return newToken.AccessToken, err
}

func (u *AuthUsecase) Logout(refreshToken string) error {
	claims, err := u.jwtUtils.ValidateToken(refreshToken)
	if err != nil {
		return errors.New("invalid refresh token")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return errors.New("invalid refresh token claims")
	}

	// Invalidate refresh token
	return u.userRepo.InvalidateRefreshToken(int64(userID))
}
