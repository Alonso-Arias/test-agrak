package product

import (
	"context"

	_ "github.com/Alonso-Arias/test-agrak/errors"
	"github.com/Alonso-Arias/test-agrak/log"
	"github.com/Alonso-Arias/test-agrak/services/model"
)

var loggerf = log.LoggerJSON().WithField("package", "services")

type ProductService struct {
}

type FindAllProductsResponse struct {
	Products []model.Product `json:"products"`
}

func (ps ProductService) FindAllProducts(ctx context.Context) (FindAllProductsResponse, error) {
	return FindAllProductsResponse{}, nil
}
