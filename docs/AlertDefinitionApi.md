# \AlertDefinitionApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertDefinitionAddIntegrationToEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionAddIntegrationToEscalationLevel) | **Post** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration | Adds an integration to a specified escalation level.
[**AlertDefinitionAddMonitorGroupToAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionAddMonitorGroupToAlertDefinition) | **Post** /AlertDefinition/{alertDefinitionGuid}/Member/MonitorGroup/{monitorGroupGuid} | Adds a monitor group to the specified alert definition.
[**AlertDefinitionAddMonitorToAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionAddMonitorToAlertDefinition) | **Post** /AlertDefinition/{alertDefinitionGuid}/Member/Monitor/{monitorGuid} | Adds a monitor to the specified alert definition.
[**AlertDefinitionAddOperatorGroupToEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionAddOperatorGroupToEscalationLevel) | **Post** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Member/OperatorGroup/{operatorGroupGuid} | Adds an operator group to the specified escalation level.
[**AlertDefinitionAddOperatorToEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionAddOperatorToEscalationLevel) | **Post** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Member/Operator/{operatorGuid} | Adds an operator to the specified escalation level.
[**AlertDefinitionCreateAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionCreateAlertDefinition) | **Post** /AlertDefinition | Creates a new alert definition.
[**AlertDefinitionCreateAuthorizationForAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionCreateAuthorizationForAlertDefinition) | **Post** /AlertDefinition/{alertDefinitionGuid}/Authorizations | Create authorizations for alert definition If the wanted authorizations requires other authorizations, these will be added as well
[**AlertDefinitionDeleteAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionDeleteAlertDefinition) | **Delete** /AlertDefinition/{alertDefinitionGuid} | Deletes an existing alert definition.
[**AlertDefinitionDeleteAuthorizationForAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionDeleteAuthorizationForAlertDefinition) | **Delete** /AlertDefinition/{alertDefinitionGuid}/Authorizations/{authorizationGuid} | Delete alert definition authorization for alert definition
[**AlertDefinitionGetAllAlertDefinitions**](AlertDefinitionApi.md#AlertDefinitionGetAllAlertDefinitions) | **Get** /AlertDefinition | Gets a list of all alert definitions.
[**AlertDefinitionGetAllEscalationLevelIntegrations**](AlertDefinitionApi.md#AlertDefinitionGetAllEscalationLevelIntegrations) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration | Gets all integrations for a specified escalation level.
[**AlertDefinitionGetAllEscalationLevels**](AlertDefinitionApi.md#AlertDefinitionGetAllEscalationLevels) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel | Gets all escalation level information for the specified alert definition.
[**AlertDefinitionGetAllMembers**](AlertDefinitionApi.md#AlertDefinitionGetAllMembers) | **Get** /AlertDefinition/{alertDefinitionGuid}/Member | Gets a list of all monitor and monitor group guids for the specified alert definition.
[**AlertDefinitionGetAuthorizationsOfAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionGetAuthorizationsOfAlertDefinition) | **Get** /AlertDefinition/{alertDefinitionGuid}/Authorizations | Get authorizations of alert definition
[**AlertDefinitionGetEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionGetEscalationLevel) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId} | Gets the escalation level information of the specified alert definition.
[**AlertDefinitionGetEscalationLevelIntegration**](AlertDefinitionApi.md#AlertDefinitionGetEscalationLevelIntegration) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration/{integrationGuid} | Gets a single integration for a specified escalation level.
[**AlertDefinitionGetEscalationLevelOperator**](AlertDefinitionApi.md#AlertDefinitionGetEscalationLevelOperator) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Member | Gets the operator and operator group guids for the specified escalation level.
[**AlertDefinitionGetSpecifiedAlertDefinitions**](AlertDefinitionApi.md#AlertDefinitionGetSpecifiedAlertDefinitions) | **Get** /AlertDefinition/{alertDefinitionGuid} | Gets the specified alert definition.
[**AlertDefinitionPatchAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionPatchAlertDefinition) | **Patch** /AlertDefinition/{alertDefinitionGuid} | Partially updates the definition for the specified alert definition.
[**AlertDefinitionPatchAlertDefinitionEscalation**](AlertDefinitionApi.md#AlertDefinitionPatchAlertDefinitionEscalation) | **Patch** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId} | Partially updates the escalation level for the specified alert definition.
[**AlertDefinitionPutAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionPutAlertDefinition) | **Put** /AlertDefinition/{alertDefinitionGuid} | Updates the definition for the specified alert definition.
[**AlertDefinitionPutAlertDefinitionEscalation**](AlertDefinitionApi.md#AlertDefinitionPutAlertDefinitionEscalation) | **Put** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId} | Updates the escalation level for the specified alert definition.
[**AlertDefinitionRemoveIntegrationFromEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionRemoveIntegrationFromEscalationLevel) | **Delete** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration/{integrationGuid} | Removes an integration from a specified escalation level.
[**AlertDefinitionRemoveMonitorFromAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionRemoveMonitorFromAlertDefinition) | **Delete** /AlertDefinition/{alertDefinitionGuid}/Member/Monitor/{monitorGuid} | Removes a monitor for the specified alert definition.
[**AlertDefinitionRemoveMonitorGroupFromAlertDefinition**](AlertDefinitionApi.md#AlertDefinitionRemoveMonitorGroupFromAlertDefinition) | **Delete** /AlertDefinition/{alertDefinitionGuid}/Member/MonitorGroup/{monitorGroupGuid} | Removes a monitor group for the specified alert definition.
[**AlertDefinitionRemoveOperatorFromEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionRemoveOperatorFromEscalationLevel) | **Delete** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Member/Operator/{operatorGuid} | Removes an operator for the specified escalation level.
[**AlertDefinitionRemoveOperatorGroupFromEscalationLevel**](AlertDefinitionApi.md#AlertDefinitionRemoveOperatorGroupFromEscalationLevel) | **Delete** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Member/OperatorGroup/{operatorGroupGuid} | Removes an operator group for the specified escalation level.
[**AlertDefinitionUpdateIntegrationForEscalationWithPatch**](AlertDefinitionApi.md#AlertDefinitionUpdateIntegrationForEscalationWithPatch) | **Patch** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration/{integrationGuid} | Partially updates an integration for a specified escalation level.
[**AlertDefinitionUpdateIntegrationForEscalationWithPut**](AlertDefinitionApi.md#AlertDefinitionUpdateIntegrationForEscalationWithPut) | **Put** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration/{integrationGuid} | Updates an integration for a specified escalation level.


# **AlertDefinitionAddIntegrationToEscalationLevel**
> Integration AlertDefinitionAddIntegrationToEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, optional)
Adds an integration to a specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 
 **optional** | ***AlertDefinitionApiAlertDefinitionAddIntegrationToEscalationLevelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertDefinitionApiAlertDefinitionAddIntegrationToEscalationLevelOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **escalationLevelIntegration** | [**optional.Interface of EscalationLevelIntegration**](EscalationLevelIntegration.md)| The integration to add. | 

### Return type

[**Integration**](Integration.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionAddMonitorGroupToAlertDefinition**
> AlertDefinitionMonitorGroup AlertDefinitionAddMonitorGroupToAlertDefinition(ctx, alertDefinitionGuid, monitorGroupGuid)
Adds a monitor group to the specified alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition to modify. | 
  **monitorGroupGuid** | **string**| The Guid of the monitor group to add. | 

### Return type

[**AlertDefinitionMonitorGroup**](AlertDefinitionMonitorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionAddMonitorToAlertDefinition**
> AlertDefinitionMonitor AlertDefinitionAddMonitorToAlertDefinition(ctx, alertDefinitionGuid, monitorGuid)
Adds a monitor to the specified alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition to modify. | 
  **monitorGuid** | **string**| The Guid of the monitor to add. | 

### Return type

[**AlertDefinitionMonitor**](AlertDefinitionMonitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionAddOperatorGroupToEscalationLevel**
> AlertDefinitionOperatorGroup AlertDefinitionAddOperatorGroupToEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, operatorGroupGuid)
Adds an operator group to the specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 
  **operatorGroupGuid** | **string**| The Guid of the operator group to add. | 

### Return type

[**AlertDefinitionOperatorGroup**](AlertDefinitionOperatorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionAddOperatorToEscalationLevel**
> AlertDefinitionOperator AlertDefinitionAddOperatorToEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, operatorGuid)
Adds an operator to the specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 
  **operatorGuid** | **string**| The Guid of the operator to add. | 

### Return type

[**AlertDefinitionOperator**](AlertDefinitionOperator.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionCreateAlertDefinition**
> AlertDefinition AlertDefinitionCreateAlertDefinition(ctx, optional)
Creates a new alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertDefinitionApiAlertDefinitionCreateAlertDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertDefinitionApiAlertDefinitionCreateAlertDefinitionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertDefinition** | [**optional.Interface of AlertDefinition**](AlertDefinition.md)| The details of the alert definition to create. | 

### Return type

[**AlertDefinition**](AlertDefinition.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionCreateAuthorizationForAlertDefinition**
> []AlertDefinitionAuthorization AlertDefinitionCreateAuthorizationForAlertDefinition(ctx, alertDefinitionGuid, optional)
Create authorizations for alert definition If the wanted authorizations requires other authorizations, these will be added as well

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The alert definition GUID | 
 **optional** | ***AlertDefinitionApiAlertDefinitionCreateAuthorizationForAlertDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertDefinitionApiAlertDefinitionCreateAuthorizationForAlertDefinitionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertDefinitionAuthorization** | [**optional.Interface of AlertDefinitionAuthorization**](AlertDefinitionAuthorization.md)| Authorization to add | 

### Return type

[**[]AlertDefinitionAuthorization**](AlertDefinitionAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionDeleteAlertDefinition**
> AlertDefinitionDeleteAlertDefinition(ctx, alertDefinitionGuid)
Deletes an existing alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition to remove. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionDeleteAuthorizationForAlertDefinition**
> AlertDefinitionDeleteAuthorizationForAlertDefinition(ctx, alertDefinitionGuid, authorizationGuid)
Delete alert definition authorization for alert definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The alert definition GUID | 
  **authorizationGuid** | **string**| The authorization GUID that needs to be deleted | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionGetAllAlertDefinitions**
> []AlertDefinition AlertDefinitionGetAllAlertDefinitions(ctx, )
Gets a list of all alert definitions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AlertDefinition**](AlertDefinition.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionGetAllEscalationLevelIntegrations**
> []Integration AlertDefinitionGetAllEscalationLevelIntegrations(ctx, alertDefinitionGuid, escalationLevelId)
Gets all integrations for a specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 

### Return type

[**[]Integration**](Integration.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionGetAllEscalationLevels**
> []EscalationLevel AlertDefinitionGetAllEscalationLevels(ctx, alertDefinitionGuid)
Gets all escalation level information for the specified alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition for which to return all escalation levels. | 

### Return type

[**[]EscalationLevel**](EscalationLevel.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionGetAllMembers**
> []AlertDefinitionMember AlertDefinitionGetAllMembers(ctx, alertDefinitionGuid)
Gets a list of all monitor and monitor group guids for the specified alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition for which to return the members. | 

### Return type

[**[]AlertDefinitionMember**](AlertDefinitionMember.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionGetAuthorizationsOfAlertDefinition**
> []AlertDefinitionAuthorization AlertDefinitionGetAuthorizationsOfAlertDefinition(ctx, alertDefinitionGuid)
Get authorizations of alert definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The alert definition GUID | 

### Return type

[**[]AlertDefinitionAuthorization**](AlertDefinitionAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionGetEscalationLevel**
> EscalationLevel AlertDefinitionGetEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId)
Gets the escalation level information of the specified alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 

### Return type

[**EscalationLevel**](EscalationLevel.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionGetEscalationLevelIntegration**
> Integration AlertDefinitionGetEscalationLevelIntegration(ctx, alertDefinitionGuid, escalationLevelId, integrationGuid)
Gets a single integration for a specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 
  **integrationGuid** | **string**| The Guid of the integration. | 

### Return type

[**Integration**](Integration.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionGetEscalationLevelOperator**
> []AlertEscalationLevelMember AlertDefinitionGetEscalationLevelOperator(ctx, alertDefinitionGuid, escalationLevelId)
Gets the operator and operator group guids for the specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 

### Return type

[**[]AlertEscalationLevelMember**](AlertEscalationLevelMember.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionGetSpecifiedAlertDefinitions**
> AlertDefinition AlertDefinitionGetSpecifiedAlertDefinitions(ctx, alertDefinitionGuid)
Gets the specified alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 

### Return type

[**AlertDefinition**](AlertDefinition.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionPatchAlertDefinition**
> AlertDefinitionPatchAlertDefinition(ctx, alertDefinitionGuid, optional)
Partially updates the definition for the specified alert definition.

This methods accepts parts of an alert definition. Fields that do not require changes can be omitted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition that should be updated. | 
 **optional** | ***AlertDefinitionApiAlertDefinitionPatchAlertDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertDefinitionApiAlertDefinitionPatchAlertDefinitionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertDefinition** | [**optional.Interface of AlertDefinition**](AlertDefinition.md)| The partial definition for the alert definition that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionPatchAlertDefinitionEscalation**
> AlertDefinitionPatchAlertDefinitionEscalation(ctx, alertDefinitionGuid, escalationLevelId, optional)
Partially updates the escalation level for the specified alert definition.

This methods only accepts a complete alert definition where all fields are specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition that should be updated. | 
  **escalationLevelId** | **int32**| The level number of the escalation that should be updated. | 
 **optional** | ***AlertDefinitionApiAlertDefinitionPatchAlertDefinitionEscalationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertDefinitionApiAlertDefinitionPatchAlertDefinitionEscalationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **escalationLevel** | [**optional.Interface of EscalationLevel**](EscalationLevel.md)| The escalation level for the alert definition that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionPutAlertDefinition**
> AlertDefinitionPutAlertDefinition(ctx, alertDefinitionGuid, optional)
Updates the definition for the specified alert definition.

This methods only accepts a complete alert definition where all fields are specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition that should be updated. | 
 **optional** | ***AlertDefinitionApiAlertDefinitionPutAlertDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertDefinitionApiAlertDefinitionPutAlertDefinitionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertDefinition** | [**optional.Interface of AlertDefinition**](AlertDefinition.md)| The partial definition for the alert definition that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionPutAlertDefinitionEscalation**
> AlertDefinitionPutAlertDefinitionEscalation(ctx, alertDefinitionGuid, escalationLevelId, optional)
Updates the escalation level for the specified alert definition.

This methods only accepts a complete alert definition where all fields are specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition that should be updated. | 
  **escalationLevelId** | **int32**| The level number of the escalation that should be updated. | 
 **optional** | ***AlertDefinitionApiAlertDefinitionPutAlertDefinitionEscalationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertDefinitionApiAlertDefinitionPutAlertDefinitionEscalationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **escalationLevel** | [**optional.Interface of EscalationLevel**](EscalationLevel.md)| The escalation level for the alert definition that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionRemoveIntegrationFromEscalationLevel**
> AlertDefinitionRemoveIntegrationFromEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, integrationGuid)
Removes an integration from a specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 
  **integrationGuid** | **string**| The Guid of the integration to remove. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionRemoveMonitorFromAlertDefinition**
> AlertDefinitionRemoveMonitorFromAlertDefinition(ctx, alertDefinitionGuid, monitorGuid)
Removes a monitor for the specified alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition to modify. | 
  **monitorGuid** | **string**| The Guid of the monitor to remove. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionRemoveMonitorGroupFromAlertDefinition**
> AlertDefinitionRemoveMonitorGroupFromAlertDefinition(ctx, alertDefinitionGuid, monitorGroupGuid)
Removes a monitor group for the specified alert definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition to modify. | 
  **monitorGroupGuid** | **string**| The Guid of the monitor group to remove. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionRemoveOperatorFromEscalationLevel**
> AlertDefinitionRemoveOperatorFromEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, operatorGuid)
Removes an operator for the specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 
  **operatorGuid** | **string**| The Guid of the operator to remove. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionRemoveOperatorGroupFromEscalationLevel**
> AlertDefinitionRemoveOperatorGroupFromEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, operatorGroupGuid)
Removes an operator group for the specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 
  **operatorGroupGuid** | **string**| The Guid of the operator group to remove. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionUpdateIntegrationForEscalationWithPatch**
> AlertDefinitionUpdateIntegrationForEscalationWithPatch(ctx, alertDefinitionGuid, escalationLevelId, integrationGuid, optional)
Partially updates an integration for a specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 
  **integrationGuid** | **string**| The Guid of the integration to update. | 
 **optional** | ***AlertDefinitionApiAlertDefinitionUpdateIntegrationForEscalationWithPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertDefinitionApiAlertDefinitionUpdateIntegrationForEscalationWithPatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **escalationLevelIntegration** | [**optional.Interface of EscalationLevelIntegration**](EscalationLevelIntegration.md)| The partial definition for the integration that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlertDefinitionUpdateIntegrationForEscalationWithPut**
> AlertDefinitionUpdateIntegrationForEscalationWithPut(ctx, alertDefinitionGuid, escalationLevelId, integrationGuid, optional)
Updates an integration for a specified escalation level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionGuid** | **string**| The Guid of the alert definition. | 
  **escalationLevelId** | **int32**| The escalation level id. | 
  **integrationGuid** | **string**| The Guid of the integration to update. | 
 **optional** | ***AlertDefinitionApiAlertDefinitionUpdateIntegrationForEscalationWithPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertDefinitionApiAlertDefinitionUpdateIntegrationForEscalationWithPutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **escalationLevelIntegration** | [**optional.Interface of EscalationLevelIntegration**](EscalationLevelIntegration.md)| The definition for the integration that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

