package di

import (
	"github.com/victoraatb/clean-go/adapter/http/productservice"
	"github.com/victoraatb/clean-go/adapter/postgres"
	"github.com/victoraatb/clean-go/adapter/postgres/productrepository"
	"github.com/victoraatb/clean-go/core/domain"
	"github.com/victoraatb/clean-go/core/domain/usecase/productusecase"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigProductDI(conn postgres.PoolInterface) domain.ProductService {
	productRepository := productrepository.New(conn)
	productUseCase := productusecase.New(productRepository)
	productService := productservice.New(productUseCase)

	return productService
}
