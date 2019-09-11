# TextMagic\ContactNotesApi

All URIs are relative to *http://my.textmagic.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContactNote**](ContactNotesApi.md#CreateContactNote) | **Post** /api/v2/contacts/{id}/notes | Create a new contact note.
[**DeleteContactNote**](ContactNotesApi.md#DeleteContactNote) | **Delete** /api/v2/notes/{id} | Delete a single contact note.
[**DeleteContactNotesBulk**](ContactNotesApi.md#DeleteContactNotesBulk) | **Post** /api/v2/contacts/{id}/notes/delete | Delete contact note by given ID(s) or delete all contact notes.
[**GetContactNote**](ContactNotesApi.md#GetContactNote) | **Get** /api/v2/notes/{id} | Get a single contact note.
[**GetContactNotes**](ContactNotesApi.md#GetContactNotes) | **Get** /api/v2/contacts/{id}/notes | Fetch notes assigned to the given contact.
[**UpdateContactNote**](ContactNotesApi.md#UpdateContactNote) | **Put** /api/v2/notes/{id} | Update existing contact note.


# **CreateContactNote**
> ResourceLinkResponse CreateContactNote(ctx, createContactNoteInputObject, id, optional)
Create a new contact note.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createContactNoteInputObject** | [**CreateContactNoteInputObject**](CreateContactNoteInputObject.md)|  | 
  **id** | **int32**|  | 
 **optional** | ***CreateContactNoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateContactNoteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xIgnoreNullValues** | **optional.Bool**|  | [default to true]

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContactNote**
> DeleteContactNote(ctx, id)
Delete a single contact note.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContactNotesBulk**
> DeleteContactNotesBulk(ctx, id, deleteContactNotesBulkInputObject, optional)
Delete contact note by given ID(s) or delete all contact notes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **deleteContactNotesBulkInputObject** | [**DeleteContactNotesBulkInputObject**](DeleteContactNotesBulkInputObject.md)|  | 
 **optional** | ***DeleteContactNotesBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteContactNotesBulkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xIgnoreNullValues** | **optional.Bool**|  | [default to true]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactNote**
> ContactNote GetContactNote(ctx, id)
Get a single contact note.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ContactNote**](ContactNote.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactNotes**
> map[string]interface{} GetContactNotes(ctx, id, optional)
Fetch notes assigned to the given contact.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetContactNotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContactNotesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContactNote**
> ResourceLinkResponse UpdateContactNote(ctx, updateContactNoteInputObject, id, optional)
Update existing contact note.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateContactNoteInputObject** | [**UpdateContactNoteInputObject**](UpdateContactNoteInputObject.md)|  | 
  **id** | **int32**|  | 
 **optional** | ***UpdateContactNoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateContactNoteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xIgnoreNullValues** | **optional.Bool**|  | [default to true]

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

