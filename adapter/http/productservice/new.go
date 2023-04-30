package productservice

import "github.com/victoraatb/clean-go/core/domain"

type service struct {
	usecase domain.ProductUseCase
}

func New(usecase domain.ProductUseCase) domain.ProductService {
	return &service{
		usecase: usecase,
	}
}
