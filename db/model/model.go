package model

type Product struct {
	Sku            string
	Name           string
	Brand          string
	Size           string
	Price          int
	PrincipalImage string
}

type ProductImage struct {
	ProductsSku string
	Url         string
}

func (ProductImage) TableName() string {
	return "products_images"
}
