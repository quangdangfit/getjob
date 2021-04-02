package repositories

import (
	"context"

	"github.com/quangdangfit/getjob/app/interfaces"
	"github.com/quangdangfit/getjob/app/models"
)

type CompanyRepository struct {
	db interfaces.IDatabase
}

func NewCompanyRepository(db interfaces.IDatabase) interfaces.ICompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

func (r *CompanyRepository) GetByID(ctx context.Context, id string) (*models.Company, error) {
	return nil, nil
}

func (r *CompanyRepository) GetByEmail(ctx context.Context, email string) (*models.Company, error) {
	return nil, nil
}

func (r *CompanyRepository) Create(ctx context.Context, user *models.Company) error {
	return nil
}
