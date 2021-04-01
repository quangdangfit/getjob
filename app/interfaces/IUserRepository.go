package interfaces

import (
	"context"

	"github.com/quangdangfit/getjob/app/models"
)

type IUserRepository interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}
