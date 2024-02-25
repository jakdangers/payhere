package domain

import (
	"fmt"
	cerrors "payhere/pkg/cerrors"
	"time"
)

type ProductDTO struct {
	BaseDTO
	UserID      int             `json:"userID" validate:"required" example:"1"`
	Initial     string          `json:"initial" validate:"required" example:"ㅅㅋㄹ ㄹㄸ"`
	Category    string          `json:"category" validate:"required" example:"payhere"`
	Price       float64         `json:"price" validate:"required" example:"1000"`
	Cost        float64         `json:"cost" validate:"required" example:"500"`
	Name        string          `json:"name" validate:"required" example:"슈크림 라떼"`
	Description string          `json:"description" validate:"required" example:"슈크림 라떼 팔아요"`
	Barcode     string          `json:"barcode" validate:"required" example:"25611234"`
	ExpiryDate  time.Time       `json:"expiryDate" validate:"required" example:"2024-02-28T15:04:05Z"`
	Size        ProductSizeType `json:"size" validate:"required" enum:"small,large" example:"large"`
}

func ProductDTOFrom(domain Product) ProductDTO {
	dto := ProductDTO{
		BaseDTO: BaseDTO{
			ID:         domain.ID,
			CreateDate: domain.CreateDate,
			UpdateDate: domain.UpdateDate,
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

	return dto
}

type CreateProductRequest struct {
	UserID      int             `swaggerignore:"true"`
	Category    string          `json:"category" validate:"required" example:"payhere"`
	Price       float64         `json:"price" validate:"required" example:"1000"`
	Cost        float64         `json:"cost" validate:"required" example:"500"`
	Name        string          `json:"name" validate:"required" example:"슈크림 라떼"`
	Description string          `json:"description" validate:"required" example:"슈크림 라떼 팔아요"`
	Barcode     string          `json:"barcode" validate:"required" example:"25611234"`
	ExpiryDate  time.Time       `json:"expiryDate" validate:"required" example:"2024-02-28T15:04:05Z"`
	Size        ProductSizeType `json:"size" validate:"required" enum:"small,large" example:"large"`
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
	UserID      int              `swaggerignore:"true"`
	ID          int              `json:"id" validate:"required" example:"1"`
	Category    *string          `json:"category" validate:"omitempty" example:"payhere"`
	Price       *float64         `json:"price" validate:"omitempty" example:"1000"`
	Cost        *float64         `json:"cost" validate:"omitempty" example:"500"`
	Name        *string          `json:"name" validate:"omitempty" example:"슈크림 라떼"`
	Description *string          `json:"description" validate:"omitempty" example:"슈크림 라떼 팔아요"`
	Barcode     *string          `json:"barcode" validate:"omitempty" example:"25611234"`
	ExpiryDate  *time.Time       `json:"expiryDate" validate:"omitempty" example:"2024-02-28T15:04:05Z"`
	Size        *ProductSizeType `json:"size" validate:"omitempty" enum:"small,large" example:"large"`
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

type ListProductsParams struct {
	UserID  int
	Cursor  *int
	Name    *string
	Initial *string
}

func (lp ListProductsParams) LikeName() string {
	if lp.Name == nil {
		return ""
	}

	return fmt.Sprintf("AND name LIKE '%%%s%%'", *lp.Name)
}

func (lp ListProductsParams) LikeInitial() string {
	if lp.Initial == nil {
		return ""
	}

	return fmt.Sprintf("AND initial LIKE '%%%s%%'", *lp.Initial)
}

func (lp ListProductsParams) AfterCursor() string {
	if lp.Cursor == nil {
		return ""
	}

	return fmt.Sprintf("AND id > %d", *lp.Cursor)
}

type ListProductsRequest struct {
	UserID int
	Cursor *int    `form:"cursor"`
	Search *string `form:"search"`
}

type ListProductsResponse struct {
	Products []ProductDTO `json:"products"`
	Cursor   *int         `json:"cursor"`
}
