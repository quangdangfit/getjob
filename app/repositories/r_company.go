package repositories

import (
	"context"

	"github.com/quangdangfit/getjob/app/interfaces"
	"github.com/quangdangfit/getjob/app/models"
	"github.com/quangdangfit/getjob/pkg/errors"
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

func (r *CompanyRepository) Create(ctx context.Context, company *models.Company) error {
	if err := r.db.GetInstance().Create(&company).Error; err != nil {
		return errors.New(errors.MysqlCreate, err.Error())
	}
	return nil
}
