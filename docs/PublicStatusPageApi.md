# \PublicStatusPageApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicStatusPageAddAuthorizationToPublicStatusPage**](PublicStatusPageApi.md#PublicStatusPageAddAuthorizationToPublicStatusPage) | **Post** /PublicStatusPage/{publicStatusPageGuid}/Authorization | Creates a new authorization for the specified public status page.
[**PublicStatusPageDeletePublicStatusPage**](PublicStatusPageApi.md#PublicStatusPageDeletePublicStatusPage) | **Delete** /PublicStatusPage/{publicStatusPageGuid} | Deletes the definition of the specified public status page.
[**PublicStatusPageGetAuthorizationsForPublicStatusPage**](PublicStatusPageApi.md#PublicStatusPageGetAuthorizationsForPublicStatusPage) | **Get** /PublicStatusPage/{publicStatusPageGuid}/Authorization | Returns all authorizations for the specified public status page.
[**PublicStatusPageGetPublicStatusPage**](PublicStatusPageApi.md#PublicStatusPageGetPublicStatusPage) | **Get** /PublicStatusPage/{publicStatusPageGuid} | Returns the definition of the specified public status page.
[**PublicStatusPageGetPublicStatusPages**](PublicStatusPageApi.md#PublicStatusPageGetPublicStatusPages) | **Get** /PublicStatusPage | Returns the definition of all public status pages available in the account.
[**PublicStatusPagePatchPublicStatusPage**](PublicStatusPageApi.md#PublicStatusPagePatchPublicStatusPage) | **Patch** /PublicStatusPage/{publicStatusPageGuid} | Partially updates the definition of the specified public status page.
[**PublicStatusPagePostPublicStatusPage**](PublicStatusPageApi.md#PublicStatusPagePostPublicStatusPage) | **Post** /PublicStatusPage | Creates a new public status page.
[**PublicStatusPagePutPublicStatusPage**](PublicStatusPageApi.md#PublicStatusPagePutPublicStatusPage) | **Put** /PublicStatusPage/{publicStatusPageGuid} | Updates the definition of the specified public status page.
[**PublicStatusPageRemoveAuthorizationFromPublicStatusPage**](PublicStatusPageApi.md#PublicStatusPageRemoveAuthorizationFromPublicStatusPage) | **Delete** /PublicStatusPage/{publicStatusPageGuid}/Authorization/{authorizationGuid} | Removes an authorization from a public status page.


# **PublicStatusPageAddAuthorizationToPublicStatusPage**
> PspAuthorization PublicStatusPageAddAuthorizationToPublicStatusPage(ctx, publicStatusPageGuid, optional)
Creates a new authorization for the specified public status page.

The AuthorizationId attribute should be omitted in the request body. The newly created authorization will be returned in the response. An authorization should be granted to either an individual operator, or an operator group. Therefore, either specify the OperatorGuid attribute or the OperatorGroupGuid attribute.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publicStatusPageGuid** | **string**| The Guid of the public status page. | 
 **optional** | ***PublicStatusPageApiPublicStatusPageAddAuthorizationToPublicStatusPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicStatusPageApiPublicStatusPageAddAuthorizationToPublicStatusPageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | [**optional.Interface of PspAuthorization**](PspAuthorization.md)| The complete definition of the authorization that should be created. | 

### Return type

[**PspAuthorization**](PSPAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicStatusPageDeletePublicStatusPage**
> PublicStatusPageDeletePublicStatusPage(ctx, publicStatusPageGuid)
Deletes the definition of the specified public status page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publicStatusPageGuid** | **string**| The Guid of the public status page that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicStatusPageGetAuthorizationsForPublicStatusPage**
> []PspAuthorization PublicStatusPageGetAuthorizationsForPublicStatusPage(ctx, publicStatusPageGuid)
Returns all authorizations for the specified public status page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publicStatusPageGuid** | **string**| The Guid of the public status page. | 

### Return type

[**[]PspAuthorization**](PSPAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicStatusPageGetPublicStatusPage**
> PublicStatusPage PublicStatusPageGetPublicStatusPage(ctx, publicStatusPageGuid)
Returns the definition of the specified public status page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publicStatusPageGuid** | **string**| The Guid of the requested public status page. | 

### Return type

[**PublicStatusPage**](PublicStatusPage.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicStatusPageGetPublicStatusPages**
> []PublicStatusPage PublicStatusPageGetPublicStatusPages(ctx, )
Returns the definition of all public status pages available in the account.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PublicStatusPage**](PublicStatusPage.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicStatusPagePatchPublicStatusPage**
> PublicStatusPagePatchPublicStatusPage(ctx, publicStatusPageGuid, optional)
Partially updates the definition of the specified public status page.

This methods accepts parts of a public status page definition. We recommend retrieving the existing definition first (using the GET method). You can then process the changes you want to make and send back these changes only using this PATCH method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publicStatusPageGuid** | **string**| The Guid of the public status page that should be updated. | 
 **optional** | ***PublicStatusPageApiPublicStatusPagePatchPublicStatusPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicStatusPageApiPublicStatusPagePatchPublicStatusPageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicStatusPage** | [**optional.Interface of PublicStatusPage**](PublicStatusPage.md)| The partial definition for the public status page that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicStatusPagePostPublicStatusPage**
> PublicStatusPage PublicStatusPagePostPublicStatusPage(ctx, optional)
Creates a new public status page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublicStatusPageApiPublicStatusPagePostPublicStatusPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicStatusPageApiPublicStatusPagePostPublicStatusPageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicStatusPage** | [**optional.Interface of PublicStatusPage**](PublicStatusPage.md)| The complete definition for the public status page that should be updated. | 

### Return type

[**PublicStatusPage**](PublicStatusPage.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicStatusPagePutPublicStatusPage**
> PublicStatusPagePutPublicStatusPage(ctx, publicStatusPageGuid, optional)
Updates the definition of the specified public status page.

This methods only accepts a complete public status page definition. We recommend retrieving the existing definition first (using the GET method). You can then process the changes you want to make and send back the updated definition using this PUT method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publicStatusPageGuid** | **string**| The Guid of the public status page that should be updated. | 
 **optional** | ***PublicStatusPageApiPublicStatusPagePutPublicStatusPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicStatusPageApiPublicStatusPagePutPublicStatusPageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicStatusPage** | [**optional.Interface of PublicStatusPage**](PublicStatusPage.md)| The complete definition for the public status page that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicStatusPageRemoveAuthorizationFromPublicStatusPage**
> PublicStatusPageRemoveAuthorizationFromPublicStatusPage(ctx, publicStatusPageGuid, authorizationGuid)
Removes an authorization from a public status page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publicStatusPageGuid** | **string**| The Guid of the public status page. | 
  **authorizationGuid** | **string**| The Guid of the authorization. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

