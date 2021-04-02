package interfaces

import (
	"context"

	"github.com/quangdangfit/getjob/app/models"
)

type IUserRepository interface {
	// Read
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)

	// Write
	Create(ctx context.Context, user *models.User) error
}
