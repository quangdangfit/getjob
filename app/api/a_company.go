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

// CompanyAPI handle company api
type CompanyAPI struct {
	service interfaces.ICompanyService
}

// NewCompanyAPI return new CompanyAPI
func NewCompanyAPI(service interfaces.ICompanyService) *CompanyAPI {
	return &CompanyAPI{
		service: service,
	}
}

func (a *CompanyAPI) Create(c *gin.Context) wrapper.Response {
	var params *schema.CompanyCreateParams
	if err := c.ShouldBindJSON(&params); err != nil {
		logger.Error(err.Error())
		return wrapper.Res(http.StatusBadRequest, errors.New(errors.InvalidArgument, err.Error()), nil)
	}

	validator := validation.New()
	if err := validator.ValidateStruct(params); err != nil {
		return wrapper.Res(http.StatusBadRequest, errors.New(errors.InvalidArgument, err.Error()), nil)
	}

	company, err := a.service.Create(c, params)
	if err != nil {
		logger.Error(err.Error())
		return wrapper.Res(http.StatusBadRequest, err, nil)
	}

	return wrapper.Res(http.StatusOK, errors.New(errors.OK, "OK"), company.ToSchema())
}
