package product

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"k8s.io/utils/pointer"
	"net/http"
	"net/http/httptest"
	"net/url"
	"payhere/config"
	"payhere/domain"
	"payhere/internal/auth_token"
	"payhere/mocks"
	"testing"
	"time"
)

type productControllerTestSuite struct {
	router            *gin.Engine
	cfg               *config.Config
	autRepository     *mocks.AuthTokenRepository
	productService    *mocks.ProductService
	productController domain.ProductController
}

func setupProductControllerTestSuite(t *testing.T) productControllerTestSuite {
	var us productControllerTestSuite

	gin.SetMode(gin.TestMode)
	us.router = gin.Default()
	us.autRepository = mocks.NewAuthTokenRepository(t)
	us.productService = mocks.NewProductService(t)
	us.cfg = &config.Config{
		Auth: config.Auth{
			Secret: "payhere_test_secret",
		},
	}

	us.productController = NewProductController(us.productService)
	RegisterRoutes(
		us.router, us.productController,
		us.autRepository,
		us.cfg,
	)

	return us
}

func Test_productController_CreateProduct(t *testing.T) {
	expiryDate := time.Date(2025, time.June, 10, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		body func() *bytes.Reader
		mock func(ts productControllerTestSuite)
		code int
	}{
		{
			name: "PASS - 상품 생성 성공",
			body: func() *bytes.Reader {
				req := domain.CreateProductRequest{
					UserID:      1,
					Category:    "payhere",
					Price:       1000,
					Cost:        500,
					Name:        "test_product",
					Description: "test_description",
					Barcode:     "1234567890",
					ExpiryDate:  expiryDate,
					Size:        domain.ProductSizeTypeSmall,
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
				ts.productService.EXPECT().CreateProduct(mock.Anything, domain.CreateProductRequest{
					UserID:      1,
					Category:    "payhere",
					Price:       1000,
					Cost:        500,
					Name:        "test_product",
					Description: "test_description",
					Barcode:     "1234567890",
					ExpiryDate:  expiryDate,
					Size:        domain.ProductSizeTypeSmall,
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "FAIL - 음수의 가격",
			body: func() *bytes.Reader {
				req := domain.CreateProductRequest{
					UserID:      1,
					Category:    "payhere",
					Price:       -1000,
					Cost:        500,
					Name:        "test_product",
					Description: "test_description",
					Barcode:     "1234567890",
					ExpiryDate:  expiryDate,
					Size:        domain.ProductSizeTypeSmall,
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 음수의 원가",
			body: func() *bytes.Reader {
				req := domain.CreateProductRequest{
					UserID:      1,
					Category:    "payhere",
					Price:       1000,
					Cost:        -500,
					Name:        "test_product",
					Description: "test_description",
					Barcode:     "1234567890",
					ExpiryDate:  expiryDate,
					Size:        domain.ProductSizeTypeSmall,
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 비어있는 이름",
			body: func() *bytes.Reader {
				req := domain.CreateProductRequest{
					UserID:      1,
					Category:    "payhere",
					Price:       1000,
					Cost:        500,
					Name:        "",
					Description: "test_description",
					Barcode:     "1234567890",
					ExpiryDate:  expiryDate,
					Size:        domain.ProductSizeTypeSmall,
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 비어있는 설명",
			body: func() *bytes.Reader {
				req := domain.CreateProductRequest{
					UserID:      1,
					Category:    "category",
					Price:       1000,
					Cost:        500,
					Name:        "test_product",
					Description: "",
					Barcode:     "1234567890",
					ExpiryDate:  expiryDate,
					Size:        domain.ProductSizeTypeSmall,
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 비어있는 바코드",
			body: func() *bytes.Reader {
				req := domain.CreateProductRequest{
					UserID:      1,
					Category:    "payhere",
					Price:       1000,
					Cost:        500,
					Name:        "test_product",
					Description: "test_description",
					Barcode:     "",
					ExpiryDate:  expiryDate,
					Size:        domain.ProductSizeTypeSmall,
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 비어있는 유효기간",
			body: func() *bytes.Reader {
				req := domain.CreateProductRequest{
					UserID:      1,
					Category:    "",
					Price:       1000,
					Cost:        500,
					Name:        "test_product",
					Description: "test_description",
					Barcode:     "1234567890",
					ExpiryDate:  time.Time{},
					Size:        domain.ProductSizeTypeSmall,
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 비어있는 사이즈 입력",
			body: func() *bytes.Reader {
				req := domain.CreateProductRequest{
					UserID:      1,
					Category:    "",
					Price:       1000,
					Cost:        500,
					Name:        "test_product",
					Description: "test_description",
					Barcode:     "1234567890",
					ExpiryDate:  expiryDate,
					Size:        "",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},

			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 잘못 된 사이즈 입력",
			body: func() *bytes.Reader {
				req := domain.CreateProductRequest{
					UserID:      1,
					Category:    "",
					Price:       1000,
					Cost:        500,
					Name:        "test_product",
					Description: "test_description",
					Barcode:     "1234567890",
					ExpiryDate:  expiryDate,
					Size:        "payhere",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupProductControllerTestSuite(t)
			tt.mock(ts)
			req, _ := http.NewRequest(http.MethodPost, "/products", tt.body())
			req.Header.Set("Content-Type", "application/json")
			token, _ := auth_token.CreateAccessToken(domain.User{
				Base: domain.Base{
					ID: 1,
				},
			}, ts.cfg.Auth.Secret, time.Now().UTC().Add(time.Hour*time.Duration(24)))
			req.Header.Set("Authorization", "Bearer "+token)

			// when
			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			// then
			assert.Equal(t, tt.code, rec.Code)
			ts.productService.AssertExpectations(t)
		})
	}
}

func Test_productController_GetProduct(t *testing.T) {
	tests := []struct {
		name string
		path func() string
		mock func(ts productControllerTestSuite)
		code int
	}{
		{
			name: "PASS - 유효한 상품 ID",
			path: func() string {
				path, _ := url.JoinPath("/products", "100")
				return path
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
				ts.productService.EXPECT().GetProduct(mock.Anything, domain.GetProductRequest{
					UserID:    1,
					ProductID: 100,
				}).Return(domain.GetProductResponse{}, nil).Once()
			},
			code: http.StatusOK,
		},
		{
			name: "FAIL - 유효하지 않은 상품ID",
			path: func() string {
				path, _ := url.JoinPath("/products", "/payhere")
				return path
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// path
			ts := setupProductControllerTestSuite(t)
			tt.mock(ts)
			req, _ := http.NewRequest(http.MethodGet, tt.path(), nil)
			token, _ := auth_token.CreateAccessToken(domain.User{
				Base: domain.Base{
					ID: 1,
				},
			}, ts.cfg.Auth.Secret, time.Now().UTC().Add(time.Hour*time.Duration(24)))
			req.Header.Set("Authorization", "Bearer "+token)

			// when
			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			// then
			ts.productService.AssertExpectations(t)
			assert.Equal(t, tt.code, rec.Code)
		})
	}
}

func Test_productController_PatchProduct(t *testing.T) {
	expiryDate := time.Date(2025, time.June, 10, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		body func() *bytes.Reader
		mock func(ts productControllerTestSuite)
		code int
	}{
		{
			name: "PASS - 상품 전체 수정",
			body: func() *bytes.Reader {
				req := domain.PatchProductRequest{
					ID:          1,
					Category:    pointer.String("payhere"),
					Price:       pointer.Float64(1000),
					Cost:        pointer.Float64(500),
					Name:        pointer.String("이름을 수정"),
					Description: pointer.String("modify description"),
					Barcode:     pointer.String("barcode"),
					ExpiryDate:  &expiryDate,
					Size:        domain.ProductSizeTypeLarge.ToPointer(),
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
				ts.productService.EXPECT().PatchProduct(mock.Anything, domain.PatchProductRequest{
					UserID:      1,
					ID:          1,
					Category:    pointer.String("payhere"),
					Price:       pointer.Float64(1000),
					Cost:        pointer.Float64(500),
					Name:        pointer.String("이름을 수정"),
					Description: pointer.String("modify description"),
					Barcode:     pointer.String("barcode"),
					ExpiryDate:  &expiryDate,
					Size:        domain.ProductSizeTypeLarge.ToPointer(),
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "PASS - Category 수정",
			body: func() *bytes.Reader {
				req := domain.PatchProductRequest{
					ID:       1,
					Category: pointer.String("payhere"),
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
				ts.productService.EXPECT().PatchProduct(mock.Anything, domain.PatchProductRequest{
					UserID:   1,
					ID:       1,
					Category: pointer.String("payhere"),
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "PASS - 가격 수정",
			body: func() *bytes.Reader {
				req := domain.PatchProductRequest{
					ID:    1,
					Price: pointer.Float64(1000),
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
				ts.productService.EXPECT().PatchProduct(mock.Anything, domain.PatchProductRequest{
					UserID: 1,
					ID:     1,
					Price:  pointer.Float64(1000),
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "PASS - 원가 수정",
			body: func() *bytes.Reader {
				req := domain.PatchProductRequest{
					ID:   1,
					Cost: pointer.Float64(500),
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
				ts.productService.EXPECT().PatchProduct(mock.Anything, domain.PatchProductRequest{
					UserID: 1,
					ID:     1,
					Cost:   pointer.Float64(500),
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "PASS - 이름 수정",
			body: func() *bytes.Reader {
				req := domain.PatchProductRequest{
					ID:   1,
					Name: pointer.String("이름을 수정"),
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
				ts.productService.EXPECT().PatchProduct(mock.Anything, domain.PatchProductRequest{
					UserID: 1,
					ID:     1,
					Name:   pointer.String("이름을 수정"),
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "PASS - 설명 수정",
			body: func() *bytes.Reader {
				req := domain.PatchProductRequest{
					ID:          1,
					Description: pointer.String("modify description"),
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
				ts.productService.EXPECT().PatchProduct(mock.Anything, domain.PatchProductRequest{
					UserID:      1,
					ID:          1,
					Description: pointer.String("modify description"),
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "PASS - 바코드 수정",
			body: func() *bytes.Reader {
				req := domain.PatchProductRequest{
					ID:      1,
					Barcode: pointer.String("barcode"),
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
				ts.productService.EXPECT().PatchProduct(mock.Anything, domain.PatchProductRequest{
					UserID:  1,
					ID:      1,
					Barcode: pointer.String("barcode"),
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "PASS - 유통기한 수정",
			body: func() *bytes.Reader {
				req := domain.PatchProductRequest{
					ID:         1,
					ExpiryDate: &expiryDate,
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
				ts.productService.EXPECT().PatchProduct(mock.Anything, domain.PatchProductRequest{
					UserID:     1,
					ID:         1,
					ExpiryDate: &expiryDate,
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "PASS - 사이즈 수정",
			body: func() *bytes.Reader {
				req := domain.PatchProductRequest{
					ID:   1,
					Size: domain.ProductSizeTypeLarge.ToPointer(),
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
				ts.productService.EXPECT().PatchProduct(mock.Anything, domain.PatchProductRequest{
					UserID: 1,
					ID:     1,
					Size:   domain.ProductSizeTypeLarge.ToPointer(),
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "FAIL - 비어있는 카테고리 수정",
			body: func() *bytes.Reader {
				req := domain.PatchProductRequest{
					ID:       1,
					Category: pointer.String(""),
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 음수의 가격 수정",
			body: func() *bytes.Reader {
				req := domain.PatchProductRequest{
					ID:    1,
					Price: pointer.Float64(-1000),
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 음수의 원가 수정",
			body: func() *bytes.Reader {
				req := domain.PatchProductRequest{
					ID:   1,
					Cost: pointer.Float64(-1000),
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 비어있는 이름 수정",
			body: func() *bytes.Reader {
				req := domain.PatchProductRequest{
					ID:   1,
					Name: pointer.String(""),
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 비어있는 설명 수정",
			body: func() *bytes.Reader {
				req := domain.PatchProductRequest{
					ID:          1,
					Description: pointer.String(""),
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 비어있는 바코드 수정",
			body: func() *bytes.Reader {
				req := domain.PatchProductRequest{
					ID:      1,
					Barcode: pointer.String(""),
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 비어있는 유통기한 수정",
			body: func() *bytes.Reader {
				emptyTime := time.Time{}
				req := domain.PatchProductRequest{
					ID:         1,
					ExpiryDate: &emptyTime,
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 빈 문자열 사이즈 수정",
			body: func() *bytes.Reader {
				emptySizeType := domain.ProductSizeType("")
				req := domain.PatchProductRequest{
					ID:   1,
					Size: &emptySizeType,
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupProductControllerTestSuite(t)
			tt.mock(ts)
			req, _ := http.NewRequest(http.MethodPatch, "/products", tt.body())
			req.Header.Set("Content-Type", "application/json")
			token, _ := auth_token.CreateAccessToken(domain.User{
				Base: domain.Base{
					ID: 1,
				},
			}, ts.cfg.Auth.Secret, time.Now().UTC().Add(time.Hour*time.Duration(24)))
			req.Header.Set("Authorization", "Bearer "+token)

			// when
			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			// then
			assert.Equal(t, tt.code, rec.Code)
			ts.productService.AssertExpectations(t)
		})
	}
}

func Test_productController_DeleteProduct(t *testing.T) {
	tests := []struct {
		name string
		path func() string
		mock func(ts productControllerTestSuite)
		code int
	}{
		{
			name: "PASS - 상품 삭제 성공",
			path: func() string {
				path, _ := url.JoinPath("/products", "100")
				return path
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
				ts.productService.EXPECT().DeleteProduct(mock.Anything, domain.DeleteProductRequest{
					UserID: 1,
					ID:     100,
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "FAIL - 잘못된 상품 ID",
			path: func() string {
				path, _ := url.JoinPath("/products", "-1")
				return path
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
			},
			code: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupProductControllerTestSuite(t)
			tt.mock(ts)
			req, _ := http.NewRequest(http.MethodDelete, tt.path(), nil)
			token, _ := auth_token.CreateAccessToken(domain.User{
				Base: domain.Base{
					ID: 1,
				},
			}, ts.cfg.Auth.Secret, time.Now().UTC().Add(time.Hour*time.Duration(24)))
			req.Header.Set("Authorization", "Bearer "+token)

			// when
			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			// then
			assert.Equal(t, tt.code, rec.Code)
			ts.productService.AssertExpectations(t)
		})
	}
}

func Test_productController_ListProducts(t *testing.T) {
	tests := []struct {
		name  string
		query func() string
		mock  func(ts productControllerTestSuite)
		code  int
	}{
		{
			name: "PASS - 유효한 커서와 검색어",
			query: func() string {
				params := url.Values{}
				params.Add("cursor", "1")
				params.Add("search", "슈크림 라떼")
				return params.Encode()
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
				ts.productService.EXPECT().ListProducts(mock.Anything, domain.ListProductsRequest{
					UserID: 1,
					Cursor: pointer.Int(1),
					Search: pointer.String("슈크림 라떼"),
				}).Return(domain.ListProductsResponse{}, nil).Once()
			},
			code: http.StatusOK,
		},
		{
			name: "PASS - 커서만 입력",
			query: func() string {
				params := url.Values{}
				params.Add("cursor", "1")
				return params.Encode()
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
				ts.productService.EXPECT().ListProducts(mock.Anything, domain.ListProductsRequest{
					UserID: 1,
					Cursor: pointer.Int(1),
				}).Return(domain.ListProductsResponse{}, nil).Once()
			},
			code: http.StatusOK,
		},
		{
			name: "PASS - 검색어만 입력",
			query: func() string {
				params := url.Values{}
				params.Add("search", "슈크림 라떼")
				return params.Encode()
			},
			mock: func(ts productControllerTestSuite) {
				ts.autRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(
					mock.Anything,
					mock.MatchedBy(func(params domain.FindByUserIDAndJwtTokenParams) bool { return params.UserID == 1 }),
				).Return(domain.AuthToken{
					ExpirationTime: time.Now().UTC().Add(time.Hour * time.Duration(24)),
					Active:         true,
				}, nil).Once()
				ts.productService.EXPECT().ListProducts(mock.Anything, domain.ListProductsRequest{
					UserID: 1,
					Search: pointer.String("슈크림 라떼"),
				}).Return(domain.ListProductsResponse{}, nil).Once()
			},
			code: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupProductControllerTestSuite(t)
			tt.mock(ts)
			req, _ := http.NewRequest(http.MethodGet, "/products", nil)
			req.URL.RawQuery = tt.query()
			token, _ := auth_token.CreateAccessToken(domain.User{
				Base: domain.Base{
					ID: 1,
				},
			}, ts.cfg.Auth.Secret, time.Now().UTC().Add(time.Hour*time.Duration(24)))
			req.Header.Set("Authorization", "Bearer "+token)

			// when
			rec := httptest.NewRecorder()
			t.Logf("Request URL: %s", req.URL.String())
			ts.router.ServeHTTP(rec, req)

			// then
			assert.Equal(t, tt.code, rec.Code)
			ts.productService.AssertExpectations(t)
		})
	}
}
