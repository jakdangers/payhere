package product

import (
	"context"
	"payhere/domain"
	cerrors "payhere/pkg/cerror"
)

type productService struct {
	userRepository    domain.UserRepository
	productRepository domain.ProductRepository
}

func NewProductService(
	userRepository domain.UserRepository,
	productRepository domain.ProductRepository,
) *productService {
	return &productService{
		userRepository:    userRepository,
		productRepository: productRepository,
	}
}

var _ domain.ProductService = (*productService)(nil)

var hangulCHO = []string{"ㄱ", "ㄲ", "ㄴ", "ㄷ", "ㄸ", "ㄹ", "ㅁ", "ㅂ", "ㅃ", "ㅅ", "ㅆ", "ㅇ", "ㅈ", "ㅉ", "ㅊ", "ㅋ", "ㅌ", "ㅍ", "ㅎ"}

const (
	hangulBASE = rune('가')
	hangulEND  = rune('힣')
)

func (ps productService) CreateProduct(ctx context.Context, req domain.CreateProductRequest) error {
	const op cerrors.Op = "product/service/createProduct"

	_, err := ps.productRepository.CreateProduct(ctx, domain.Product{
		UserID:      req.UserID,
		Initial:     ExtractChosung(req.Name),
		Category:    req.Category,
		Price:       req.Price,
		Cost:        req.Cost,
		Name:        req.Name,
		Description: req.Description,
		Barcode:     req.Barcode,
		ExpiryDate:  req.ExpiryDate,
		Size:        req.Size,
	})
	if err != nil {
		return cerrors.E(op, cerrors.Internal, err, "상품을 생성하는 중에 에러가 발생했습니다.")
	}

	return nil
}

func (ps productService) GetProduct(ctx context.Context, req domain.GetProductRequest) (domain.GetProductResponse, error) {
	const op cerrors.Op = "product/service/getProduct"

	product, err := ps.productRepository.GetProduct(ctx, req.ProductID)
	if err != nil {
		return domain.GetProductResponse{}, cerrors.E(op, cerrors.Internal, err, "상품을 조회하는 중에 에러가 발생했습니다.")
	}
	if product == nil {
		return domain.GetProductResponse{}, cerrors.E(op, cerrors.NotExist, "상품을 찾을 수 없습니다.")
	}
	if product.UserID != req.UserID {
		return domain.GetProductResponse{}, cerrors.E(op, cerrors.Permission, "상품을 조회할 권한이 없습니다.")
	}

	return domain.GetProductResponse{
		Product: domain.ProductDTOFrom(*product),
	}, nil
}

func (ps productService) PatchProduct(ctx context.Context, req domain.PatchProductRequest) error {
	const op cerrors.Op = "product/service/patchProduct"

	product, err := ps.productRepository.GetProduct(ctx, req.ID)
	if err != nil {
		return cerrors.E(op, cerrors.Internal, err, "상품을 조회하는 중에 에러가 발생했습니다.")
	}
	if product == nil {
		return cerrors.E(op, cerrors.NotExist, "상품을 찾을 수 없습니다.")
	}
	if product.UserID != req.UserID {
		return cerrors.E(op, cerrors.Permission, "상품을 수정할 권한이 없습니다.")
	}

	if req.Category != nil {
		product.Category = *req.Category
	}
	if req.Price != nil {
		product.Price = *req.Price
	}
	if req.Cost != nil {
		product.Cost = *req.Cost
	}
	if req.Name != nil {
		product.Name = *req.Name
		product.Initial = ExtractChosung(*req.Name)
	}
	if req.Description != nil {
		product.Description = *req.Description
	}
	if req.Barcode != nil {
		product.Barcode = *req.Barcode
	}
	if req.ExpiryDate != nil {
		product.ExpiryDate = *req.ExpiryDate
	}
	if req.Size != nil {
		product.Size = *req.Size
	}

	if err := ps.productRepository.UpdateProduct(ctx, *product); err != nil {
		return cerrors.E(op, cerrors.Internal, err, "상품을 수정하는 중에 에러가 발생했습니다.")
	}

	return nil
}

func (ps productService) DeleteProduct(ctx context.Context, req domain.DeleteProductRequest) error {
	const op cerrors.Op = "product/service/DeleteProduct"

	product, err := ps.productRepository.GetProduct(ctx, req.ID)
	if err != nil {
		return cerrors.E(op, cerrors.Internal, err, "상품을 조회하는 중에 에러가 발생했습니다.")
	}
	if product == nil {
		return cerrors.E(op, cerrors.NotExist, "상품을 찾을 수 없습니다.")
	}
	if product.UserID != req.UserID {
		return cerrors.E(op, cerrors.Permission, "상품을 삭제할 권한이 없습니다.")
	}

	if err := ps.productRepository.DeleteProduct(ctx, req.ID); err != nil {
		return cerrors.E(op, cerrors.Internal, err, "상품을 삭제하는 중에 에러가 발생했습니다.")
	}

	return nil
}

func (ps productService) ListProducts(ctx context.Context, req domain.ListProductsRequest) (domain.ListProductsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func ExtractChosung(s string) string {
	result := ""
	for _, c := range []rune(s) {
		if c >= hangulBASE && c <= hangulEND {
			offset := c - hangulBASE
			choIndex := int(offset / 588)
			result += hangulCHO[choIndex]
		} else {
			result += string(c)
		}
	}
	return result
}
