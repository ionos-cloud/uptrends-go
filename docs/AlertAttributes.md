# AlertAttributes

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertType** | [***AlertAttributesAlertType**](AlertAttributes_AlertType.md) |  | [default to null]
**MonitorGuid** | **string** | The monitor identifier. | [default to null]
**Timestamp** | **interface{}** | Date/time stamp of the alert. | [default to null]
**FirstError** | **interface{}** | Date/time stamp of the first monitor check. | [default to null]
**MonitorCheckId** | **int64** | The Id of the monitor check that triggered this alert. | [default to null]
**FirstErrorMonitorCheckId** | **int64** | The Id of the first monitor check error that led to this alert. | [default to null]
**ErrorDescription** | **string** | A text value that describes the error that was found, or OK if no error was found. | [optional] [default to null]
**ErrorMessage** | **string** | Any additional error information, if available. | [optional] [default to null]
**IncidentKey** | **string** | The incident key of this alert. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


