package product

import (
	"context"

	"github.com/Alonso-Arias/test-agrak/db/dao"
	errs "github.com/Alonso-Arias/test-agrak/errors"
	"github.com/Alonso-Arias/test-agrak/log"
	"github.com/Alonso-Arias/test-agrak/services/model"
	"gorm.io/gorm"
)

var loggerf = log.LoggerJSON().WithField("package", "services")

type ProductService struct {
}

type FindAllProductsResponse struct {
	Products []model.Product `json:"products"`
}

func (ps ProductService) FindAllProducts(ctx context.Context) (FindAllProductsResponse, error) {
	log := loggerf.WithField("service", "ProductService").WithField("func", "FindAllProducts")

	productDao := dao.NewProductDAO()

	products, err := productDao.FindAll(ctx)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("problems with getting products")
		return FindAllProductsResponse{}, err
	} else if err == gorm.ErrRecordNotFound {
		return FindAllProductsResponse{}, errs.ProductsNotFound
	}

	results := []model.Product{}

	for _, v := range products {
		product := model.Product{
			Sku:              v.Sku,
			Name:             v.Name,
			Brand:            v.Brand,
			Size:             v.Size,
			Price:            v.Price,
			PrincipalImage:   v.PrincipalImage,
			AdditionalImages: []string{},
		}
		results = append(results, product)
	}

	return FindAllProductsResponse{Products: results}, nil
}
