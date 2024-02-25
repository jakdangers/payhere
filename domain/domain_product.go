package domain

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product Product) (int, error)
	GetProduct(ctx context.Context, productID int) (*Product, error)
	UpdateProduct(ctx context.Context, product Product) error
	DeleteProduct(ctx context.Context, productID int) error
	ListProducts(ctx context.Context, params ListProductsParams) ([]Product, error)
}

type ProductService interface {
	CreateProduct(ctx context.Context, req CreateProductRequest) error
	GetProduct(ctx context.Context, req GetProductRequest) (GetProductResponse, error)
	PatchProduct(ctx context.Context, req PatchProductRequest) error
	DeleteProduct(ctx context.Context, req DeleteProductRequest) error
	ListProducts(ctx context.Context, req ListProductsRequest) (ListProductsResponse, error)
}

type ProductController interface {
	CreateProduct(c *gin.Context)
	GetProduct(c *gin.Context)
	PatchProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	ListProducts(c *gin.Context)
}

type ProductSizeType string

func (p ProductSizeType) ToPointer() *ProductSizeType {
	return &p
}

const (
	ProductSizeTypeSmall ProductSizeType = "small"
	ProductSizeTypeLarge ProductSizeType = "large"
)

type Product struct {
	Base
	UserID      int
	Initial     string
	Category    string
	Price       float64
	Cost        float64
	Name        string
	Description string
	Barcode     string
	ExpiryDate  time.Time
	Size        ProductSizeType
}
