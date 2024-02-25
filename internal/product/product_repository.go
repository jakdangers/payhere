package product

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"payhere/domain"
	cerrors "payhere/pkg/cerrors"
	"time"
)

type productRepository struct {
	sqlDB *sql.DB
}

func NewProductRepository(sqlDB *sql.DB) *productRepository {
	return &productRepository{
		sqlDB: sqlDB,
	}
}

var _ domain.ProductRepository = (*productRepository)(nil)

func (pr productRepository) CreateProduct(ctx context.Context, product domain.Product) (int, error) {
	const op cerrors.Op = "product/productRepository/CreateProduct"

	result, err := pr.sqlDB.ExecContext(
		ctx,
		createProductQuery,
		product.UserID,
		product.Initial,
		product.Category,
		product.Price,
		product.Cost,
		product.Name,
		product.Description,
		product.Barcode,
		product.ExpiryDate,
		product.Size,
	)
	if err != nil {
		return 0, cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	productID, err := result.LastInsertId()
	if err != nil {
		return 0, cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	return int(productID), nil
}

func (pr productRepository) GetProduct(ctx context.Context, productID int) (*domain.Product, error) {
	const op cerrors.Op = "product/productRepository/GetProduct"
	var product domain.Product

	err := pr.sqlDB.QueryRowContext(ctx, findProductByIDQuery, productID).
		Scan(
			&product.ID,
			&product.CreateDate,
			&product.UpdateDate,
			&product.DeleteDate,
			&product.UserID,
			&product.Initial,
			&product.Category,
			&product.Price,
			&product.Cost,
			&product.Name,
			&product.Description,
			&product.Barcode,
			&product.ExpiryDate,
			&product.Size,
		)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	return &product, nil
}

func (pr productRepository) UpdateProduct(ctx context.Context, product domain.Product) error {
	const op cerrors.Op = "product/productRepository/UpdateProduct"

	_, err := pr.sqlDB.ExecContext(
		ctx,
		updateProductQuery,
		product.Initial,
		product.Category,
		product.Price,
		product.Cost,
		product.Name,
		product.Description,
		product.Barcode,
		product.ExpiryDate,
		product.Size,
		product.ID,
	)
	if err != nil {
		return cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	return nil
}

func (pr productRepository) DeleteProduct(ctx context.Context, productID int) error {
	const op cerrors.Op = "product/productRepository/DeleteProduct"

	_, err := pr.sqlDB.ExecContext(ctx, deleteProductQuery, time.Now().UTC(), productID)
	if err != nil {
		return cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	return nil
}

func (pr productRepository) ListProducts(ctx context.Context, params domain.ListProductsParams) ([]domain.Product, error) {
	const op cerrors.Op = "product/productRepository/ListProducts"

	var products []domain.Product

	query := fmt.Sprintf(listProductsQuery,
		params.LikeInitial(),
		params.LikeName(),
		params.AfterCursor(),
	)

	rows, err := pr.sqlDB.QueryContext(ctx, query, params.UserID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.Product
		err := rows.Scan(
			&product.ID,
			&product.CreateDate,
			&product.UpdateDate,
			&product.DeleteDate,
			&product.UserID,
			&product.Initial,
			&product.Category,
			&product.Price,
			&product.Cost,
			&product.Name,
			&product.Description,
			&product.Barcode,
			&product.ExpiryDate,
			&product.Size,
		)
		if err != nil {
			return nil, cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
		}
		products = append(products, product)
	}

	return products, nil
}
