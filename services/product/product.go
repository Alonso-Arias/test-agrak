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
	productImageDao := dao.NewProductImageDAO()

	products, err := productDao.FindAll(ctx)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("problems with getting products")
		return FindAllProductsResponse{}, err
	} else if err == gorm.ErrRecordNotFound {
		return FindAllProductsResponse{}, errs.ProductsNotFound
	}

	results := []model.Product{}

	for _, v := range products {
		pi, err := productImageDao.FindAll(ctx, v.Sku)
		if err != nil {
			return FindAllProductsResponse{}, err
		}
		var productsImages []string
		for _, item := range pi {
			productsImages = append(productsImages, item.Url)
		}
		product := model.Product{
			Sku:              v.Sku,
			Name:             v.Name,
			Brand:            v.Brand,
			Size:             v.Size,
			Price:            v.Price,
			PrincipalImage:   v.PrincipalImage,
			AdditionalImages: productsImages,
		}
		results = append(results, product)
	}

	return FindAllProductsResponse{Products: results}, nil
}

type GetProductRequest struct {
	Sku string `json:"sku"`
}
type GetProductResponse struct {
	Product model.Product `json:"product"`
}

func (ps ProductService) GetProduct(ctx context.Context, in GetProductRequest) (GetProductResponse, error) {
	log := loggerf.WithField("service", "ProductService").WithField("func", "GetProduct")

	productDao := dao.NewProductDAO()
	productImageDao := dao.NewProductImageDAO()

	v, err := productDao.Get(ctx, in.Sku)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("problems with getting products")
		return GetProductResponse{}, err
	} else if err == gorm.ErrRecordNotFound {
		return GetProductResponse{}, errs.ProductsNotFound
	}

	pi, err := productImageDao.FindAll(ctx, v.Sku)
	if err != nil {
		return GetProductResponse{}, err
	}
	var productsImages []string
	for _, item := range pi {
		productsImages = append(productsImages, item.Url)
	}
	product := model.Product{
		Sku:              v.Sku,
		Name:             v.Name,
		Brand:            v.Brand,
		Size:             v.Size,
		Price:            v.Price,
		PrincipalImage:   v.PrincipalImage,
		AdditionalImages: productsImages,
	}

	return GetProductResponse{Product: product}, nil
}
