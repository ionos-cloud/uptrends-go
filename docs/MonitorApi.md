# \MonitorApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MonitorCleanupMaintenancePeriods**](MonitorApi.md#MonitorCleanupMaintenancePeriods) | **Post** /Monitor/{monitorGuid}/MaintenancePeriod/Cleanup/{beforeDate} | Clears out all one-time maintenance periods for the specified monitor older than the specified date
[**MonitorCloneMonitor**](MonitorApi.md#MonitorCloneMonitor) | **Post** /Monitor/{monitorGuid}/Clone | Creates a clone (duplicate) of the specified monitor.
[**MonitorCreateAuthorizationForMonitor**](MonitorApi.md#MonitorCreateAuthorizationForMonitor) | **Post** /Monitor/{monitorGuid}/Authorizations | Create monitor authorizations for monitor If the wanted authorizations requires other authorizations, these will be added as well
[**MonitorCreateMaintenancePeriodForMonitor**](MonitorApi.md#MonitorCreateMaintenancePeriodForMonitor) | **Post** /Monitor/{monitorGuid}/MaintenancePeriod | Saves the new maintenance period provided for the specified monitor
[**MonitorDeleteAuthorizationForMonitorGroup**](MonitorApi.md#MonitorDeleteAuthorizationForMonitorGroup) | **Delete** /Monitor/{monitorGuid}/Authorizations/{authorizationGuid} | Delete monitor authorization for monitor
[**MonitorDeleteMaintenancePeriodFromMonitor**](MonitorApi.md#MonitorDeleteMaintenancePeriodFromMonitor) | **Delete** /Monitor/{monitorGuid}/MaintenancePeriod/{maintenancePeriodId} | Deletes the specified maintenance period from the specified monitor
[**MonitorDeleteMonitor**](MonitorApi.md#MonitorDeleteMonitor) | **Delete** /Monitor/{monitorGuid} | Deletes the specified monitor.
[**MonitorGetAllMaintenancePeriodsForMonitor**](MonitorApi.md#MonitorGetAllMaintenancePeriodsForMonitor) | **Get** /Monitor/{monitorGuid}/MaintenancePeriod | Finds all maintenance periods for a monitor.
[**MonitorGetAuthorizationsOfMonitor**](MonitorApi.md#MonitorGetAuthorizationsOfMonitor) | **Get** /Monitor/{monitorGuid}/Authorizations | Get monitor authorizations of monitor
[**MonitorGetMonitor**](MonitorApi.md#MonitorGetMonitor) | **Get** /Monitor/{monitorGuid} | Returns the definition of the specified monitor.
[**MonitorGetMonitorGroups**](MonitorApi.md#MonitorGetMonitorGroups) | **Get** /Monitor/{monitorGuid}/MonitorGroup | Returns the Guid of each monitor group where the specified monitor is a member of.
[**MonitorGetMonitors**](MonitorApi.md#MonitorGetMonitors) | **Get** /Monitor | Returns the definition of all monitors available in the account.
[**MonitorGetMonitorsByMonitorGroup**](MonitorApi.md#MonitorGetMonitorsByMonitorGroup) | **Get** /Monitor/MonitorGroup/{monitorGroupGuid} | Returns the definition of all monitors available in the account that are a member of the specified monitor group.
[**MonitorPatchMonitor**](MonitorApi.md#MonitorPatchMonitor) | **Patch** /Monitor/{monitorGuid} | Partially updates the definition of the specified monitor.
[**MonitorPostMonitor**](MonitorApi.md#MonitorPostMonitor) | **Post** /Monitor | Creates a new monitor.
[**MonitorPutMonitor**](MonitorApi.md#MonitorPutMonitor) | **Put** /Monitor/{monitorGuid} | Updates the definition of the specified monitor.
[**MonitorUpdateMaintenancePeriodForMonitor**](MonitorApi.md#MonitorUpdateMaintenancePeriodForMonitor) | **Put** /Monitor/{monitorGuid}/MaintenancePeriod/{maintenancePeriodId} | Updates the specified maintenance schedule for the specified monitor


# **MonitorCleanupMaintenancePeriods**
> MonitorCleanupMaintenancePeriods(ctx, monitorGuid, beforeDate)
Clears out all one-time maintenance periods for the specified monitor older than the specified date

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**|  | 
  **beforeDate** | [**interface{}**](.md)| A string representing the date, formatted as \&quot;yyyy-MM-dd\&quot; | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorCloneMonitor**
> Monitor MonitorCloneMonitor(ctx, monitorGuid, optional)
Creates a clone (duplicate) of the specified monitor.

Upon creation, the new monitor will be inactive. This allows you to make the necessary changes before you activate it. All other settings will be transferred to the new monitor as-is.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The guid of the monitor you want to clone. | 
 **optional** | ***MonitorApiMonitorCloneMonitorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorApiMonitorCloneMonitorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeMaintenancePeriods** | **optional.Bool**| Whether or not to also copy the maintenance periods into the clone. | [default to true]
 **includeMonitorGroups** | **optional.Bool**| Whether or not to also copy the monitor group memberships into the clone. | [default to true]

### Return type

[**Monitor**](Monitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorCreateAuthorizationForMonitor**
> []MonitorAuthorization MonitorCreateAuthorizationForMonitor(ctx, monitorGuid, optional)
Create monitor authorizations for monitor If the wanted authorizations requires other authorizations, these will be added as well

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The monitor GUID | 
 **optional** | ***MonitorApiMonitorCreateAuthorizationForMonitorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorApiMonitorCreateAuthorizationForMonitorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitorAuthorization** | [**optional.Interface of MonitorAuthorization**](MonitorAuthorization.md)| Authorization to add | 

### Return type

[**[]MonitorAuthorization**](MonitorAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorCreateMaintenancePeriodForMonitor**
> MaintenancePeriod MonitorCreateMaintenancePeriodForMonitor(ctx, monitorGuid, optional)
Saves the new maintenance period provided for the specified monitor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**|  | 
 **optional** | ***MonitorApiMonitorCreateMaintenancePeriodForMonitorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorApiMonitorCreateMaintenancePeriodForMonitorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maintenancePeriod** | [**optional.Interface of MaintenancePeriod**](MaintenancePeriod.md)|  | 

### Return type

[**MaintenancePeriod**](MaintenancePeriod.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorDeleteAuthorizationForMonitorGroup**
> MonitorDeleteAuthorizationForMonitorGroup(ctx, monitorGuid, authorizationGuid)
Delete monitor authorization for monitor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The monitor GUID | 
  **authorizationGuid** | **string**| The authorization GUID that needs to be deleted | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorDeleteMaintenancePeriodFromMonitor**
> MonitorDeleteMaintenancePeriodFromMonitor(ctx, monitorGuid, maintenancePeriodId)
Deletes the specified maintenance period from the specified monitor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**|  | 
  **maintenancePeriodId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorDeleteMonitor**
> MonitorDeleteMonitor(ctx, monitorGuid)
Deletes the specified monitor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The guid of the monitor you want to delete. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGetAllMaintenancePeriodsForMonitor**
> []MaintenancePeriod MonitorGetAllMaintenancePeriodsForMonitor(ctx, monitorGuid)
Finds all maintenance periods for a monitor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The guid of the monitor you want to find the maintenance periods of. | 

### Return type

[**[]MaintenancePeriod**](MaintenancePeriod.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGetAuthorizationsOfMonitor**
> []MonitorAuthorization MonitorGetAuthorizationsOfMonitor(ctx, monitorGuid)
Get monitor authorizations of monitor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The monitor GUID | 

### Return type

[**[]MonitorAuthorization**](MonitorAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGetMonitor**
> Monitor MonitorGetMonitor(ctx, monitorGuid, optional)
Returns the definition of the specified monitor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The Guid of the requested monitor. | 
 **optional** | ***MonitorApiMonitorGetMonitorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorApiMonitorGetMonitorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Provide the option to only retrieve the requested fields. E.g. \&quot;Name,IsActive\&quot;. | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGetMonitorGroups**
> []string MonitorGetMonitorGroups(ctx, monitorGuid)
Returns the Guid of each monitor group where the specified monitor is a member of.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The Guid of the requested monitor. | 

### Return type

**[]string**

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGetMonitors**
> []Monitor MonitorGetMonitors(ctx, optional)
Returns the definition of all monitors available in the account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MonitorApiMonitorGetMonitorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorApiMonitorGetMonitorsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Provide the option to only retrieve the requested fields. E.g. \&quot;Name,IsActive\&quot;. | 

### Return type

[**[]Monitor**](Monitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGetMonitorsByMonitorGroup**
> []Monitor MonitorGetMonitorsByMonitorGroup(ctx, monitorGroupGuid, optional)
Returns the definition of all monitors available in the account that are a member of the specified monitor group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGroupGuid** | **string**| The Guid of the requested monitor group to retrieve the monitors of. | 
 **optional** | ***MonitorApiMonitorGetMonitorsByMonitorGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorApiMonitorGetMonitorsByMonitorGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Provide the option to only retrieve the requested fields. E.g. \&quot;Name,IsActive\&quot;. | 

### Return type

[**[]Monitor**](Monitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorPatchMonitor**
> MonitorPatchMonitor(ctx, monitorGuid, optional)
Partially updates the definition of the specified monitor.

This methods accepts parts of a monitor definition. We recommend retrieving the existing definition first (using the GET method). You can then process the changes you want to make and send back these changes only using this PATCH method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The Guid of the monitor that should be updated. | 
 **optional** | ***MonitorApiMonitorPatchMonitorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorApiMonitorPatchMonitorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitor** | [**optional.Interface of Monitor**](Monitor.md)| The partial definition for the monitor that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorPostMonitor**
> Monitor MonitorPostMonitor(ctx, optional)
Creates a new monitor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MonitorApiMonitorPostMonitorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorApiMonitorPostMonitorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **monitor** | [**optional.Interface of Monitor**](Monitor.md)| The complete definition of the monitor that should be created. | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorPutMonitor**
> MonitorPutMonitor(ctx, monitorGuid, optional)
Updates the definition of the specified monitor.

This methods only accepts a complete monitor definition. We recommend retrieving the existing definition first (using the GET method). You can then process the changes you want to make and send back the updated definition using this PUT method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**| The Guid of the monitor that should be updated. | 
 **optional** | ***MonitorApiMonitorPutMonitorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorApiMonitorPutMonitorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitor** | [**optional.Interface of Monitor**](Monitor.md)| The complete definition for the monitor that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorUpdateMaintenancePeriodForMonitor**
> MonitorUpdateMaintenancePeriodForMonitor(ctx, monitorGuid, maintenancePeriodId, optional)
Updates the specified maintenance schedule for the specified monitor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorGuid** | **string**|  | 
  **maintenancePeriodId** | **int32**|  | 
 **optional** | ***MonitorApiMonitorUpdateMaintenancePeriodForMonitorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorApiMonitorUpdateMaintenancePeriodForMonitorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **maintenancePeriod** | [**optional.Interface of MaintenancePeriod**](MaintenancePeriod.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

