// internal/service/auth.go
package service

import (
	"context"
	"errors"
	"ums/internal/domain/model"
	"ums/internal/repository"
	"ums/internal/utils"
)

type AuthService interface {
	Register(ctx context.Context, user *model.User) error
	Login(ctx context.Context, username, password string) (string, *model.User, error)
}

type authService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo repository.UserRepository, jwtSecret string) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *authService) Register(ctx context.Context, user *model.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.userRepo.Create(ctx, user)
}

func (s *authService) Login(ctx context.Context, username, password string) (string, *model.User, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", nil, errors.New("invalid credentials")
	}
	token, err := utils.GenerateJWT(user.ID, user.Role, s.jwtSecret)
	return token, user, err
}
