package interfaces

import (
	"context"

	"github.com/quangdangfit/getjob/app/models"
	"github.com/quangdangfit/getjob/app/schema"
)

type IUserService interface {
	Register(ctx context.Context, params *schema.UserRegisterParams) (*models.User, error)
}
