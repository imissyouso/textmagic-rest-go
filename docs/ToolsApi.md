# TextMagic\ToolsApi

All URIs are relative to *http://my.textmagic.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DoCarrierLookup**](ToolsApi.md#DoCarrierLookup) | **Get** /api/v2/lookups/{phone} | Carrier Lookup
[**DoEmailLookup**](ToolsApi.md#DoEmailLookup) | **Get** /api/v2/email-lookups/{email} | Validate Email address using Email Lookup tool


# **DoCarrierLookup**
> DoCarrierLookupResponse DoCarrierLookup(ctx, phone, optional)
Carrier Lookup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **phone** | **string**|  | 
 **optional** | ***DoCarrierLookupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DoCarrierLookupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **country** | **optional.String**| Country code for local formatted numbers | [default to US]

### Return type

[**DoCarrierLookupResponse**](DoCarrierLookupResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoEmailLookup**
> DoEmailLookupResponse DoEmailLookup(ctx, email)
Validate Email address using Email Lookup tool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**|  | 

### Return type

[**DoEmailLookupResponse**](DoEmailLookupResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

