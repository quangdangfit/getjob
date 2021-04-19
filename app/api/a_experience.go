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

// ExperienceAPI handle experience api
type ExperienceAPI struct {
	service interfaces.IExperienceService
}

// NewExperienceAPI return new ExperienceAPI
func NewExperienceAPI(service interfaces.IExperienceService) *ExperienceAPI {
	return &ExperienceAPI{
		service: service,
	}
}

func (a *ExperienceAPI) Create(c *gin.Context) wrapper.Response {
	var params *schema.ExperienceCreateParams
	if err := c.ShouldBindJSON(&params); err != nil {
		logger.Error(err.Error())
		return wrapper.Res(http.StatusBadRequest, errors.New(errors.InvalidArgument, err.Error()), nil)
	}

	validator := validation.New()
	if err := validator.ValidateStruct(params); err != nil {
		logger.Error(err)
		return wrapper.Res(http.StatusBadRequest, errors.New(errors.InvalidArgument, err.Error()), nil)
	}

	var userID = c.GetString(UserIDContextKey)
	if userID == "" {
		return wrapper.Res(http.StatusUnauthorized, errors.New(errors.PermissionDenied, "cannot get user id from context"), nil)
	}

	experience, err := a.service.Create(c, userID, params)
	if err != nil {
		logger.Error(err.Error())
		return wrapper.Res(http.StatusBadRequest, err, nil)
	}

	return wrapper.Res(http.StatusOK, errors.New(errors.OK, "OK"), experience.ToSchema())
}
