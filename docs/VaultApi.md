# \VaultApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VaultCreateAuthorizationForVaultSection**](VaultApi.md#VaultCreateAuthorizationForVaultSection) | **Post** /VaultSection/{vaultSectionGuid}/Authorization | Creates a new authorization for the specified vault section.
[**VaultCreateNewVaultItem**](VaultApi.md#VaultCreateNewVaultItem) | **Post** /VaultItem | Creates a new vault item.
[**VaultCreateNewVaultSection**](VaultApi.md#VaultCreateNewVaultSection) | **Post** /VaultSection | Creates a new vault section.
[**VaultDeleteAuthorizationForVaultSection**](VaultApi.md#VaultDeleteAuthorizationForVaultSection) | **Delete** /VaultSection/{vaultSectionGuid}/Authorization/{authorizationGuid} | Deletes the specified authorization for the specified vault section.
[**VaultDeleteVaultItem**](VaultApi.md#VaultDeleteVaultItem) | **Delete** /VaultItem/{vaultItemGuid} | Deletes the specified vault item.
[**VaultDeleteVaultSection**](VaultApi.md#VaultDeleteVaultSection) | **Delete** /VaultSection/{vaultSectionGuid} | Deletes the specified vault section.
[**VaultGetAllVaultItems**](VaultApi.md#VaultGetAllVaultItems) | **Get** /VaultItem | Returns all vault items.
[**VaultGetAllVaultSections**](VaultApi.md#VaultGetAllVaultSections) | **Get** /VaultSection | Returns all vault sections.
[**VaultGetAuthorizationsForVaultSection**](VaultApi.md#VaultGetAuthorizationsForVaultSection) | **Get** /VaultSection/{vaultSectionGuid}/Authorization | Returns all authorizations for the specified vault section.
[**VaultGetVaultItem**](VaultApi.md#VaultGetVaultItem) | **Get** /VaultItem/{vaultItemGuid} | Returns the specified vault item.
[**VaultGetVaultSection**](VaultApi.md#VaultGetVaultSection) | **Get** /VaultSection/{vaultSectionGuid} | Returns the specified vault section.
[**VaultPartiallyUpdateVaultItem**](VaultApi.md#VaultPartiallyUpdateVaultItem) | **Patch** /VaultItem/{vaultItemGuid} | Partially updates the specified vault item.
[**VaultUpdateVaultItem**](VaultApi.md#VaultUpdateVaultItem) | **Put** /VaultItem/{vaultItemGuid} | Updates the specified vault item.
[**VaultUpdateVaultSection**](VaultApi.md#VaultUpdateVaultSection) | **Put** /VaultSection/{vaultSectionGuid} | Updates the specified vault section.


# **VaultCreateAuthorizationForVaultSection**
> VaultSectionAuthorization VaultCreateAuthorizationForVaultSection(ctx, vaultSectionGuid, optional)
Creates a new authorization for the specified vault section.

The AuthorizationId attribute should be omitted in the request body. The newly created authorization will be returned in the response. An authorization should be granted to either an individual operator, or an operator group. Therefore, either specify the OperatorGuid attribute or the OperatorGroupGuid attribute.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultSectionGuid** | **string**| The Guid of the vault section for which to create the new authorization. | 
 **optional** | ***VaultApiVaultCreateAuthorizationForVaultSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VaultApiVaultCreateAuthorizationForVaultSectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | [**optional.Interface of VaultSectionAuthorization**](VaultSectionAuthorization.md)|  | 

### Return type

[**VaultSectionAuthorization**](VaultSectionAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultCreateNewVaultItem**
> VaultItem VaultCreateNewVaultItem(ctx, optional)
Creates a new vault item.

The VaultItemGuid field should be omitted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VaultApiVaultCreateNewVaultItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VaultApiVaultCreateNewVaultItemOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **item** | [**optional.Interface of VaultItem**](VaultItem.md)| The item to create | 

### Return type

[**VaultItem**](VaultItem.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultCreateNewVaultSection**
> VaultSection VaultCreateNewVaultSection(ctx, optional)
Creates a new vault section.

When a new vault section is created, the user that created the section is granted View and Edit authorizations to that section. The VaultSectionGuid attribute should be omitted in the request body. The Guid of the newly created section will be returned in the response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VaultApiVaultCreateNewVaultSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VaultApiVaultCreateNewVaultSectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **section** | [**optional.Interface of VaultSection**](VaultSection.md)| The attributes of the vault section that should be created. | 

### Return type

[**VaultSection**](VaultSection.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultDeleteAuthorizationForVaultSection**
> VaultDeleteAuthorizationForVaultSection(ctx, vaultSectionGuid, authorizationGuid)
Deletes the specified authorization for the specified vault section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultSectionGuid** | **string**| The Guid of the vault section for which the authorization should be deleted. | 
  **authorizationGuid** | **string**| The Guid of the authorization that should be deleted. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultDeleteVaultItem**
> VaultDeleteVaultItem(ctx, vaultItemGuid)
Deletes the specified vault item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultItemGuid** | **string**| The Guid of the vault item that should be deleted. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultDeleteVaultSection**
> VaultDeleteVaultSection(ctx, vaultSectionGuid)
Deletes the specified vault section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultSectionGuid** | **string**| The Guid of the vault section that should be deleted. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultGetAllVaultItems**
> []VaultItem VaultGetAllVaultItems(ctx, )
Returns all vault items.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]VaultItem**](VaultItem.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultGetAllVaultSections**
> []VaultSection VaultGetAllVaultSections(ctx, )
Returns all vault sections.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]VaultSection**](VaultSection.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultGetAuthorizationsForVaultSection**
> VaultSectionAuthorization VaultGetAuthorizationsForVaultSection(ctx, vaultSectionGuid)
Returns all authorizations for the specified vault section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultSectionGuid** | **string**| The Guid of the vault section for which to return authorizations. | 

### Return type

[**VaultSectionAuthorization**](VaultSectionAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultGetVaultItem**
> VaultItem VaultGetVaultItem(ctx, vaultItemGuid)
Returns the specified vault item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultItemGuid** | **string**| The Guid of the requested vault item. | 

### Return type

[**VaultItem**](VaultItem.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultGetVaultSection**
> VaultSection VaultGetVaultSection(ctx, vaultSectionGuid)
Returns the specified vault section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultSectionGuid** | **string**| The Guid of the requested vault section. | 

### Return type

[**VaultSection**](VaultSection.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultPartiallyUpdateVaultItem**
> VaultPartiallyUpdateVaultItem(ctx, vaultItemGuid, optional)
Partially updates the specified vault item.

The vault item type cannot be changed with this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultItemGuid** | **string**| The Guid of the vault item that should be updated. | 
 **optional** | ***VaultApiVaultPartiallyUpdateVaultItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VaultApiVaultPartiallyUpdateVaultItemOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **item** | [**optional.Interface of VaultItem**](VaultItem.md)| Part of the definition of the vault item that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultUpdateVaultItem**
> VaultUpdateVaultItem(ctx, vaultItemGuid, optional)
Updates the specified vault item.

Only complete definitions are accepted. Fields not specified will be NULLed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultItemGuid** | **string**| The Guid of the vault item that should be updated. | 
 **optional** | ***VaultApiVaultUpdateVaultItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VaultApiVaultUpdateVaultItemOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **item** | [**optional.Interface of VaultItem**](VaultItem.md)| The complete definition of the vault item that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VaultUpdateVaultSection**
> VaultUpdateVaultSection(ctx, vaultSectionGuid, optional)
Updates the specified vault section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultSectionGuid** | **string**| The Guid of the vault section that should be updated. | 
 **optional** | ***VaultApiVaultUpdateVaultSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VaultApiVaultUpdateVaultSectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **item** | [**optional.Interface of VaultSection**](VaultSection.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

