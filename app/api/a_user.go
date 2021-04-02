package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/logger"
	"github.com/quangdangfit/gocommon/validation"

	"github.com/quangdangfit/getjob/app/interfaces"
	"github.com/quangdangfit/getjob/app/schema"
	"github.com/quangdangfit/getjob/pkg/errors"
	"github.com/quangdangfit/getjob/pkg/http/wrapper"
)

// UserAPI handle user api
type UserAPI struct {
	service interfaces.IUserService
}

// NewCandidateAPI return new UserAPI
func NewCandidateAPI(service interfaces.IUserService) *UserAPI {
	return &UserAPI{
		service: service,
	}
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
	rs, err := a.service.Register(ctx, params)
	if err != nil {
		logger.Error(err.Error())
		return wrapper.Res(http.StatusBadRequest, err, nil)
	}

	return wrapper.Res(http.StatusOK, errors.New(errors.ECOK, "OK"), rs.ToSchema())
}
