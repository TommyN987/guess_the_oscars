package service

import (
	"context"
	"errors"

	"github.com/TommyN987/guess_the_oscars/backend/internal/domain"
	"github.com/TommyN987/guess_the_oscars/backend/internal/repository"
	"github.com/jackc/pgx/v5"
)

type Service interface {
	CheckHealth(ctx context.Context) (string, error)
	RegisterUser(ctx context.Context, user domain.User) error
	LoginUser(ctx context.Context, email, password string) (string, error)
	GetAllCategories(ctx context.Context) ([]domain.Category, error)
	GetNominationsByCategory(ctx context.Context, categoryID int) (domain.Category, []domain.Nomination, error)
}

type DefaultService struct {
	repo repository.Repository
}

func NewDefaultService(repo repository.Repository) *DefaultService {
	return &DefaultService{repo: repo}
}

func (s *DefaultService) CheckHealth(ctx context.Context) (string, error) {
	return s.repo.CheckHealth(ctx)
}

func (s *DefaultService) RegisterUser(ctx context.Context, user domain.User) error {
	// Check if the user already exists
	_, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err == nil {
		// No error means the user exists
		return errors.New("User already exists.")
	} else if err != pgx.ErrNoRows {
		// An unexpected error occurred (e.g., database issue)
		return err
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return s.repo.CreateUser(ctx, user)
}

func (s *DefaultService) LoginUser(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("Invalid email or password...!")
	}

	if err := checkPasswordHash(password, user.Password); err != nil {
		return "", errors.New("Invalid email or password.")
	}

	return generateJWT(user)
}

func (s *DefaultService) GetAllCategories(ctx context.Context) ([]domain.Category, error) {
	return s.repo.GetAllCategories(ctx)
}

func (s *DefaultService) GetNominationsByCategory(ctx context.Context, categoryID int) (domain.Category, []domain.Nomination, error) {
	return s.repo.GetNominationsByCategoryID(ctx, categoryID)
}
