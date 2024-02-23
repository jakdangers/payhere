package product

import (
	"go.uber.org/fx"
	"payhere/domain"
)

var Module = fx.Module(
	"product", fx.Provide(
		fx.Annotate(NewProductRepository, fx.As(new(domain.ProductRepository))),
		fx.Annotate(NewProductService, fx.As(new(domain.ProductService))),
		fx.Annotate(NewProductController, fx.As(new(domain.ProductController))),
	),
)
