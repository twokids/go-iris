package services

import (
	"go-iris/datamodels"
	"go-iris/repositories"
)

type IProductService interface {
	GetProductByID(int64) (*datamodels.Product, error)
	GetAllProduct() ([]*datamodels.Product, error)
	DeleteProductByID(int64) bool
	InsertProduct(product *datamodels.Product) (int64, error)
	UpdateProduct(product *datamodels.Product) error
}

type ProductService struct {
	productRepository repositories.IProduct
}

func NewProductService(repository repositories.IProduct) IProductService {
	return &ProductService{repository}
}

func (c *ProductService) GetProductByID(productID int64) (*datamodels.Product, error) {
	return c.productRepository.SelectByKey(productID)
}

func (c *ProductService) GetAllProduct() ([]*datamodels.Product, error) {
	return c.productRepository.SelectAll()
}

func (c *ProductService) DeleteProductByID(productID int64) bool {
	return c.productRepository.Delete(productID)
}

func (c *ProductService) InsertProduct(product *datamodels.Product) (int64, error) {
	return c.productRepository.Insert(product), nil
}

func (c *ProductService) UpdateProduct(product *datamodels.Product) error {
	return c.productRepository.Update(product)
}
