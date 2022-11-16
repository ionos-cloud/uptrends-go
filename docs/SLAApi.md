# \SLAApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlaCreateSla**](SLAApi.md#SlaCreateSla) | **Post** /Sla | Creates a new SLA.
[**SlaDeleteExclusionPeriod**](SLAApi.md#SlaDeleteExclusionPeriod) | **Delete** /Sla/{slaGuid}/ExclusionPeriod/{exclusionPeriodId} | Deletes the specified exclusion period for the specified SLA.
[**SlaDeleteSla**](SLAApi.md#SlaDeleteSla) | **Delete** /Sla/{slaGuid} | Deletes the specified SLA.
[**SlaGetExclusionPeriod**](SLAApi.md#SlaGetExclusionPeriod) | **Get** /Sla/{slaGuid}/ExclusionPeriod/{exclusionPeriodId} | Gets the specified exclusion period for the specified SLA.
[**SlaGetExclusionPeriods**](SLAApi.md#SlaGetExclusionPeriods) | **Get** /Sla/{slaGuid}/ExclusionPeriod | Gets a list of all exclusion periods for the specified SLA.
[**SlaGetSla**](SLAApi.md#SlaGetSla) | **Get** /Sla/{slaGuid} | Gets the specified SLA definition.
[**SlaGetSlas**](SLAApi.md#SlaGetSlas) | **Get** /Sla | Gets a list of all SLA definitions.
[**SlaPatchExclusionPeriod**](SLAApi.md#SlaPatchExclusionPeriod) | **Patch** /Sla/{slaGuid}/ExclusionPeriod/{exclusionPeriodId} | Partially updates the specified exclusion period for the specified SLA.
[**SlaPatchSla**](SLAApi.md#SlaPatchSla) | **Patch** /Sla/{slaGuid} | Partially updates the definition of the specified SLA.
[**SlaPostExclusionPeriod**](SLAApi.md#SlaPostExclusionPeriod) | **Post** /Sla/{slaGuid}/ExclusionPeriod | Creates a new exclusion period for the specified SLA.
[**SlaPutExclusionPeriod**](SLAApi.md#SlaPutExclusionPeriod) | **Put** /Sla/{slaGuid}/ExclusionPeriod/{exclusionPeriodId} | Updates the specified exclusion period for the specified SLA.
[**SlaPutSla**](SLAApi.md#SlaPutSla) | **Put** /Sla/{slaGuid} | Updates the definition of the specified SLA.


# **SlaCreateSla**
> Sla SlaCreateSla(ctx, optional)
Creates a new SLA.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SLAApiSlaCreateSlaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLAApiSlaCreateSlaOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sla** | [**optional.Interface of Sla**](Sla.md)| The complete SLA definition that should be created. | 

### Return type

[**Sla**](Sla.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SlaDeleteExclusionPeriod**
> SlaDeleteExclusionPeriod(ctx, slaGuid, exclusionPeriodId)
Deletes the specified exclusion period for the specified SLA.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaGuid** | **string**| The Guid of the sla definition. | 
  **exclusionPeriodId** | **int32**| The id of the exclusion period. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SlaDeleteSla**
> SlaDeleteSla(ctx, slaGuid)
Deletes the specified SLA.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaGuid** | **string**| The Guid of the SLA definition that should be deleted. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SlaGetExclusionPeriod**
> ExclusionPeriod SlaGetExclusionPeriod(ctx, slaGuid, exclusionPeriodId)
Gets the specified exclusion period for the specified SLA.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaGuid** | **string**| The Guid of the SLA definition. | 
  **exclusionPeriodId** | **int32**| The id of the exclusion period. | 

### Return type

[**ExclusionPeriod**](ExclusionPeriod.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SlaGetExclusionPeriods**
> []ExclusionPeriod SlaGetExclusionPeriods(ctx, slaGuid)
Gets a list of all exclusion periods for the specified SLA.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaGuid** | **string**| The Guid of the SLA definition. | 

### Return type

[**[]ExclusionPeriod**](ExclusionPeriod.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SlaGetSla**
> Sla SlaGetSla(ctx, slaGuid)
Gets the specified SLA definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaGuid** | **string**| The Guid of the SLA definition. | 

### Return type

[**Sla**](Sla.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SlaGetSlas**
> []Sla SlaGetSlas(ctx, )
Gets a list of all SLA definitions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Sla**](Sla.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SlaPatchExclusionPeriod**
> SlaPatchExclusionPeriod(ctx, slaGuid, exclusionPeriodId, optional)
Partially updates the specified exclusion period for the specified SLA.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaGuid** | **string**| The Guid of the SLA definition. | 
  **exclusionPeriodId** | **int32**| The id of the exclusion period. | 
 **optional** | ***SLAApiSlaPatchExclusionPeriodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLAApiSlaPatchExclusionPeriodOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exclusionPeriod** | [**optional.Interface of ExclusionPeriod**](ExclusionPeriod.md)| The complete definition of the exclusion period. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SlaPatchSla**
> SlaPatchSla(ctx, slaGuid, optional)
Partially updates the definition of the specified SLA.

This methods accepts parts of a SLA definition. We recommend retrieving the existing definition first (using the GET method). You can then process the changes you want to make and send back these changes only using this PATCH method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaGuid** | **string**| The Guid of the SLA that should be updated. | 
 **optional** | ***SLAApiSlaPatchSlaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLAApiSlaPatchSlaOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sla** | [**optional.Interface of Sla**](Sla.md)| The partial definition for the SLA that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SlaPostExclusionPeriod**
> ExclusionPeriod SlaPostExclusionPeriod(ctx, slaGuid, optional)
Creates a new exclusion period for the specified SLA.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaGuid** | **string**| The Guid of the SLA definition. | 
 **optional** | ***SLAApiSlaPostExclusionPeriodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLAApiSlaPostExclusionPeriodOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exclusionPeriod** | [**optional.Interface of ExclusionPeriod**](ExclusionPeriod.md)| The complete definition of the exclusion period. | 

### Return type

[**ExclusionPeriod**](ExclusionPeriod.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SlaPutExclusionPeriod**
> SlaPutExclusionPeriod(ctx, slaGuid, exclusionPeriodId, optional)
Updates the specified exclusion period for the specified SLA.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaGuid** | **string**| The Guid of the SLA definition. | 
  **exclusionPeriodId** | **int32**| The id of the exclusion period. | 
 **optional** | ***SLAApiSlaPutExclusionPeriodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLAApiSlaPutExclusionPeriodOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exclusionPeriod** | [**optional.Interface of ExclusionPeriod**](ExclusionPeriod.md)| The complete definition of the exclusion period. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SlaPutSla**
> SlaPutSla(ctx, slaGuid, optional)
Updates the definition of the specified SLA.

This methods only accepts a complete SLA definition. We recommend retrieving the existing definition first (using the GET method). You can then process the changes you want to make and send back the updated definition using this PUT method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaGuid** | **string**| The Guid of the SLA that should be updated. | 
 **optional** | ***SLAApiSlaPutSlaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLAApiSlaPutSlaOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sla** | [**optional.Interface of Sla**](Sla.md)| The complete definition for the SLA that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

