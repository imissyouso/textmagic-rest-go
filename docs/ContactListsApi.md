# TextMagic\ContactListsApi

All URIs are relative to *http://my.textmagic.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignContactsToList**](ContactListsApi.md#AssignContactsToList) | **Put** /api/v2/lists/{id}/contacts | Assign contacts to the specified list.
[**ClearAndAssignContactsToList**](ContactListsApi.md#ClearAndAssignContactsToList) | **Post** /api/v2/lists/{id}/contacts | Reset list members to the specified contacts.
[**CreateList**](ContactListsApi.md#CreateList) | **Post** /api/v2/lists | Create a new list from the submitted data.
[**DeleteContactsFromList**](ContactListsApi.md#DeleteContactsFromList) | **Delete** /api/v2/lists/{id}/contacts | Unassign contacts from the specified list.
[**DeleteList**](ContactListsApi.md#DeleteList) | **Delete** /api/v2/lists/{id} | Delete a single list.
[**DeleteListAvatar**](ContactListsApi.md#DeleteListAvatar) | **Delete** /api/v2/lists/{id}/avatar | Delete an avatar for the list.
[**DeleteListContactsBulk**](ContactListsApi.md#DeleteListContactsBulk) | **Post** /api/v2/lists/{id}/contacts/delete | Delete contact from list by given ID(s) or all contacts from list.
[**DeleteListsBulk**](ContactListsApi.md#DeleteListsBulk) | **Post** /api/v2/lists/delete | Delete list by given ID(s) or delete all lists.
[**GetContactsByListId**](ContactListsApi.md#GetContactsByListId) | **Get** /api/v2/lists/{id}/contacts | Fetch user contacts by given group id.
[**GetList**](ContactListsApi.md#GetList) | **Get** /api/v2/lists/{id} | Get a single list.
[**GetListContactsIds**](ContactListsApi.md#GetListContactsIds) | **Get** /api/v2/lists/{id}/contacts/ids | Fetch all contacts IDs belonging to the list with ID.
[**GetListsOfContact**](ContactListsApi.md#GetListsOfContact) | **Get** /api/v2/contacts/{id}/lists | Return lists which contact belongs to.
[**GetUserLists**](ContactListsApi.md#GetUserLists) | **Get** /api/v2/lists | Get all user lists.
[**SearchLists**](ContactListsApi.md#SearchLists) | **Get** /api/v2/lists/search | Find contact lists by given parameters.
[**UpdateList**](ContactListsApi.md#UpdateList) | **Put** /api/v2/lists/{id} | Update existing list.
[**UploadListAvatar**](ContactListsApi.md#UploadListAvatar) | **Post** /api/v2/lists/{id}/avatar | Add an avatar for the list.


# **AssignContactsToList**
> ResourceLinkResponse AssignContactsToList(ctx, assignContactsToListInputObject, id)
Assign contacts to the specified list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assignContactsToListInputObject** | [**AssignContactsToListInputObject**](AssignContactsToListInputObject.md)| Contact ID(s), separated by comma or &#39;all&#39; to add all contacts belonging to the current user | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearAndAssignContactsToList**
> ResourceLinkResponse ClearAndAssignContactsToList(ctx, clearAndAssignContactsToListInputObject, id)
Reset list members to the specified contacts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clearAndAssignContactsToListInputObject** | [**ClearAndAssignContactsToListInputObject**](ClearAndAssignContactsToListInputObject.md)| Contact ID(s), separated by comma or &#39;all&#39; to add all contacts belonging to the current user | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateList**
> ResourceLinkResponse CreateList(ctx, createListInputObject, optional)
Create a new list from the submitted data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createListInputObject** | [**CreateListInputObject**](CreateListInputObject.md)|  | 
 **optional** | ***CreateListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateListOpts struct

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

# **DeleteContactsFromList**
> DeleteContactsFromList(ctx, deleteContacsFromListObject, id)
Unassign contacts from the specified list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteContacsFromListObject** | [**DeleteContacsFromListObject**](DeleteContacsFromListObject.md)| Contact ID(s), separated by comma | 
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteList**
> DeleteList(ctx, id)
Delete a single list.

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

# **DeleteListAvatar**
> DeleteListAvatar(ctx, id)
Delete an avatar for the list.

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
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListContactsBulk**
> DeleteListContactsBulk(ctx, deleteListContactsBulkInputObject, id, optional)
Delete contact from list by given ID(s) or all contacts from list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteListContactsBulkInputObject** | [**DeleteListContactsBulkInputObject**](DeleteListContactsBulkInputObject.md)|  | 
  **id** | **int32**|  | 
 **optional** | ***DeleteListContactsBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteListContactsBulkOpts struct

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

# **DeleteListsBulk**
> DeleteListsBulk(ctx, deleteListsBulkInputObject, optional)
Delete list by given ID(s) or delete all lists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteListsBulkInputObject** | [**DeleteListsBulkInputObject**](DeleteListsBulkInputObject.md)|  | 
 **optional** | ***DeleteListsBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteListsBulkOpts struct

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

# **GetContactsByListId**
> map[string]interface{} GetContactsByListId(ctx, id, optional)
Fetch user contacts by given group id.

A useful synonym for \"contacts/search\" command with provided \"listId\" parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Given group Id. | 
 **optional** | ***GetContactsByListIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContactsByListIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetList**
> Group GetList(ctx, id)
Get a single list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**Group**](Group.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListContactsIds**
> GetListContactsIdsResponse GetListContactsIds(ctx, id)
Fetch all contacts IDs belonging to the list with ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**GetListContactsIdsResponse**](GetListContactsIdsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsOfContact**
> map[string]interface{} GetListsOfContact(ctx, id, optional)
Return lists which contact belongs to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetListsOfContactOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetListsOfContactOpts struct

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

# **GetUserLists**
> map[string]interface{} GetUserLists(ctx, optional)
Get all user lists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUserListsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUserListsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]
 **favoriteOnly** | **optional.Int32**| Return only favorite lists | [default to 0]
 **onlyMine** | **optional.Int32**| Return only current user lists | [default to 0]

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchLists**
> map[string]interface{} SearchLists(ctx, optional)
Find contact lists by given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchListsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchListsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **ids** | **optional.String**| Find lists by ID(s) | 
 **query** | **optional.String**| Find lists by specified search query | 
 **onlyMine** | **optional.Int32**| Return only current user lists | [default to 0]
 **onlyDefault** | **optional.Int32**| Return only default lists | [default to 0]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateList**
> ResourceLinkResponse UpdateList(ctx, id, optional)
Update existing list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***UpdateListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateListObject** | [**optional.Interface of UpdateListObject**](UpdateListObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadListAvatar**
> ResourceLinkResponse UploadListAvatar(ctx, image, id)
Add an avatar for the list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **image** | ***os.File**| List avatar. Should be PNG or JPG file not more than 10 MB | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

