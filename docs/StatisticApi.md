# TextMagic\StatisticApi

All URIs are relative to *http://my.textmagic.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMessagingCounters**](StatisticApi.md#GetMessagingCounters) | **Get** /api/v2/stats/messaging/data | Return counters for messaging data views.
[**GetMessagingStat**](StatisticApi.md#GetMessagingStat) | **Get** /api/v2/stats/messaging | Return messaging statistics.


# **GetMessagingCounters**
> GetMessagingCountersResponse GetMessagingCounters(ctx, )
Return counters for messaging data views.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetMessagingCountersResponse**](GetMessagingCountersResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagingStat**
> GetMessagingStatResponse GetMessagingStat(ctx, optional)
Return messaging statistics.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetMessagingStatOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMessagingStatOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **by** | **optional.String**| Group results by specified period: off, day, month or year. Default is off | [default to off]
 **start** | **optional.Int32**| Start date in unix timestamp format. Default is 7 days ago | 
 **end** | **optional.String**| End date in unix timestamp format. Default is now | 

### Return type

[**GetMessagingStatResponse**](GetMessagingStatResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

