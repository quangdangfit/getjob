package interfaces

import (
	"context"

	"github.com/quangdangfit/getjob/app/models"
)

type IExperienceRepository interface {
	GetByID(ctx context.Context, id string) (*models.Experience, error)
	ListByUserID(ctx context.Context, userID string) (*models.Experiences, error)
	Create(ctx context.Context, experience *models.Experience) error
}
