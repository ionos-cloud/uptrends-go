# \MonitorGroupUptrends

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MonitorGroupAddMaintenancePeriodToAllMembers**](MonitorGroupUptrends.md#MonitorGroupAddMaintenancePeriodToAllMembers) | **Post** /MonitorGroup/{monitorGroupGuid}/AddMaintenancePeriodToAllMembers | Adds the provided maintenance period to all monitors in the group specified
[**MonitorGroupAddMonitorToMonitorGroup**](MonitorGroupUptrends.md#MonitorGroupAddMonitorToMonitorGroup) | **Post** /MonitorGroup/{monitorGroupGuid}/Member/{monitorGuid} | Adds the specified monitor to the monitor group
[**MonitorGroupCreateAuthorizationForMonitorGroup**](MonitorGroupUptrends.md#MonitorGroupCreateAuthorizationForMonitorGroup) | **Post** /MonitorGroup/{monitorGroupGuid}/Authorizations | Create monitor authorizations for monitor group If the wanted authorizations requires other authorizations, these will be added as well
[**MonitorGroupCreateMonitorGroup**](MonitorGroupUptrends.md#MonitorGroupCreateMonitorGroup) | **Post** /MonitorGroup | Creates a new monitor group
[**MonitorGroupDeleteAuthorizationForMonitorGroup**](MonitorGroupUptrends.md#MonitorGroupDeleteAuthorizationForMonitorGroup) | **Delete** /MonitorGroup/{monitorGroupGuid}/Authorizations/{authorizationGuid} | Delete monitor authorization for monitor group
[**MonitorGroupDeleteMonitorGroup**](MonitorGroupUptrends.md#MonitorGroupDeleteMonitorGroup) | **Delete** /MonitorGroup/{monitorGroupGuid} | Deletes the specified monitor group
[**MonitorGroupGetAllMonitorGroups**](MonitorGroupUptrends.md#MonitorGroupGetAllMonitorGroups) | **Get** /MonitorGroup | Gets all monitor groups
[**MonitorGroupGetAuthorizationsOfMonitorGroup**](MonitorGroupUptrends.md#MonitorGroupGetAuthorizationsOfMonitorGroup) | **Get** /MonitorGroup/{monitorGroupGuid}/Authorizations | Get monitor authorizations of monitor group
[**MonitorGroupGetMonitorGroup**](MonitorGroupUptrends.md#MonitorGroupGetMonitorGroup) | **Get** /MonitorGroup/{monitorGroupGuid} | Gets the details of a monitor group
[**MonitorGroupGetMonitorGroupMembers**](MonitorGroupUptrends.md#MonitorGroupGetMonitorGroupMembers) | **Get** /MonitorGroup/{monitorGroupGuid}/Member | Gets a list of all members of a monitor group
[**MonitorGroupRemoveMonitorFromMonitorGroup**](MonitorGroupUptrends.md#MonitorGroupRemoveMonitorFromMonitorGroup) | **Delete** /MonitorGroup/{monitorGroupGuid}/Member/{monitorGuid} | Removes the specified monitor from the monitor group
[**MonitorGroupStartAllMonitorAlertsInGroup**](MonitorGroupUptrends.md#MonitorGroupStartAllMonitorAlertsInGroup) | **Post** /MonitorGroup/{monitorGroupGuid}/StartAllMonitorAlerts | Starts alerting for all monitors that are a member of the monitor group specified by the monitor group GUID
[**MonitorGroupStartAllMonitorsInGroup**](MonitorGroupUptrends.md#MonitorGroupStartAllMonitorsInGroup) | **Post** /MonitorGroup/{monitorGroupGuid}/StartAllMonitors | Starts all monitors that are a member of the monitor group specified by the monitor group GUID
[**MonitorGroupStopAllMonitorAlertsInGroup**](MonitorGroupUptrends.md#MonitorGroupStopAllMonitorAlertsInGroup) | **Post** /MonitorGroup/{monitorGroupGuid}/StopAllMonitorAlerts | Stops alerting for all monitors that are a member of the monitor group specified by the monitor group GUID
[**MonitorGroupStopAllMonitorsInGroup**](MonitorGroupUptrends.md#MonitorGroupStopAllMonitorsInGroup) | **Post** /MonitorGroup/{monitorGroupGuid}/StopAllMonitors | Stops all monitors that are a member of the monitor group specified by the monitor group GUID
[**MonitorGroupUpdateMonitorGroup**](MonitorGroupUptrends.md#MonitorGroupUpdateMonitorGroup) | **Put** /MonitorGroup/{monitorGroupGuid} | Updates the monitor group with the Guid specified



## MonitorGroupAddMaintenancePeriodToAllMembers

> MonitorGroupAddMaintenancePeriodToAllMembers(ctx, monitorGroupGuid).MaintenancePeriod(maintenancePeriod).Execute()

Adds the provided maintenance period to all monitors in the group specified

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | 
    maintenancePeriod := *openapiclient.NewMaintenancePeriod(int32(123), map[string][]openapiclient.ScheduleMode{ ... }, map[string][]openapiclient.MaintenanceTypes{ ... }) // MaintenancePeriod |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorGroupUptrends.MonitorGroupAddMaintenancePeriodToAllMembers(context.Background(), monitorGroupGuid).MaintenancePeriod(maintenancePeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupUptrends.MonitorGroupAddMaintenancePeriodToAllMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupAddMaintenancePeriodToAllMembersRequest struct via the builder pattern


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


## MonitorGroupAddMonitorToMonitorGroup

> MonitorGroupAddMonitorToMonitorGroup(ctx, monitorGroupGuid, monitorGuid).Execute()

Adds the specified monitor to the monitor group

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to add the monitor to
    monitorGuid := "monitorGuid_example" // string | The monitor Guid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorGroupUptrends.MonitorGroupAddMonitorToMonitorGroup(context.Background(), monitorGroupGuid, monitorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupUptrends.MonitorGroupAddMonitorToMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the monitor group to add the monitor to | 
**monitorGuid** | **string** | The monitor Guid | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupAddMonitorToMonitorGroupRequest struct via the builder pattern


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


## MonitorGroupCreateAuthorizationForMonitorGroup

> []MonitorGroupAuthorization MonitorGroupCreateAuthorizationForMonitorGroup(ctx, monitorGroupGuid).MonitorGroupAuthorization(monitorGroupAuthorization).Execute()

Create monitor authorizations for monitor group If the wanted authorizations requires other authorizations, these will be added as well

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The monitor group GUID
    monitorGroupAuthorization := *openapiclient.NewMonitorGroupAuthorization(map[string][]openapiclient.MonitorGroupAuthorizationType{ ... }) // MonitorGroupAuthorization | Authorization to add (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorGroupUptrends.MonitorGroupCreateAuthorizationForMonitorGroup(context.Background(), monitorGroupGuid).MonitorGroupAuthorization(monitorGroupAuthorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupUptrends.MonitorGroupCreateAuthorizationForMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorGroupCreateAuthorizationForMonitorGroup`: []MonitorGroupAuthorization
    fmt.Fprintf(os.Stdout, "Response from `MonitorGroupUptrends.MonitorGroupCreateAuthorizationForMonitorGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The monitor group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupCreateAuthorizationForMonitorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitorGroupAuthorization** | [**MonitorGroupAuthorization**](MonitorGroupAuthorization.md) | Authorization to add | 

### Return type

[**[]MonitorGroupAuthorization**](MonitorGroupAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorGroupCreateMonitorGroup

> MonitorGroup MonitorGroupCreateMonitorGroup(ctx).MonitorGroup(monitorGroup).Execute()

Creates a new monitor group

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
    monitorGroup := *openapiclient.NewMonitorGroup(false) // MonitorGroup | The MonitorGroup object to be created (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorGroupUptrends.MonitorGroupCreateMonitorGroup(context.Background()).MonitorGroup(monitorGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupUptrends.MonitorGroupCreateMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorGroupCreateMonitorGroup`: MonitorGroup
    fmt.Fprintf(os.Stdout, "Response from `MonitorGroupUptrends.MonitorGroupCreateMonitorGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupCreateMonitorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **monitorGroup** | [**MonitorGroup**](MonitorGroup.md) | The MonitorGroup object to be created | 

### Return type

[**MonitorGroup**](MonitorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorGroupDeleteAuthorizationForMonitorGroup

> MonitorGroupDeleteAuthorizationForMonitorGroup(ctx, monitorGroupGuid, authorizationGuid).Execute()

Delete monitor authorization for monitor group

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The monitor group GUID
    authorizationGuid := "authorizationGuid_example" // string | The authorization GUID that needs to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorGroupUptrends.MonitorGroupDeleteAuthorizationForMonitorGroup(context.Background(), monitorGroupGuid, authorizationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupUptrends.MonitorGroupDeleteAuthorizationForMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The monitor group GUID | 
**authorizationGuid** | **string** | The authorization GUID that needs to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupDeleteAuthorizationForMonitorGroupRequest struct via the builder pattern


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


## MonitorGroupDeleteMonitorGroup

> MonitorGroupDeleteMonitorGroup(ctx, monitorGroupGuid).Execute()

Deletes the specified monitor group

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorGroupUptrends.MonitorGroupDeleteMonitorGroup(context.Background(), monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupUptrends.MonitorGroupDeleteMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the monitor group to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupDeleteMonitorGroupRequest struct via the builder pattern


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


## MonitorGroupGetAllMonitorGroups

> []MonitorGroup MonitorGroupGetAllMonitorGroups(ctx).Execute()

Gets all monitor groups

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorGroupUptrends.MonitorGroupGetAllMonitorGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupUptrends.MonitorGroupGetAllMonitorGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorGroupGetAllMonitorGroups`: []MonitorGroup
    fmt.Fprintf(os.Stdout, "Response from `MonitorGroupUptrends.MonitorGroupGetAllMonitorGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupGetAllMonitorGroupsRequest struct via the builder pattern


### Return type

[**[]MonitorGroup**](MonitorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorGroupGetAuthorizationsOfMonitorGroup

> []MonitorGroupAuthorization MonitorGroupGetAuthorizationsOfMonitorGroup(ctx, monitorGroupGuid).Execute()

Get monitor authorizations of monitor group

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The monitor group GUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorGroupUptrends.MonitorGroupGetAuthorizationsOfMonitorGroup(context.Background(), monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupUptrends.MonitorGroupGetAuthorizationsOfMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorGroupGetAuthorizationsOfMonitorGroup`: []MonitorGroupAuthorization
    fmt.Fprintf(os.Stdout, "Response from `MonitorGroupUptrends.MonitorGroupGetAuthorizationsOfMonitorGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The monitor group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupGetAuthorizationsOfMonitorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MonitorGroupAuthorization**](MonitorGroupAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorGroupGetMonitorGroup

> MonitorGroup MonitorGroupGetMonitorGroup(ctx, monitorGroupGuid).Execute()

Gets the details of a monitor group

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to be retrieved

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorGroupUptrends.MonitorGroupGetMonitorGroup(context.Background(), monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupUptrends.MonitorGroupGetMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorGroupGetMonitorGroup`: MonitorGroup
    fmt.Fprintf(os.Stdout, "Response from `MonitorGroupUptrends.MonitorGroupGetMonitorGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the monitor group to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupGetMonitorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitorGroup**](MonitorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorGroupGetMonitorGroupMembers

> MonitorGroupMember MonitorGroupGetMonitorGroupMembers(ctx, monitorGroupGuid).Execute()

Gets a list of all members of a monitor group

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to retrieve the members for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorGroupUptrends.MonitorGroupGetMonitorGroupMembers(context.Background(), monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupUptrends.MonitorGroupGetMonitorGroupMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorGroupGetMonitorGroupMembers`: MonitorGroupMember
    fmt.Fprintf(os.Stdout, "Response from `MonitorGroupUptrends.MonitorGroupGetMonitorGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the monitor group to retrieve the members for | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupGetMonitorGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitorGroupMember**](MonitorGroupMember.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorGroupRemoveMonitorFromMonitorGroup

> MonitorGroupRemoveMonitorFromMonitorGroup(ctx, monitorGroupGuid, monitorGuid).Execute()

Removes the specified monitor from the monitor group

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to remove the monitor from
    monitorGuid := "monitorGuid_example" // string | The monitor Guid to be removed

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorGroupUptrends.MonitorGroupRemoveMonitorFromMonitorGroup(context.Background(), monitorGroupGuid, monitorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupUptrends.MonitorGroupRemoveMonitorFromMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the monitor group to remove the monitor from | 
**monitorGuid** | **string** | The monitor Guid to be removed | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupRemoveMonitorFromMonitorGroupRequest struct via the builder pattern


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


## MonitorGroupStartAllMonitorAlertsInGroup

> MonitorGroupStartAllMonitorAlertsInGroup(ctx, monitorGroupGuid).Execute()

Starts alerting for all monitors that are a member of the monitor group specified by the monitor group GUID

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The monitor group GUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorGroupUptrends.MonitorGroupStartAllMonitorAlertsInGroup(context.Background(), monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupUptrends.MonitorGroupStartAllMonitorAlertsInGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The monitor group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupStartAllMonitorAlertsInGroupRequest struct via the builder pattern


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


## MonitorGroupStartAllMonitorsInGroup

> MonitorGroupStartAllMonitorsInGroup(ctx, monitorGroupGuid).Execute()

Starts all monitors that are a member of the monitor group specified by the monitor group GUID

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The monitor group GUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorGroupUptrends.MonitorGroupStartAllMonitorsInGroup(context.Background(), monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupUptrends.MonitorGroupStartAllMonitorsInGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The monitor group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupStartAllMonitorsInGroupRequest struct via the builder pattern


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


## MonitorGroupStopAllMonitorAlertsInGroup

> MonitorGroupStopAllMonitorAlertsInGroup(ctx, monitorGroupGuid).Execute()

Stops alerting for all monitors that are a member of the monitor group specified by the monitor group GUID

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The monitor group GUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorGroupUptrends.MonitorGroupStopAllMonitorAlertsInGroup(context.Background(), monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupUptrends.MonitorGroupStopAllMonitorAlertsInGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The monitor group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupStopAllMonitorAlertsInGroupRequest struct via the builder pattern


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


## MonitorGroupStopAllMonitorsInGroup

> MonitorGroupStopAllMonitorsInGroup(ctx, monitorGroupGuid).Execute()

Stops all monitors that are a member of the monitor group specified by the monitor group GUID

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The monitor group GUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorGroupUptrends.MonitorGroupStopAllMonitorsInGroup(context.Background(), monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupUptrends.MonitorGroupStopAllMonitorsInGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The monitor group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupStopAllMonitorsInGroupRequest struct via the builder pattern


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


## MonitorGroupUpdateMonitorGroup

> MonitorGroupUpdateMonitorGroup(ctx, monitorGroupGuid).MonitorGroup(monitorGroup).Execute()

Updates the monitor group with the Guid specified

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
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to be updated
    monitorGroup := *openapiclient.NewMonitorGroup(false) // MonitorGroup | The monitor group to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorGroupUptrends.MonitorGroupUpdateMonitorGroup(context.Background(), monitorGroupGuid).MonitorGroup(monitorGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorGroupUptrends.MonitorGroupUpdateMonitorGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorGroupGuid** | **string** | The Guid of the monitor group to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorGroupUpdateMonitorGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitorGroup** | [**MonitorGroup**](MonitorGroup.md) | The monitor group to be updated | 

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

