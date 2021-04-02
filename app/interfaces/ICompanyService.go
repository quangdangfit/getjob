package interfaces

import (
	"context"

	"github.com/quangdangfit/getjob/app/models"
)

type ICompanyService interface {
	GetByID(ctx context.Context, id string) (*models.Company, error)
}
