package models

// CreateProductDto representa a criação de um produto
type CreateProductDto struct {
	Name        string
	Description string
	Category    Category
	Variations  []CreateProductVariationDto
	Available   bool
}

// CreateProductVariationDto representa a criação de uma variação de produto
type CreateProductVariationDto struct {
	SizeName    string
	Description string
	Price       float64
	Avaliable   bool
}

type RecoveryProductDto struct {
	Id          uint64
	Name        string
	Description string
	Category    Category
	Variations  []RecoveryProductVariationDto
	Available   bool
}

type RecoveryProductVariationDto struct {
	Id          uint64
	SizeName    string
	Description string
	Price       float64
	Available   bool
}

type UpdateProductDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
}

type UpdateProductVariationDto struct {
	SizeName    string
	Description string
	Price       float64
	Available   bool
}
