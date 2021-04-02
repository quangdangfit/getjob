package repositories

import (
	"context"

	"github.com/quangdangfit/getjob/app/interfaces"
	"github.com/quangdangfit/getjob/app/models"
	"github.com/quangdangfit/getjob/pkg/errors"
)

// ExperienceRepository experience repository
type ExperienceRepository struct {
	db interfaces.IDatabase
}

// NewExperienceRepository return new IExperienceRepository interface
func NewExperienceRepository(db interfaces.IDatabase) interfaces.IExperienceRepository {
	return &ExperienceRepository{
		db: db,
	}
}

// GetByID get experience by id
func (r *ExperienceRepository) GetByID(ctx context.Context, id string) (*models.Experience, error) {
	var experience models.Experience
	if err := r.db.GetInstance().Where("id = ?", id).First(&experience).Error; err != nil {
		return nil, errors.New(errors.MysqlRead, err.Error())
	}
	return &experience, nil
}

// ListByUserID get list experiences by userID
func (r *ExperienceRepository) ListByUserID(ctx context.Context, userID string) (*models.Experiences, error) {
	var experience models.Experiences
	if err := r.db.GetInstance().Where("user_id = ?", userID).Find(&experience).Error; err != nil {
		return nil, errors.New(errors.MysqlRead, err.Error())
	}
	return &experience, nil
}

// Create create new experience
func (r *ExperienceRepository) Create(ctx context.Context, experience *models.Experience) error {
	if err := r.db.GetInstance().Create(&experience).Error; err != nil {
		return errors.New(errors.MysqlCreate, err.Error())
	}
	return nil
}
