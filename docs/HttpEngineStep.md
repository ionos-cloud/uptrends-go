# HttpEngineStep

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepName** | **string** | The name of the step | [optional] [default to null]
**Url** | **string** | Url the step was executed on | [optional] [default to null]
**HttpStatusCode** | **string** | The HTTP status code returned | [optional] [default to null]
**HttpMethod** | **string** | Http method used in this step | [optional] [default to null]
**HttpStatusDescription** | **string** | Step description | [optional] [default to null]
**ResponseCompleted** | **bool** | Did the response complete? | [default to null]
**StepExecuted** | **bool** | Was this step executed? | [default to null]
**AssertionResultsInfo** | [***HttpEngineStepAssertionResultsInfo**](HttpEngineStep_AssertionResultsInfo.md) |  | [optional] [default to null]
**TotalTime** | **int64** | Number of milliseconds it took for this step to succeed | [default to null]
**ResponseHeaders** | **string** | Response headers | [optional] [default to null]
**ResponseBody** | **string** | Response body | [optional] [default to null]
**RequestHeaders** | **string** | Request headers send | [optional] [default to null]
**RequestBody** | **string** | Request body send | [optional] [default to null]
**ResolvedIpAddress** | **string** | Resolved IP address | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


