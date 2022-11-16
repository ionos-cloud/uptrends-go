# PageElement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Index of the element in the waterfall context | [default to null]
**StartTime** | **int32** | Starting time in milliseconds | [default to null]
**QueueTime** | **int32** | Number of milliseconds this element has been queueing, when appropriate. | [default to null]
**ResolveTime** | **int32** | Number of milliseconds needed to perform the DNS query for this element, when appropriate. | [default to null]
**ConnectTime** | **int32** | Number of milliseconds needed to establish a connection. | [default to null]
**StaleTime** | **int32** | Number of milliseconds the connection was stale | [default to null]
**HttpsHandshakeTime** | **int32** | Number of milliseconds needed for the HTTPS handshake | [default to null]
**SendTime** | **int32** | Number of milliseconds it took to send data | [default to null]
**WaitTime** | **int32** | Number of milliseconds the connection was in waiting state | [default to null]
**ReceiveTime** | **int32** | Number of milliseconds it took to retrieve the data | [default to null]
**TimeoutTime** | **int32** | Number of milliseconds the connection was timed-out  | [default to null]
**TotalTime** | **int32** | Total number of milliseconds it took for the connection to complete | [default to null]
**HttpStatusCode** | **int32** | The Http status code | [default to null]
**Url** | **string** | The requested resource url | [optional] [default to null]
**TotalBytes** | **int32** | Total number of downloaded bytes | [default to null]
**ElementType** | **string** | Requested resource element type, can be HTML, scripts, CSS, images, frames, objects, data and other | [optional] [default to null]
**RequestHeaders** | **string** | The HTTP request headers used | [optional] [default to null]
**ResponseHeaders** | **string** | The HTTP response headers retrieved | [optional] [default to null]
**ResolvedIpAddress** | **interface{}** | The IP address that was found for the specified domain name as part of this monitor check. | [optional] [default to null]
**GroupIds** | **[]int32** |  | [optional] [default to null]
**UrlIsBlocked** | **bool** | Was the Url excluded from waterfall (timing) data by the user? | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


