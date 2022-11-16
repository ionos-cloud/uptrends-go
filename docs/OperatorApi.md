# \OperatorApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperatorAddDutyPeriodForOperator**](OperatorApi.md#OperatorAddDutyPeriodForOperator) | **Post** /Operator/{operatorGuid}/DutySchedule | Adds a duty schedule to the specified operator.
[**OperatorCreateOperator**](OperatorApi.md#OperatorCreateOperator) | **Post** /Operator | Creates a new operator.
[**OperatorDeleteAuthorizationForOperator**](OperatorApi.md#OperatorDeleteAuthorizationForOperator) | **Delete** /Operator/{operatorGuid}/Authorization/{authorizationType} | Removes the specified authorization of this operator.
[**OperatorDeleteDutyScheduleFromOperator**](OperatorApi.md#OperatorDeleteDutyScheduleFromOperator) | **Delete** /Operator/{operatorGuid}/DutySchedule/{dutyScheduleId} | Deletes the specified duty schedule of the specified operator.
[**OperatorDeleteOperator**](OperatorApi.md#OperatorDeleteOperator) | **Delete** /Operator/{operatorGuid} | Deletes an existing operator.
[**OperatorGetAllOperators**](OperatorApi.md#OperatorGetAllOperators) | **Get** /Operator | Gets a list of all operators.
[**OperatorGetAuthorizationsForOperator**](OperatorApi.md#OperatorGetAuthorizationsForOperator) | **Get** /Operator/{operatorGuid}/Authorization | Gets all authorizations for the specified operator.
[**OperatorGetDutyScheduleForOperator**](OperatorApi.md#OperatorGetDutyScheduleForOperator) | **Get** /Operator/{operatorGuid}/DutySchedule | Gets the duty schedules for an specified operator.
[**OperatorGetOperator**](OperatorApi.md#OperatorGetOperator) | **Get** /Operator/{operatorGuid} | Gets the details of the operator with the provided OperatorGuid.
[**OperatorGetOperatorGroupsForOperator**](OperatorApi.md#OperatorGetOperatorGroupsForOperator) | **Get** /Operator/{operatorGuid}/OperatorGroup | Gets a list of all operator groups for the specified operator.
[**OperatorPostAuthorizationForOperator**](OperatorApi.md#OperatorPostAuthorizationForOperator) | **Post** /Operator/{operatorGuid}/Authorization/{authorizationType} | Assigns the specified authorization to this operator.
[**OperatorUpdateDutyPeriodForOperator**](OperatorApi.md#OperatorUpdateDutyPeriodForOperator) | **Put** /Operator/{operatorGuid}/DutySchedule/{dutyScheduleId} | Updates the specified duty schedule of the specified operator.
[**OperatorUpdateOperator**](OperatorApi.md#OperatorUpdateOperator) | **Put** /Operator/{operatorGuid} | Updates an existing operator.
[**OperatorUpdateOperatorWithPatch**](OperatorApi.md#OperatorUpdateOperatorWithPatch) | **Patch** /Operator/{operatorGuid} | Updates an existing operator.


# **OperatorAddDutyPeriodForOperator**
> OperatorAddDutyPeriodForOperator(ctx, operatorGuid, optional)
Adds a duty schedule to the specified operator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGuid** | **string**| The Guid of the operator to add the duty schedule to | 
 **optional** | ***OperatorApiOperatorAddDutyPeriodForOperatorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperatorApiOperatorAddDutyPeriodForOperatorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schedule** | [**optional.Interface of OperatorDutySchedule**](OperatorDutySchedule.md)| The duty schedule to add | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorCreateOperator**
> Operator OperatorCreateOperator(ctx, optional)
Creates a new operator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OperatorApiOperatorCreateOperatorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperatorApiOperatorCreateOperatorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uptrendsOperator** | [**optional.Interface of Operator**](Operator.md)| The details of the operator to create | 

### Return type

[**Operator**](Operator.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorDeleteAuthorizationForOperator**
> OperatorDeleteAuthorizationForOperator(ctx, operatorGuid, authorizationType)
Removes the specified authorization of this operator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGuid** | **string**| The Guid of the operator | 
  **authorizationType** | **string**| The type of authorization | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorDeleteDutyScheduleFromOperator**
> OperatorDeleteDutyScheduleFromOperator(ctx, operatorGuid, dutyScheduleId)
Deletes the specified duty schedule of the specified operator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGuid** | **string**|  | 
  **dutyScheduleId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorDeleteOperator**
> OperatorDeleteOperator(ctx, operatorGuid)
Deletes an existing operator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGuid** | **string**| The Guid of the operator to delete | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGetAllOperators**
> []Operator OperatorGetAllOperators(ctx, )
Gets a list of all operators.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Operator**](Operator.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGetAuthorizationsForOperator**
> []OperatorAuthorizationType OperatorGetAuthorizationsForOperator(ctx, operatorGuid)
Gets all authorizations for the specified operator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGuid** | **string**| The Guid of the operator | 

### Return type

[**[]OperatorAuthorizationType**](OperatorAuthorizationType.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGetDutyScheduleForOperator**
> []OperatorDutySchedule OperatorGetDutyScheduleForOperator(ctx, operatorGuid)
Gets the duty schedules for an specified operator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGuid** | **string**| The Guid of the operator to get the duty schedule for | 

### Return type

[**[]OperatorDutySchedule**](OperatorDutySchedule.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGetOperator**
> Operator OperatorGetOperator(ctx, operatorGuid)
Gets the details of the operator with the provided OperatorGuid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGuid** | **string**| The Guid of the operator for which to retrieve the details | 

### Return type

[**Operator**](Operator.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGetOperatorGroupsForOperator**
> []OperatorMember OperatorGetOperatorGroupsForOperator(ctx, operatorGuid)
Gets a list of all operator groups for the specified operator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGuid** | **string**| The Guid of the operator for which to retrieve the operator group guids | 

### Return type

[**[]OperatorMember**](OperatorMember.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorPostAuthorizationForOperator**
> OperatorPostAuthorizationForOperator(ctx, operatorGuid, authorizationType)
Assigns the specified authorization to this operator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGuid** | **string**| The Guid of the operator | 
  **authorizationType** | **string**| The type of authorization | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorUpdateDutyPeriodForOperator**
> OperatorUpdateDutyPeriodForOperator(ctx, operatorGuid, dutyScheduleId, optional)
Updates the specified duty schedule of the specified operator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGuid** | **string**|  | 
  **dutyScheduleId** | **int32**|  | 
 **optional** | ***OperatorApiOperatorUpdateDutyPeriodForOperatorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperatorApiOperatorUpdateDutyPeriodForOperatorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **schedule** | [**optional.Interface of OperatorDutySchedule**](OperatorDutySchedule.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorUpdateOperator**
> OperatorUpdateOperator(ctx, operatorGuid, optional)
Updates an existing operator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGuid** | **string**| The Guid of the operator to update | 
 **optional** | ***OperatorApiOperatorUpdateOperatorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperatorApiOperatorUpdateOperatorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uptrendsOperator** | [**optional.Interface of Operator**](Operator.md)| The updated details of the operator | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorUpdateOperatorWithPatch**
> OperatorUpdateOperatorWithPatch(ctx, operatorGuid, optional)
Updates an existing operator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGuid** | **string**| The Guid of the operator to update | 
 **optional** | ***OperatorApiOperatorUpdateOperatorWithPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperatorApiOperatorUpdateOperatorWithPatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uptrendsOperator** | [**optional.Interface of Operator**](Operator.md)| The updated details of the operator | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

