package model

type Product struct {
	Sku              string   `json:"sku"`
	Name             string   `json:"name"`
	Brand            string   `json:"brand"`
	Size             string   `json:"size"`
	Price            string   `json:"price"`
	PrincipalImage   string   `json:"principalImage"`
	AdditionalImages []string `json:"additionalImages"`
}
