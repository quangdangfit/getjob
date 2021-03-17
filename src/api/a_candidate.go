package api

import (
	"github.com/quangdangfit/getjob/src/interfaces"
)

type CandidateAPI struct {
	service interfaces.ICandidateService
}

func NewCandidateAPI(service interfaces.ICandidateService) *CandidateAPI {
	return &CandidateAPI{
		service: service,
	}
}
