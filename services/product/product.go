package product

import (
	"context"

	"github.com/Alonso-Arias/test-agrak/db/dao"
	md "github.com/Alonso-Arias/test-agrak/db/model"
	errs "github.com/Alonso-Arias/test-agrak/errors"
	"github.com/Alonso-Arias/test-agrak/log"
	"github.com/Alonso-Arias/test-agrak/services/model"
	"gopkg.in/dealancer/validate.v2"
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

	if in.Sku == "" {
		return GetProductResponse{}, errs.BadRequest
	}

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

type DeleteProductRequest struct {
	Sku string `json:"sku"`
}
type DeleteProductResponse struct {
}

func (ps ProductService) DeleteProduct(ctx context.Context, in DeleteProductRequest) (DeleteProductResponse, error) {
	log := loggerf.WithField("service", "ProductService").WithField("func", "GetProduct")

	if in.Sku == "" {
		return DeleteProductResponse{}, errs.BadRequest
	}
	productDao := dao.NewProductDAO()

	_, err := productDao.Get(ctx, in.Sku)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("problems with getting products")
		return DeleteProductResponse{}, err
	} else if err == gorm.ErrRecordNotFound {
		return DeleteProductResponse{}, errs.ProductsNotFound
	}

	err = productDao.Delete(ctx, in.Sku)
	if err != nil {
		return DeleteProductResponse{}, err
	}

	return DeleteProductResponse{}, nil
}

type UpdateProductRequest struct {
	Product model.Product `json:"product"`
}
type UpdateProductResponse struct {
}

func (ps ProductService) UpdateProduct(ctx context.Context, in UpdateProductRequest) (UpdateProductResponse, error) {
	log := loggerf.WithField("service", "ProductService").WithField("func", "GetProduct")

	// validates input request
	if err := validate.Validate(in); err != nil {
		log.WithError(err).Error("validates problems")
		return UpdateProductResponse{}, errs.BadRequest
	}

	productDao := dao.NewProductDAO()
	productImageDao := dao.NewProductImageDAO()

	_, err := productDao.Get(ctx, in.Product.Sku)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("problems with getting products")
		return UpdateProductResponse{}, err
	} else if err == gorm.ErrRecordNotFound {
		return UpdateProductResponse{}, errs.ProductsNotFound
	}

	err = productDao.Update(ctx, md.Product(md.Product{Sku: in.Product.Sku, Name: in.Product.Name, Brand: in.Product.Brand, Size: in.Product.Size, Price: in.Product.Price}))
	if err != nil {
		return UpdateProductResponse{}, err
	}

	for _, v := range in.Product.AdditionalImages {
		err = productImageDao.UpdateProductImages(ctx, in.Product.Sku, md.ProductImage{Url: v})
		if err != nil {
			return UpdateProductResponse{}, err
		}
	}

	return UpdateProductResponse{}, nil
}

type SaveProductRequest struct {
	Product model.Product `json:"product"`
}
type SaveProductResponse struct {
}

func (ps ProductService) SaveProduct(ctx context.Context, in SaveProductRequest) (SaveProductResponse, error) {
	log := loggerf.WithField("service", "ProductService").WithField("func", "GetProduct")

	// validates input request
	if err := validate.Validate(in); err != nil {
		log.WithError(err).Error("validates problems")
		return SaveProductResponse{}, errs.BadRequest
	}

	productDao := dao.NewProductDAO()
	productImageDao := dao.NewProductImageDAO()

	// validacion de producto
	p, err := productDao.Get(ctx, in.Product.Sku)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("problems with getting products")
		return SaveProductResponse{}, err
	} else if err == gorm.ErrRecordNotFound {
		return SaveProductResponse{}, errs.ProductsNotFound
	}

	// valida si ya esta o no el producto a guardar
	if p.Sku == in.Product.Sku {
		return SaveProductResponse{}, errs.ProductAlreadySaved
	}

	err = productDao.Save(ctx, md.Product(md.Product{Sku: in.Product.Sku, Name: in.Product.Name, Brand: in.Product.Brand, Size: in.Product.Size, Price: in.Product.Price}))
	if err != nil {
		return SaveProductResponse{}, err
	}

	for _, v := range in.Product.AdditionalImages {
		err = productImageDao.UpdateProductImages(ctx, in.Product.Sku, md.ProductImage{Url: v, ProductsSku: in.Product.Sku})
		if err != nil {
			return SaveProductResponse{}, err
		}
	}

	return SaveProductResponse{}, nil
}
