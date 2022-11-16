# \SLAUptrends

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlaCreateSla**](SLAUptrends.md#SlaCreateSla) | **Post** /Sla | Creates a new SLA.
[**SlaDeleteExclusionPeriod**](SLAUptrends.md#SlaDeleteExclusionPeriod) | **Delete** /Sla/{slaGuid}/ExclusionPeriod/{exclusionPeriodId} | Deletes the specified exclusion period for the specified SLA.
[**SlaDeleteSla**](SLAUptrends.md#SlaDeleteSla) | **Delete** /Sla/{slaGuid} | Deletes the specified SLA.
[**SlaGetExclusionPeriod**](SLAUptrends.md#SlaGetExclusionPeriod) | **Get** /Sla/{slaGuid}/ExclusionPeriod/{exclusionPeriodId} | Gets the specified exclusion period for the specified SLA.
[**SlaGetExclusionPeriods**](SLAUptrends.md#SlaGetExclusionPeriods) | **Get** /Sla/{slaGuid}/ExclusionPeriod | Gets a list of all exclusion periods for the specified SLA.
[**SlaGetSla**](SLAUptrends.md#SlaGetSla) | **Get** /Sla/{slaGuid} | Gets the specified SLA definition.
[**SlaGetSlas**](SLAUptrends.md#SlaGetSlas) | **Get** /Sla | Gets a list of all SLA definitions.
[**SlaPatchExclusionPeriod**](SLAUptrends.md#SlaPatchExclusionPeriod) | **Patch** /Sla/{slaGuid}/ExclusionPeriod/{exclusionPeriodId} | Partially updates the specified exclusion period for the specified SLA.
[**SlaPatchSla**](SLAUptrends.md#SlaPatchSla) | **Patch** /Sla/{slaGuid} | Partially updates the definition of the specified SLA.
[**SlaPostExclusionPeriod**](SLAUptrends.md#SlaPostExclusionPeriod) | **Post** /Sla/{slaGuid}/ExclusionPeriod | Creates a new exclusion period for the specified SLA.
[**SlaPutExclusionPeriod**](SLAUptrends.md#SlaPutExclusionPeriod) | **Put** /Sla/{slaGuid}/ExclusionPeriod/{exclusionPeriodId} | Updates the specified exclusion period for the specified SLA.
[**SlaPutSla**](SLAUptrends.md#SlaPutSla) | **Put** /Sla/{slaGuid} | Updates the definition of the specified SLA.



## SlaCreateSla

> Sla SlaCreateSla(ctx).Sla(sla).Execute()

Creates a new SLA.

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
    sla := *openapiclient.NewSla() // Sla | The complete SLA definition that should be created. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SLAUptrends.SlaCreateSla(context.Background()).Sla(sla).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLAUptrends.SlaCreateSla``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlaCreateSla`: Sla
    fmt.Fprintf(os.Stdout, "Response from `SLAUptrends.SlaCreateSla`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlaCreateSlaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sla** | [**Sla**](Sla.md) | The complete SLA definition that should be created. | 

### Return type

[**Sla**](Sla.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlaDeleteExclusionPeriod

> SlaDeleteExclusionPeriod(ctx, slaGuid, exclusionPeriodId).Execute()

Deletes the specified exclusion period for the specified SLA.

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
    slaGuid := "slaGuid_example" // string | The Guid of the sla definition.
    exclusionPeriodId := int32(56) // int32 | The id of the exclusion period.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SLAUptrends.SlaDeleteExclusionPeriod(context.Background(), slaGuid, exclusionPeriodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLAUptrends.SlaDeleteExclusionPeriod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaGuid** | **string** | The Guid of the sla definition. | 
**exclusionPeriodId** | **int32** | The id of the exclusion period. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlaDeleteExclusionPeriodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlaDeleteSla

> SlaDeleteSla(ctx, slaGuid).Execute()

Deletes the specified SLA.

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
    slaGuid := "slaGuid_example" // string | The Guid of the SLA definition that should be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SLAUptrends.SlaDeleteSla(context.Background(), slaGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLAUptrends.SlaDeleteSla``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaGuid** | **string** | The Guid of the SLA definition that should be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlaDeleteSlaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlaGetExclusionPeriod

> ExclusionPeriod SlaGetExclusionPeriod(ctx, slaGuid, exclusionPeriodId).Execute()

Gets the specified exclusion period for the specified SLA.

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
    slaGuid := "slaGuid_example" // string | The Guid of the SLA definition.
    exclusionPeriodId := int32(56) // int32 | The id of the exclusion period.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SLAUptrends.SlaGetExclusionPeriod(context.Background(), slaGuid, exclusionPeriodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLAUptrends.SlaGetExclusionPeriod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlaGetExclusionPeriod`: ExclusionPeriod
    fmt.Fprintf(os.Stdout, "Response from `SLAUptrends.SlaGetExclusionPeriod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaGuid** | **string** | The Guid of the SLA definition. | 
**exclusionPeriodId** | **int32** | The id of the exclusion period. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlaGetExclusionPeriodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExclusionPeriod**](ExclusionPeriod.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlaGetExclusionPeriods

> []ExclusionPeriod SlaGetExclusionPeriods(ctx, slaGuid).Execute()

Gets a list of all exclusion periods for the specified SLA.

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
    slaGuid := "slaGuid_example" // string | The Guid of the SLA definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SLAUptrends.SlaGetExclusionPeriods(context.Background(), slaGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLAUptrends.SlaGetExclusionPeriods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlaGetExclusionPeriods`: []ExclusionPeriod
    fmt.Fprintf(os.Stdout, "Response from `SLAUptrends.SlaGetExclusionPeriods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaGuid** | **string** | The Guid of the SLA definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlaGetExclusionPeriodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ExclusionPeriod**](ExclusionPeriod.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlaGetSla

> Sla SlaGetSla(ctx, slaGuid).Execute()

Gets the specified SLA definition.

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
    slaGuid := "slaGuid_example" // string | The Guid of the SLA definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SLAUptrends.SlaGetSla(context.Background(), slaGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLAUptrends.SlaGetSla``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlaGetSla`: Sla
    fmt.Fprintf(os.Stdout, "Response from `SLAUptrends.SlaGetSla`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaGuid** | **string** | The Guid of the SLA definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlaGetSlaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sla**](Sla.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlaGetSlas

> []Sla SlaGetSlas(ctx).Execute()

Gets a list of all SLA definitions.

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SLAUptrends.SlaGetSlas(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLAUptrends.SlaGetSlas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlaGetSlas`: []Sla
    fmt.Fprintf(os.Stdout, "Response from `SLAUptrends.SlaGetSlas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlaGetSlasRequest struct via the builder pattern


### Return type

[**[]Sla**](Sla.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlaPatchExclusionPeriod

> SlaPatchExclusionPeriod(ctx, slaGuid, exclusionPeriodId).ExclusionPeriod(exclusionPeriod).Execute()

Partially updates the specified exclusion period for the specified SLA.

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
    slaGuid := "slaGuid_example" // string | The Guid of the SLA definition.
    exclusionPeriodId := int32(56) // int32 | The id of the exclusion period.
    exclusionPeriod := *openapiclient.NewExclusionPeriod() // ExclusionPeriod | The complete definition of the exclusion period. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SLAUptrends.SlaPatchExclusionPeriod(context.Background(), slaGuid, exclusionPeriodId).ExclusionPeriod(exclusionPeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLAUptrends.SlaPatchExclusionPeriod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaGuid** | **string** | The Guid of the SLA definition. | 
**exclusionPeriodId** | **int32** | The id of the exclusion period. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlaPatchExclusionPeriodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exclusionPeriod** | [**ExclusionPeriod**](ExclusionPeriod.md) | The complete definition of the exclusion period. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlaPatchSla

> SlaPatchSla(ctx, slaGuid).Sla(sla).Execute()

Partially updates the definition of the specified SLA.



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
    slaGuid := "slaGuid_example" // string | The Guid of the SLA that should be updated.
    sla := *openapiclient.NewSla() // Sla | The partial definition for the SLA that should be updated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SLAUptrends.SlaPatchSla(context.Background(), slaGuid).Sla(sla).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLAUptrends.SlaPatchSla``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaGuid** | **string** | The Guid of the SLA that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlaPatchSlaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sla** | [**Sla**](Sla.md) | The partial definition for the SLA that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlaPostExclusionPeriod

> ExclusionPeriod SlaPostExclusionPeriod(ctx, slaGuid).ExclusionPeriod(exclusionPeriod).Execute()

Creates a new exclusion period for the specified SLA.

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
    slaGuid := "slaGuid_example" // string | The Guid of the SLA definition.
    exclusionPeriod := *openapiclient.NewExclusionPeriod() // ExclusionPeriod | The complete definition of the exclusion period. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SLAUptrends.SlaPostExclusionPeriod(context.Background(), slaGuid).ExclusionPeriod(exclusionPeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLAUptrends.SlaPostExclusionPeriod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlaPostExclusionPeriod`: ExclusionPeriod
    fmt.Fprintf(os.Stdout, "Response from `SLAUptrends.SlaPostExclusionPeriod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaGuid** | **string** | The Guid of the SLA definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlaPostExclusionPeriodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exclusionPeriod** | [**ExclusionPeriod**](ExclusionPeriod.md) | The complete definition of the exclusion period. | 

### Return type

[**ExclusionPeriod**](ExclusionPeriod.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlaPutExclusionPeriod

> SlaPutExclusionPeriod(ctx, slaGuid, exclusionPeriodId).ExclusionPeriod(exclusionPeriod).Execute()

Updates the specified exclusion period for the specified SLA.

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
    slaGuid := "slaGuid_example" // string | The Guid of the SLA definition.
    exclusionPeriodId := int32(56) // int32 | The id of the exclusion period.
    exclusionPeriod := *openapiclient.NewExclusionPeriod() // ExclusionPeriod | The complete definition of the exclusion period. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SLAUptrends.SlaPutExclusionPeriod(context.Background(), slaGuid, exclusionPeriodId).ExclusionPeriod(exclusionPeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLAUptrends.SlaPutExclusionPeriod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaGuid** | **string** | The Guid of the SLA definition. | 
**exclusionPeriodId** | **int32** | The id of the exclusion period. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlaPutExclusionPeriodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exclusionPeriod** | [**ExclusionPeriod**](ExclusionPeriod.md) | The complete definition of the exclusion period. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlaPutSla

> SlaPutSla(ctx, slaGuid).Sla(sla).Execute()

Updates the definition of the specified SLA.



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
    slaGuid := "slaGuid_example" // string | The Guid of the SLA that should be updated.
    sla := *openapiclient.NewSla() // Sla | The complete definition for the SLA that should be updated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SLAUptrends.SlaPutSla(context.Background(), slaGuid).Sla(sla).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLAUptrends.SlaPutSla``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaGuid** | **string** | The Guid of the SLA that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlaPutSlaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sla** | [**Sla**](Sla.md) | The complete definition for the SLA that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

