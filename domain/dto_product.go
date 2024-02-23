package domain

import (
	cerrors "payhere/pkg/cerror"
	"time"
)

type ProductDTO struct {
	BaseDTO
	UserID      int             `json:"userID"`
	Initial     string          `json:"initial"`
	Category    string          `json:"category"`
	Price       float64         `json:"price"`
	Cost        float64         `json:"cost"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Barcode     string          `json:"barcode"`
	ExpiryDate  time.Time       `json:"expiryDate"`
	Size        ProductSizeType `json:"size"`
}

func ProductDTOFrom(domain Product) ProductDTO {
	dto := ProductDTO{
		BaseDTO: BaseDTO{
			ID:          domain.ID,
			CreatedDate: domain.CreatedDate,
			UpdatedDate: domain.UpdatedDate,
		},
		UserID:      domain.UserID,
		Initial:     domain.Initial,
		Category:    domain.Category,
		Price:       domain.Price,
		Cost:        domain.Cost,
		Name:        domain.Name,
		Description: domain.Description,
		Barcode:     domain.Barcode,
		ExpiryDate:  domain.ExpiryDate,
		Size:        domain.Size,
	}

	if domain.DeletedDate.Valid {
		dto.DeletedDate = &domain.DeletedDate.Time
	}

	return dto
}

type CreateProductRequest struct {
	UserID      int             `json:"userID"`
	Category    string          `json:"category"`
	Price       float64         `json:"price"`
	Cost        float64         `json:"cost"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Barcode     string          `json:"barcode"`
	ExpiryDate  time.Time       `json:"expiryDate"`
	Size        ProductSizeType `json:"size"`
}

func (req CreateProductRequest) Validate() error {
	var op cerrors.Op = "domain/CreateProductRequest.Validate"

	if req.Category == "" {
		return cerrors.E(op, cerrors.Invalid, "카테고리를 확인해주세요.")
	}

	if req.Price < 0 {
		return cerrors.E(op, cerrors.Invalid, "가격을 확인해주세요.")
	}

	if req.Cost < 0 {
		return cerrors.E(op, cerrors.Invalid, "원가를 확인해주세요.")
	}

	if req.Name == "" {
		return cerrors.E(op, cerrors.Invalid, "상품명을 확인해주세요.")
	}

	if req.Description == "" {
		return cerrors.E(op, cerrors.Invalid, "상품 설명을 확인해주세요.")
	}

	if req.Barcode == "" {
		return cerrors.E(op, cerrors.Invalid, "바코드를 확인해주세요.")
	}

	if req.ExpiryDate.IsZero() {
		return cerrors.E(op, cerrors.Invalid, "유통기한을 확인해주세요.")
	}

	if req.Size != ProductSizeTypeSmall && req.Size != ProductSizeTypeLarge {
		return cerrors.E(op, cerrors.Invalid, "상품 사이즈를 확인해주세요.")
	}

	return nil
}

type GetProductRequest struct {
	UserID    int `json:"userID"`
	ProductID int `json:"productID" uri:"productID"`
}

func (req GetProductRequest) Validate() error {
	var op cerrors.Op = "domain/GetProductRequest.Validate"

	if req.ProductID <= 0 {
		return cerrors.E(op, cerrors.Invalid, "상품 ID를 입력해주세요.")
	}

	return nil
}

type GetProductResponse struct {
	Product ProductDTO `json:"product"`
}

type PatchProductRequest struct {
	UserID      int
	ID          int              `json:"id"`
	Category    *string          `json:"category"`
	Price       *float64         `json:"price"`
	Cost        *float64         `json:"cost"`
	Name        *string          `json:"name"`
	Description *string          `json:"description"`
	Barcode     *string          `json:"barcode"`
	ExpiryDate  *time.Time       `json:"expiryDate"`
	Size        *ProductSizeType `json:"size"`
}

func (req PatchProductRequest) Validate() error {
	const op cerrors.Op = "domain/PatchProductRequest.Validate"

	if req.Category != nil && *req.Category == "" {
		return cerrors.E(op, cerrors.Invalid, "카테고리를 확인해주세요.")
	}

	if req.Price != nil && *req.Price < 0 {
		return cerrors.E(op, cerrors.Invalid, "가격을 확인해주세요.")
	}

	if req.Cost != nil && *req.Cost < 0 {
		return cerrors.E(op, cerrors.Invalid, "원가를 확인해주세요.")
	}

	if req.Name != nil && *req.Name == "" {
		return cerrors.E(op, cerrors.Invalid, "상품명을 확인해주세요.")
	}

	if req.Description != nil && *req.Description == "" {
		return cerrors.E(op, cerrors.Invalid, "상품 설명을 확인해주세요.")
	}

	if req.Barcode != nil && *req.Barcode == "" {
		return cerrors.E(op, cerrors.Invalid, "바코드를 확인해주세요.")
	}

	if req.ExpiryDate != nil && req.ExpiryDate.IsZero() {
		return cerrors.E(op, cerrors.Invalid, "유통기한을 확인해주세요.")
	}

	if req.Size != nil && *req.Size != ProductSizeTypeSmall && *req.Size != ProductSizeTypeLarge {
		return cerrors.E(op, cerrors.Invalid, "상품 사이즈를 확인해주세요.")
	}

	return nil
}

type DeleteProductRequest struct {
	UserID int
	ID     int `uri:"productID"`
}

func (req DeleteProductRequest) Validate() error {
	const op cerrors.Op = "domain/DeleteProductRequest.Validate"

	if req.ID <= 0 {
		return cerrors.E(op, cerrors.Invalid, "상품 ID를 확인해주세요.")
	}

	return nil
}

type ListProductsRequest struct {
}

type ListProductsResponse struct {
}
