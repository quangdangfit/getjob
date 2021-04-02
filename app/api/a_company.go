package api

import (
	"github.com/quangdangfit/getjob/app/interfaces"
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
