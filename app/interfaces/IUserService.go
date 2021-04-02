package interfaces

import (
	"context"

	"github.com/quangdangfit/getjob/app/models"
	"github.com/quangdangfit/getjob/app/schema"
)

// IUserService interface UserService
type IUserService interface {
	Login(ctx context.Context, params *schema.UserLoginParams) (*models.User, error)
	Register(ctx context.Context, params *schema.UserRegisterParams) (*models.User, error)
}
