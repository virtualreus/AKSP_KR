package service

import (
	"context"
	"errors"
	"fmt"

	"auth-service/internal/models"
	"auth-service/internal/repository"
	"auth-service/pkg/jwtutil"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailExists    = errors.New("email already exists")
	ErrInvalidAuth    = errors.New("invalid email or password")
	ErrUserNotFound   = errors.New("user not found")
	minPasswordLength = 6
)

type AuthService struct {
	repo      *repository.UserRepository
	jwtSecret string
}

func NewAuthService(repo *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{repo: repo, jwtSecret: jwtSecret}
}

func (s *AuthService) Register(ctx context.Context, email, password, firstName, lastName string) (*models.User, string, error) {
	if len(password) < minPasswordLength {
		return nil, "", fmt.Errorf("password too short")
	}

	existing, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, "", err
	}
	if existing != nil {
		return nil, "", ErrEmailExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", fmt.Errorf("hash password: %w", err)
	}

	user := &models.User{
		Email:        email,
		PasswordHash: string(hash),
		FirstName:    firstName,
		LastName:     lastName,
	}
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, "", err
	}

	token, err := jwtutil.Generate(s.jwtSecret, user.ID, user.Email)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*models.User, string, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, "", err
	}
	if user == nil {
		return nil, "", ErrInvalidAuth
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, "", ErrInvalidAuth
	}

	token, err := jwtutil.Generate(s.jwtSecret, user.ID, user.Email)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *AuthService) GetUser(ctx context.Context, id string) (*models.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

