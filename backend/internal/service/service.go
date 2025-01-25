package service

import (
	"context"

	"github.com/TommyN987/guess_the_oscars/backend/internal/domain"
	"github.com/TommyN987/guess_the_oscars/backend/internal/repository"
)

type Service interface {
	CheckHealth(ctx context.Context) (string, error)
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

func (s *DefaultService) GetAllCategories(ctx context.Context) ([]domain.Category, error) {
	return s.repo.GetAllCategories(ctx)
}

func (s *DefaultService) GetNominationsByCategory(ctx context.Context, categoryID int) (domain.Category, []domain.Nomination, error) {
	return s.repo.GetNominationsByCategoryID(ctx, categoryID)
}
