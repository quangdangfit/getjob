package services

import "github.com/quangdangfit/getjob/src/interfaces"

type CandidateService struct {
	repo interfaces.ICandidateRepository
}

func NewCandidateService(repo interfaces.ICandidateRepository) interfaces.ICandidateService {
	return &CandidateService{
		repo: repo,
	}
}
