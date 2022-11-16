# \ScheduledReportApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScheduledReportCreateScheduledReport**](ScheduledReportApi.md#ScheduledReportCreateScheduledReport) | **Post** /ScheduledReport | Creates a new scheduled report.
[**ScheduledReportDeleteSpecifiedScheduledReport**](ScheduledReportApi.md#ScheduledReportDeleteSpecifiedScheduledReport) | **Delete** /ScheduledReport/{scheduledReportGuid} | Delete the specified scheduled report.
[**ScheduledReportGetAllScheduledReports**](ScheduledReportApi.md#ScheduledReportGetAllScheduledReports) | **Get** /ScheduledReport | Returns data for all scheduled reports.
[**ScheduledReportGetSpecifiedScheduledReport**](ScheduledReportApi.md#ScheduledReportGetSpecifiedScheduledReport) | **Get** /ScheduledReport/{scheduledReportGuid} | Returns data for the specified scheduled report.
[**ScheduledReportPartiallyUpdateScheduledReport**](ScheduledReportApi.md#ScheduledReportPartiallyUpdateScheduledReport) | **Patch** /ScheduledReport/{scheduledReportGuid} | Partially update the specified scheduled report.
[**ScheduledReportUpdateScheduledReport**](ScheduledReportApi.md#ScheduledReportUpdateScheduledReport) | **Put** /ScheduledReport/{scheduledReportGuid} | Update the specified scheduled report.


# **ScheduledReportCreateScheduledReport**
> ScheduledReport ScheduledReportCreateScheduledReport(ctx, optional)
Creates a new scheduled report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScheduledReportApiScheduledReportCreateScheduledReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScheduledReportApiScheduledReportCreateScheduledReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scheduledReport** | [**optional.Interface of ScheduledReport**](ScheduledReport.md)| The details for the scheduled report. | 

### Return type

[**ScheduledReport**](ScheduledReport.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScheduledReportDeleteSpecifiedScheduledReport**
> ScheduledReportDeleteSpecifiedScheduledReport(ctx, scheduledReportGuid)
Delete the specified scheduled report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scheduledReportGuid** | **string**| The guid of the specified scheduled report. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScheduledReportGetAllScheduledReports**
> []ScheduledReport ScheduledReportGetAllScheduledReports(ctx, )
Returns data for all scheduled reports.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ScheduledReport**](ScheduledReport.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScheduledReportGetSpecifiedScheduledReport**
> ScheduledReport ScheduledReportGetSpecifiedScheduledReport(ctx, scheduledReportGuid)
Returns data for the specified scheduled report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scheduledReportGuid** | **string**| The guid of the specified scheduled report. | 

### Return type

[**ScheduledReport**](ScheduledReport.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScheduledReportPartiallyUpdateScheduledReport**
> ScheduledReport ScheduledReportPartiallyUpdateScheduledReport(ctx, scheduledReportGuid, optional)
Partially update the specified scheduled report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scheduledReportGuid** | **string**| The guid of the specified scheduled report. | 
 **optional** | ***ScheduledReportApiScheduledReportPartiallyUpdateScheduledReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScheduledReportApiScheduledReportPartiallyUpdateScheduledReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduledReport** | [**optional.Interface of ScheduledReport**](ScheduledReport.md)| The details for the scheduled report. | 

### Return type

[**ScheduledReport**](ScheduledReport.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScheduledReportUpdateScheduledReport**
> ScheduledReport ScheduledReportUpdateScheduledReport(ctx, scheduledReportGuid, optional)
Update the specified scheduled report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scheduledReportGuid** | **string**| The guid of the specified scheduled report. | 
 **optional** | ***ScheduledReportApiScheduledReportUpdateScheduledReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScheduledReportApiScheduledReportUpdateScheduledReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduledReport** | [**optional.Interface of ScheduledReport**](ScheduledReport.md)| The details for the scheduled report. | 

### Return type

[**ScheduledReport**](ScheduledReport.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

