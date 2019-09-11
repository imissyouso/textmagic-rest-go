# TextMagic\CommonApi

All URIs are relative to *http://my.textmagic.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCountries**](CommonApi.md#GetCountries) | **Get** /api/v2/countries | Return list of countries.
[**GetState**](CommonApi.md#GetState) | **Get** /api/v2/state | Get current entities state
[**GetTimezones**](CommonApi.md#GetTimezones) | **Get** /api/v2/timezones | Return all available timezone IDs.
[**GetVersions**](CommonApi.md#GetVersions) | **Get** /api/v2/versions | Get minimal valid apps versions
[**Ping**](CommonApi.md#Ping) | **Get** /api/v2/ping | Just does a pong.


# **GetCountries**
> []Country GetCountries(ctx, )
Return list of countries.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Country**](Country.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetState**
> GetStateResponse GetState(ctx, )
Get current entities state

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetStateResponse**](GetStateResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimezones**
> interface{} GetTimezones(ctx, optional)
Return all available timezone IDs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTimezonesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTimezonesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **full** | **optional.Int32**| Return full info about timezones in array (0 or 1). Default is 0 | [default to 0]

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVersions**
> GetVersionsResponse GetVersions(ctx, )
Get minimal valid apps versions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetVersionsResponse**](GetVersionsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ping**
> PingResponse Ping(ctx, )
Just does a pong.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PingResponse**](PingResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

