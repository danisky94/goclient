# \GvmApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GvmsGvmAlterarEstadoPost**](GvmApi.md#GvmsGvmAlterarEstadoPost) | **Post** /Gvms/Gvm/AlterarEstado | 
[**GvmsGvmDownloadImageGet**](GvmApi.md#GvmsGvmDownloadImageGet) | **Get** /Gvms/Gvm/DownloadImage | 
[**GvmsGvmFinalizarVendaGet**](GvmApi.md#GvmsGvmFinalizarVendaGet) | **Get** /Gvms/Gvm/FinalizarVenda | 
[**GvmsGvmGetGvmGet**](GvmApi.md#GvmsGvmGetGvmGet) | **Get** /Gvms/Gvm/GetGvm | 
[**GvmsGvmGetJwtTestePost**](GvmApi.md#GvmsGvmGetJwtTestePost) | **Post** /Gvms/Gvm/GetJwtTeste | 
[**GvmsGvmGetVendaGet**](GvmApi.md#GvmsGvmGetVendaGet) | **Get** /Gvms/Gvm/GetVenda | 
[**GvmsGvmLoginPost**](GvmApi.md#GvmsGvmLoginPost) | **Post** /Gvms/Gvm/Login | 
[**GvmsGvmRegistarEventoPost**](GvmApi.md#GvmsGvmRegistarEventoPost) | **Post** /Gvms/Gvm/RegistarEvento | 
[**GvmsGvmRegistarRaspberryPiPost**](GvmApi.md#GvmsGvmRegistarRaspberryPiPost) | **Post** /Gvms/Gvm/RegistarRaspberryPi | 



## GvmsGvmAlterarEstadoPost

> Resultado GvmsGvmAlterarEstadoPost(ctx).RefCompartimento(refCompartimento).Estado(estado).Execute()



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
    refCompartimento := int64(789) // int64 |  (optional)
    estado := openapiclient.EstadosCompartimentoEstado(0) // EstadosCompartimentoEstado |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GvmApi.GvmsGvmAlterarEstadoPost(context.Background()).RefCompartimento(refCompartimento).Estado(estado).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GvmApi.GvmsGvmAlterarEstadoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GvmsGvmAlterarEstadoPost`: Resultado
    fmt.Fprintf(os.Stdout, "Response from `GvmApi.GvmsGvmAlterarEstadoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGvmsGvmAlterarEstadoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refCompartimento** | **int64** |  | 
 **estado** | [**EstadosCompartimentoEstado**](EstadosCompartimentoEstado.md) |  | 

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


## GvmsGvmDownloadImageGet

> GvmsGvmDownloadImageGet(ctx).RefProduto(refProduto).Execute()



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
    resp, r, err := api_client.GvmApi.GvmsGvmDownloadImageGet(context.Background()).RefProduto(refProduto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GvmApi.GvmsGvmDownloadImageGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGvmsGvmDownloadImageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refProduto** | **int64** |  | 

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


## GvmsGvmFinalizarVendaGet

> Resultado GvmsGvmFinalizarVendaGet(ctx).RefVenda(refVenda).RefCompartimentoProduto(refCompartimentoProduto).RefCompartimentoVasilhame(refCompartimentoVasilhame).Execute()



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
    refVenda := int64(789) // int64 |  (optional)
    refCompartimentoProduto := int64(789) // int64 |  (optional)
    refCompartimentoVasilhame := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GvmApi.GvmsGvmFinalizarVendaGet(context.Background()).RefVenda(refVenda).RefCompartimentoProduto(refCompartimentoProduto).RefCompartimentoVasilhame(refCompartimentoVasilhame).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GvmApi.GvmsGvmFinalizarVendaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GvmsGvmFinalizarVendaGet`: Resultado
    fmt.Fprintf(os.Stdout, "Response from `GvmApi.GvmsGvmFinalizarVendaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGvmsGvmFinalizarVendaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refVenda** | **int64** |  | 
 **refCompartimentoProduto** | **int64** |  | 
 **refCompartimentoVasilhame** | **int64** |  | 

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


## GvmsGvmGetGvmGet

> GvmDTOResultado GvmsGvmGetGvmGet(ctx).IdentificadorPI(identificadorPI).Execute()



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
    identificadorPI := "identificadorPI_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GvmApi.GvmsGvmGetGvmGet(context.Background()).IdentificadorPI(identificadorPI).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GvmApi.GvmsGvmGetGvmGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GvmsGvmGetGvmGet`: GvmDTOResultado
    fmt.Fprintf(os.Stdout, "Response from `GvmApi.GvmsGvmGetGvmGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGvmsGvmGetGvmGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identificadorPI** | **string** |  | 

### Return type

[**GvmDTOResultado**](GvmDTOResultado.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GvmsGvmGetJwtTestePost

> GvmResultado GvmsGvmGetJwtTestePost(ctx).Execute()



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
    resp, r, err := api_client.GvmApi.GvmsGvmGetJwtTestePost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GvmApi.GvmsGvmGetJwtTestePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GvmsGvmGetJwtTestePost`: GvmResultado
    fmt.Fprintf(os.Stdout, "Response from `GvmApi.GvmsGvmGetJwtTestePost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGvmsGvmGetJwtTestePostRequest struct via the builder pattern


### Return type

[**GvmResultado**](GvmResultado.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GvmsGvmGetVendaGet

> CustomVendaResultado GvmsGvmGetVendaGet(ctx).TokenVenda(tokenVenda).Execute()



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
    tokenVenda := "tokenVenda_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GvmApi.GvmsGvmGetVendaGet(context.Background()).TokenVenda(tokenVenda).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GvmApi.GvmsGvmGetVendaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GvmsGvmGetVendaGet`: CustomVendaResultado
    fmt.Fprintf(os.Stdout, "Response from `GvmApi.GvmsGvmGetVendaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGvmsGvmGetVendaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenVenda** | **string** |  | 

### Return type

[**CustomVendaResultado**](CustomVendaResultado.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GvmsGvmLoginPost

> StringResultado GvmsGvmLoginPost(ctx).IdendificadorPi(idendificadorPi).Password(password).Execute()



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
    idendificadorPi := "idendificadorPi_example" // string |  (optional)
    password := "password_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GvmApi.GvmsGvmLoginPost(context.Background()).IdendificadorPi(idendificadorPi).Password(password).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GvmApi.GvmsGvmLoginPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GvmsGvmLoginPost`: StringResultado
    fmt.Fprintf(os.Stdout, "Response from `GvmApi.GvmsGvmLoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGvmsGvmLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idendificadorPi** | **string** |  | 
 **password** | **string** |  | 

### Return type

[**StringResultado**](StringResultado.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GvmsGvmRegistarEventoPost

> EventoDTOResultado GvmsGvmRegistarEventoPost(ctx).EventoDTO(eventoDTO).Execute()



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
    eventoDTO := *openapiclient.NewEventoDTO() // EventoDTO |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GvmApi.GvmsGvmRegistarEventoPost(context.Background()).EventoDTO(eventoDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GvmApi.GvmsGvmRegistarEventoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GvmsGvmRegistarEventoPost`: EventoDTOResultado
    fmt.Fprintf(os.Stdout, "Response from `GvmApi.GvmsGvmRegistarEventoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGvmsGvmRegistarEventoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventoDTO** | [**EventoDTO**](EventoDTO.md) |  | 

### Return type

[**EventoDTOResultado**](EventoDTOResultado.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GvmsGvmRegistarRaspberryPiPost

> RaspberryPiDTOResultado GvmsGvmRegistarRaspberryPiPost(ctx).Identificador(identificador).Execute()



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
    identificador := "identificador_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GvmApi.GvmsGvmRegistarRaspberryPiPost(context.Background()).Identificador(identificador).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GvmApi.GvmsGvmRegistarRaspberryPiPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GvmsGvmRegistarRaspberryPiPost`: RaspberryPiDTOResultado
    fmt.Fprintf(os.Stdout, "Response from `GvmApi.GvmsGvmRegistarRaspberryPiPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGvmsGvmRegistarRaspberryPiPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identificador** | **string** |  | 

### Return type

[**RaspberryPiDTOResultado**](RaspberryPiDTOResultado.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

