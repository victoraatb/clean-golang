package productusecase

import (
	"github.com/victoraatb/clean-go/core/domain"
	"github.com/victoraatb/clean-go/core/dto"
)

func (usecase usecase) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination, error) {
	products, err := usecase.repository.Fetch(paginationRequest)
	if err != nil {
		return nil, err
	}
	return products, nil
}
