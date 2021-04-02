package repositories

import (
	"context"

	"github.com/quangdangfit/getjob/app/interfaces"
	"github.com/quangdangfit/getjob/app/models"
	"github.com/quangdangfit/getjob/pkg/errors"
)

type UserRepository struct {
	db interfaces.IDatabase
}

func NewUserRepository(db interfaces.IDatabase) interfaces.IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	return nil, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.db.GetInstance().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New(errors.ECMysqlRead, err.Error())
	}
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	if err := r.db.GetInstance().Create(&user).Error; err != nil {
		return errors.New(errors.ECMysqlCreate, err.Error())
	}
	return nil
}
