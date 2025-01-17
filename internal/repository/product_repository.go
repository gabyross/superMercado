package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gabyross/superMercado/internal/domain"
)

type ProductRepository interface {
	GetAllProducts() ([]domain.Product, error)
	GetProductByID(id int) (domain.Product, error)
	CreateProduct(product domain.Product) error
	UpdateProduct(id int, product domain.Product) error
	PatchProduct(id int, product domain.Product) error
	DeleteProduct(id int) error
	GetProductsByPriceGreaterThan(price float64) []domain.Product
}

type productRepositoryImpl struct {
	products []domain.Product
}

func NewProductRepository(fileName string) (ProductRepository, error) {
	repo := &productRepositoryImpl{}
	err := repo.loadProducts(fileName)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (pr *productRepositoryImpl) loadProducts(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&pr.products)

	if err != nil {
		return err
	}

	return nil
}

func (pr *productRepositoryImpl) getNextID() int {
	maxID := 0
	for _, product := range pr.products {
		if product.ID > maxID {
			maxID = product.ID
		}
	}

	return maxID + 1
}

// CreateProduct implements ProductRepository.
func (pr *productRepositoryImpl) CreateProduct(p domain.Product) error {
	p.ID = pr.getNextID()
	pr.products = append(pr.products, p)
	return nil
}

// DeleteProduct implements ProductRepository.
func (pr *productRepositoryImpl) DeleteProduct(id int) error {
	for i, product := range pr.products {
		if product.ID == id {
			pr.products = append(pr.products[:i], pr.products[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Product not found")
}

// GetAllProducts implements ProductRepository.
func (pr *productRepositoryImpl) GetAllProducts() ([]domain.Product, error) {
	if pr.products == nil {
		return nil, fmt.Errorf("no products found")
	}
	return pr.products, nil
}

// GetProductByID implements ProductRepository.
func (pr *productRepositoryImpl) GetProductByID(id int) (domain.Product, error) {
	for _, product := range pr.products {
		if product.ID == id {
			return product, nil
		}
	}
	return domain.Product{}, fmt.Errorf("product not found")
}

// PatchProduct implements ProductRepository.
func (pr *productRepositoryImpl) PatchProduct(id int, product domain.Product) error {
	for i, p := range pr.products {
		if p.ID == id {

			if product.Name != "" {
				pr.products[i].Name = product.Name
			}
			if product.Quantity > 0 {
				pr.products[i].Quantity = product.Quantity
			}
			if product.CodeValue != "" {
				pr.products[i].CodeValue = product.CodeValue
			}
			if product.IsPublished != false {
				pr.products[i].IsPublished = product.IsPublished
			}
			if product.Expiration != "" {
				pr.products[i].Expiration = product.Expiration
			}
			if product.Price > 0 {
				pr.products[i].Price = product.Price
			}
			return nil
		}
	}
	return fmt.Errorf("product not found")
}

// UpdateProduct implements ProductRepository.
func (pr *productRepositoryImpl) UpdateProduct(id int, product domain.Product) error {
	for i, p := range pr.products {
		if p.ID == id {
			product.ID = id
			pr.products[i] = product
			return nil
		}
	}
	return fmt.Errorf("product not found")
}

// GetProductsByPriceGreaterThan implements ProductRepository.
func (pr *productRepositoryImpl) GetProductsByPriceGreaterThan(price float64) (filteredProducts []domain.Product) {
	for _, product := range pr.products {
		if product.Price >= price {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}
