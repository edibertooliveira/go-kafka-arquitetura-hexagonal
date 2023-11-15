package usecase

import "go-kafka-arquitetura-hexagonal /internal/domain/entity"

type CreateProductInputDTO struct {
	Name  string
	Price float64
}

type CreateProductOutputDTO struct {
	Name  string
	Price float64
	ID    string
}

type CreateProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewCreateProductUseCase(productRepository entity.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{ProductRepository: productRepository}
}

func (u *CreateProductUseCase) Execute(input CreateProductInputDTO) (*CreateProductOutputDTO, error) {
	product := entity.NewProduct(input.Name, input.Price)

	err := u.ProductRepository.Create(product)

	if err != nil {
		return nil, err
	}

	return &CreateProductOutputDTO{
		Name:  product.Name,
		Price: product.Price,
		ID:    product.ID,
	}, nil
}
