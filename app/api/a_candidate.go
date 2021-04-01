package api

import (
	"github.com/quangdangfit/getjob/app/interfaces"
)

type CandidateAPI struct {
	service interfaces.ICandidateService
}

func NewCandidateAPI(service interfaces.ICandidateService) *CandidateAPI {
	return &CandidateAPI{
		service: service,
	}
}
