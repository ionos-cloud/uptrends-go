# \MonitorUptrends

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MonitorCleanupMaintenancePeriods**](MonitorUptrends.md#MonitorCleanupMaintenancePeriods) | **Post** /Monitor/{monitorGuid}/MaintenancePeriod/Cleanup/{beforeDate} | Clears out all one-time maintenance periods for the specified monitor older than the specified date
[**MonitorCloneMonitor**](MonitorUptrends.md#MonitorCloneMonitor) | **Post** /Monitor/{monitorGuid}/Clone | Creates a clone (duplicate) of the specified monitor.
[**MonitorCreateAuthorizationForMonitor**](MonitorUptrends.md#MonitorCreateAuthorizationForMonitor) | **Post** /Monitor/{monitorGuid}/Authorizations | Create monitor authorizations for monitor If the wanted authorizations requires other authorizations, these will be added as well
[**MonitorCreateMaintenancePeriodForMonitor**](MonitorUptrends.md#MonitorCreateMaintenancePeriodForMonitor) | **Post** /Monitor/{monitorGuid}/MaintenancePeriod | Saves the new maintenance period provided for the specified monitor
[**MonitorDeleteAuthorizationForMonitorGroup**](MonitorUptrends.md#MonitorDeleteAuthorizationForMonitorGroup) | **Delete** /Monitor/{monitorGuid}/Authorizations/{authorizationGuid} | Delete monitor authorization for monitor
[**MonitorDeleteMaintenancePeriodFromMonitor**](MonitorUptrends.md#MonitorDeleteMaintenancePeriodFromMonitor) | **Delete** /Monitor/{monitorGuid}/MaintenancePeriod/{maintenancePeriodId} | Deletes the specified maintenance period from the specified monitor
[**MonitorDeleteMonitor**](MonitorUptrends.md#MonitorDeleteMonitor) | **Delete** /Monitor/{monitorGuid} | Deletes the specified monitor.
[**MonitorGetAllMaintenancePeriodsForMonitor**](MonitorUptrends.md#MonitorGetAllMaintenancePeriodsForMonitor) | **Get** /Monitor/{monitorGuid}/MaintenancePeriod | Finds all maintenance periods for a monitor.
[**MonitorGetAuthorizationsOfMonitor**](MonitorUptrends.md#MonitorGetAuthorizationsOfMonitor) | **Get** /Monitor/{monitorGuid}/Authorizations | Get monitor authorizations of monitor
[**MonitorGetMonitor**](MonitorUptrends.md#MonitorGetMonitor) | **Get** /Monitor/{monitorGuid} | Returns the definition of the specified monitor.
[**MonitorGetMonitorGroups**](MonitorUptrends.md#MonitorGetMonitorGroups) | **Get** /Monitor/{monitorGuid}/MonitorGroup | Returns the Guid of each monitor group where the specified monitor is a member of.
[**MonitorGetMonitors**](MonitorUptrends.md#MonitorGetMonitors) | **Get** /Monitor | Returns the definition of all monitors available in the account.
[**MonitorGetMonitorsByMonitorGroup**](MonitorUptrends.md#MonitorGetMonitorsByMonitorGroup) | **Get** /Monitor/MonitorGroup/{monitorGroupGuid} | Returns the definition of all monitors available in the account that are a member of the specified monitor group.
[**MonitorPatchMonitor**](MonitorUptrends.md#MonitorPatchMonitor) | **Patch** /Monitor/{monitorGuid} | Partially updates the definition of the specified monitor.
[**MonitorPostMonitor**](MonitorUptrends.md#MonitorPostMonitor) | **Post** /Monitor | Creates a new monitor.
[**MonitorPutMonitor**](MonitorUptrends.md#MonitorPutMonitor) | **Put** /Monitor/{monitorGuid} | Updates the definition of the specified monitor.
[**MonitorUpdateMaintenancePeriodForMonitor**](MonitorUptrends.md#MonitorUpdateMaintenancePeriodForMonitor) | **Put** /Monitor/{monitorGuid}/MaintenancePeriod/{maintenancePeriodId} | Updates the specified maintenance schedule for the specified monitor



## MonitorCleanupMaintenancePeriods

> MonitorCleanupMaintenancePeriods(ctx, monitorGuid, beforeDate).Execute()

Clears out all one-time maintenance periods for the specified monitor older than the specified date

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    monitorGuid := "monitorGuid_example" // string | 
    beforeDate := TODO // interface{} | A string representing the date, formatted as \"yyyy-MM-dd\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorCleanupMaintenancePeriods(context.Background(), monitorGuid, beforeDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorCleanupMaintenancePeriods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** |  | 
**beforeDate** | [**interface{}**](.md) | A string representing the date, formatted as \&quot;yyyy-MM-dd\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorCleanupMaintenancePeriodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorCloneMonitor

> Monitor MonitorCloneMonitor(ctx, monitorGuid).IncludeMaintenancePeriods(includeMaintenancePeriods).IncludeMonitorGroups(includeMonitorGroups).Execute()

Creates a clone (duplicate) of the specified monitor.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    monitorGuid := "monitorGuid_example" // string | The guid of the monitor you want to clone.
    includeMaintenancePeriods := true // bool | Whether or not to also copy the maintenance periods into the clone. (optional) (default to true)
    includeMonitorGroups := true // bool | Whether or not to also copy the monitor group memberships into the clone. (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorCloneMonitor(context.Background(), monitorGuid).IncludeMaintenancePeriods(includeMaintenancePeriods).IncludeMonitorGroups(includeMonitorGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorCloneMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorCloneMonitor`: Monitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorUptrends.MonitorCloneMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** | The guid of the monitor you want to clone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorCloneMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeMaintenancePeriods** | **bool** | Whether or not to also copy the maintenance periods into the clone. | [default to true]
 **includeMonitorGroups** | **bool** | Whether or not to also copy the monitor group memberships into the clone. | [default to true]

### Return type

[**Monitor**](Monitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorCreateAuthorizationForMonitor

> []MonitorAuthorization MonitorCreateAuthorizationForMonitor(ctx, monitorGuid).MonitorAuthorization(monitorAuthorization).Execute()

Create monitor authorizations for monitor If the wanted authorizations requires other authorizations, these will be added as well

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    monitorGuid := "monitorGuid_example" // string | The monitor GUID
    monitorAuthorization := *openapiclient.NewMonitorAuthorization(map[string][]openapiclient.MonitorAuthorizationType{ ... }) // MonitorAuthorization | Authorization to add (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorCreateAuthorizationForMonitor(context.Background(), monitorGuid).MonitorAuthorization(monitorAuthorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorCreateAuthorizationForMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorCreateAuthorizationForMonitor`: []MonitorAuthorization
    fmt.Fprintf(os.Stdout, "Response from `MonitorUptrends.MonitorCreateAuthorizationForMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** | The monitor GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorCreateAuthorizationForMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitorAuthorization** | [**MonitorAuthorization**](MonitorAuthorization.md) | Authorization to add | 

### Return type

[**[]MonitorAuthorization**](MonitorAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorCreateMaintenancePeriodForMonitor

> MaintenancePeriod MonitorCreateMaintenancePeriodForMonitor(ctx, monitorGuid).MaintenancePeriod(maintenancePeriod).Execute()

Saves the new maintenance period provided for the specified monitor

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    monitorGuid := "monitorGuid_example" // string | 
    maintenancePeriod := *openapiclient.NewMaintenancePeriod(int32(123), map[string][]openapiclient.ScheduleMode{ ... }, map[string][]openapiclient.MaintenanceTypes{ ... }) // MaintenancePeriod |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorCreateMaintenancePeriodForMonitor(context.Background(), monitorGuid).MaintenancePeriod(maintenancePeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorCreateMaintenancePeriodForMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorCreateMaintenancePeriodForMonitor`: MaintenancePeriod
    fmt.Fprintf(os.Stdout, "Response from `MonitorUptrends.MonitorCreateMaintenancePeriodForMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorCreateMaintenancePeriodForMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maintenancePeriod** | [**MaintenancePeriod**](MaintenancePeriod.md) |  | 

### Return type

[**MaintenancePeriod**](MaintenancePeriod.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorDeleteAuthorizationForMonitorGroup

> MonitorDeleteAuthorizationForMonitorGroup(ctx, monitorGuid, authorizationGuid).Execute()

Delete monitor authorization for monitor

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    monitorGuid := "monitorGuid_example" // string | The monitor GUID
    authorizationGuid := "authorizationGuid_example" // string | The authorization GUID that needs to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorDeleteAuthorizationForMonitorGroup(context.Background(), monitorGuid, authorizationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorDeleteAuthorizationForMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** | The monitor GUID | 
**authorizationGuid** | **string** | The authorization GUID that needs to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorDeleteAuthorizationForMonitorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorDeleteMaintenancePeriodFromMonitor

> MonitorDeleteMaintenancePeriodFromMonitor(ctx, monitorGuid, maintenancePeriodId).Execute()

Deletes the specified maintenance period from the specified monitor

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    monitorGuid := "monitorGuid_example" // string | 
    maintenancePeriodId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorDeleteMaintenancePeriodFromMonitor(context.Background(), monitorGuid, maintenancePeriodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorDeleteMaintenancePeriodFromMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** |  | 
**maintenancePeriodId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorDeleteMaintenancePeriodFromMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorDeleteMonitor

> MonitorDeleteMonitor(ctx, monitorGuid).Execute()

Deletes the specified monitor.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    monitorGuid := "monitorGuid_example" // string | The guid of the monitor you want to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorDeleteMonitor(context.Background(), monitorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorDeleteMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** | The guid of the monitor you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorDeleteMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorGetAllMaintenancePeriodsForMonitor

> []MaintenancePeriod MonitorGetAllMaintenancePeriodsForMonitor(ctx, monitorGuid).Execute()

Finds all maintenance periods for a monitor.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    monitorGuid := "monitorGuid_example" // string | The guid of the monitor you want to find the maintenance periods of.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorGetAllMaintenancePeriodsForMonitor(context.Background(), monitorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorGetAllMaintenancePeriodsForMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorGetAllMaintenancePeriodsForMonitor`: []MaintenancePeriod
    fmt.Fprintf(os.Stdout, "Response from `MonitorUptrends.MonitorGetAllMaintenancePeriodsForMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** | The guid of the monitor you want to find the maintenance periods of. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGetAllMaintenancePeriodsForMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MaintenancePeriod**](MaintenancePeriod.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorGetAuthorizationsOfMonitor

> []MonitorAuthorization MonitorGetAuthorizationsOfMonitor(ctx, monitorGuid).Execute()

Get monitor authorizations of monitor

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    monitorGuid := "monitorGuid_example" // string | The monitor GUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorGetAuthorizationsOfMonitor(context.Background(), monitorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorGetAuthorizationsOfMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorGetAuthorizationsOfMonitor`: []MonitorAuthorization
    fmt.Fprintf(os.Stdout, "Response from `MonitorUptrends.MonitorGetAuthorizationsOfMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** | The monitor GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGetAuthorizationsOfMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MonitorAuthorization**](MonitorAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorGetMonitor

> Monitor MonitorGetMonitor(ctx, monitorGuid).Filter(filter).Execute()

Returns the definition of the specified monitor.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    monitorGuid := "monitorGuid_example" // string | The Guid of the requested monitor.
    filter := "filter_example" // string | Provide the option to only retrieve the requested fields. E.g. \"Name,IsActive\". (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorGetMonitor(context.Background(), monitorGuid).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorGetMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorGetMonitor`: Monitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorUptrends.MonitorGetMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** | The Guid of the requested monitor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGetMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Provide the option to only retrieve the requested fields. E.g. \&quot;Name,IsActive\&quot;. | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorGetMonitorGroups

> []string MonitorGetMonitorGroups(ctx, monitorGuid).Execute()

Returns the Guid of each monitor group where the specified monitor is a member of.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    monitorGuid := "monitorGuid_example" // string | The Guid of the requested monitor.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorGetMonitorGroups(context.Background(), monitorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorGetMonitorGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorGetMonitorGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `MonitorUptrends.MonitorGetMonitorGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** | The Guid of the requested monitor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGetMonitorGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorGetMonitors

> []Monitor MonitorGetMonitors(ctx).Filter(filter).Execute()

Returns the definition of all monitors available in the account.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Provide the option to only retrieve the requested fields. E.g. \"Name,IsActive\". (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorGetMonitors(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorGetMonitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorGetMonitors`: []Monitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorUptrends.MonitorGetMonitors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGetMonitorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Provide the option to only retrieve the requested fields. E.g. \&quot;Name,IsActive\&quot;. | 

### Return type

[**[]Monitor**](Monitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorGetMonitorsByMonitorGroup

> []Monitor MonitorGetMonitorsByMonitorGroup(ctx, monitorGroupGuid).Filter(filter).Execute()

Returns the definition of all monitors available in the account that are a member of the specified monitor group.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the requested monitor group to retrieve the monitors of.
    filter := "filter_example" // string | Provide the option to only retrieve the requested fields. E.g. \"Name,IsActive\". (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorGetMonitorsByMonitorGroup(context.Background(), monitorGroupGuid).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorGetMonitorsByMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorGetMonitorsByMonitorGroup`: []Monitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorUptrends.MonitorGetMonitorsByMonitorGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the requested monitor group to retrieve the monitors of. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGetMonitorsByMonitorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Provide the option to only retrieve the requested fields. E.g. \&quot;Name,IsActive\&quot;. | 

### Return type

[**[]Monitor**](Monitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorPatchMonitor

> MonitorPatchMonitor(ctx, monitorGuid).Monitor(monitor).Execute()

Partially updates the definition of the specified monitor.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    monitorGuid := "monitorGuid_example" // string | The Guid of the monitor that should be updated.
    monitor := *openapiclient.NewMonitor() // Monitor | The partial definition for the monitor that should be updated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorPatchMonitor(context.Background(), monitorGuid).Monitor(monitor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorPatchMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** | The Guid of the monitor that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorPatchMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitor** | [**Monitor**](Monitor.md) | The partial definition for the monitor that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorPostMonitor

> Monitor MonitorPostMonitor(ctx).Monitor(monitor).Execute()

Creates a new monitor.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    monitor := *openapiclient.NewMonitor() // Monitor | The complete definition of the monitor that should be created. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorPostMonitor(context.Background()).Monitor(monitor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorPostMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorPostMonitor`: Monitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorUptrends.MonitorPostMonitor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitorPostMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **monitor** | [**Monitor**](Monitor.md) | The complete definition of the monitor that should be created. | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorPutMonitor

> MonitorPutMonitor(ctx, monitorGuid).Monitor(monitor).Execute()

Updates the definition of the specified monitor.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    monitorGuid := "monitorGuid_example" // string | The Guid of the monitor that should be updated.
    monitor := *openapiclient.NewMonitor() // Monitor | The complete definition for the monitor that should be updated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorPutMonitor(context.Background(), monitorGuid).Monitor(monitor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorPutMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** | The Guid of the monitor that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorPutMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitor** | [**Monitor**](Monitor.md) | The complete definition for the monitor that should be updated. | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorUpdateMaintenancePeriodForMonitor

> MonitorUpdateMaintenancePeriodForMonitor(ctx, monitorGuid, maintenancePeriodId).MaintenancePeriod(maintenancePeriod).Execute()

Updates the specified maintenance schedule for the specified monitor

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    monitorGuid := "monitorGuid_example" // string | 
    maintenancePeriodId := int32(56) // int32 | 
    maintenancePeriod := *openapiclient.NewMaintenancePeriod(int32(123), map[string][]openapiclient.ScheduleMode{ ... }, map[string][]openapiclient.MaintenanceTypes{ ... }) // MaintenancePeriod |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorUptrends.MonitorUpdateMaintenancePeriodForMonitor(context.Background(), monitorGuid, maintenancePeriodId).MaintenancePeriod(maintenancePeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorUptrends.MonitorUpdateMaintenancePeriodForMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGuid** | **string** |  | 
**maintenancePeriodId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorUpdateMaintenancePeriodForMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **maintenancePeriod** | [**MaintenancePeriod**](MaintenancePeriod.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

