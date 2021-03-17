package api

import (
	"github.com/quangdangfit/getjob/src/interfaces"
	"github.com/quangdangfit/getjob/src/services"
)

type CandidateAPI struct {
	service interfaces.ICandidateService
}

func NewCandidateAPI(service services.CandidateService) *CandidateAPI {
	return &CandidateAPI{
		service: service,
	}
}
