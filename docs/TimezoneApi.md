# \TimezoneApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TimezoneGetAllTimezones**](TimezoneApi.md#TimezoneGetAllTimezones) | **Get** /Timezone | Gets all timezones available.
[**TimezoneGetTimezoneById**](TimezoneApi.md#TimezoneGetTimezoneById) | **Get** /Timezone/{timezoneId} | Gets the timezone with the specified Id.


# **TimezoneGetAllTimezones**
> []Timezone TimezoneGetAllTimezones(ctx, )
Gets all timezones available.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Timezone**](Timezone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TimezoneGetTimezoneById**
> Timezone TimezoneGetTimezoneById(ctx, timezoneId)
Gets the timezone with the specified Id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timezoneId** | **int32**|  | 

### Return type

[**Timezone**](Timezone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

