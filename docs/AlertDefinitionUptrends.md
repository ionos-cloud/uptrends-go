# \AlertDefinitionUptrends

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertDefinitionAddIntegrationToEscalationLevel**](AlertDefinitionUptrends.md#AlertDefinitionAddIntegrationToEscalationLevel) | **Post** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration | Adds an integration to a specified escalation level.
[**AlertDefinitionAddMonitorGroupToAlertDefinition**](AlertDefinitionUptrends.md#AlertDefinitionAddMonitorGroupToAlertDefinition) | **Post** /AlertDefinition/{alertDefinitionGuid}/Member/MonitorGroup/{monitorGroupGuid} | Adds a monitor group to the specified alert definition.
[**AlertDefinitionAddMonitorToAlertDefinition**](AlertDefinitionUptrends.md#AlertDefinitionAddMonitorToAlertDefinition) | **Post** /AlertDefinition/{alertDefinitionGuid}/Member/Monitor/{monitorGuid} | Adds a monitor to the specified alert definition.
[**AlertDefinitionAddOperatorGroupToEscalationLevel**](AlertDefinitionUptrends.md#AlertDefinitionAddOperatorGroupToEscalationLevel) | **Post** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Member/OperatorGroup/{operatorGroupGuid} | Adds an operator group to the specified escalation level.
[**AlertDefinitionAddOperatorToEscalationLevel**](AlertDefinitionUptrends.md#AlertDefinitionAddOperatorToEscalationLevel) | **Post** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Member/Operator/{operatorGuid} | Adds an operator to the specified escalation level.
[**AlertDefinitionCreateAlertDefinition**](AlertDefinitionUptrends.md#AlertDefinitionCreateAlertDefinition) | **Post** /AlertDefinition | Creates a new alert definition.
[**AlertDefinitionCreateAuthorizationForAlertDefinition**](AlertDefinitionUptrends.md#AlertDefinitionCreateAuthorizationForAlertDefinition) | **Post** /AlertDefinition/{alertDefinitionGuid}/Authorizations | Create authorizations for alert definition If the wanted authorizations requires other authorizations, these will be added as well
[**AlertDefinitionDeleteAlertDefinition**](AlertDefinitionUptrends.md#AlertDefinitionDeleteAlertDefinition) | **Delete** /AlertDefinition/{alertDefinitionGuid} | Deletes an existing alert definition.
[**AlertDefinitionDeleteAuthorizationForAlertDefinition**](AlertDefinitionUptrends.md#AlertDefinitionDeleteAuthorizationForAlertDefinition) | **Delete** /AlertDefinition/{alertDefinitionGuid}/Authorizations/{authorizationGuid} | Delete alert definition authorization for alert definition
[**AlertDefinitionGetAllAlertDefinitions**](AlertDefinitionUptrends.md#AlertDefinitionGetAllAlertDefinitions) | **Get** /AlertDefinition | Gets a list of all alert definitions.
[**AlertDefinitionGetAllEscalationLevelIntegrations**](AlertDefinitionUptrends.md#AlertDefinitionGetAllEscalationLevelIntegrations) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration | Gets all integrations for a specified escalation level.
[**AlertDefinitionGetAllEscalationLevels**](AlertDefinitionUptrends.md#AlertDefinitionGetAllEscalationLevels) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel | Gets all escalation level information for the specified alert definition.
[**AlertDefinitionGetAllMembers**](AlertDefinitionUptrends.md#AlertDefinitionGetAllMembers) | **Get** /AlertDefinition/{alertDefinitionGuid}/Member | Gets a list of all monitor and monitor group guids for the specified alert definition.
[**AlertDefinitionGetAuthorizationsOfAlertDefinition**](AlertDefinitionUptrends.md#AlertDefinitionGetAuthorizationsOfAlertDefinition) | **Get** /AlertDefinition/{alertDefinitionGuid}/Authorizations | Get authorizations of alert definition
[**AlertDefinitionGetEscalationLevel**](AlertDefinitionUptrends.md#AlertDefinitionGetEscalationLevel) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId} | Gets the escalation level information of the specified alert definition.
[**AlertDefinitionGetEscalationLevelIntegration**](AlertDefinitionUptrends.md#AlertDefinitionGetEscalationLevelIntegration) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration/{integrationGuid} | Gets a single integration for a specified escalation level.
[**AlertDefinitionGetEscalationLevelOperator**](AlertDefinitionUptrends.md#AlertDefinitionGetEscalationLevelOperator) | **Get** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Member | Gets the operator and operator group guids for the specified escalation level.
[**AlertDefinitionGetSpecifiedAlertDefinitions**](AlertDefinitionUptrends.md#AlertDefinitionGetSpecifiedAlertDefinitions) | **Get** /AlertDefinition/{alertDefinitionGuid} | Gets the specified alert definition.
[**AlertDefinitionPatchAlertDefinition**](AlertDefinitionUptrends.md#AlertDefinitionPatchAlertDefinition) | **Patch** /AlertDefinition/{alertDefinitionGuid} | Partially updates the definition for the specified alert definition.
[**AlertDefinitionPatchAlertDefinitionEscalation**](AlertDefinitionUptrends.md#AlertDefinitionPatchAlertDefinitionEscalation) | **Patch** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId} | Partially updates the escalation level for the specified alert definition.
[**AlertDefinitionPutAlertDefinition**](AlertDefinitionUptrends.md#AlertDefinitionPutAlertDefinition) | **Put** /AlertDefinition/{alertDefinitionGuid} | Updates the definition for the specified alert definition.
[**AlertDefinitionPutAlertDefinitionEscalation**](AlertDefinitionUptrends.md#AlertDefinitionPutAlertDefinitionEscalation) | **Put** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId} | Updates the escalation level for the specified alert definition.
[**AlertDefinitionRemoveIntegrationFromEscalationLevel**](AlertDefinitionUptrends.md#AlertDefinitionRemoveIntegrationFromEscalationLevel) | **Delete** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration/{integrationGuid} | Removes an integration from a specified escalation level.
[**AlertDefinitionRemoveMonitorFromAlertDefinition**](AlertDefinitionUptrends.md#AlertDefinitionRemoveMonitorFromAlertDefinition) | **Delete** /AlertDefinition/{alertDefinitionGuid}/Member/Monitor/{monitorGuid} | Removes a monitor for the specified alert definition.
[**AlertDefinitionRemoveMonitorGroupFromAlertDefinition**](AlertDefinitionUptrends.md#AlertDefinitionRemoveMonitorGroupFromAlertDefinition) | **Delete** /AlertDefinition/{alertDefinitionGuid}/Member/MonitorGroup/{monitorGroupGuid} | Removes a monitor group for the specified alert definition.
[**AlertDefinitionRemoveOperatorFromEscalationLevel**](AlertDefinitionUptrends.md#AlertDefinitionRemoveOperatorFromEscalationLevel) | **Delete** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Member/Operator/{operatorGuid} | Removes an operator for the specified escalation level.
[**AlertDefinitionRemoveOperatorGroupFromEscalationLevel**](AlertDefinitionUptrends.md#AlertDefinitionRemoveOperatorGroupFromEscalationLevel) | **Delete** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Member/OperatorGroup/{operatorGroupGuid} | Removes an operator group for the specified escalation level.
[**AlertDefinitionUpdateIntegrationForEscalationWithPatch**](AlertDefinitionUptrends.md#AlertDefinitionUpdateIntegrationForEscalationWithPatch) | **Patch** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration/{integrationGuid} | Partially updates an integration for a specified escalation level.
[**AlertDefinitionUpdateIntegrationForEscalationWithPut**](AlertDefinitionUptrends.md#AlertDefinitionUpdateIntegrationForEscalationWithPut) | **Put** /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Integration/{integrationGuid} | Updates an integration for a specified escalation level.



## AlertDefinitionAddIntegrationToEscalationLevel

> Integration AlertDefinitionAddIntegrationToEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId).EscalationLevelIntegration(escalationLevelIntegration).Execute()

Adds an integration to a specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    escalationLevelIntegration := *openapiclient.NewEscalationLevelIntegration() // EscalationLevelIntegration | The integration to add. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionAddIntegrationToEscalationLevel(context.Background(), alertDefinitionGuid, escalationLevelId).EscalationLevelIntegration(escalationLevelIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionAddIntegrationToEscalationLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionAddIntegrationToEscalationLevel`: Integration
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionUptrends.AlertDefinitionAddIntegrationToEscalationLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionAddIntegrationToEscalationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **escalationLevelIntegration** | [**EscalationLevelIntegration**](EscalationLevelIntegration.md) | The integration to add. | 

### Return type

[**Integration**](Integration.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionAddMonitorGroupToAlertDefinition

> AlertDefinitionMonitorGroup AlertDefinitionAddMonitorGroupToAlertDefinition(ctx, alertDefinitionGuid, monitorGroupGuid).Execute()

Adds a monitor group to the specified alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition to modify.
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to add.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionAddMonitorGroupToAlertDefinition(context.Background(), alertDefinitionGuid, monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionAddMonitorGroupToAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionAddMonitorGroupToAlertDefinition`: AlertDefinitionMonitorGroup
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionUptrends.AlertDefinitionAddMonitorGroupToAlertDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition to modify. | 
**monitorGroupGuid** | **string** | The Guid of the monitor group to add. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionAddMonitorGroupToAlertDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlertDefinitionMonitorGroup**](AlertDefinitionMonitorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionAddMonitorToAlertDefinition

> AlertDefinitionMonitor AlertDefinitionAddMonitorToAlertDefinition(ctx, alertDefinitionGuid, monitorGuid).Execute()

Adds a monitor to the specified alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition to modify.
    monitorGuid := "monitorGuid_example" // string | The Guid of the monitor to add.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionAddMonitorToAlertDefinition(context.Background(), alertDefinitionGuid, monitorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionAddMonitorToAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionAddMonitorToAlertDefinition`: AlertDefinitionMonitor
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionUptrends.AlertDefinitionAddMonitorToAlertDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition to modify. | 
**monitorGuid** | **string** | The Guid of the monitor to add. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionAddMonitorToAlertDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlertDefinitionMonitor**](AlertDefinitionMonitor.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionAddOperatorGroupToEscalationLevel

> AlertDefinitionOperatorGroup AlertDefinitionAddOperatorGroupToEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, operatorGroupGuid).Execute()

Adds an operator group to the specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    operatorGroupGuid := "operatorGroupGuid_example" // string | The Guid of the operator group to add.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionAddOperatorGroupToEscalationLevel(context.Background(), alertDefinitionGuid, escalationLevelId, operatorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionAddOperatorGroupToEscalationLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionAddOperatorGroupToEscalationLevel`: AlertDefinitionOperatorGroup
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionUptrends.AlertDefinitionAddOperatorGroupToEscalationLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 
**operatorGroupGuid** | **string** | The Guid of the operator group to add. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionAddOperatorGroupToEscalationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AlertDefinitionOperatorGroup**](AlertDefinitionOperatorGroup.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionAddOperatorToEscalationLevel

> AlertDefinitionOperator AlertDefinitionAddOperatorToEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, operatorGuid).Execute()

Adds an operator to the specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    operatorGuid := "operatorGuid_example" // string | The Guid of the operator to add.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionAddOperatorToEscalationLevel(context.Background(), alertDefinitionGuid, escalationLevelId, operatorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionAddOperatorToEscalationLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionAddOperatorToEscalationLevel`: AlertDefinitionOperator
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionUptrends.AlertDefinitionAddOperatorToEscalationLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 
**operatorGuid** | **string** | The Guid of the operator to add. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionAddOperatorToEscalationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AlertDefinitionOperator**](AlertDefinitionOperator.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionCreateAlertDefinition

> AlertDefinition AlertDefinitionCreateAlertDefinition(ctx).AlertDefinition(alertDefinition).Execute()

Creates a new alert definition.

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
    alertDefinition := *openapiclient.NewAlertDefinition() // AlertDefinition | The details of the alert definition to create. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionCreateAlertDefinition(context.Background()).AlertDefinition(alertDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionCreateAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionCreateAlertDefinition`: AlertDefinition
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionUptrends.AlertDefinitionCreateAlertDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionCreateAlertDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertDefinition** | [**AlertDefinition**](AlertDefinition.md) | The details of the alert definition to create. | 

### Return type

[**AlertDefinition**](AlertDefinition.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionCreateAuthorizationForAlertDefinition

> []AlertDefinitionAuthorization AlertDefinitionCreateAuthorizationForAlertDefinition(ctx, alertDefinitionGuid).AlertDefinitionAuthorization(alertDefinitionAuthorization).Execute()

Create authorizations for alert definition If the wanted authorizations requires other authorizations, these will be added as well

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The alert definition GUID
    alertDefinitionAuthorization := *openapiclient.NewAlertDefinitionAuthorization(map[string][]openapiclient.AlertDefinitionAuthorizationType{ ... }) // AlertDefinitionAuthorization | Authorization to add (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionCreateAuthorizationForAlertDefinition(context.Background(), alertDefinitionGuid).AlertDefinitionAuthorization(alertDefinitionAuthorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionCreateAuthorizationForAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionCreateAuthorizationForAlertDefinition`: []AlertDefinitionAuthorization
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionUptrends.AlertDefinitionCreateAuthorizationForAlertDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The alert definition GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionCreateAuthorizationForAlertDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertDefinitionAuthorization** | [**AlertDefinitionAuthorization**](AlertDefinitionAuthorization.md) | Authorization to add | 

### Return type

[**[]AlertDefinitionAuthorization**](AlertDefinitionAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionDeleteAlertDefinition

> AlertDefinitionDeleteAlertDefinition(ctx, alertDefinitionGuid).Execute()

Deletes an existing alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition to remove.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionDeleteAlertDefinition(context.Background(), alertDefinitionGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionDeleteAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionDeleteAlertDefinitionRequest struct via the builder pattern


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


## AlertDefinitionDeleteAuthorizationForAlertDefinition

> AlertDefinitionDeleteAuthorizationForAlertDefinition(ctx, alertDefinitionGuid, authorizationGuid).Execute()

Delete alert definition authorization for alert definition

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The alert definition GUID
    authorizationGuid := "authorizationGuid_example" // string | The authorization GUID that needs to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionDeleteAuthorizationForAlertDefinition(context.Background(), alertDefinitionGuid, authorizationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionDeleteAuthorizationForAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The alert definition GUID | 
**authorizationGuid** | **string** | The authorization GUID that needs to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionDeleteAuthorizationForAlertDefinitionRequest struct via the builder pattern


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


## AlertDefinitionGetAllAlertDefinitions

> []AlertDefinition AlertDefinitionGetAllAlertDefinitions(ctx).Execute()

Gets a list of all alert definitions.

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
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionGetAllAlertDefinitions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionGetAllAlertDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetAllAlertDefinitions`: []AlertDefinition
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionUptrends.AlertDefinitionGetAllAlertDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetAllAlertDefinitionsRequest struct via the builder pattern


### Return type

[**[]AlertDefinition**](AlertDefinition.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionGetAllEscalationLevelIntegrations

> []Integration AlertDefinitionGetAllEscalationLevelIntegrations(ctx, alertDefinitionGuid, escalationLevelId).Execute()

Gets all integrations for a specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionGetAllEscalationLevelIntegrations(context.Background(), alertDefinitionGuid, escalationLevelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionGetAllEscalationLevelIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetAllEscalationLevelIntegrations`: []Integration
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionUptrends.AlertDefinitionGetAllEscalationLevelIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetAllEscalationLevelIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Integration**](Integration.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionGetAllEscalationLevels

> []EscalationLevel AlertDefinitionGetAllEscalationLevels(ctx, alertDefinitionGuid).Execute()

Gets all escalation level information for the specified alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition for which to return all escalation levels.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionGetAllEscalationLevels(context.Background(), alertDefinitionGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionGetAllEscalationLevels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetAllEscalationLevels`: []EscalationLevel
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionUptrends.AlertDefinitionGetAllEscalationLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition for which to return all escalation levels. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetAllEscalationLevelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EscalationLevel**](EscalationLevel.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionGetAllMembers

> []AlertDefinitionMember AlertDefinitionGetAllMembers(ctx, alertDefinitionGuid).Execute()

Gets a list of all monitor and monitor group guids for the specified alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition for which to return the members.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionGetAllMembers(context.Background(), alertDefinitionGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionGetAllMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetAllMembers`: []AlertDefinitionMember
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionUptrends.AlertDefinitionGetAllMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition for which to return the members. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetAllMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AlertDefinitionMember**](AlertDefinitionMember.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionGetAuthorizationsOfAlertDefinition

> []AlertDefinitionAuthorization AlertDefinitionGetAuthorizationsOfAlertDefinition(ctx, alertDefinitionGuid).Execute()

Get authorizations of alert definition

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The alert definition GUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionGetAuthorizationsOfAlertDefinition(context.Background(), alertDefinitionGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionGetAuthorizationsOfAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetAuthorizationsOfAlertDefinition`: []AlertDefinitionAuthorization
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionUptrends.AlertDefinitionGetAuthorizationsOfAlertDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The alert definition GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetAuthorizationsOfAlertDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AlertDefinitionAuthorization**](AlertDefinitionAuthorization.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionGetEscalationLevel

> EscalationLevel AlertDefinitionGetEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId).Execute()

Gets the escalation level information of the specified alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionGetEscalationLevel(context.Background(), alertDefinitionGuid, escalationLevelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionGetEscalationLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetEscalationLevel`: EscalationLevel
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionUptrends.AlertDefinitionGetEscalationLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetEscalationLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EscalationLevel**](EscalationLevel.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionGetEscalationLevelIntegration

> Integration AlertDefinitionGetEscalationLevelIntegration(ctx, alertDefinitionGuid, escalationLevelId, integrationGuid).Execute()

Gets a single integration for a specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    integrationGuid := "integrationGuid_example" // string | The Guid of the integration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionGetEscalationLevelIntegration(context.Background(), alertDefinitionGuid, escalationLevelId, integrationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionGetEscalationLevelIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetEscalationLevelIntegration`: Integration
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionUptrends.AlertDefinitionGetEscalationLevelIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 
**integrationGuid** | **string** | The Guid of the integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetEscalationLevelIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Integration**](Integration.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionGetEscalationLevelOperator

> []AlertEscalationLevelMember AlertDefinitionGetEscalationLevelOperator(ctx, alertDefinitionGuid, escalationLevelId).Execute()

Gets the operator and operator group guids for the specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionGetEscalationLevelOperator(context.Background(), alertDefinitionGuid, escalationLevelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionGetEscalationLevelOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetEscalationLevelOperator`: []AlertEscalationLevelMember
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionUptrends.AlertDefinitionGetEscalationLevelOperator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetEscalationLevelOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]AlertEscalationLevelMember**](AlertEscalationLevelMember.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionGetSpecifiedAlertDefinitions

> AlertDefinition AlertDefinitionGetSpecifiedAlertDefinitions(ctx, alertDefinitionGuid).Execute()

Gets the specified alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionGetSpecifiedAlertDefinitions(context.Background(), alertDefinitionGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionGetSpecifiedAlertDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertDefinitionGetSpecifiedAlertDefinitions`: AlertDefinition
    fmt.Fprintf(os.Stdout, "Response from `AlertDefinitionUptrends.AlertDefinitionGetSpecifiedAlertDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionGetSpecifiedAlertDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertDefinition**](AlertDefinition.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionPatchAlertDefinition

> AlertDefinitionPatchAlertDefinition(ctx, alertDefinitionGuid).AlertDefinition(alertDefinition).Execute()

Partially updates the definition for the specified alert definition.



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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition that should be updated.
    alertDefinition := *openapiclient.NewAlertDefinition() // AlertDefinition | The partial definition for the alert definition that should be updated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionPatchAlertDefinition(context.Background(), alertDefinitionGuid).AlertDefinition(alertDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionPatchAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionPatchAlertDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertDefinition** | [**AlertDefinition**](AlertDefinition.md) | The partial definition for the alert definition that should be updated. | 

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


## AlertDefinitionPatchAlertDefinitionEscalation

> AlertDefinitionPatchAlertDefinitionEscalation(ctx, alertDefinitionGuid, escalationLevelId).EscalationLevel(escalationLevel).Execute()

Partially updates the escalation level for the specified alert definition.



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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition that should be updated.
    escalationLevelId := int32(56) // int32 | The level number of the escalation that should be updated.
    escalationLevel := *openapiclient.NewEscalationLevel() // EscalationLevel | The escalation level for the alert definition that should be updated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionPatchAlertDefinitionEscalation(context.Background(), alertDefinitionGuid, escalationLevelId).EscalationLevel(escalationLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionPatchAlertDefinitionEscalation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition that should be updated. | 
**escalationLevelId** | **int32** | The level number of the escalation that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionPatchAlertDefinitionEscalationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **escalationLevel** | [**EscalationLevel**](EscalationLevel.md) | The escalation level for the alert definition that should be updated. | 

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


## AlertDefinitionPutAlertDefinition

> AlertDefinitionPutAlertDefinition(ctx, alertDefinitionGuid).AlertDefinition(alertDefinition).Execute()

Updates the definition for the specified alert definition.



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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition that should be updated.
    alertDefinition := *openapiclient.NewAlertDefinition() // AlertDefinition | The partial definition for the alert definition that should be updated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionPutAlertDefinition(context.Background(), alertDefinitionGuid).AlertDefinition(alertDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionPutAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionPutAlertDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertDefinition** | [**AlertDefinition**](AlertDefinition.md) | The partial definition for the alert definition that should be updated. | 

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


## AlertDefinitionPutAlertDefinitionEscalation

> AlertDefinitionPutAlertDefinitionEscalation(ctx, alertDefinitionGuid, escalationLevelId).EscalationLevel(escalationLevel).Execute()

Updates the escalation level for the specified alert definition.



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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition that should be updated.
    escalationLevelId := int32(56) // int32 | The level number of the escalation that should be updated.
    escalationLevel := *openapiclient.NewEscalationLevel() // EscalationLevel | The escalation level for the alert definition that should be updated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionPutAlertDefinitionEscalation(context.Background(), alertDefinitionGuid, escalationLevelId).EscalationLevel(escalationLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionPutAlertDefinitionEscalation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition that should be updated. | 
**escalationLevelId** | **int32** | The level number of the escalation that should be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionPutAlertDefinitionEscalationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **escalationLevel** | [**EscalationLevel**](EscalationLevel.md) | The escalation level for the alert definition that should be updated. | 

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


## AlertDefinitionRemoveIntegrationFromEscalationLevel

> AlertDefinitionRemoveIntegrationFromEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, integrationGuid).Execute()

Removes an integration from a specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    integrationGuid := "integrationGuid_example" // string | The Guid of the integration to remove.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionRemoveIntegrationFromEscalationLevel(context.Background(), alertDefinitionGuid, escalationLevelId, integrationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionRemoveIntegrationFromEscalationLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 
**integrationGuid** | **string** | The Guid of the integration to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionRemoveIntegrationFromEscalationLevelRequest struct via the builder pattern


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


## AlertDefinitionRemoveMonitorFromAlertDefinition

> AlertDefinitionRemoveMonitorFromAlertDefinition(ctx, alertDefinitionGuid, monitorGuid).Execute()

Removes a monitor for the specified alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition to modify.
    monitorGuid := "monitorGuid_example" // string | The Guid of the monitor to remove.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionRemoveMonitorFromAlertDefinition(context.Background(), alertDefinitionGuid, monitorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionRemoveMonitorFromAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition to modify. | 
**monitorGuid** | **string** | The Guid of the monitor to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionRemoveMonitorFromAlertDefinitionRequest struct via the builder pattern


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


## AlertDefinitionRemoveMonitorGroupFromAlertDefinition

> AlertDefinitionRemoveMonitorGroupFromAlertDefinition(ctx, alertDefinitionGuid, monitorGroupGuid).Execute()

Removes a monitor group for the specified alert definition.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition to modify.
    monitorGroupGuid := "monitorGroupGuid_example" // string | The Guid of the monitor group to remove.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionRemoveMonitorGroupFromAlertDefinition(context.Background(), alertDefinitionGuid, monitorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionRemoveMonitorGroupFromAlertDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition to modify. | 
**monitorGroupGuid** | **string** | The Guid of the monitor group to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionRemoveMonitorGroupFromAlertDefinitionRequest struct via the builder pattern


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


## AlertDefinitionRemoveOperatorFromEscalationLevel

> AlertDefinitionRemoveOperatorFromEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, operatorGuid).Execute()

Removes an operator for the specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    operatorGuid := "operatorGuid_example" // string | The Guid of the operator to remove.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionRemoveOperatorFromEscalationLevel(context.Background(), alertDefinitionGuid, escalationLevelId, operatorGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionRemoveOperatorFromEscalationLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 
**operatorGuid** | **string** | The Guid of the operator to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionRemoveOperatorFromEscalationLevelRequest struct via the builder pattern


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


## AlertDefinitionRemoveOperatorGroupFromEscalationLevel

> AlertDefinitionRemoveOperatorGroupFromEscalationLevel(ctx, alertDefinitionGuid, escalationLevelId, operatorGroupGuid).Execute()

Removes an operator group for the specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    operatorGroupGuid := "operatorGroupGuid_example" // string | The Guid of the operator group to remove.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionRemoveOperatorGroupFromEscalationLevel(context.Background(), alertDefinitionGuid, escalationLevelId, operatorGroupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionRemoveOperatorGroupFromEscalationLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 
**operatorGroupGuid** | **string** | The Guid of the operator group to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionRemoveOperatorGroupFromEscalationLevelRequest struct via the builder pattern


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


## AlertDefinitionUpdateIntegrationForEscalationWithPatch

> AlertDefinitionUpdateIntegrationForEscalationWithPatch(ctx, alertDefinitionGuid, escalationLevelId, integrationGuid).EscalationLevelIntegration(escalationLevelIntegration).Execute()

Partially updates an integration for a specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    integrationGuid := "integrationGuid_example" // string | The Guid of the integration to update.
    escalationLevelIntegration := *openapiclient.NewEscalationLevelIntegration() // EscalationLevelIntegration | The partial definition for the integration that should be updated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionUpdateIntegrationForEscalationWithPatch(context.Background(), alertDefinitionGuid, escalationLevelId, integrationGuid).EscalationLevelIntegration(escalationLevelIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionUpdateIntegrationForEscalationWithPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 
**integrationGuid** | **string** | The Guid of the integration to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionUpdateIntegrationForEscalationWithPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **escalationLevelIntegration** | [**EscalationLevelIntegration**](EscalationLevelIntegration.md) | The partial definition for the integration that should be updated. | 

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


## AlertDefinitionUpdateIntegrationForEscalationWithPut

> AlertDefinitionUpdateIntegrationForEscalationWithPut(ctx, alertDefinitionGuid, escalationLevelId, integrationGuid).EscalationLevelIntegration(escalationLevelIntegration).Execute()

Updates an integration for a specified escalation level.

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
    alertDefinitionGuid := "alertDefinitionGuid_example" // string | The Guid of the alert definition.
    escalationLevelId := int32(56) // int32 | The escalation level id.
    integrationGuid := "integrationGuid_example" // string | The Guid of the integration to update.
    escalationLevelIntegration := *openapiclient.NewEscalationLevelIntegration() // EscalationLevelIntegration | The definition for the integration that should be updated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertDefinitionUptrends.AlertDefinitionUpdateIntegrationForEscalationWithPut(context.Background(), alertDefinitionGuid, escalationLevelId, integrationGuid).EscalationLevelIntegration(escalationLevelIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertDefinitionUptrends.AlertDefinitionUpdateIntegrationForEscalationWithPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinitionGuid** | **string** | The Guid of the alert definition. | 
**escalationLevelId** | **int32** | The escalation level id. | 
**integrationGuid** | **string** | The Guid of the integration to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertDefinitionUpdateIntegrationForEscalationWithPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **escalationLevelIntegration** | [**EscalationLevelIntegration**](EscalationLevelIntegration.md) | The definition for the integration that should be updated. | 

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

