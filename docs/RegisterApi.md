# \RegisterApi

All URIs are relative to *https://api.uptrends.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegisterPost**](RegisterApi.md#RegisterPost) | **Post** /Register | Creates a new API account.


# **RegisterPost**
> RegistrationResponse RegisterPost(ctx, optional)
Creates a new API account.

This method requires that you specify the username and password of an Uptrends operator login as authentication. When registration is successful, please save the UserName and Password fields from the response; these are your new API credentials.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegisterApiRegisterPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegisterApiRegisterPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **optional.String**| An optional description for the new API account, e.g. \&quot;API\&quot;. If this is empty, it will be defaulted to \&quot;API\&quot; | [default to API]
 **operatorGuid** | **optional.String**| The operator guid for which the new API account needs to be created. Leave empty to create an API acount for your own operator. | 

### Return type

[**RegistrationResponse**](RegistrationResponse.md)

### Authorization

[user-basicauth](../README.md#user-basicauth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

