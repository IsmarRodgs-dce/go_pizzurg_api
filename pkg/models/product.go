package models

type Category int

const (
	// Pizza representa um estado.
	Pizza Category = iota

	// Hamburguer representa um estado.
	Hamburguer
)

//Product representa um produto que será utilizado.
type Product struct {
	ID          uint64             `json:"id,omitempty"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Category    Category           `json:"category"`
	Variations  []ProductVariation `json:"productVariations"`
	Available   bool               `json:"available"`
}

//ProductVeriation representa um produto que será utilizado.
type ProductVariation struct {
	ID          uint64  `json:"id,omitempty"`
	ProductId   uint64  `json:"productId"`
	SizeName    string  `json:"sizename"`
	Description string  `json:"description"`
	Available   bool    `json:"available"`
	Price       float64 `json:"price"`
}
