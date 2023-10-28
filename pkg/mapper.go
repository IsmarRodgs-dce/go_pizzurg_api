package pkg

import (
	"pizzurg/pkg/models"
)

// Função para mapear um objeto Product para um objeto RecoveryProductDto.
func mapProductToRecoveryProductDto(product models.Product) *models.RecoveryProductDto {
	recoveryProductDto := models.RecoveryProductDto{
		Name:        product.Name,
		Description: product.Description,
		Category:    product.Category,
		Available:   product.Available,
	}

	for _, variation := range product.Variations {
		recoveryProductDto.ProductVariations = append(recoveryProductDto.ProductVariations, *mapProductVariationToRecoveryProductVariationDto(variation))
	}

	return &recoveryProductDto
}
func mapProductsToRecoveryProductsDto(products []models.Product) *[]models.RecoveryProductDto {
	recoveryProductsDto := []models.RecoveryProductDto{}
	for _, product := range products {
		recoveryProductDto := models.RecoveryProductDto{
			Name:        product.Name,
			Description: product.Description,
			Category:    product.Category,
			Available:   product.Available,
		}

		for _, variation := range product.Variations {
			recoveryProductDto.ProductVariations = append(recoveryProductDto.ProductVariations, *mapProductVariationToRecoveryProductVariationDto(variation))
		}
		recoveryProductsDto = append(recoveryProductsDto, recoveryProductDto)

	}

	return &recoveryProductsDto
}

// Função para mapear uma lista de objetos ProductVariation para uma lista de RecoveryProductVariationDto.
func mapProductsVariationToRecoveryProductsVariationDto(productVariations []models.ProductVariation) *[]models.RecoveryProductVariationDto {
	var recoveryVariations []models.RecoveryProductVariationDto

	for _, variation := range productVariations {
		recoveryVariations = append(recoveryVariations, *mapProductVariationToRecoveryProductVariationDto(variation))
	}

	return &recoveryVariations
}

// Função para mapear um objeto ProductVariation para um objeto RecoveryProductVariationDto.
func mapProductVariationToRecoveryProductVariationDto(productVariation models.ProductVariation) *models.RecoveryProductVariationDto {
	recoveryVariation := models.RecoveryProductVariationDto{
		SizeName:    productVariation.SizeName,
		Description: productVariation.Description,
		Price:       productVariation.Price,
		Available:   productVariation.Available,
	}

	return &recoveryVariation
}

// Função para mapear um objeto RecoveryProductVariationDto de volta para um objeto ProductVariation.
func mapCreateProductVariationDtoToProductVariation(productId uint64, productVariationDto models.CreateProductVariationDto) *models.ProductVariation {
	productVariation := models.ProductVariation{
		ProductId:   productId,
		SizeName:    productVariationDto.SizeName,
		Description: productVariationDto.Description,
		Price:       productVariationDto.Price,
		Available:   productVariationDto.Avaliable,
	}

	return &productVariation
}

func mapCreateProductDtoToProduct(createProductDto models.CreateProductDto) *models.Product {

	var productVariations []models.ProductVariation
	for _, productVariationDto := range createProductDto.Variations {
		variation := &mapCreateProductVariationDtoToProductVariation(productVariationDto)
		productVariations = append(productVariations, variation)
	}

	product := &models.Product{
		Name:              createProductDto.Name,
		Description:       createProductDto.Description,
		Category:          createProductDto.Category,
		ProductVariations: productVariations,
		Available:         createProductDto.Available,
	}

}
