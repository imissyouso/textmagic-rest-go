# TextMagic\ContactsApi

All URIs are relative to *http://my.textmagic.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockContact**](ContactsApi.md#BlockContact) | **Post** /api/v2/contacts/block | Block contact from inbound and outbound communication by phone number.
[**CreateContact**](ContactsApi.md#CreateContact) | **Post** /api/v2/contacts | Create a new contact from the submitted data.
[**DeleteAllContacts**](ContactsApi.md#DeleteAllContacts) | **Delete** /api/v2/contact/all | Delete all contacts.
[**DeleteContact**](ContactsApi.md#DeleteContact) | **Delete** /api/v2/contacts/{id} | Delete a single contact.
[**DeleteContactAvatar**](ContactsApi.md#DeleteContactAvatar) | **Delete** /api/v2/contacts/{id}/avatar | Delete an avatar for the contact.
[**DeleteContactsByIds**](ContactsApi.md#DeleteContactsByIds) | **Post** /api/v2/contacts/delete | Delete contact by given ID(s) or delete all contacts.
[**GetBlockedContacts**](ContactsApi.md#GetBlockedContacts) | **Get** /api/v2/contacts/block/list | Get blocked contacts.
[**GetContact**](ContactsApi.md#GetContact) | **Get** /api/v2/contacts/{id} | Get a single contact.
[**GetContactByPhone**](ContactsApi.md#GetContactByPhone) | **Get** /api/v2/contacts/phone/{phone} | Get a single contact by phone number.
[**GetContactIfBlocked**](ContactsApi.md#GetContactIfBlocked) | **Get** /api/v2/contacts/block/phone | Check is that phone number blocked
[**GetContacts**](ContactsApi.md#GetContacts) | **Get** /api/v2/contacts | Get all user contacts.
[**GetContactsAutocomplete**](ContactsApi.md#GetContactsAutocomplete) | **Get** /api/v2/contacts/autocomplete | Get contacts autocomplete suggestions by given search term.
[**GetFavourites**](ContactsApi.md#GetFavourites) | **Get** /api/v2/contacts/favorite | Get favorite contacts and lists.
[**GetUnsubscribedContact**](ContactsApi.md#GetUnsubscribedContact) | **Get** /api/v2/unsubscribers/{id} | Get a single unsubscribed contact.
[**GetUnsubscribers**](ContactsApi.md#GetUnsubscribers) | **Get** /api/v2/unsubscribers | Get all contact have unsubscribed from your communication.
[**SearchContacts**](ContactsApi.md#SearchContacts) | **Get** /api/v2/contacts/search | Find user contacts by given parameters.
[**UnblockContact**](ContactsApi.md#UnblockContact) | **Post** /api/v2/contacts/unblock | Unblock contact by phone number.
[**UnblockContactsBulk**](ContactsApi.md#UnblockContactsBulk) | **Post** /api/v2/contacts/unblock/bulk | Unblock several contacts by blocked contact ids or unblock all contacts
[**UnsubscribeContact**](ContactsApi.md#UnsubscribeContact) | **Post** /api/v2/unsubscribers | Unsubscribe contact from your communication by phone number.
[**UpdateContact**](ContactsApi.md#UpdateContact) | **Put** /api/v2/contacts/{id} | Update existing contact.
[**UploadContactAvatar**](ContactsApi.md#UploadContactAvatar) | **Post** /api/v2/contacts/{id}/avatar | Add an avatar for the contact.


# **BlockContact**
> ResourceLinkResponse BlockContact(ctx, blockContactInputObject, optional)
Block contact from inbound and outbound communication by phone number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockContactInputObject** | [**BlockContactInputObject**](BlockContactInputObject.md)|  | 
 **optional** | ***BlockContactOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlockContactOpts struct

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

# **CreateContact**
> ResourceLinkResponse CreateContact(ctx, createContactInputObject, optional)
Create a new contact from the submitted data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createContactInputObject** | [**CreateContactInputObject**](CreateContactInputObject.md)|  | 
 **optional** | ***CreateContactOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateContactOpts struct

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

# **DeleteAllContacts**
> DeleteAllContacts(ctx, )
Delete all contacts.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContact**
> DeleteContact(ctx, id)
Delete a single contact.

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

# **DeleteContactAvatar**
> DeleteContactAvatar(ctx, id)
Delete an avatar for the contact.

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

# **DeleteContactsByIds**
> DeleteContactsByIds(ctx, deleteContactsByIdsInputObject, optional)
Delete contact by given ID(s) or delete all contacts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteContactsByIdsInputObject** | [**DeleteContactsByIdsInputObject**](DeleteContactsByIdsInputObject.md)|  | 
 **optional** | ***DeleteContactsByIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteContactsByIdsOpts struct

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

# **GetBlockedContacts**
> map[string]interface{} GetBlockedContacts(ctx, optional)
Get blocked contacts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBlockedContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBlockedContactsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **query** | **optional.String**| Find blocked contacts by specified search query | 
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

# **GetContact**
> Contact GetContact(ctx, id)
Get a single contact.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The contact id | 

### Return type

[**Contact**](Contact.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactByPhone**
> Contact GetContactByPhone(ctx, phone)
Get a single contact by phone number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **phone** | **string**|  | 

### Return type

[**Contact**](Contact.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactIfBlocked**
> Contact GetContactIfBlocked(ctx, phone)
Check is that phone number blocked

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **phone** | **string**| Phone number to check | 

### Return type

[**Contact**](Contact.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContacts**
> map[string]interface{} GetContacts(ctx, optional)
Get all user contacts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContactsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **shared** | **optional.Int32**| Should shared contacts to be included | [default to 0]
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

# **GetContactsAutocomplete**
> []GetContactsAutocompleteResponse GetContactsAutocomplete(ctx, query, optional)
Get contacts autocomplete suggestions by given search term.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| Find recipients by specified search query | 
 **optional** | ***GetContactsAutocompleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContactsAutocompleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **lists** | **optional.Int32**| Should lists be returned or not | [default to 0]

### Return type

[**[]GetContactsAutocompleteResponse**](GetContactsAutocompleteResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFavourites**
> map[string]interface{} GetFavourites(ctx, optional)
Get favorite contacts and lists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetFavouritesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFavouritesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **query** | **optional.String**| Find contacts or lists by specified search query | [default to A]

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUnsubscribedContact**
> UnsubscribedContact GetUnsubscribedContact(ctx, id)
Get a single unsubscribed contact.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**UnsubscribedContact**](UnsubscribedContact.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUnsubscribers**
> map[string]interface{} GetUnsubscribers(ctx, optional)
Get all contact have unsubscribed from your communication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUnsubscribersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUnsubscribersOpts struct

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

# **SearchContacts**
> map[string]interface{} SearchContacts(ctx, optional)
Find user contacts by given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchContactsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **shared** | **optional.Int32**| Should shared contacts to be included | [default to 0]
 **ids** | **optional.String**| Find contact by ID(s) | 
 **listId** | **optional.Int32**| Find contact by List ID | 
 **includeBlocked** | **optional.Int32**| Should blocked contacts to be included | 
 **query** | **optional.String**| Find contacts by specified search query | 
 **local** | **optional.Int32**| Treat phone number passed in &#39;query&#39; field as local. Default is 0 | [default to 0]
 **country** | **optional.String**| 2-letter ISO country code for local phone numbers, used when &#39;local&#39; is set to true. Default is account country | 
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

# **UnblockContact**
> UnblockContact(ctx, unblockContactInputObject, optional)
Unblock contact by phone number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unblockContactInputObject** | [**UnblockContactInputObject**](UnblockContactInputObject.md)|  | 
 **optional** | ***UnblockContactOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UnblockContactOpts struct

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

# **UnblockContactsBulk**
> UnblockContactsBulk(ctx, unblockContactsBulkInputObject, optional)
Unblock several contacts by blocked contact ids or unblock all contacts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unblockContactsBulkInputObject** | [**UnblockContactsBulkInputObject**](UnblockContactsBulkInputObject.md)|  | 
 **optional** | ***UnblockContactsBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UnblockContactsBulkOpts struct

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

# **UnsubscribeContact**
> ResourceLinkResponse UnsubscribeContact(ctx, unsubscribeContactInputObject, optional)
Unsubscribe contact from your communication by phone number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unsubscribeContactInputObject** | [**UnsubscribeContactInputObject**](UnsubscribeContactInputObject.md)|  | 
 **optional** | ***UnsubscribeContactOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UnsubscribeContactOpts struct

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

# **UpdateContact**
> ResourceLinkResponse UpdateContact(ctx, updateContactInputObject, id, optional)
Update existing contact.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateContactInputObject** | [**UpdateContactInputObject**](UpdateContactInputObject.md)|  | 
  **id** | **int32**|  | 
 **optional** | ***UpdateContactOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateContactOpts struct

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

# **UploadContactAvatar**
> ResourceLinkResponse UploadContactAvatar(ctx, image, id)
Add an avatar for the contact.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **image** | ***os.File**| Contact avatar. Should be PNG or JPG file not more than 10 MB | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

