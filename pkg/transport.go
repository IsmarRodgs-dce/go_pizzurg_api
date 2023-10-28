package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"pizzurg/pkg/models"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	"github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrBadRouting = errors.New("mapeamento inconsistente entre rota e manipulador (erro de programador)")
)

// func DecodeRegisterProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
//     var req registerProductRequest
//     if err := http.DecodeJSONRequest(r, &req); err != nil {
//         return nil, err
//     }
//     return req, nil
// }

// func EncodeRegisterProductResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
//     resp := response.(registerProductResponse)
//     return http.EncodeJSONResponse(w, resp)
// }

func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	endpoints := MakeServerEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}
	//PostProductEndpoint
	r.Methods("POST").Path("/products/").Handler(httptransport.NewServer(
		endpoints.PostProductEndpoint,
		decodeCreateProductRequest,
		encodeResponse,
		options...,
	))
	//PostProductVariationEndpoint
	r.Methods("POST").Path("/products/{productId}/variation").Handler(httptransport.NewServer(
		endpoints.PostProductVariationEndpoint,
		decodeCreateProductVariationRequest,
		encodeResponse,
		options...,
	))
	// // GetProductsEndpoint
	// r.Methods("GET").Path("/products/").Handler(httptransport.NewServer(
	//     endpoints.GetProductsEndpoint,
	//     nil,
	//     encodeResponse,
	//     options...,
	// ))
	// // GetProductByIdEndpoint
	// r.Methods("GET").Path("/products/{productId}").Handler(httptransport.NewServer(
	//     endpoints.GetProductByIdEndpoint,
	//     decodeGetProductByIdRequest,
	//     encodeResponse,
	//     options...,
	// ))
	// //UpdateProductEndpoint
	// r.Methods("PATCH").Path("/products/{productId}").Handler(httptransport.NewServer(
	//     endpoints.UpdateProductEndpoint,
	//     decodeUpdateProductRequest,
	//     encodeResponse,
	//     options...,
	// ))
	// //UpdateProductVariationEndpoint
	// r.Methods("PUT").Path("/products/{productId}/variation/{productVariationId}").Handler(httptransport.NewServer(
	//     endpoints.UpdateProductVariationEndpoint,
	//     decodeUpdateProductVariationRequest,
	//     encodeResponse,
	//     options...,
	// ))

	// // Rota para o endpoint DeleteProductByIdEndpoint
	// r.Methods("DELETE").Path("/products/{productId}").Handler(httptransport.NewServer(
	//     endpoints.DeleteProductByIdEndpoint,
	//     decodeDeleteProductByIdRequest,
	//     encodeResponse,
	//     options...,
	// ))

	// // Rota para o endpoint DeleteProductVariationByIdEndpoint
	// r.Methods("DELETE").Path("/products/{productId}/variation/{productVariationId}").Handler(httptransport.NewServer(
	//     endpoints.DeleteProductVariationByIdEndpoint,
	//     decodeDeleteProductVariationByIdRequest,
	//     encodeResponse,
	//     options...,
	// ))
	return r
}

func decodeCreateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request models.CreateProductDto
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func decodeCreateProductVariationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request models.CreateProductVariationDto // Substitua models.CreateProductVariationDto pelo tipo de dados real usado em sua aplicação
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func decodeGetProductByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["productId"]
	if !ok {
		return nil, ErrBadRouting
	}
	return request{ID: id}, nil
}
func decodeUpdateProductVariationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	variationId := vars["variationId"] // Supondo que você tenha uma rota que inclua o ID da variação

	var updateVariationRequest models.UpdateProductVariationDto

	if err := json.NewDecoder(r.Body).Decode(&updateVariationRequest); err != nil {
		return nil, err
	}

	request := UpdateProductVariationRequestWithIDs{
		ProductID:   productId,
		VariationID: variationId,
		UpdateData:  updateVariationRequest,
	}

	return request, nil
}
func decodeDeleteProductByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	request := DeleteProductRequest{
		ProductID: productId,
	}
	return request, nil
}
func decodeDeleteProductVariationByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	variationId := vars["variationId"]
	request := DeleteProductVariationRequest{
		ProductID:   productId,
		VariationID: variationId,
	}
	return request, nil
}
func decodeUpdateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	var updateProductRequest models.UpdateProductDto

	if err := json.NewDecoder(r.Body).Decode(&updateProductRequest); err != nil {
		return nil, err
	}
	request := UpdateProductRequestWithID{
		ProductID:  productId,
		UpdateData: updateProductRequest,
	}
	return request, nil
}

type errorer interface {
	error() error
}

func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrAlreadyExists, ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
