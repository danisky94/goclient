# \PetroliferaApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PetroliferasPetroliferaAnularVendaPost**](PetroliferaApi.md#PetroliferasPetroliferaAnularVendaPost) | **Post** /Petroliferas/Petrolifera/AnularVenda | 
[**PetroliferasPetroliferaCriarVendaPost**](PetroliferaApi.md#PetroliferasPetroliferaCriarVendaPost) | **Post** /Petroliferas/Petrolifera/CriarVenda | 
[**PetroliferasPetroliferaDownloadImageGet**](PetroliferaApi.md#PetroliferasPetroliferaDownloadImageGet) | **Get** /Petroliferas/Petrolifera/DownloadImage | 
[**PetroliferasPetroliferaGetDetalheLocalGet**](PetroliferaApi.md#PetroliferasPetroliferaGetDetalheLocalGet) | **Get** /Petroliferas/Petrolifera/GetDetalheLocal | 
[**PetroliferasPetroliferaGetGvmsGet**](PetroliferaApi.md#PetroliferasPetroliferaGetGvmsGet) | **Get** /Petroliferas/Petrolifera/GetGvms | 
[**PetroliferasPetroliferaGetLocaisGet**](PetroliferaApi.md#PetroliferasPetroliferaGetLocaisGet) | **Get** /Petroliferas/Petrolifera/GetLocais | 
[**PetroliferasPetroliferaGetProdutoDetalheGet**](PetroliferaApi.md#PetroliferasPetroliferaGetProdutoDetalheGet) | **Get** /Petroliferas/Petrolifera/GetProdutoDetalhe | 
[**PetroliferasPetroliferaGetProdutosByTipoGet**](PetroliferaApi.md#PetroliferasPetroliferaGetProdutosByTipoGet) | **Get** /Petroliferas/Petrolifera/GetProdutosByTipo | 
[**PetroliferasPetroliferaGetProdutosGet**](PetroliferaApi.md#PetroliferasPetroliferaGetProdutosGet) | **Get** /Petroliferas/Petrolifera/GetProdutos | 



## PetroliferasPetroliferaAnularVendaPost

> Resultado PetroliferasPetroliferaAnularVendaPost(ctx).Token(token).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    token := "token_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PetroliferaApi.PetroliferasPetroliferaAnularVendaPost(context.Background()).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetroliferaApi.PetroliferasPetroliferaAnularVendaPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PetroliferasPetroliferaAnularVendaPost`: Resultado
    fmt.Fprintf(os.Stdout, "Response from `PetroliferaApi.PetroliferasPetroliferaAnularVendaPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPetroliferasPetroliferaAnularVendaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** |  | 

### Return type

[**Resultado**](Resultado.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PetroliferasPetroliferaCriarVendaPost

> VendaPetroliferaListResultado PetroliferasPetroliferaCriarVendaPost(ctx).VendaPetrolifera(vendaPetrolifera).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    vendaPetrolifera := []openapiclient.VendaPetrolifera{*openapiclient.NewVendaPetrolifera()} // []VendaPetrolifera |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PetroliferaApi.PetroliferasPetroliferaCriarVendaPost(context.Background()).VendaPetrolifera(vendaPetrolifera).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetroliferaApi.PetroliferasPetroliferaCriarVendaPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PetroliferasPetroliferaCriarVendaPost`: VendaPetroliferaListResultado
    fmt.Fprintf(os.Stdout, "Response from `PetroliferaApi.PetroliferasPetroliferaCriarVendaPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPetroliferasPetroliferaCriarVendaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vendaPetrolifera** | [**[]VendaPetrolifera**](VendaPetrolifera.md) |  | 

### Return type

[**VendaPetroliferaListResultado**](VendaPetroliferaListResultado.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PetroliferasPetroliferaDownloadImageGet

> PetroliferasPetroliferaDownloadImageGet(ctx).RefImagem(refImagem).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    refImagem := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PetroliferaApi.PetroliferasPetroliferaDownloadImageGet(context.Background()).RefImagem(refImagem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetroliferaApi.PetroliferasPetroliferaDownloadImageGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPetroliferasPetroliferaDownloadImageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refImagem** | **int64** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PetroliferasPetroliferaGetDetalheLocalGet

> LocalDTO PetroliferasPetroliferaGetDetalheLocalGet(ctx).RefProduto(refProduto).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    refProduto := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PetroliferaApi.PetroliferasPetroliferaGetDetalheLocalGet(context.Background()).RefProduto(refProduto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetroliferaApi.PetroliferasPetroliferaGetDetalheLocalGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PetroliferasPetroliferaGetDetalheLocalGet`: LocalDTO
    fmt.Fprintf(os.Stdout, "Response from `PetroliferaApi.PetroliferasPetroliferaGetDetalheLocalGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPetroliferasPetroliferaGetDetalheLocalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refProduto** | **int64** |  | 

### Return type

[**LocalDTO**](LocalDTO.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PetroliferasPetroliferaGetGvmsGet

> GvmDTOListResultado PetroliferasPetroliferaGetGvmsGet(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PetroliferaApi.PetroliferasPetroliferaGetGvmsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetroliferaApi.PetroliferasPetroliferaGetGvmsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PetroliferasPetroliferaGetGvmsGet`: GvmDTOListResultado
    fmt.Fprintf(os.Stdout, "Response from `PetroliferaApi.PetroliferasPetroliferaGetGvmsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPetroliferasPetroliferaGetGvmsGetRequest struct via the builder pattern


### Return type

[**GvmDTOListResultado**](GvmDTOListResultado.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PetroliferasPetroliferaGetLocaisGet

> LocalDTOListResultado PetroliferasPetroliferaGetLocaisGet(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PetroliferaApi.PetroliferasPetroliferaGetLocaisGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetroliferaApi.PetroliferasPetroliferaGetLocaisGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PetroliferasPetroliferaGetLocaisGet`: LocalDTOListResultado
    fmt.Fprintf(os.Stdout, "Response from `PetroliferaApi.PetroliferasPetroliferaGetLocaisGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPetroliferasPetroliferaGetLocaisGetRequest struct via the builder pattern


### Return type

[**LocalDTOListResultado**](LocalDTOListResultado.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PetroliferasPetroliferaGetProdutoDetalheGet

> ProdutoDTO PetroliferasPetroliferaGetProdutoDetalheGet(ctx).RefProduto(refProduto).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    refProduto := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PetroliferaApi.PetroliferasPetroliferaGetProdutoDetalheGet(context.Background()).RefProduto(refProduto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetroliferaApi.PetroliferasPetroliferaGetProdutoDetalheGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PetroliferasPetroliferaGetProdutoDetalheGet`: ProdutoDTO
    fmt.Fprintf(os.Stdout, "Response from `PetroliferaApi.PetroliferasPetroliferaGetProdutoDetalheGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPetroliferasPetroliferaGetProdutoDetalheGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refProduto** | **int64** |  | 

### Return type

[**ProdutoDTO**](ProdutoDTO.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PetroliferasPetroliferaGetProdutosByTipoGet

> ProdutoDTOListResultado PetroliferasPetroliferaGetProdutosByTipoGet(ctx).Tipo(tipo).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tipo := openapiclient.TipoProduto(0) // TipoProduto |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PetroliferaApi.PetroliferasPetroliferaGetProdutosByTipoGet(context.Background()).Tipo(tipo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetroliferaApi.PetroliferasPetroliferaGetProdutosByTipoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PetroliferasPetroliferaGetProdutosByTipoGet`: ProdutoDTOListResultado
    fmt.Fprintf(os.Stdout, "Response from `PetroliferaApi.PetroliferasPetroliferaGetProdutosByTipoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPetroliferasPetroliferaGetProdutosByTipoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tipo** | [**TipoProduto**](TipoProduto.md) |  | 

### Return type

[**ProdutoDTOListResultado**](ProdutoDTOListResultado.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PetroliferasPetroliferaGetProdutosGet

> ProdutoDTOListResultado PetroliferasPetroliferaGetProdutosGet(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PetroliferaApi.PetroliferasPetroliferaGetProdutosGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetroliferaApi.PetroliferasPetroliferaGetProdutosGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PetroliferasPetroliferaGetProdutosGet`: ProdutoDTOListResultado
    fmt.Fprintf(os.Stdout, "Response from `PetroliferaApi.PetroliferasPetroliferaGetProdutosGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPetroliferasPetroliferaGetProdutosGetRequest struct via the builder pattern


### Return type

[**ProdutoDTOListResultado**](ProdutoDTOListResultado.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

