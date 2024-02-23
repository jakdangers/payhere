package product

import (
	"context"
	"database/sql"
	"errors"
	"payhere/domain"
	cerrors "payhere/pkg/cerror"
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
	const op cerrors.Op = "product/productRepository/createProduct"

	createProductQuery := "INSERT INTO `products` (user_id, initial, category, price, cost, name, description, barcode, expiry_date, size) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
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
			&product.CreatedDate,
			&product.UpdatedDate,
			&product.DeletedDate,
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
	const op cerrors.Op = "product/productRepository/updateProduct"

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
	const op cerrors.Op = "product/productRepository/deleteProduct"

	_, err := pr.sqlDB.ExecContext(ctx, deleteProductQuery, time.Now().UTC(), productID)
	if err != nil {
		return cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	return nil
}

func (pr productRepository) ListProducts(ctx context.Context) ([]domain.Product, error) {
	//TODO implement me
	panic("implement me")
}
