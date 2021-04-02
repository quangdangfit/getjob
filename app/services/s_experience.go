package services

import (
	"context"

	"github.com/quangdangfit/getjob/app/interfaces"
	"github.com/quangdangfit/getjob/app/models"
	"github.com/quangdangfit/getjob/app/schema"
	"github.com/quangdangfit/getjob/pkg/utils"
)

// ExperienceService handle company logic
type ExperienceService struct {
	repo interfaces.IExperienceRepository
}

// NewExperienceService return new IExperienceService interface
func NewExperienceService(repo interfaces.IExperienceRepository) interfaces.IExperienceService {
	return &ExperienceService{
		repo: repo,
	}
}

// GetByID handle logic get experience by id
func (s *ExperienceService) GetByID(ctx context.Context, id string) (*models.Experience, error) {
	company, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return company, nil
}

// ListByUserID handle logic get list experience by userID
func (s *ExperienceService) ListByUserID(ctx context.Context, userID string) (*models.Experiences, error) {
	company, err := s.repo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return company, nil
}

// GetByID handle logic create new experience for user
func (s *ExperienceService) Create(ctx context.Context, params *schema.ExperienceCreateParams) (*models.Experience, error) {
	var company *models.Experience
	err := utils.Copy(&company, &params)
	if err != nil {
		return nil, err
	}

	err = s.repo.Create(ctx, company)
	if err != nil {
		return nil, err
	}

	return company, nil
}
