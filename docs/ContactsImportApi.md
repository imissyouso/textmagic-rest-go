# TextMagic\ContactsImportApi

All URIs are relative to *http://my.textmagic.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContactImportSessionProgress**](ContactsImportApi.md#GetContactImportSessionProgress) | **Get** /api/v2/contacts/import/progress/{id} | Get contact import session progress.


# **GetContactImportSessionProgress**
> GetContactImportSessionProgressResponse GetContactImportSessionProgress(ctx, id)
Get contact import session progress.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**GetContactImportSessionProgressResponse**](GetContactImportSessionProgressResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

