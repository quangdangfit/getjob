package repositories

import "github.com/quangdangfit/getjob/src/interfaces"

type CandidateRepository struct {
	db interfaces.IDatabase
}

func NewCandidateRepository(db interfaces.IDatabase) interfaces.ICandidateRepository {
	return &CandidateRepository{
		db: db,
	}
}
