package productusecase_test

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/victoraatb/clean-go/core/domain"
	"github.com/victoraatb/clean-go/core/domain/mocks"
	"github.com/victoraatb/clean-go/core/domain/usecase/productusecase"
	"github.com/victoraatb/clean-go/core/dto"
)

func TestCreate(t *testing.T) {
	fakeRequestProduct := dto.CreateProductRequest{}
	fakeDBProduct := domain.Product{}
	faker.FakeData(&fakeRequestProduct)
	faker.FakeData(&fakeDBProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().Create(&fakeRequestProduct).Return(&fakeDBProduct, nil)

	sut := productusecase.New(mockProductRepository)
	product, err := sut.Create(&fakeRequestProduct)

	require.Nil(t, err)
	require.NotEmpty(t, product.ID)
	require.Equal(t, product.Name, fakeDBProduct.Name)
	require.Equal(t, product.Price, fakeDBProduct.Price)
	require.Equal(t, product.Description, fakeDBProduct.Description)

}

func TestCreate_Error(t *testing.T) {
	fakeRequestProduct := dto.CreateProductRequest{}
	faker.FakeData(&fakeRequestProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().Create(&fakeRequestProduct).Return(nil, fmt.Errorf("ANY ERRO"))

	sut := productusecase.New(mockProductRepository)
	product, err := sut.Create(&fakeRequestProduct)

	require.NotNil(t, err)
	require.Nil(t, product)
}
