# \RUMUptrends

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RumGetRumMetricsForAllWebsites**](RUMUptrends.md#RumGetRumMetricsForAllWebsites) | **Get** /Rum/Website/Metrics | Returns all metrics of all RUM websites.
[**RumGetRumWebsiteMetrics**](RUMUptrends.md#RumGetRumWebsiteMetrics) | **Get** /Rum/Website/{rumWebsiteGuid}/Metrics | Returns all metrics of the specified RUM website.
[**RumGetRumWebsites**](RUMUptrends.md#RumGetRumWebsites) | **Get** /Rum/Website | Returns the definition of all RUM websites available in the account.



## RumGetRumMetricsForAllWebsites

> []RumWebsiteWithMetricValues RumGetRumMetricsForAllWebsites(ctx).WebsiteFilter(websiteFilter).Start(start).End(end).PresetPeriod(presetPeriod).Execute()

Returns all metrics of all RUM websites.

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
    websiteFilter := "websiteFilter_example" // string | A search pattern to filter for matching RUM website descriptions. Wildcards (? and *) are allowed. To filter for multiple patterns, use the | symbol as separator. (optional)
    start := TODO // interface{} | The start of a custom period (can't be used together with the PresetPeriod parameter) (optional)
    end := TODO // interface{} | The end of a custom period (optional)
    presetPeriod := "presetPeriod_example" // string | The requested time period. (optional) (default to "Last24Hours")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMUptrends.RumGetRumMetricsForAllWebsites(context.Background()).WebsiteFilter(websiteFilter).Start(start).End(end).PresetPeriod(presetPeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMUptrends.RumGetRumMetricsForAllWebsites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RumGetRumMetricsForAllWebsites`: []RumWebsiteWithMetricValues
    fmt.Fprintf(os.Stdout, "Response from `RUMUptrends.RumGetRumMetricsForAllWebsites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRumGetRumMetricsForAllWebsitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **websiteFilter** | **string** | A search pattern to filter for matching RUM website descriptions. Wildcards (? and *) are allowed. To filter for multiple patterns, use the | symbol as separator. | 
 **start** | [**interface{}**](interface{}.md) | The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | [**interface{}**](interface{}.md) | The end of a custom period | 
 **presetPeriod** | **string** | The requested time period. | [default to &quot;Last24Hours&quot;]

### Return type

[**[]RumWebsiteWithMetricValues**](RumWebsiteWithMetricValues.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RumGetRumWebsiteMetrics

> []RumMetricValues RumGetRumWebsiteMetrics(ctx, rumWebsiteGuid).Start(start).End(end).PresetPeriod(presetPeriod).Execute()

Returns all metrics of the specified RUM website.

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
    rumWebsiteGuid := "rumWebsiteGuid_example" // string | 
    start := TODO // interface{} | The start of a custom period (can't be used together with the PresetPeriod parameter) (optional)
    end := TODO // interface{} | The end of a custom period (optional)
    presetPeriod := "presetPeriod_example" // string | The requested time period. (optional) (default to "Last24Hours")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMUptrends.RumGetRumWebsiteMetrics(context.Background(), rumWebsiteGuid).Start(start).End(end).PresetPeriod(presetPeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMUptrends.RumGetRumWebsiteMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RumGetRumWebsiteMetrics`: []RumMetricValues
    fmt.Fprintf(os.Stdout, "Response from `RUMUptrends.RumGetRumWebsiteMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rumWebsiteGuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRumGetRumWebsiteMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | [**interface{}**](interface{}.md) | The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | [**interface{}**](interface{}.md) | The end of a custom period | 
 **presetPeriod** | **string** | The requested time period. | [default to &quot;Last24Hours&quot;]

### Return type

[**[]RumMetricValues**](RumMetricValues.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RumGetRumWebsites

> []RumWebsite RumGetRumWebsites(ctx).Execute()

Returns the definition of all RUM websites available in the account.

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
    resp, r, err := apiClient.RUMUptrends.RumGetRumWebsites(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMUptrends.RumGetRumWebsites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RumGetRumWebsites`: []RumWebsite
    fmt.Fprintf(os.Stdout, "Response from `RUMUptrends.RumGetRumWebsites`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRumGetRumWebsitesRequest struct via the builder pattern


### Return type

[**[]RumWebsite**](RumWebsite.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

