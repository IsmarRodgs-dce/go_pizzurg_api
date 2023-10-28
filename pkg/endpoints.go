package pkg

import (
	"context"
	"pizzurg/pkg/models"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints contém os endpoints sua aplicação
type Endpoints struct {
	PostProductEndpoint                endpoint.Endpoint
	PostProductVariationEndpoint       endpoint.Endpoint
	GetProductsEndpoint                endpoint.Endpoint
	GetProductByIdEndpoint             endpoint.Endpoint
	UpdateProductEndpoint              endpoint.Endpoint
	UpdateProductVariationEndpoint     endpoint.Endpoint
	DeleteProductByIdEndpoint          endpoint.Endpoint
	DeleteProductVariationByIdEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(pse ProductService) Endpoints {
	return Endpoints{
		PostProductEndpoint:                MakePostProductEndpoint(pse),
		PostProductVariationEndpoint:       MakePostProductVariationEndpoint(pse),
		GetProductsEndpoint:                MakeGetProductsEndpoint(pse),
		GetProductByIdEndpoint:             MakeGetProductByIdEndpoint(pse),
		UpdateProductEndpoint:              MakeUpdateProductEndpoint(pse),
		UpdateProductVariationEndpoint:     MakeUpdateProductVariationEndpoint(pse),
		DeleteProductByIdEndpoint:          MakeDeleteProductByIdEndpoint(pse),
		DeleteProductVariationByIdEndpoint: MakeDeleteProductVariationByIdEndpoint(pse),
	}
}

func MakePostProductEndpoint(pse ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.CreateProductDto)
		response, err := pse.CreateProduct(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}
func MakePostProductVariationEndpoint(pse ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.CreateProductVariationDto)
		response, err := pse.CreateProductVariation(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}

func MakeGetProductsEndpoint(pse ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response, err := pse.GetProducts(ctx)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}

func MakeGetProductByIdEndpoint(pse ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(int64)
		response, err := pse.GetProductById(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}

func MakeUpdateProductEndpoint(pse ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response, err := pse.UpdateProduct(ctx, request.ProductID, req.UpdateData)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}

func MakeUpdateProductVariationEndpoint(pse ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.UpdateProductVariationDto)
		response, err := pse.UpdateProductVariation(ctx, req.ProductID, req.ProductVariationID, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}

func MakeDeleteProductByIdEndpoint(pse ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(int64)
		err := pse.DeleteProductById(ctx, req)
		if err != nil {
			return nil, err
		}
		return "Produto excluído com sucesso", nil
	}
}

func MakeDeleteProductVariationByIdEndpoint(pse ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.DeleteProductVariationDto)
		err := pse.DeleteProductVariationById(ctx, req.ProductID, req.ProductVariationID)
		if err != nil {
			return nil, err
		}
		return "Variação de produto excluída com sucesso", nil
	}
}

func (e Endpoints) PostProduct(ctx context.Context, createProductDto models.CreateProductDto) (models.RecoveryProductDto, error) {
	request := postProductRequest{CreateProductDto: createProductDto}
	response, err := e.PostProductEndpoint(ctx, request)
	if err != nil {
		return models.RecoveryProductDto{}, err
	}
	resp := response.(postProductResponse)
	return resp.RecoveryProductDto, resp.Err
}

// PostProfile implements Service. Primarily useful in a client.
func (e Endpoints) PostProfile(ctx context.Context, p Profile) error {
	request := postProfileRequest{Profile: p}
	response, err := e.PostProfileEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(postProfileResponse)
	return resp.Err
}

// GetProfile implements Service. Primarily useful in a client.
func (e Endpoints) GetProfile(ctx context.Context, id string) (Profile, error) {
	request := getProfileRequest{ID: id}
	response, err := e.GetProfileEndpoint(ctx, request)
	if err != nil {
		return Profile{}, err
	}
	resp := response.(getProfileResponse)
	return resp.Profile, resp.Err
}

// PutProfile implements Service. Primarily useful in a client.
func (e Endpoints) PutProfile(ctx context.Context, id string, p Profile) error {
	request := putProfileRequest{ID: id, Profile: p}
	response, err := e.PutProfileEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(putProfileResponse)
	return resp.Err
}

// PatchProfile implements Service. Primarily useful in a client.
func (e Endpoints) PatchProfile(ctx context.Context, id string, p Profile) error {
	request := patchProfileRequest{ID: id, Profile: p}
	response, err := e.PatchProfileEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(patchProfileResponse)
	return resp.Err
}

// DeleteProfile implements Service. Primarily useful in a client.
func (e Endpoints) DeleteProfile(ctx context.Context, id string) error {
	request := deleteProfileRequest{ID: id}
	response, err := e.DeleteProfileEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(deleteProfileResponse)
	return resp.Err
}

// GetAddresses implements Service. Primarily useful in a client.
func (e Endpoints) GetAddresses(ctx context.Context, profileID string) ([]Address, error) {
	request := getAddressesRequest{ProfileID: profileID}
	response, err := e.GetAddressesEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	resp := response.(getAddressesResponse)
	return resp.Addresses, resp.Err
}

// GetAddress implements Service. Primarily useful in a client.
func (e Endpoints) GetAddress(ctx context.Context, profileID string, addressID string) (Address, error) {
	request := getAddressRequest{ProfileID: profileID, AddressID: addressID}
	response, err := e.GetAddressEndpoint(ctx, request)
	if err != nil {
		return Address{}, err
	}
	resp := response.(getAddressResponse)
	return resp.Address, resp.Err
}

// PostAddress implements Service. Primarily useful in a client.
func (e Endpoints) PostAddress(ctx context.Context, profileID string, a Address) error {
	request := postAddressRequest{ProfileID: profileID, Address: a}
	response, err := e.PostAddressEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(postAddressResponse)
	return resp.Err
}

// DeleteAddress implements Service. Primarily useful in a client.
func (e Endpoints) DeleteAddress(ctx context.Context, profileID string, addressID string) error {
	request := deleteAddressRequest{ProfileID: profileID, AddressID: addressID}
	response, err := e.DeleteAddressEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(deleteAddressResponse)
	return resp.Err
}
