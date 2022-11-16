# \RUMApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RumGetRumMetricsForAllWebsites**](RUMApi.md#RumGetRumMetricsForAllWebsites) | **Get** /Rum/Website/Metrics | Returns all metrics of all RUM websites.
[**RumGetRumWebsiteMetrics**](RUMApi.md#RumGetRumWebsiteMetrics) | **Get** /Rum/Website/{rumWebsiteGuid}/Metrics | Returns all metrics of the specified RUM website.
[**RumGetRumWebsites**](RUMApi.md#RumGetRumWebsites) | **Get** /Rum/Website | Returns the definition of all RUM websites available in the account.


# **RumGetRumMetricsForAllWebsites**
> []RumWebsiteWithMetricValues RumGetRumMetricsForAllWebsites(ctx, optional)
Returns all metrics of all RUM websites.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RUMApiRumGetRumMetricsForAllWebsitesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RUMApiRumGetRumMetricsForAllWebsitesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **websiteFilter** | **optional.String**| A search pattern to filter for matching RUM website descriptions. Wildcards (? and *) are allowed. To filter for multiple patterns, use the | symbol as separator. | 
 **start** | [**optional.Interface of interface{}**](.md)| The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | [**optional.Interface of interface{}**](.md)| The end of a custom period | 
 **presetPeriod** | **optional.String**| The requested time period. | [default to Last24Hours]

### Return type

[**[]RumWebsiteWithMetricValues**](RumWebsiteWithMetricValues.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RumGetRumWebsiteMetrics**
> []RumMetricValues RumGetRumWebsiteMetrics(ctx, rumWebsiteGuid, optional)
Returns all metrics of the specified RUM website.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rumWebsiteGuid** | **string**|  | 
 **optional** | ***RUMApiRumGetRumWebsiteMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RUMApiRumGetRumWebsiteMetricsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | [**optional.Interface of interface{}**](.md)| The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | [**optional.Interface of interface{}**](.md)| The end of a custom period | 
 **presetPeriod** | **optional.String**| The requested time period. | [default to Last24Hours]

### Return type

[**[]RumMetricValues**](RumMetricValues.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RumGetRumWebsites**
> []RumWebsite RumGetRumWebsites(ctx, )
Returns the definition of all RUM websites available in the account.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]RumWebsite**](RumWebsite.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

