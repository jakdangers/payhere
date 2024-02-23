package product

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"payhere/config"
	"payhere/domain"
	cerrors "payhere/pkg/cerror"
	"payhere/pkg/router"
	"time"
)

// RegisterRoutes TODO : authRepository을 직접 사용하지 않도록 변경
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

func (pc productController) ListProducts(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
