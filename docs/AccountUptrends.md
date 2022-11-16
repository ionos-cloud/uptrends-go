# \AccountUptrends

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountGetAccountStatistics**](AccountUptrends.md#AccountGetAccountStatistics) | **Get** /Account | Returns the account statistics.
[**AccountGetSettings**](AccountUptrends.md#AccountGetSettings) | **Get** /Account/Settings | Returns account general settings.



## AccountGetAccountStatistics

> AccountStatistics AccountGetAccountStatistics(ctx).Execute()

Returns the account statistics.

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
    resp, r, err := apiClient.AccountUptrends.AccountGetAccountStatistics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUptrends.AccountGetAccountStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountGetAccountStatistics`: AccountStatistics
    fmt.Fprintf(os.Stdout, "Response from `AccountUptrends.AccountGetAccountStatistics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountGetAccountStatisticsRequest struct via the builder pattern


### Return type

[**AccountStatistics**](AccountStatistics.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountGetSettings

> AccountSettings AccountGetSettings(ctx).Execute()

Returns account general settings.

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
    resp, r, err := apiClient.AccountUptrends.AccountGetSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUptrends.AccountGetSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountGetSettings`: AccountSettings
    fmt.Fprintf(os.Stdout, "Response from `AccountUptrends.AccountGetSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountGetSettingsRequest struct via the builder pattern


### Return type

[**AccountSettings**](AccountSettings.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

