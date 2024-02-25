package product

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"payhere/config"
	"payhere/domain"
	cerrors "payhere/pkg/cerrors"
	"payhere/pkg/router"
	"time"
)

func RegisterRoutes(e *gin.Engine, controller domain.ProductController, authTokenRepository domain.AuthTokenRepository, cfg *config.Config) {
	products := e.Group("/products")
	{
		products.POST("", router.JWTMiddleware(cfg.Auth.Secret, authTokenRepository), controller.CreateProduct)
		products.GET("/:productID", router.JWTMiddleware(cfg.Auth.Secret, authTokenRepository), controller.GetProduct)
		products.PATCH("", router.JWTMiddleware(cfg.Auth.Secret, authTokenRepository), controller.PatchProduct)
		products.DELETE("/:productID", router.JWTMiddleware(cfg.Auth.Secret, authTokenRepository), controller.DeleteProduct)
		products.GET("", router.JWTMiddleware(cfg.Auth.Secret, authTokenRepository), controller.ListProducts)
	}
}

type productController struct {
	productService domain.ProductService
}

func NewProductController(service domain.ProductService) *productController {
	return &productController{
		productService: service,
	}
}

var _ domain.ProductController = (*productController)(nil)

// CreateProduct
// @Summary 상품 생성
// @Description 필수 항목을 입력하여 상품을 생성합니다.
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param CreateProductRequest body domain.CreateProductRequest true "상품 생성 요청"
// @Success 204
// @Router /products [post]
func (pc productController) CreateProduct(c *gin.Context) {
	var req domain.CreateProductRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	userID, err := router.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}
	req.UserID = userID

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	if err := pc.productService.CreateProduct(ctx, req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.Status(http.StatusNoContent)
}

// GetProduct
// @Summary 단일 상품 조회
// @Description 상품 ID로 상품을 조회합니다. (단 자신의 상품만 조회 가능)
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Param id path int true "상품 ID"
// @Success 200 {object} domain.GetProductResponse "상품 상세 정보"
// @Router /products/{id} [get]
func (pc productController) GetProduct(c *gin.Context) {
	var req domain.GetProductRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	userID, err := router.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}
	req.UserID = userID

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := pc.productService.GetProduct(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(domain.PayhereResponseFrom(http.StatusOK, res))
}

// PatchProduct
// @Summary 전체 또는 부분 상품 수정
// @Description 특정 상품 ID로 상품을 전체 또는 부분 수정합니다. (단 자신의 상품만 수정 가능)
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param PatchProductRequest body domain.PatchProductRequest true "상품 수정 요청"
// @Success 204
// @Router /products [patch]
func (pc productController) PatchProduct(c *gin.Context) {
	var req domain.PatchProductRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	userID, err := router.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}
	req.UserID = userID

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	if err := pc.productService.PatchProduct(ctx, req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteProduct
// @Summary 상품 삭제
// @Description 상품 ID로 상품을 삭제합니다. (단 자신의 상품만 삭제 가능)
// @Tags Product
// @Produce json
// @Param id path int true "제품 ID"
// @Security BearerAuth
// @Success 204
// @Router /products/{id} [delete]
func (pc productController) DeleteProduct(c *gin.Context) {
	var req domain.DeleteProductRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	userID, err := router.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}
	req.UserID = userID

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	if err := pc.productService.DeleteProduct(ctx, req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.Status(http.StatusNoContent)
}

// ListProducts
// @Summary 상품 목록 조회
// @Description 상품 목록을 조회합니다.
// @Tags Product
// @Produce json
// @Param cursor query int false "커서"
// @Param search query string false "검색어"
// @Security BearerAuth
// @Success 200 {object} domain.ListProductsResponse "상품 목록"
// @Router /products [get]
func (pc productController) ListProducts(c *gin.Context) {
	var req domain.ListProductsRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	userID, err := router.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}
	req.UserID = userID

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := pc.productService.ListProducts(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(domain.PayhereResponseFrom(http.StatusOK, res))
}
