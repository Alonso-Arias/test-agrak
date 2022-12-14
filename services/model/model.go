package model

// swagger:model Product
type Product struct {
	Sku              string   `json:"sku" validate:"empty=false"`
	Name             string   `json:"name" validate:"empty=false"`
	Brand            string   `json:"brand" validate:"empty=false"`
	Size             string   `json:"size,omitempty"`
	Price            string   `json:"price" validate:"empty=false"`
	PrincipalImage   string   `json:"principalImage" validate:"empty=false"`
	AdditionalImages []string `json:"additionalImages,omitempty"`
}
