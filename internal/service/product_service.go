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
	SearchProductByPriceGreaterThan(float64) ([]domain.Product)
	CreateProduct(domain.Product) error
	UpdateProduct(int, domain.Product) error
	PatchProduct(int, domain.Product) error
	DeleteProduct(int) error
}

// productServiceImpl is a concrete implementation of ProductService
type productServiceImpl struct {
	repository repository.ProductRepository
}

// NewProductService creates a new ProductService with the necessary dependencies
func NewProductService(repository repository.ProductRepository) ProductService {
	return &productServiceImpl{
		repository: repository,
	}
}

func (ps *productServiceImpl) GetAllProducts() ([]domain.Product, error) {
	product, err := ps.repository.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("error getting products: %w", err)
	}
	return product, nil
}

func (ps *productServiceImpl) GetProductByID(id int) (domain.Product, error) {
	product, err := ps.repository.GetProductByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (ps *productServiceImpl) SearchProductByPriceGreaterThan(priceGt float64) ([]domain.Product) {
	return ps.repository.GetProductsByPriceGreaterThan(priceGt)
}

func (ps *productServiceImpl) CreateProduct(product domain.Product) error {
	return ps.repository.CreateProduct(product)
}

// UpdateProduct implements ProductService.
func (p *productServiceImpl) UpdateProduct(id int, product domain.Product) error {
	err := p.repository.UpdateProduct(id, product)
	if err != nil {
		return err
	}
	return nil
}

// PatchProduct implements ProductService.
func (p *productServiceImpl) PatchProduct(id int, product domain.Product) error {
	err := p.repository.PatchProduct(id, product)
	if err != nil {
		return err
	}
	return nil
}

// DeleteProduct implements ProductService.
func (p *productServiceImpl) DeleteProduct(id int) error {
	return p.repository.DeleteProduct(id)
}
