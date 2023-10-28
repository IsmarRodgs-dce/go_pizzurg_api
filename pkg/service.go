package pkg

import (
	"context"
	repository "pizzurg/internal/repositories"
	"pizzurg/pkg/models"
)

// service implements the product service
type productservice struct {
	repositoryProduct          repository.RepositoryProduct
	repositoryProductVariation repository.RepositoryProductVariations
}

type ProductService interface {
	CreateProduct(ctx context.Context, createProductDto models.CreateProductDto) (models.RecoveryProductDto, error)
	CreateProductVariation(ctx context.Context, productId int64, createProductVariationDto models.CreateProductVariationDto) (models.RecoveryProductDto, error)
	GetProducts(ctx context.Context) ([]models.RecoveryProductDto, error)
	GetProductById(ctx context.Context, productId int64) (models.RecoveryProductDto, error)
	UpdateProduct(ctx context.Context, productId int64, updateProductDto models.UpdateProductDto) (models.RecoveryProductDto, error)
	UpdateProductVariation(ctx context.Context, productId, productVariationId int64, updateProductVariationDto models.UpdateProductVariationDto) (models.RecoveryProductDto, error)
	DeleteProductById(ctx context.Context, productId int64) error
	DeleteProductVariationById(ctx context.Context, productId, productVariationId int64) error
}

//Em toda função será colocada um Log em seu início para demarcar, mas isso ainda pode ser usado por uma chamada de Logger do middleware.
//logger := log.With(s.logger, "method", "Create")

// NewProductService cria uma instância do ProductService
func NewProductService(repositoryProduct repository.RepositoryProduct, repositoryProductVariation repository.RepositoryProductVariations) *productservice {
	return &productservice{
		repositoryProduct:          repositoryProduct,
		repositoryProductVariation: repositoryProductVariation,
	}
}

// CreateProduct cria um novo produto
func (s *productservice) CreateProduct(ctx context.Context, createProductDto models.CreateProductDto) (*models.RecoveryProductDto, error) {
	// Converte a lista de ProductVariationDto em uma lista de ProductVariation
	var productVariations []models.ProductVariation
	for _, productVariationDto := range createProductDto.Variations {
		productVariation := mapProductVariationToRecoveryProductVariationDto(productVariationDto)
		productVariations = append(productVariations, productVariation)
	}
	// // Cria um produto
	// product := &models.Product{
	// 	Name:              createProductDto.Name,
	// 	Description:       createProductDto.Description,
	// 	Category:          strings.ToUpper(string(createProductDto.Category)),
	// 	ProductVariations: productVariations,
	// 	Available:         createProductDto.Available,
	// }
	productSaved, err := s.repositoryProduct.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	return &mapProductToRecoveryProductDto(productSaved), nil
}
func (s *productservice) CreateProductVariation(ctx context.Context, productId int64, productVariationDto CreateProductVariationDto) (*RecoveryProductDto, error) {
	productVariation := mapCreateProductVariationDtoToProductVariation(productId, productVariationDto)
	recoveryProductVariation, err := s.repositoryProduct.CreateProductVariation(ctx, productVariation)
	if err != nil {
		return nil, err
	}
	return &mapProductVariationToRecoveryProductVariationDto(recoveryProductVariation), nil
}
func (s *productservice) GetProducts(ctx context.Context) ([]RecoveryProductDto, error) {
	products, err := s.repositoryProduct.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return mapProductToRecoveryProductDto(products), nil
}
func (s *productservice) GetProductById(ctx context.Context, productId int) (*RecoveryProductDto, error) {
	product, err := s.repositoryProduct.GetProductById(ctx, productId)
	if err != nil {
		return nil, err
	}

	return &mapProductToRecoveryProductDto(product), nil
}

func (s *productService) UpdateProduct(ctx context.Context, productId int, updateProductDto UpdateProductDto) (*RecoveryProductDto, error) {
	productUpdated, err := s.repositoryProduct.UpdateProduct(ctx, productId, updateProductDto)
	if err != nil {
		return nil, err
	}
	return &mapProductToRecoveryProductDto(productUpdated), nil
}

func (r *productService) UpdateProductVariation(ctx context.Context, productId, productVariationId int, updateProductVariationDto UpdateProductVariationDto) (*RecoveryProductDto, error) {
	productVariationUpdated, err := s.repositoryProduct.UpdateProductVariation(ctx, productId, productVariationId, updateProductVariationDto)
	if err != nil {
		return nil, err
	}
	return &mapProductVariationToRecoveryProductVariationDto(productVariationUpdated), nil
}

func (s *productService) DeleteProductById(ctx context.Context, productId int) error {
	err := s.repositoryProduct.DeleteProductById(ctx, productId)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) DeleteProductVariationById(ctx context.Context, productId, productVariationId int) error {
	err := s.repositoryProduct.DeleteProductVariationById(ctx, productId, productVariationId)
	if err != nil {
		return err
	}
	return nil
}
