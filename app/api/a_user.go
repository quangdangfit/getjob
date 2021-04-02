package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/jwt"
	"github.com/quangdangfit/gocommon/logger"
	"github.com/quangdangfit/gocommon/validation"

	"github.com/quangdangfit/getjob/app/interfaces"
	"github.com/quangdangfit/getjob/app/schema"
	"github.com/quangdangfit/getjob/pkg/errors"
	"github.com/quangdangfit/getjob/pkg/http/wrapper"
)

const (
	UserIDContextKey = "id"
)

// UserAPI handle user api
type UserAPI struct {
	auth    jwt.IJWTAuth
	service interfaces.IUserService
}

// NewUserAPI return new UserAPI
func NewUserAPI(auth jwt.IJWTAuth, service interfaces.IUserService) *UserAPI {
	return &UserAPI{
		auth:    auth,
		service: service,
	}
}

// Login handle api user login
func (a *UserAPI) Login(c *gin.Context) wrapper.Response {
	var params *schema.UserLoginParams
	if err := c.ShouldBindJSON(&params); err != nil {
		logger.Error(err.Error())
		return wrapper.Res(http.StatusBadRequest, errors.New(errors.InvalidArgument, err.Error()), nil)
	}

	validator := validation.New()
	if err := validator.ValidateStruct(params); err != nil {
		return wrapper.Res(http.StatusBadRequest, errors.New(errors.InvalidArgument, err.Error()), nil)
	}

	user, err := a.service.Login(c, params)
	if err != nil {
		logger.Error(err.Error())
		return wrapper.Res(http.StatusBadRequest, err, nil)
	}

	token, expiredTime := a.auth.GenerateToken(user.ID)
	rs := schema.UserLoginRes{
		Token:       token,
		ExpiredTime: expiredTime.Format(time.RFC3339Nano),
	}
	return wrapper.Res(http.StatusOK, errors.New(errors.OK, "OK"), rs)
}

// Register handle api register new user
func (a *UserAPI) Register(c *gin.Context) wrapper.Response {
	var params *schema.UserRegisterParams
	if err := c.ShouldBindJSON(&params); err != nil {
		logger.Error(err.Error())
		return wrapper.Res(http.StatusBadRequest, errors.New(errors.InvalidArgument, err.Error()), nil)
	}

	validator := validation.New()
	if err := validator.ValidateStruct(params); err != nil {
		return wrapper.Res(http.StatusBadRequest, errors.New(errors.InvalidArgument, err.Error()), nil)
	}

	user, err := a.service.Register(c, params)
	if err != nil {
		logger.Error(err.Error())
		return wrapper.Res(http.StatusBadRequest, err, nil)
	}

	return wrapper.Res(http.StatusOK, errors.New(errors.OK, "OK"), user.ToSchema())
}

// GetProfile handle api get personal profile
func (a *UserAPI) GetProfile(c *gin.Context) wrapper.Response {
	var id = c.GetString(UserIDContextKey)
	if id == "" {
		return wrapper.Res(http.StatusBadRequest, errors.New(errors.InvalidArgument, "id is required"), nil)
	}

	user, err := a.service.GetByID(c, id)
	if err != nil {
		logger.Error(err.Error())
		return wrapper.Res(http.StatusBadRequest, err, nil)
	}

	return wrapper.Res(http.StatusOK, errors.New(errors.OK, "OK"), user.ToSchema())
}
