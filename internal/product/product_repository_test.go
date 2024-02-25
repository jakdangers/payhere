package product

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"payhere/domain"
	"testing"
	"time"
)

type productRepositoryTestSuite struct {
	sqlDB             *sql.DB
	sqlMock           sqlmock.Sqlmock
	productRepository domain.ProductRepository
}

func setupUserRepositoryTestSuite() productRepositoryTestSuite {
	var us productRepositoryTestSuite

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	us.sqlDB = mockDB
	us.sqlMock = mock

	if err != nil {
		panic(err)
	}

	us.productRepository = NewProductRepository(mockDB)

	return us
}

func Test_productRepository_CreateProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		product domain.Product
	}

	expiryDate := time.Date(2025, time.June, 10, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		args    args
		mock    func(ts productRepositoryTestSuite)
		want    int
		wantErr bool
	}{
		{
			name: "PASS - 상품 생성",
			args: args{
				ctx: context.Background(),
				product: domain.Product{
					UserID:      1,
					Initial:     "ㅅㅋㄹ ㄹㄸ",
					Category:    "payhere",
					Price:       1000,
					Cost:        500,
					Name:        "슈크림 라떼",
					Description: "description",
					Barcode:     "barcode",
					ExpiryDate:  expiryDate,
					Size:        domain.ProductSizeTypeSmall,
				},
			},
			mock: func(ts productRepositoryTestSuite) {
				ts.sqlMock.ExpectExec("INSERT INTO products").
					WithArgs(
						1,
						"ㅅㅋㄹ ㄹㄸ",
						"payhere",
						float64(1000),
						float64(500),
						"슈크림 라떼",
						"description",
						"barcode",
						expiryDate,
						domain.ProductSizeTypeSmall,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:    1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupUserRepositoryTestSuite()
			tt.mock(ts)

			// when
			got, err := ts.productRepository.CreateProduct(tt.args.ctx, tt.args.product)

			// then
			if ts.sqlMock.ExpectationsWereMet() != nil {
				t.Errorf("there were unfulfilled expectations: %s", ts.sqlMock.ExpectationsWereMet())
			}
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_productRepository_GetProduct(t *testing.T) {
	type args struct {
		ctx       context.Context
		productID int
	}

	createDate := time.Now()
	updateDate := time.Now()
	expiryDate := time.Date(2025, time.June, 10, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		args    args
		mock    func(ts productRepositoryTestSuite)
		want    *domain.Product
		wantErr bool
	}{
		{
			name: "PASS - 상품 조회 성공",
			args: args{
				ctx:       context.Background(),
				productID: 100,
			},
			mock: func(ts productRepositoryTestSuite) {
				query := `SELECT id, create_date, update_date, delete_date, user_id, initial, category, price, cost, name, description, barcode, expiry_date, size FROM products`
				columns := []string{"id", "create_date", "update_date", "delete_date", "user_id", "initial", "category", "price", "cost", "name", "description", "barcode", "expiry_date", "size"}
				rows := sqlmock.NewRows(columns).AddRow(100, createDate, updateDate, nil, 1, "ㅅㅋㄹ ㄹㄸ", "payhere", 1000, 500, "슈크림 라떼", "description", "barcode", expiryDate, domain.ProductSizeTypeSmall)
				ts.sqlMock.ExpectQuery(query).WithArgs(100).WillReturnRows(rows)
			},
			want: &domain.Product{
				Base: domain.Base{
					ID:         100,
					CreateDate: createDate,
					UpdateDate: updateDate,
					DeleteDate: sql.NullTime{
						Time:  time.Time{},
						Valid: false,
					},
				},
				UserID:      1,
				Initial:     "ㅅㅋㄹ ㄹㄸ",
				Category:    "payhere",
				Price:       1000,
				Cost:        500,
				Name:        "슈크림 라떼",
				Description: "description",
				Barcode:     "barcode",
				ExpiryDate:  expiryDate,
				Size:        domain.ProductSizeTypeSmall,
			},
			wantErr: false,
		},
		{
			name: "FAIL - 존재하지 않는 productID로 조회",
			args: args{
				ctx:       context.Background(),
				productID: 100,
			},
			mock: func(ts productRepositoryTestSuite) {
				query := `SELECT id, create_date, update_date, delete_date, user_id, initial, category, price, cost, name, description, barcode, expiry_date, size FROM products`
				ts.sqlMock.ExpectQuery(query).WithArgs(100).WillReturnError(sql.ErrNoRows)
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupUserRepositoryTestSuite()
			tt.mock(ts)

			// when
			got, err := ts.productRepository.GetProduct(tt.args.ctx, tt.args.productID)

			// then
			assert.Equal(t, tt.want, got)
			if ts.sqlMock.ExpectationsWereMet() != nil {
				t.Errorf("there were unfulfilled expectations: %s", ts.sqlMock.ExpectationsWereMet())
			}
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_productRepository_UpdateProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		product domain.Product
	}

	expiryDate := time.Date(2025, time.June, 10, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		args    args
		mock    func(ts productRepositoryTestSuite)
		wantErr bool
	}{
		{
			name: "PASS - 상품 수정 성공",
			args: args{
				ctx: context.Background(),
				product: domain.Product{
					Base: domain.Base{
						ID: 100,
					},
					UserID:      1,
					Initial:     "ㅅㅈ ㄹㄸ",
					Category:    "modified category",
					Price:       1000,
					Cost:        2000,
					Name:        "수정 라떼",
					Description: "modified description",
					Barcode:     "modified barcode",
					ExpiryDate:  expiryDate,
					Size:        domain.ProductSizeTypeSmall,
				},
			},
			mock: func(ts productRepositoryTestSuite) {
				ts.sqlMock.ExpectExec("UPDATE products").
					WithArgs(
						"ㅅㅈ ㄹㄸ",
						"modified category",
						float64(1000),
						float64(2000),
						"수정 라떼",
						"modified description",
						"modified barcode",
						expiryDate,
						domain.ProductSizeTypeSmall,
						100,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupUserRepositoryTestSuite()
			tt.mock(ts)

			// when
			err := ts.productRepository.UpdateProduct(tt.args.ctx, tt.args.product)

			// then
			if ts.sqlMock.ExpectationsWereMet() != nil {
				t.Errorf("there were unfulfilled expectations: %s", ts.sqlMock.ExpectationsWereMet())
			}
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_productRepository_DeleteProduct(t *testing.T) {
	type args struct {
		ctx       context.Context
		productID int
	}

	tests := []struct {
		name    string
		args    args
		mock    func(ts productRepositoryTestSuite)
		wantErr bool
	}{
		{
			name: "PASS - 상품 삭제",
			args: args{
				ctx:       context.Background(),
				productID: 1,
			},
			mock: func(ts productRepositoryTestSuite) {
				ts.sqlMock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupUserRepositoryTestSuite()
			tt.mock(ts)

			// when
			err := ts.productRepository.DeleteProduct(tt.args.ctx, tt.args.productID)

			// then
			if ts.sqlMock.ExpectationsWereMet() != nil {
				t.Errorf("there were unfulfilled expectations: %s", ts.sqlMock.ExpectationsWereMet())
			}
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_productRepository_ListProducts(t *testing.T) {
	type args struct {
		ctx    context.Context
		params domain.ListProductsParams
	}

	createDate := time.Now()
	updateDate := time.Now()
	expiryDate := time.Date(2025, time.June, 10, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		args    args
		mock    func(ts productRepositoryTestSuite)
		want    []domain.Product
		wantErr bool
	}{
		{
			name: "PASS - 조건 없는 상품 조회 성공",
			args: args{
				ctx: context.Background(),
				params: domain.ListProductsParams{
					UserID: 1,
				},
			},
			mock: func(ts productRepositoryTestSuite) {
				query := `SELECT id, create_date, update_date, delete_date, user_id, initial, category, price, cost, name, description, barcode, expiry_date, size FROM products`
				columns := []string{"id", "create_date", "update_date", "delete_date", "user_id", "initial", "category", "price", "cost", "name", "description", "barcode", "expiry_date", "size"}
				rows := sqlmock.NewRows(columns).AddRow(100, createDate, updateDate, nil, 1, "ㅅㅋㄹ ㄹㄸ", "payhere", 1000, 500, "슈크림 라떼", "description", "barcode", expiryDate, domain.ProductSizeTypeSmall)
				ts.sqlMock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
			},
			want: []domain.Product{
				{
					Base: domain.Base{
						ID:         100,
						CreateDate: createDate,
						UpdateDate: updateDate,
						DeleteDate: sql.NullTime{
							Time:  time.Time{},
							Valid: false,
						},
					},
					UserID:      1,
					Initial:     "ㅅㅋㄹ ㄹㄸ",
					Category:    "payhere",
					Price:       1000,
					Cost:        500,
					Name:        "슈크림 라떼",
					Description: "description",
					Barcode:     "barcode",
					ExpiryDate:  expiryDate,
					Size:        domain.ProductSizeTypeSmall,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupUserRepositoryTestSuite()
			tt.mock(ts)

			// whenR
			got, err := ts.productRepository.ListProducts(tt.args.ctx, tt.args.params)

			// then
			assert.Equal(t, tt.want, got)
			if ts.sqlMock.ExpectationsWereMet() != nil {
				t.Errorf("there were unfulfilled expectations: %s", ts.sqlMock.ExpectationsWereMet())
			}
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}
