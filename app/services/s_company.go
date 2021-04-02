package services

import (
	"context"

	"github.com/quangdangfit/getjob/app/interfaces"
	"github.com/quangdangfit/getjob/app/models"
)

// CompanyService handle company logic
type CompanyService struct {
	repo interfaces.ICompanyRepository
}

// NewCompanyService return new ICompanyService interface
func NewCompanyService(repo interfaces.ICompanyRepository) interfaces.ICompanyService {
	return &CompanyService{
		repo: repo,
	}
}

// GetByID handle logic get company by id
func (s *CompanyService) GetByID(ctx context.Context, id string) (*models.Company, error) {
	company, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return company, nil
}
