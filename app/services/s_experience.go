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
	repo     interfaces.IExperienceRepository
	userRepo interfaces.IUserRepository
}

// NewExperienceService return new IExperienceService interface
func NewExperienceService(repo interfaces.IExperienceRepository, userRepo interfaces.IUserRepository) interfaces.IExperienceService {
	return &ExperienceService{
		repo:     repo,
		userRepo: userRepo,
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

// Create handle logic create new experience for user
func (s *ExperienceService) Create(ctx context.Context, userID string, params *schema.ExperienceCreateParams) (*models.Experience, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var experience *models.Experience
	err = utils.Copy(&experience, &params)
	if err != nil {
		return nil, err
	}

	experience.UserID = user.ID
	err = s.repo.Create(ctx, experience)
	if err != nil {
		return nil, err
	}

	// Update current experience for user
	user.CompanyID = experience.CompanyID
	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return experience, nil
}
