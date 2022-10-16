package service

import(
	"errors"
	"go-unit-test/entity"
	"go-unit=test/repository"
)

type ProductService struct {
	Repository repository.ProductService
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("prduct not found")
	}
	return product, nil
}
