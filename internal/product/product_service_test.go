package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"k8s.io/utils/pointer"
	"payhere/domain"
	"payhere/mocks"
	"testing"
	"time"
)

type productServiceTestSuite struct {
	userRepository    *mocks.UserRepository
	productRepository *mocks.ProductRepository
	productService    domain.ProductService
}

func setupUserServiceTestSuite(t *testing.T) productServiceTestSuite {
	var us productServiceTestSuite

	us.userRepository = mocks.NewUserRepository(t)
	us.productRepository = mocks.NewProductRepository(t)
	us.productService = NewProductService(
		us.userRepository,
		us.productRepository,
	)

	return us
}

func Test_productService_CreateProduct(t *testing.T) {
	type args struct {
		ctx context.Context
		req domain.CreateProductRequest
	}

	tests := []struct {
		name    string
		args    args
		mock    func(ts productServiceTestSuite)
		wantErr bool
	}{
		{
			name: "PASS - 상품 생성 ",
			args: args{
				ctx: context.Background(),
				req: domain.CreateProductRequest{
					UserID:      1,
					Category:    "category",
					Price:       1000,
					Cost:        500,
					Name:        "슈크림 라떼",
					Description: "description",
					Barcode:     "barcode",
					ExpiryDate:  time.Date(2025, time.June, 10, 0, 0, 0, 0, time.UTC),
					Size:        domain.ProductSizeTypeSmall,
				},
			},
			mock: func(ts productServiceTestSuite) {
				ts.productRepository.On("CreateProduct", context.Background(), domain.Product{
					UserID:      1,
					Category:    "category",
					Initial:     "ㅅㅋㄹ ㄹㄸ",
					Price:       1000,
					Cost:        500,
					Name:        "슈크림 라떼",
					Description: "description",
					Barcode:     "barcode",
					ExpiryDate:  time.Date(2025, time.June, 10, 0, 0, 0, 0, time.UTC),
					Size:        domain.ProductSizeTypeSmall,
				}).Return(0, nil)

			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupUserServiceTestSuite(t)
			tt.mock(ts)

			// when
			err := ts.productService.CreateProduct(tt.args.ctx, tt.args.req)

			// then
			ts.userRepository.AssertExpectations(t)
			ts.productRepository.AssertExpectations(t)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_productService_GetProduct(t *testing.T) {
	type args struct {
		ctx context.Context
		req domain.GetProductRequest
	}

	tests := []struct {
		name    string
		args    args
		mock    func(ts productServiceTestSuite)
		want    domain.GetProductResponse
		wantErr bool
	}{
		{
			name: "PASS - 상품 조회 성공",
			args: args{
				ctx: context.Background(),
				req: domain.GetProductRequest{
					UserID:    1,
					ProductID: 100,
				},
			},
			mock: func(ts productServiceTestSuite) {
				ts.productRepository.EXPECT().GetProduct(mock.Anything, 100).Return(&domain.Product{
					Base: domain.Base{
						ID: 100,
					},
					UserID:      1,
					Category:    "category",
					Initial:     "ㅅㅋㄹ ㄹㄸ",
					Price:       1000,
					Cost:        500,
					Name:        "슈크림 라떼",
					Description: "description",
					Barcode:     "barcode",
					ExpiryDate:  time.Date(2025, time.June, 10, 0, 0, 0, 0, time.UTC),
					Size:        domain.ProductSizeTypeSmall,
				}, nil)
			},
			want: domain.GetProductResponse{
				Product: domain.ProductDTO{
					BaseDTO: domain.BaseDTO{
						ID: 100,
					},
					UserID:      1,
					Category:    "category",
					Initial:     "ㅅㅋㄹ ㄹㄸ",
					Price:       1000,
					Cost:        500,
					Name:        "슈크림 라떼",
					Description: "description",
					Barcode:     "barcode",
					ExpiryDate:  time.Date(2025, time.June, 10, 0, 0, 0, 0, time.UTC),
					Size:        domain.ProductSizeTypeSmall,
				},
			},
			wantErr: false,
		},
		{
			name: "FAIL - 상품을 찾을 수 없는 경우",
			args: args{
				ctx: context.Background(),
				req: domain.GetProductRequest{
					UserID:    1,
					ProductID: 100,
				},
			},
			mock: func(ts productServiceTestSuite) {
				ts.productRepository.EXPECT().GetProduct(mock.Anything, 100).Return(nil, nil).Once()
			},
			want:    domain.GetProductResponse{},
			wantErr: true,
		},
		{
			name: "FAIL - 상품을 조회 할 권한이 없는 경우",
			args: args{
				ctx: context.Background(),
				req: domain.GetProductRequest{
					UserID:    1,
					ProductID: 100,
				},
			},
			mock: func(ts productServiceTestSuite) {
				ts.productRepository.EXPECT().GetProduct(mock.Anything, 100).Return(&domain.Product{
					Base: domain.Base{
						ID: 100,
					},
					UserID:      2,
					Category:    "category",
					Initial:     "ㅅㅋㄹ ㄹㄸ",
					Price:       1000,
					Cost:        500,
					Name:        "슈크림 라떼",
					Description: "description",
					Barcode:     "barcode",
					ExpiryDate:  time.Date(2025, time.June, 10, 0, 0, 0, 0, time.UTC),
					Size:        domain.ProductSizeTypeSmall,
				}, nil).Once()
			},
			want:    domain.GetProductResponse{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupUserServiceTestSuite(t)
			tt.mock(ts)

			// when
			got, err := ts.productService.GetProduct(tt.args.ctx, tt.args.req)

			// then
			ts.userRepository.AssertExpectations(t)
			ts.productRepository.AssertExpectations(t)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_productService_PatchProduct(t *testing.T) {
	type args struct {
		ctx context.Context
		req domain.PatchProductRequest
	}

	tests := []struct {
		name    string
		args    args
		mock    func(ts productServiceTestSuite)
		wantErr bool
	}{
		{
			name: "PASS - 상품 수정",
			args: args{
				ctx: nil,
				req: domain.PatchProductRequest{
					UserID:      2,
					ID:          100,
					Category:    pointer.String("modified category"),
					Price:       pointer.Float64(2000),
					Cost:        pointer.Float64(1000),
					Name:        pointer.String("수정된 모카"),
					Description: pointer.String("modified description"),
					Barcode:     pointer.String("modified barcode"),
					ExpiryDate: func() *time.Time {
						t := time.Date(2026, time.June, 10, 0, 0, 0, 0, time.UTC)
						return &t
					}(),
					Size: domain.ProductSizeTypeLarge.ToPointer(),
				},
			},
			mock: func(ts productServiceTestSuite) {
				ts.productRepository.EXPECT().GetProduct(mock.Anything, 100).Return(&domain.Product{
					Base: domain.Base{
						ID: 100,
					},
					UserID:      2,
					Category:    "original category",
					Initial:     "ㅇㄹㅈㄴ ㄹㄸ",
					Price:       1000,
					Cost:        500,
					Name:        "오리지널 라떼",
					Description: "description",
					Barcode:     "barcode",
					ExpiryDate:  time.Date(2025, time.June, 10, 0, 0, 0, 0, time.UTC),
					Size:        domain.ProductSizeTypeSmall,
				}, nil).Once()
				ts.productRepository.EXPECT().UpdateProduct(mock.Anything, domain.Product{
					Base: domain.Base{
						ID: 100,
					},
					UserID:      2,
					Initial:     "ㅅㅈㄷ ㅁㅋ",
					Category:    "modified category",
					Price:       2000,
					Cost:        1000,
					Name:        "수정된 모카",
					Description: "modified description",
					Barcode:     "modified barcode",
					ExpiryDate:  time.Date(2026, time.June, 10, 0, 0, 0, 0, time.UTC),
					Size:        domain.ProductSizeTypeLarge,
				}).Return(nil).Once()
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupUserServiceTestSuite(t)
			tt.mock(ts)

			// when
			err := ts.productService.PatchProduct(tt.args.ctx, tt.args.req)

			// then
			ts.userRepository.AssertExpectations(t)
			ts.productRepository.AssertExpectations(t)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_productService_DeleteProduct(t *testing.T) {
	type args struct {
		ctx context.Context
		req domain.DeleteProductRequest
	}

	tests := []struct {
		name    string
		args    args
		mock    func(ts productServiceTestSuite)
		wantErr bool
	}{
		{
			name: "PASS - 상품 삭제",
			args: args{
				ctx: context.Background(),
				req: domain.DeleteProductRequest{
					UserID: 1,
					ID:     100,
				},
			},
			mock: func(ts productServiceTestSuite) {
				ts.productRepository.EXPECT().GetProduct(mock.Anything, 100).Return(&domain.Product{
					Base: domain.Base{
						ID: 100,
					},
					UserID:      1,
					Category:    "original category",
					Initial:     "ㅇㄹㅈㄴ ㄹㄸ",
					Price:       1000,
					Cost:        500,
					Name:        "오리지널 라떼",
					Description: "description",
					Barcode:     "barcode",
					ExpiryDate:  time.Date(2025, time.June, 10, 0, 0, 0, 0, time.UTC),
					Size:        domain.ProductSizeTypeSmall,
				}, nil).Once()
				ts.productRepository.EXPECT().DeleteProduct(mock.Anything, 100).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name: "PASS - 상품을 찾을 수 없는 경우",
			args: args{
				ctx: context.Background(),
				req: domain.DeleteProductRequest{
					UserID: 1,
					ID:     100,
				},
			},
			mock: func(ts productServiceTestSuite) {
				ts.productRepository.EXPECT().GetProduct(mock.Anything, 100).Return(nil, nil).Once()
			},
			wantErr: true,
		},
		{
			name: "PASS - 상품을 삭제할 권한이 없는 경우",
			args: args{
				ctx: context.Background(),
				req: domain.DeleteProductRequest{
					UserID: 1,
					ID:     100,
				},
			},
			mock: func(ts productServiceTestSuite) {
				ts.productRepository.EXPECT().GetProduct(mock.Anything, 100).Return(&domain.Product{
					Base: domain.Base{
						ID: 100,
					},
					UserID:      2,
					Category:    "original category",
					Initial:     "ㅇㄹㅈㄴ ㄹㄸ",
					Price:       1000,
					Cost:        500,
					Name:        "오리지널 라떼",
					Description: "description",
					Barcode:     "barcode",
					ExpiryDate:  time.Date(2025, time.June, 10, 0, 0, 0, 0, time.UTC),
					Size:        domain.ProductSizeTypeSmall,
				}, nil).Once()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupUserServiceTestSuite(t)
			tt.mock(ts)

			// when
			err := ts.productService.DeleteProduct(tt.args.ctx, tt.args.req)

			// then
			ts.userRepository.AssertExpectations(t)
			ts.productRepository.AssertExpectations(t)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_ExtractChosung(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "한글만 입력",
			input:    "한글만 입력",
			expected: "ㅎㄱㅁ ㅇㄹ",
		},
		{
			name:     "한글 영어 혼합",
			input:    "한글abc혼합",
			expected: "ㅎㄱabcㅎㅎ",
		},
		{
			name:     "한글 영어 특수문자",
			input:    "한글abc!@#",
			expected: "ㅎㄱabc!@#",
		},
		{
			name:     "한글 자음만",
			input:    "ㄱㄴㄷ",
			expected: "ㄱㄴㄷ",
		},
		{
			name:     "영어만 입력",
			input:    "onlyenglish",
			expected: "onlyenglish",
		},
		{
			name:     "숫자만 입력",
			input:    "1234567890",
			expected: "1234567890",
		},
		{
			name:     "띄어쓰기도 포함",
			input:    "한 글  테 스 트",
			expected: "ㅎ ㄱ  ㅌ ㅅ ㅌ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractChosung(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
