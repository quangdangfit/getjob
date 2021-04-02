package interfaces

import (
	"context"

	"github.com/quangdangfit/getjob/app/models"
)

type ICompanyRepository interface {
	// Read
	GetByID(ctx context.Context, id string) (*models.Company, error)
	GetByEmail(ctx context.Context, email string) (*models.Company, error)

	// Write
	Create(ctx context.Context, user *models.Company) error
}
