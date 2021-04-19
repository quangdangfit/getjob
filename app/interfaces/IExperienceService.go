package interfaces

import (
	"context"

	"github.com/quangdangfit/getjob/app/models"
	"github.com/quangdangfit/getjob/app/schema"
)

type IExperienceService interface {
	GetByID(ctx context.Context, id string) (*models.Experience, error)
	ListByUserID(ctx context.Context, userID string) (*models.Experiences, error)
	Create(ctx context.Context, userID string, params *schema.ExperienceCreateParams) (*models.Experience, error)
}
