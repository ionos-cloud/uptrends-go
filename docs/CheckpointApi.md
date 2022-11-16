# \CheckpointApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckpointGetAllCheckpoints**](CheckpointApi.md#CheckpointGetAllCheckpoints) | **Get** /Checkpoint | Returns all the checkpoints that are not deleted
[**CheckpointGetCheckpoint**](CheckpointApi.md#CheckpointGetCheckpoint) | **Get** /Checkpoint/{checkpointId} | Returns the specified checkpoint, deleted or not
[**CheckpointRegionGetAllCheckpointRegions**](CheckpointApi.md#CheckpointRegionGetAllCheckpointRegions) | **Get** /CheckpointRegion | Returns all the checkpoint regions.
[**CheckpointRegionGetCheckpointRegionCheckpoints**](CheckpointApi.md#CheckpointRegionGetCheckpointRegionCheckpoints) | **Get** /CheckpointRegion/{checkpointRegionId}/Checkpoint | Returns the checkpoints for the specified checkpoint region.
[**CheckpointRegionGetSpecifiedCheckpointRegion**](CheckpointApi.md#CheckpointRegionGetSpecifiedCheckpointRegion) | **Get** /CheckpointRegion/{checkpointRegionId} | Returns the specified checkpoint region.
[**CheckpointServerGetAllServerIpv4Addresses**](CheckpointApi.md#CheckpointServerGetAllServerIpv4Addresses) | **Get** /Checkpoint/Server/Ipv4 | Anonymous call that returns all the IPv4 addresses of all the checkpoint servers.
[**CheckpointServerGetAllServerIpv6Addresses**](CheckpointApi.md#CheckpointServerGetAllServerIpv6Addresses) | **Get** /Checkpoint/Server/Ipv6 | Anonymous call that returns all the IPv6 addresses of all the checkpoint servers.
[**CheckpointServerGetServer**](CheckpointApi.md#CheckpointServerGetServer) | **Get** /Checkpoint/Server/{serverId} | Returns the requested checkpoint server.
[**PrivateCheckpointHealthGetSpecifiedPrivateCheckpointHealth**](CheckpointApi.md#PrivateCheckpointHealthGetSpecifiedPrivateCheckpointHealth) | **Get** /PrivateCheckpointHealth | Returns the status of the specified private checkpoints.


# **CheckpointGetAllCheckpoints**
> CheckpointListResponse CheckpointGetAllCheckpoints(ctx, )
Returns all the checkpoints that are not deleted

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CheckpointListResponse**](CheckpointListResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckpointGetCheckpoint**
> CheckpointResponse CheckpointGetCheckpoint(ctx, checkpointId)
Returns the specified checkpoint, deleted or not

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **checkpointId** | **int32**| The Id of the requested checkpoint. | 

### Return type

[**CheckpointResponse**](CheckpointResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckpointRegionGetAllCheckpointRegions**
> []CheckpointRegion CheckpointRegionGetAllCheckpointRegions(ctx, )
Returns all the checkpoint regions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]CheckpointRegion**](CheckpointRegion.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckpointRegionGetCheckpointRegionCheckpoints**
> []Checkpoint CheckpointRegionGetCheckpointRegionCheckpoints(ctx, checkpointRegionId)
Returns the checkpoints for the specified checkpoint region.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **checkpointRegionId** | **int32**| The id for the specified checkpoint region. | 

### Return type

[**[]Checkpoint**](Checkpoint.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckpointRegionGetSpecifiedCheckpointRegion**
> CheckpointRegion CheckpointRegionGetSpecifiedCheckpointRegion(ctx, checkpointRegionId)
Returns the specified checkpoint region.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **checkpointRegionId** | **int32**| The id for the specified checkpoint region. | 

### Return type

[**CheckpointRegion**](CheckpointRegion.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckpointServerGetAllServerIpv4Addresses**
> ListStringResponse CheckpointServerGetAllServerIpv4Addresses(ctx, )
Anonymous call that returns all the IPv4 addresses of all the checkpoint servers.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListStringResponse**](ListStringResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckpointServerGetAllServerIpv6Addresses**
> ListStringResponse CheckpointServerGetAllServerIpv6Addresses(ctx, )
Anonymous call that returns all the IPv6 addresses of all the checkpoint servers.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListStringResponse**](ListStringResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckpointServerGetServer**
> CheckpoinServerResponse CheckpointServerGetServer(ctx, serverId)
Returns the requested checkpoint server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **int32**| The Id of the requested server. | 

### Return type

[**CheckpoinServerResponse**](CheckpoinServerResponse.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCheckpointHealthGetSpecifiedPrivateCheckpointHealth**
> CheckpointsHealth PrivateCheckpointHealthGetSpecifiedPrivateCheckpointHealth(ctx, optional)
Returns the status of the specified private checkpoints.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CheckpointApiPrivateCheckpointHealthGetSpecifiedPrivateCheckpointHealthOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CheckpointApiPrivateCheckpointHealthGetSpecifiedPrivateCheckpointHealthOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkpointName** | **optional.String**| The name/filter for the specified private checkpoints. | [default to ]

### Return type

[**CheckpointsHealth**](CheckpointsHealth.md)

### Authorization

[basicauth](../README.md#basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

