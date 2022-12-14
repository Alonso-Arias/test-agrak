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

// FindAll -
func (pd *ProductDAOImpl) Delete(ctx context.Context, sku string) error {

	log := loggerf.WithField("struct", "ProductDAOImpl").WithField("function", "Get")

	db := base.GetDB()

	// inits tx
	err := db.Transaction(func(tx *gorm.DB) error {

		product := model.Product{}

		err := db.Where("sku = ?", sku).Delete(&product).Error
		if err != nil {
			log.WithError(err).Error("problems with deleting product")
			return err
		}

		productImages := model.ProductImage{}

		err = db.Where("products_sku = ?", sku).Delete(&productImages).Error
		if err != nil {
			log.WithError(err).Error("problems with deleting products images url")
			return err
		}

		return nil
	})

	if err != nil {
		log.WithError(err).Error("fails to save order")
		return err
	}

	log.Infof("DEBUG : Saved Sucessfull\n")

	return nil

}

func (pd *ProductDAOImpl) Update(ctx context.Context, product model.Product) error {

	log := loggerf.WithField("struct", "ProductDAOImpl").WithField("function", "Update")

	db := base.GetDB()

	tx := db.Model(&product).
		Where("sku = ?", product.Sku).
		Updates(map[string]interface{}{
			"name":            gorm.Expr("IF(? = '', name, ?)", product.Name, product.Name),
			"brand":           gorm.Expr("IF(? = '', brand, ?)", product.Brand, product.Brand),
			"size":            gorm.Expr("IF(? = '', size, ?)", product.Size, product.Size),
			"price":           gorm.Expr("IF(? = '', price, ?)", product.Price, product.Price),
			"principal_image": gorm.Expr("IF(? = '', principal_image, ?)", product.PrincipalImage, product.PrincipalImage),
		})

	if tx.Error != nil {
		log.Debugf("%v", tx.Error)
		return tx.Error
	}

	return nil
}

func (pd *ProductDAOImpl) Save(ctx context.Context, product model.Product) error {

	log := loggerf.WithField("struct", "ProductDAOImpl").WithField("function", "Save")

	db := base.GetDB()

	err := db.Create(&product)

	if err.Error != nil {
		log.Debugf("%v", err.Error)
		return err.Error
	}

	log.Infof("Save product Sucessfull\n")

	return nil

}
