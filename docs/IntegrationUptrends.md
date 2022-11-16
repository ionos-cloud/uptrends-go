# \IntegrationUptrends

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationCreateAuthorizationForIntegration**](IntegrationUptrends.md#IntegrationCreateAuthorizationForIntegration) | **Post** /Integration/{integrationGuid}/Authorizations | Create authorizations for integration If the wanted authorizations requires other authorizations, these will be added as well
[**IntegrationDeleteAuthorizationForIntegration**](IntegrationUptrends.md#IntegrationDeleteAuthorizationForIntegration) | **Delete** /Integration/{integrationGuid}/Authorizations/{authorizationGuid} | Delete integration authorization for integration
[**IntegrationGetAllIntegrations**](IntegrationUptrends.md#IntegrationGetAllIntegrations) | **Get** /Integration | Gets a list of all integrations.
[**IntegrationGetAuthorizationsOfIntegration**](IntegrationUptrends.md#IntegrationGetAuthorizationsOfIntegration) | **Get** /Integration/{integrationGuid}/Authorizations | Get authorizations of integration



## IntegrationCreateAuthorizationForIntegration

> []IntegrationAuthorization IntegrationCreateAuthorizationForIntegration(ctx, integrationGuid).IntegrationAuthorization(integrationAuthorization).Execute()

Create authorizations for integration If the wanted authorizations requires other authorizations, these will be added as well

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
    integrationGuid := "integrationGuid_example" // string | The integration GUID
    integrationAuthorization := *openapiclient.NewIntegrationAuthorization(map[string][]openapiclient.IntegrationAuthorizationType{ ... }) // IntegrationAuthorization | Authorization to add (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationUptrends.IntegrationCreateAuthorizationForIntegration(context.Background(), integrationGuid).IntegrationAuthorization(integrationAuthorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationUptrends.IntegrationCreateAuthorizationForIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationCreateAuthorizationForIntegration`: []IntegrationAuthorization
    fmt.Fprintf(os.Stdout, "Response from `IntegrationUptrends.IntegrationCreateAuthorizationForIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationGuid** | **string** | The integration GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationCreateAuthorizationForIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationAuthorization** | [**IntegrationAuthorization**](IntegrationAuthorization.md) | Authorization to add | 

### Return type

[**[]IntegrationAuthorization**](IntegrationAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationDeleteAuthorizationForIntegration

> IntegrationDeleteAuthorizationForIntegration(ctx, integrationGuid, authorizationGuid).Execute()

Delete integration authorization for integration

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
    integrationGuid := "integrationGuid_example" // string | The integration GUID
    authorizationGuid := "authorizationGuid_example" // string | The authorization GUID that needs to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationUptrends.IntegrationDeleteAuthorizationForIntegration(context.Background(), integrationGuid, authorizationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationUptrends.IntegrationDeleteAuthorizationForIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationGuid** | **string** | The integration GUID | 
**authorizationGuid** | **string** | The authorization GUID that needs to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationDeleteAuthorizationForIntegrationRequest struct via the builder pattern


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


## IntegrationGetAllIntegrations

> []Integration IntegrationGetAllIntegrations(ctx).Execute()

Gets a list of all integrations.

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
    resp, r, err := apiClient.IntegrationUptrends.IntegrationGetAllIntegrations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationUptrends.IntegrationGetAllIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationGetAllIntegrations`: []Integration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationUptrends.IntegrationGetAllIntegrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationGetAllIntegrationsRequest struct via the builder pattern


### Return type

[**[]Integration**](Integration.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationGetAuthorizationsOfIntegration

> []IntegrationAuthorization IntegrationGetAuthorizationsOfIntegration(ctx, integrationGuid).Execute()

Get authorizations of integration

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
    integrationGuid := "integrationGuid_example" // string | The intregration GUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationUptrends.IntegrationGetAuthorizationsOfIntegration(context.Background(), integrationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationUptrends.IntegrationGetAuthorizationsOfIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationGetAuthorizationsOfIntegration`: []IntegrationAuthorization
    fmt.Fprintf(os.Stdout, "Response from `IntegrationUptrends.IntegrationGetAuthorizationsOfIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationGuid** | **string** | The intregration GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationGetAuthorizationsOfIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IntegrationAuthorization**](IntegrationAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

