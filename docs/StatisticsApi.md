# \StatisticsApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatisticsGetMonitorGroupStatistics**](StatisticsApi.md#StatisticsGetMonitorGroupStatistics) | **Get** /Statistics/MonitorGroup/{monitorGroupGuid} | Gets the monitor group statistics.
[**StatisticsGetMonitorStatistics**](StatisticsApi.md#StatisticsGetMonitorStatistics) | **Get** /Statistics/Monitor/{monitorGuid} | Gets the monitor statistics.


# **StatisticsGetMonitorGroupStatistics**
> StatisticsResponse StatisticsGetMonitorGroupStatistics(ctx, monitorGroupGuid, optional)
Gets the monitor group statistics.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGroupGuid** | **string**| The Guid of the monitor group. | 
 **optional** | ***StatisticsApiStatisticsGetMonitorGroupStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsApiStatisticsGetMonitorGroupStatisticsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| The filter for the requested response fields. E.g. \&quot;Alerts,SlaTarget\&quot;. | 
 **start** | [**optional.Interface of interface{}**](.md)| The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | [**optional.Interface of interface{}**](.md)| The end of a custom period | 
 **presetPeriod** | **optional.String**| The requested time period. | [default to Last24Hours]

### Return type

[**StatisticsResponse**](StatisticsResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsGetMonitorStatistics**
> StatisticsResponse StatisticsGetMonitorStatistics(ctx, monitorGuid, optional)
Gets the monitor statistics.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The Guid of the monitor. | 
 **optional** | ***StatisticsApiStatisticsGetMonitorStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsApiStatisticsGetMonitorStatisticsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| The filter for the requested response fields. E.g. \&quot;Alerts,SlaTarget\&quot;. | 
 **start** | [**optional.Interface of interface{}**](.md)| The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | [**optional.Interface of interface{}**](.md)| The end of a custom period | 
 **presetPeriod** | **optional.String**| The requested time period. | [default to Last24Hours]

### Return type

[**StatisticsResponse**](StatisticsResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

