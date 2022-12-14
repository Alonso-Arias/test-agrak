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

func (pid *ProductImageDAOImpl) UpdateProductImages(ctx context.Context, sku string, productImages model.ProductImage) error {

	log := loggerf.WithField("struct", "ProductDAOImpl").WithField("function", "UpdateProductImages")

	db := base.GetDB()

	tx := db.Model(&productImages).
		Where("products_sku = ?", sku).
		Updates(map[string]interface{}{
			"url": gorm.Expr("IF(? = '', url, ?)", productImages.Url, productImages.Url),
		})

	if tx.Error != nil {
		log.Debugf("%v", tx.Error)
		return tx.Error
	}

	return nil
}

func (pid *ProductImageDAOImpl) SaveProductImages(ctx context.Context, sku string, productImages model.ProductImage) error {

	log := loggerf.WithField("struct", "ProductDAOImpl").WithField("function", "SaveProductImages")

	db := base.GetDB()

	err := db.Create(&productImages)

	if err.Error != nil {
		log.Debugf("%v", err.Error)
		return err.Error
	}

	log.Infof("Save product Images Sucessfull\n")

	return nil
}
