# \DashboardApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DashboardCloneDashboard**](DashboardApi.md#DashboardCloneDashboard) | **Post** /Dashboard/{dashboardGuid}/Clone | Clone the specified dashboard.
[**DashboardDeleteDashboard**](DashboardApi.md#DashboardDeleteDashboard) | **Delete** /Dashboard/{dashboardGuid} | Delete the specified dashboard.
[**DashboardGetAllDashboards**](DashboardApi.md#DashboardGetAllDashboards) | **Get** /Dashboard | Returns data for all dashboards.
[**DashboardGetOneDashboard**](DashboardApi.md#DashboardGetOneDashboard) | **Get** /Dashboard/{dashboardGuid} | Returns data for the specified dashboard.
[**DashboardPartiallyUpdateDashboard**](DashboardApi.md#DashboardPartiallyUpdateDashboard) | **Patch** /Dashboard/{dashboardGuid} | Partially update the specified dashboard.
[**DashboardUpdateDashboard**](DashboardApi.md#DashboardUpdateDashboard) | **Put** /Dashboard/{dashboardGuid} | Update the specified dashboard.


# **DashboardCloneDashboard**
> Dashboard DashboardCloneDashboard(ctx, dashboardGuid)
Clone the specified dashboard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardGuid** | **string**| The guid of the specified dashboard. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardDeleteDashboard**
> DashboardDeleteDashboard(ctx, dashboardGuid)
Delete the specified dashboard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardGuid** | **string**| The guid of the specified dashboard. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardGetAllDashboards**
> []Dashboard DashboardGetAllDashboards(ctx, )
Returns data for all dashboards.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Dashboard**](Dashboard.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardGetOneDashboard**
> Dashboard DashboardGetOneDashboard(ctx, dashboardGuid)
Returns data for the specified dashboard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardGuid** | **string**| The guid of the specified dashboard. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardPartiallyUpdateDashboard**
> DashboardPartiallyUpdateDashboard(ctx, dashboardGuid, optional)
Partially update the specified dashboard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardGuid** | **string**| The guid of the specified dashboard. | 
 **optional** | ***DashboardApiDashboardPartiallyUpdateDashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardApiDashboardPartiallyUpdateDashboardOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboard** | [**optional.Interface of Dashboard**](Dashboard.md)| The details for the dashboard. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardUpdateDashboard**
> DashboardUpdateDashboard(ctx, dashboardGuid, optional)
Update the specified dashboard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardGuid** | **string**| The guid of the specified dashboard. | 
 **optional** | ***DashboardApiDashboardUpdateDashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardApiDashboardUpdateDashboardOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboard** | [**optional.Interface of Dashboard**](Dashboard.md)| The details for the dashboard. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

