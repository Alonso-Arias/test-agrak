package model

type Product struct {
	Sku            string
	Name           string
	Brand          string
	Size           string
	Price          string
	PrincipalImage string
	ProductsImages []ProductImage `gorm:"-"`
}

type ProductImage struct {
	Url string
}

func (ProductImage) TableName() string {
	return "products_images"
}
