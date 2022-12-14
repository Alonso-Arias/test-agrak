package dao

import (
	"context"

	"github.com/Alonso-Arias/test-agrak/db/base"
	"github.com/Alonso-Arias/test-agrak/db/model"
	"gorm.io/gorm"
)

// ProductDAO - product dao interface
type ProductImageDAO interface {
	FindAll(ctx context.Context) (model.ProductImage, error)
}

// ProductDAOImpl - product dao implementation
type ProductImageDAOImpl struct {
}

// NewProductImageDAO - gets an ProductDAOImpl instance
func NewProductImageDAO() *ProductImageDAOImpl {
	return &ProductImageDAOImpl{}
}

// FindAll -
func (pid *ProductImageDAOImpl) FindAll(ctx context.Context, sku string) ([]model.ProductImage, error) {

	log := loggerf.WithField("struct", "ProductImageDAOImpl").WithField("function", "FindAll")

	db := base.GetDB()

	productsImages := []model.ProductImage{}
	err := db.Where("products_sku = ?", sku).Find(&productsImages).Error

	if err != nil {
		log.WithError(err).Error("get products images fails")
		return []model.ProductImage{}, err
	} else if productsImages == nil {
		return []model.ProductImage{}, gorm.ErrRecordNotFound
	}

	log.Debugf("%v", productsImages)

	return productsImages, nil

}
