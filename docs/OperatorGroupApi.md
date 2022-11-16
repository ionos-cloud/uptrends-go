# \OperatorGroupApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperatorGroupAddDutyScheduleToAllMembers**](OperatorGroupApi.md#OperatorGroupAddDutyScheduleToAllMembers) | **Post** /OperatorGroup/{operatorGroupGuid}/DutySchedule/AddDutyScheduleForAllMembers | Adds the provided duty schedule to all operators in the group specified
[**OperatorGroupAddOperatorToOperatorGroup**](OperatorGroupApi.md#OperatorGroupAddOperatorToOperatorGroup) | **Post** /OperatorGroup/{operatorGroupGuid}/Member/{operatorGuid} | Adds the specified operator to the operator group
[**OperatorGroupAllOperatorsInGroupOffDuty**](OperatorGroupApi.md#OperatorGroupAllOperatorsInGroupOffDuty) | **Post** /OperatorGroup/{operatorGroupGuid}/AllOperatorsOffDuty | Set the OnDuty flag to off for all operators that are a member of the operator group specified by the operator group GUID
[**OperatorGroupAllOperatorsInGroupOnDuty**](OperatorGroupApi.md#OperatorGroupAllOperatorsInGroupOnDuty) | **Post** /OperatorGroup/{operatorGroupGuid}/AllOperatorsOnDuty | Set the OnDuty flag to on for all operators that are a member of the operator group specified by the operator group GUID
[**OperatorGroupCreateOperatorGroup**](OperatorGroupApi.md#OperatorGroupCreateOperatorGroup) | **Post** /OperatorGroup | Creates a new operator group
[**OperatorGroupDeleteAuthorizationForOperatorGroup**](OperatorGroupApi.md#OperatorGroupDeleteAuthorizationForOperatorGroup) | **Delete** /OperatorGroup/{operatorGroupGuid}/Authorization/{authorizationType} | Removes the specified authorization of the operator group.
[**OperatorGroupDeleteOperatorGroup**](OperatorGroupApi.md#OperatorGroupDeleteOperatorGroup) | **Delete** /OperatorGroup/{operatorGroupGuid} | Deletes the specified operator group
[**OperatorGroupGetAllOperatorGroups**](OperatorGroupApi.md#OperatorGroupGetAllOperatorGroups) | **Get** /OperatorGroup | Gets all operator groups
[**OperatorGroupGetAuthorizationsForOperatorGroup**](OperatorGroupApi.md#OperatorGroupGetAuthorizationsForOperatorGroup) | **Get** /OperatorGroup/{operatorGroupGuid}/Authorization | Gets all authorizations for the specified operator group.
[**OperatorGroupGetOperatorGroup**](OperatorGroupApi.md#OperatorGroupGetOperatorGroup) | **Get** /OperatorGroup/{operatorGroupGuid} | Gets the details of a operator group
[**OperatorGroupGetOperatorGroupMembers**](OperatorGroupApi.md#OperatorGroupGetOperatorGroupMembers) | **Get** /OperatorGroup/{operatorGroupGuid}/Member | Gets a list of all members of an operator group
[**OperatorGroupPostAuthorizationForOperatorGroup**](OperatorGroupApi.md#OperatorGroupPostAuthorizationForOperatorGroup) | **Post** /OperatorGroup/{operatorGroupGuid}/Authorization/{authorizationType} | Assigns the specified authorization to the operator group.
[**OperatorGroupRemoveOperatorFromOperatorGroup**](OperatorGroupApi.md#OperatorGroupRemoveOperatorFromOperatorGroup) | **Delete** /OperatorGroup/{operatorGroupGuid}/Member/{operatorGuid} | Removes the specified operator from the operator group
[**OperatorGroupUpdateOperatorGroup**](OperatorGroupApi.md#OperatorGroupUpdateOperatorGroup) | **Put** /OperatorGroup/{operatorGroupGuid} | Updates the operator group with the Guid specified


# **OperatorGroupAddDutyScheduleToAllMembers**
> OperatorGroupAddDutyScheduleToAllMembers(ctx, operatorGroupGuid, optional)
Adds the provided duty schedule to all operators in the group specified

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGroupGuid** | **string**|  | 
 **optional** | ***OperatorGroupApiOperatorGroupAddDutyScheduleToAllMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperatorGroupApiOperatorGroupAddDutyScheduleToAllMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dutySchedule** | [**optional.Interface of OperatorDutySchedule**](OperatorDutySchedule.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGroupAddOperatorToOperatorGroup**
> OperatorGroupAddOperatorToOperatorGroup(ctx, operatorGroupGuid, operatorGuid)
Adds the specified operator to the operator group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGroupGuid** | **string**| The Guid of the operator group to add the operator to | 
  **operatorGuid** | **string**| The operator Guid | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGroupAllOperatorsInGroupOffDuty**
> OperatorGroupAllOperatorsInGroupOffDuty(ctx, operatorGroupGuid)
Set the OnDuty flag to off for all operators that are a member of the operator group specified by the operator group GUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGroupGuid** | **string**| The operator group GUID | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGroupAllOperatorsInGroupOnDuty**
> OperatorGroupAllOperatorsInGroupOnDuty(ctx, operatorGroupGuid)
Set the OnDuty flag to on for all operators that are a member of the operator group specified by the operator group GUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGroupGuid** | **string**| The operator group GUID | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGroupCreateOperatorGroup**
> OperatorGroup OperatorGroupCreateOperatorGroup(ctx, optional)
Creates a new operator group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OperatorGroupApiOperatorGroupCreateOperatorGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperatorGroupApiOperatorGroupCreateOperatorGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operatorGroup** | [**optional.Interface of OperatorGroup**](OperatorGroup.md)| The operatorGroup object to be created | 

### Return type

[**OperatorGroup**](OperatorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGroupDeleteAuthorizationForOperatorGroup**
> OperatorGroupDeleteAuthorizationForOperatorGroup(ctx, operatorGroupGuid, authorizationType)
Removes the specified authorization of the operator group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGroupGuid** | **string**| The Guid of the operator group | 
  **authorizationType** | **string**| The type of authorization | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGroupDeleteOperatorGroup**
> OperatorGroupDeleteOperatorGroup(ctx, operatorGroupGuid)
Deletes the specified operator group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGroupGuid** | **string**| The Guid of the operator group to be deleted | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGroupGetAllOperatorGroups**
> []OperatorGroup OperatorGroupGetAllOperatorGroups(ctx, )
Gets all operator groups

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]OperatorGroup**](OperatorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGroupGetAuthorizationsForOperatorGroup**
> []OperatorGroupAuthorizationType OperatorGroupGetAuthorizationsForOperatorGroup(ctx, operatorGroupGuid)
Gets all authorizations for the specified operator group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGroupGuid** | **string**| The Guid of the operator group | 

### Return type

[**[]OperatorGroupAuthorizationType**](OperatorGroupAuthorizationType.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGroupGetOperatorGroup**
> OperatorGroup OperatorGroupGetOperatorGroup(ctx, operatorGroupGuid)
Gets the details of a operator group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGroupGuid** | **string**| The Guid of the operator group to be retrieved | 

### Return type

[**OperatorGroup**](OperatorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGroupGetOperatorGroupMembers**
> OperatorGroupMember OperatorGroupGetOperatorGroupMembers(ctx, operatorGroupGuid)
Gets a list of all members of an operator group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGroupGuid** | **string**| The Guid of the operator group to retrieve the members for | 

### Return type

[**OperatorGroupMember**](OperatorGroupMember.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGroupPostAuthorizationForOperatorGroup**
> OperatorGroupPostAuthorizationForOperatorGroup(ctx, operatorGroupGuid, authorizationType)
Assigns the specified authorization to the operator group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGroupGuid** | **string**| The Guid of the operator group | 
  **authorizationType** | **string**| The type of authorization | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGroupRemoveOperatorFromOperatorGroup**
> OperatorGroupRemoveOperatorFromOperatorGroup(ctx, operatorGroupGuid, operatorGuid)
Removes the specified operator from the operator group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGroupGuid** | **string**| The Guid of the operator group to remove the operator from | 
  **operatorGuid** | **string**| The operator Guid to be removed | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperatorGroupUpdateOperatorGroup**
> OperatorGroupUpdateOperatorGroup(ctx, operatorGroupGuid, optional)
Updates the operator group with the Guid specified

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operatorGroupGuid** | **string**| The Guid of the operator group to be updated | 
 **optional** | ***OperatorGroupApiOperatorGroupUpdateOperatorGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperatorGroupApiOperatorGroupUpdateOperatorGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operatorGroup** | [**optional.Interface of OperatorGroup**](OperatorGroup.md)| The operator group to be updated | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

