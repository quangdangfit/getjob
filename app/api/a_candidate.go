package api

import (
	"github.com/quangdangfit/getjob/app/interfaces"
)

type CandidateAPI struct {
	service interfaces.IUserService
}

func NewCandidateAPI(service interfaces.IUserService) *CandidateAPI {
	return &CandidateAPI{
		service: service,
	}
}
