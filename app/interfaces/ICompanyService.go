package interfaces

import (
	"context"

	"github.com/quangdangfit/getjob/app/models"
	"github.com/quangdangfit/getjob/app/schema"
)

type ICompanyService interface {
	GetByID(ctx context.Context, id string) (*models.Company, error)
	Create(ctx context.Context, params *schema.CompanyCreateParams) (*models.Company, error)
}
