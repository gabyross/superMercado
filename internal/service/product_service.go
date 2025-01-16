package service

import (
	"fmt"

	"github.com/gabyross/superMercado/internal/domain"
	"github.com/gabyross/superMercado/internal/repository"
)

// ProductService is the interface that provides product methods
type ProductService interface {
	GetAllProducts() ([]domain.Product, error)
	GetProductByID(int) (domain.Product, error)
	SearchProduct(float64) ([]domain.Product, error)
	CreateProduct(domain.Product) error
	UpdateProduct(int, domain.Product) error
	PatchProduct(int, domain.Product) error
	DeleteProduct(int) error
}

// productService is a concrete implementation of ProductService
type productService struct {
	repository repository.ProductRepository
}

// NewProductService creates a new ProductService with the necessary dependencies
func NewProductService(repository repository.ProductRepository) ProductService {
	return &productService{
		repository: repository,
	}
}

func (ps *productService) GetAllProducts() ([]domain.Product, error) {
	product, err := ps.repository.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("error getting products: %w", err)
	}
	return product, nil
}

func (ps *productService) GetProductByID(id int) (domain.Product, error) {
	product, err := ps.repository.GetProductByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (ps *productService) SearchProduct(priceGt float64) ([]domain.Product, error) {
	var filteredProducts []domain.Product
	products, err := ps.repository.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("error getting all  products: %w", err)
	}
	for _, product := range products {
		if product.Price > priceGt {
			filteredProducts = append(filteredProducts, product)
		}
	}

	return filteredProducts, nil
}

func (ps *productService) CreateProduct(product domain.Product) error {
	return ps.repository.CreateProduct(product)
}

// UpdateProduct implements ProductService.
func (p *productService) UpdateProduct(id int, product domain.Product) error {
	err := p.repository.UpdateProduct(id, product)
	if err != nil {
		return err
	}
	return nil
}

// PatchProduct implements ProductService.
func (p *productService) PatchProduct(id int, product domain.Product) error {
	err := p.repository.PatchProduct(id, product)
	if err != nil {
		return err
	}
	return nil
}

// DeleteProduct implements ProductService.
func (p *productService) DeleteProduct(id int) error {
	return p.repository.DeleteProduct(id)
}
