package dao

import (
	"context"

	"github.com/Alonso-Arias/test-agrak/db/base"
	"github.com/Alonso-Arias/test-agrak/db/model"
	"github.com/Alonso-Arias/test-agrak/log"
	"gorm.io/gorm"
)

var loggerf = log.LoggerJSON().WithField("package", "dao")

// ProductDAO - product dao interface
type ProductDAO interface {
	FindAll(ctx context.Context) (model.Product, error)
}

// ProductDAOImpl - product dao implementation
type ProductDAOImpl struct {
}

// NewProductDAO - gets an ProductDAOImpl instance
func NewProductDAO() *ProductDAOImpl {
	return &ProductDAOImpl{}
}

// FindAll -
func (pd *ProductDAOImpl) FindAll(ctx context.Context) ([]model.Product, error) {

	log := loggerf.WithField("struct", "ProductDAOImpl").WithField("function", "FindAll")

	db := base.GetDB()

	products := []model.Product{}
	err := db.Find(&products).Error

	if err != nil {
		log.WithError(err).Error("get products fails")
		return []model.Product{}, err
	} else if products == nil {
		return []model.Product{}, gorm.ErrRecordNotFound
	}

	log.Debugf("%v", products)

	return products, nil

}

// FindAll -
func (pd *ProductDAOImpl) Get(ctx context.Context, sku string) (model.Product, error) {

	log := loggerf.WithField("struct", "ProductDAOImpl").WithField("function", "Get")

	db := base.GetDB()

	product := model.Product{}
	err := db.Where("SKU = ?", sku).FirstOrInit(&product).Error

	if err != nil {
		log.WithError(err).Error("get products fails")
		return model.Product{}, err
	} else if product.Sku == "" {
		return model.Product{}, gorm.ErrRecordNotFound
	}

	log.Debugf("%v", product)

	return product, nil

}
