package productrepository

import (
	"github.com/victoraatb/clean-go/adapter/postgres"
	"github.com/victoraatb/clean-go/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

func New(db postgres.PoolInterface) domain.ProductRepository {
	return &repository{
		db: db,
	}
}
