# \IntegrationApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationCreateAuthorizationForIntegration**](IntegrationApi.md#IntegrationCreateAuthorizationForIntegration) | **Post** /Integration/{integrationGuid}/Authorizations | Create authorizations for integration If the wanted authorizations requires other authorizations, these will be added as well
[**IntegrationDeleteAuthorizationForIntegration**](IntegrationApi.md#IntegrationDeleteAuthorizationForIntegration) | **Delete** /Integration/{integrationGuid}/Authorizations/{authorizationGuid} | Delete integration authorization for integration
[**IntegrationGetAllIntegrations**](IntegrationApi.md#IntegrationGetAllIntegrations) | **Get** /Integration | Gets a list of all integrations.
[**IntegrationGetAuthorizationsOfIntegration**](IntegrationApi.md#IntegrationGetAuthorizationsOfIntegration) | **Get** /Integration/{integrationGuid}/Authorizations | Get authorizations of integration


# **IntegrationCreateAuthorizationForIntegration**
> []IntegrationAuthorization IntegrationCreateAuthorizationForIntegration(ctx, integrationGuid, optional)
Create authorizations for integration If the wanted authorizations requires other authorizations, these will be added as well

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integrationGuid** | **string**| The integration GUID | 
 **optional** | ***IntegrationApiIntegrationCreateAuthorizationForIntegrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IntegrationApiIntegrationCreateAuthorizationForIntegrationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationAuthorization** | [**optional.Interface of IntegrationAuthorization**](IntegrationAuthorization.md)| Authorization to add | 

### Return type

[**[]IntegrationAuthorization**](IntegrationAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IntegrationDeleteAuthorizationForIntegration**
> IntegrationDeleteAuthorizationForIntegration(ctx, integrationGuid, authorizationGuid)
Delete integration authorization for integration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integrationGuid** | **string**| The integration GUID | 
  **authorizationGuid** | **string**| The authorization GUID that needs to be deleted | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IntegrationGetAllIntegrations**
> []Integration IntegrationGetAllIntegrations(ctx, )
Gets a list of all integrations.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Integration**](Integration.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IntegrationGetAuthorizationsOfIntegration**
> []IntegrationAuthorization IntegrationGetAuthorizationsOfIntegration(ctx, integrationGuid)
Get authorizations of integration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integrationGuid** | **string**| The intregration GUID | 

### Return type

[**[]IntegrationAuthorization**](IntegrationAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

