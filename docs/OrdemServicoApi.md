# \OrdemServicoApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GvmsOrdemServicoAbrirOrdemServicoGet**](OrdemServicoApi.md#GvmsOrdemServicoAbrirOrdemServicoGet) | **Get** /Gvms/OrdemServico/AbrirOrdemServico | 
[**GvmsOrdemServicoDepositoPost**](OrdemServicoApi.md#GvmsOrdemServicoDepositoPost) | **Post** /Gvms/OrdemServico/Deposito | 
[**GvmsOrdemServicoFinalizarOrdemServicoPost**](OrdemServicoApi.md#GvmsOrdemServicoFinalizarOrdemServicoPost) | **Post** /Gvms/OrdemServico/FinalizarOrdemServico | 
[**GvmsOrdemServicoRecolhaPost**](OrdemServicoApi.md#GvmsOrdemServicoRecolhaPost) | **Post** /Gvms/OrdemServico/Recolha | 



## GvmsOrdemServicoAbrirOrdemServicoGet

> AberturaOrdemServicoResultado GvmsOrdemServicoAbrirOrdemServicoGet(ctx).Token(token).Execute()



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
    resp, r, err := api_client.OrdemServicoApi.GvmsOrdemServicoAbrirOrdemServicoGet(context.Background()).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdemServicoApi.GvmsOrdemServicoAbrirOrdemServicoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GvmsOrdemServicoAbrirOrdemServicoGet`: AberturaOrdemServicoResultado
    fmt.Fprintf(os.Stdout, "Response from `OrdemServicoApi.GvmsOrdemServicoAbrirOrdemServicoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGvmsOrdemServicoAbrirOrdemServicoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** |  | 

### Return type

[**AberturaOrdemServicoResultado**](AberturaOrdemServicoResultado.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GvmsOrdemServicoDepositoPost

> Resultado GvmsOrdemServicoDepositoPost(ctx).RefOrdemServico(refOrdemServico).RefCompartimento(refCompartimento).RefProduto(refProduto).Execute()



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
    refOrdemServico := int64(789) // int64 |  (optional)
    refCompartimento := int64(789) // int64 |  (optional)
    refProduto := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdemServicoApi.GvmsOrdemServicoDepositoPost(context.Background()).RefOrdemServico(refOrdemServico).RefCompartimento(refCompartimento).RefProduto(refProduto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdemServicoApi.GvmsOrdemServicoDepositoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GvmsOrdemServicoDepositoPost`: Resultado
    fmt.Fprintf(os.Stdout, "Response from `OrdemServicoApi.GvmsOrdemServicoDepositoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGvmsOrdemServicoDepositoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refOrdemServico** | **int64** |  | 
 **refCompartimento** | **int64** |  | 
 **refProduto** | **int64** |  | 

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


## GvmsOrdemServicoFinalizarOrdemServicoPost

> Resultado GvmsOrdemServicoFinalizarOrdemServicoPost(ctx).RefOrdemSercico(refOrdemSercico).Execute()



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
    refOrdemSercico := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdemServicoApi.GvmsOrdemServicoFinalizarOrdemServicoPost(context.Background()).RefOrdemSercico(refOrdemSercico).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdemServicoApi.GvmsOrdemServicoFinalizarOrdemServicoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GvmsOrdemServicoFinalizarOrdemServicoPost`: Resultado
    fmt.Fprintf(os.Stdout, "Response from `OrdemServicoApi.GvmsOrdemServicoFinalizarOrdemServicoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGvmsOrdemServicoFinalizarOrdemServicoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refOrdemSercico** | **int64** |  | 

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


## GvmsOrdemServicoRecolhaPost

> Resultado GvmsOrdemServicoRecolhaPost(ctx).RefOrdemSercico(refOrdemSercico).RequestBody(requestBody).Execute()



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
    refOrdemSercico := int64(789) // int64 |  (optional)
    requestBody := []int64{int64(123)} // []int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdemServicoApi.GvmsOrdemServicoRecolhaPost(context.Background()).RefOrdemSercico(refOrdemSercico).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdemServicoApi.GvmsOrdemServicoRecolhaPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GvmsOrdemServicoRecolhaPost`: Resultado
    fmt.Fprintf(os.Stdout, "Response from `OrdemServicoApi.GvmsOrdemServicoRecolhaPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGvmsOrdemServicoRecolhaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refOrdemSercico** | **int64** |  | 
 **requestBody** | **[]int64** |  | 

### Return type

[**Resultado**](Resultado.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

