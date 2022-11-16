# \AlertApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertGetMonitorAlerts**](AlertApi.md#AlertGetMonitorAlerts) | **Get** /Alert/Monitor/{monitorGuid} | Returns alerts for a specific monitor.
[**AlertGetMonitorGroupAlerts**](AlertApi.md#AlertGetMonitorGroupAlerts) | **Get** /Alert/MonitorGroup/{monitorGroupGuid} | Returns alerts for a specific monitor group.


# **AlertGetMonitorAlerts**
> AlertResponse AlertGetMonitorAlerts(ctx, monitorGuid, optional)
Returns alerts for a specific monitor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The Guid of the monitor to get alerts for. | 
 **optional** | ***AlertApiAlertGetMonitorAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertApiAlertGetMonitorAlertsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeReminders** | **optional.Bool**| When true reminder alerts will be included in the results | [default to false]
 **cursor** | **optional.String**| A cursor value that should be used for traversing the dataset. | 
 **sorting** | **optional.String**| Sorting direction based on timestamp. | [default to Descending]
 **take** | **optional.Int32**| The number of records to return (Max value &#x3D; 100) | [default to 100]
 **start** | [**optional.Interface of interface{}**](.md)| The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | [**optional.Interface of interface{}**](.md)| The end of a custom period | 
 **presetPeriod** | **optional.String**| The requested time period. | [default to Last24Hours]

### Return type

[**AlertResponse**](AlertResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertGetMonitorGroupAlerts**
> AlertResponse AlertGetMonitorGroupAlerts(ctx, monitorGroupGuid, optional)
Returns alerts for a specific monitor group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGroupGuid** | **string**| The Guid of the monitor group to get alerts for. | 
 **optional** | ***AlertApiAlertGetMonitorGroupAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertApiAlertGetMonitorGroupAlertsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeReminders** | **optional.Bool**| When true reminder alerts will be included in the results | [default to false]
 **cursor** | **optional.String**| A cursor value that should be used for traversing the dataset. | 
 **sorting** | **optional.String**| Sorting direction based on timestamp. | [default to Descending]
 **take** | **optional.Int32**| The number of records to return (Max value &#x3D; 100) | [default to 100]
 **start** | [**optional.Interface of interface{}**](.md)| The start of a custom period (can&#39;t be used together with the PresetPeriod parameter) | 
 **end** | [**optional.Interface of interface{}**](.md)| The end of a custom period | 
 **presetPeriod** | **optional.String**| The requested time period. | [default to Last24Hours]

### Return type

[**AlertResponse**](AlertResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

