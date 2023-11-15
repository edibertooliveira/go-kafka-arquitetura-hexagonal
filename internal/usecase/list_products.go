package usecase

import "go-kafka-arquitetura-hexagonal /internal/domain/entity"

type ListProductsOutputDTO struct {
	Name  string
	Price float64
	ID    string
}

type ListProductsUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductsUseCase(productRepository entity.ProductRepository) *ListProductsUseCase {
	return &ListProductsUseCase{ProductRepository: productRepository}
}

func (u ListProductsUseCase) Execute() ([]*ListProductsOutputDTO, error) {
	products, err := u.ProductRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var productsOutput []*ListProductsOutputDTO

	for _, product := range products {
		productsOutput = append(productsOutput, &ListProductsOutputDTO{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return productsOutput, nil
}
