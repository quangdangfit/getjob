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
		return wrapper.Res(http.StatusBadRequest, errors.New(errors.ECInvalidArgument, err.Error()), nil)
	}

	validator := validation.New()
	if err := validator.ValidateStruct(params); err != nil {
		return wrapper.Res(http.StatusBadRequest, errors.New(errors.ECInvalidArgument, err.Error()), nil)
	}

	ctx := c.Request.Context()
	user, err := a.service.Login(ctx, params)
	if err != nil {
		logger.Error(err.Error())
		return wrapper.Res(http.StatusBadRequest, err, nil)
	}

	token, expiredTime := a.auth.GenerateToken(user.ID)
	rs := schema.UserLoginRes{
		Token:       token,
		ExpiredTime: expiredTime.Format(time.RFC3339Nano),
	}
	return wrapper.Res(http.StatusOK, errors.New(errors.ECOK, "OK"), rs)
}

// Register handle api register new user
func (a *UserAPI) Register(c *gin.Context) wrapper.Response {
	var params *schema.UserRegisterParams
	if err := c.ShouldBindJSON(&params); err != nil {
		logger.Error(err.Error())
		return wrapper.Res(http.StatusBadRequest, errors.New(errors.ECInvalidArgument, err.Error()), nil)
	}

	validator := validation.New()
	if err := validator.ValidateStruct(params); err != nil {
		return wrapper.Res(http.StatusBadRequest, errors.New(errors.ECInvalidArgument, err.Error()), nil)
	}

	ctx := c.Request.Context()
	user, err := a.service.Register(ctx, params)
	if err != nil {
		logger.Error(err.Error())
		return wrapper.Res(http.StatusBadRequest, err, nil)
	}

	return wrapper.Res(http.StatusOK, errors.New(errors.ECOK, "OK"), user.ToSchema())
}
