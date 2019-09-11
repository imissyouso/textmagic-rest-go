# TextMagic\OutboundMessagesApi

All URIs are relative to *http://my.textmagic.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAllOutboundMessages**](OutboundMessagesApi.md#DeleteAllOutboundMessages) | **Delete** /api/v2/message/all | Delete all messages
[**DeleteOutboundMessage**](OutboundMessagesApi.md#DeleteOutboundMessage) | **Delete** /api/v2/messages/{id} | Delete message
[**DeleteOutboundMessagesBulk**](OutboundMessagesApi.md#DeleteOutboundMessagesBulk) | **Post** /api/v2/messages/delete | Delete messages by IDs
[**GetAllOutboundMessages**](OutboundMessagesApi.md#GetAllOutboundMessages) | **Get** /api/v2/messages | Get all messages
[**GetMessagePreview**](OutboundMessagesApi.md#GetMessagePreview) | **Get** /api/v2/messages/preview | Preview message
[**GetMessagePrice**](OutboundMessagesApi.md#GetMessagePrice) | **Get** /api/v2/messages/price | Check price
[**GetMessagePrices**](OutboundMessagesApi.md#GetMessagePrices) | **Get** /api/v2/messages/prices | Get pricing
[**GetOutboundMessage**](OutboundMessagesApi.md#GetOutboundMessage) | **Get** /api/v2/messages/{id} | Get a single message
[**GetOutboundMessagesHistory**](OutboundMessagesApi.md#GetOutboundMessagesHistory) | **Get** /api/v2/history | Get history
[**SearchOutboundMessages**](OutboundMessagesApi.md#SearchOutboundMessages) | **Get** /api/v2/messages/search | Find messages
[**SendMessage**](OutboundMessagesApi.md#SendMessage) | **Post** /api/v2/messages | Send message
[**UploadMessageAttachment**](OutboundMessagesApi.md#UploadMessageAttachment) | **Post** /api/v2/messages/attachment | Upload message attachment


# **DeleteAllOutboundMessages**
> DeleteAllOutboundMessages(ctx, )
Delete all messages

Delete all messages.

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

# **DeleteOutboundMessage**
> DeleteOutboundMessage(ctx, id)
Delete message

Delete a single message.

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

# **DeleteOutboundMessagesBulk**
> DeleteOutboundMessagesBulk(ctx, deleteOutboundMessagesBulkInputObject, optional)
Delete messages by IDs

Delete outbound messages by given ID(s) or delete all outbound messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteOutboundMessagesBulkInputObject** | [**DeleteOutboundMessagesBulkInputObject**](DeleteOutboundMessagesBulkInputObject.md)|  | 
 **optional** | ***DeleteOutboundMessagesBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteOutboundMessagesBulkOpts struct

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

# **GetAllOutboundMessages**
> map[string]interface{} GetAllOutboundMessages(ctx, optional)
Get all messages

Get all user oubound messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllOutboundMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllOutboundMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **lastId** | **optional.Int32**| Filter results by ID, selecting all values lesser than the specified ID. Note that \\&#39;page\\&#39; parameter is ignored when \\&#39;lastId\\&#39; is specified | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagePreview**
> GetMessagePreviewResponse GetMessagePreview(ctx, optional)
Preview message

Get messages preview (with tags merged) up to 100 messages per session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetMessagePreviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMessagePreviewOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **optional.String**| Message text. Required if template_id is not set | 
 **templateId** | **optional.Int32**| Template used instead of message text. Required if text is not set | 
 **sendingTime** | **optional.Int32**| DEPRECATED, consider using sendingDateTime and sendingTimezone parameters instead: Optional (required with rrule set). Message sending time in unix timestamp format. Default is now | 
 **sendingDateTime** | **optional.String**| Sending time in Y-m-d H:i:s format (e.g. 2016-05-27 13:02:33). This time is relative to sendingTimezone | 
 **sendingTimezone** | **optional.String**| ID or ISO-name of timezone used for sending when sendingDateTime parameter is set. E.g. if you specify sendingDateTime &#x3D; \\\&quot;2016-05-27 13:02:33\\\&quot; and sendingTimezone &#x3D; \\\&quot;America/Buenos_Aires\\\&quot;, your message will be sent at May 27, 2016 13:02:33 Buenos Aires time, or 16:02:33 UTC. Default is account timezone | 
 **contacts** | **optional.String**| Comma separated array of contact resources id message will be sent to | 
 **lists** | **optional.String**| Comma separated array of list resources id message will be sent to | 
 **phones** | **optional.String**| Comma separated array of E.164 phone numbers message will be sent to | 
 **cutExtra** | **optional.Int32**| Should sending method cut extra characters which not fit supplied partsCount or return 400 Bad request response instead. Default is 0 | [default to 0]
 **partsCount** | **optional.Int32**| Maximum message parts count (TextMagic allows sending 1 to 6 message parts). Default is 6 | [default to 6]
 **referenceId** | **optional.Int32**| Custom message reference id which can be used in your application infrastructure | 
 **from** | **optional.String**| One of allowed Sender ID (phone number or alphanumeric sender ID). If specified Sender ID is not allowed for some destinations, a fallback default Sender ID will be used to ensure delivery | 
 **rule** | **optional.String**| iCal RRULE parameter to create recurrent scheduled messages. When used, sendingTime is mandatory as start point of sending. See https://www.textmagic.com/free-tools/rrule-generator for format details | 
 **createChat** | **optional.Int32**| Should sending method try to create new Chat(if not exist) with specified recipients. Default is 0 | [default to 0]
 **tts** | **optional.Int32**| Send Text to Speech message. Default is 0 | [default to 0]
 **local** | **optional.Int32**| Treat phone numbers passed in \\&#39;phones\\&#39; field as local. Default is 0 | [default to 0]
 **localCountry** | **optional.String**| 2-letter ISO country code for local phone numbers, used when \\&#39;local\\&#39; is set to true. Default is account country | 

### Return type

[**GetMessagePreviewResponse**](GetMessagePreviewResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagePrice**
> GetMessagePriceResponse GetMessagePrice(ctx, optional)
Check price

Check pricing for a new outbound message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetMessagePriceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMessagePriceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeBlocked** | **optional.Int32**| Should we show pricing for the blocked contacts. | [default to 0]
 **text** | **optional.String**| Message text. Required if template_id is not set | 
 **templateId** | **optional.Int32**| Template used instead of message text. Required if text is not set | 
 **sendingTime** | **optional.Int32**| DEPRECATED, consider using sendingDateTime and sendingTimezone parameters instead: Optional (required with rrule set). Message sending time in unix timestamp format. Default is now | 
 **sendingDateTime** | **optional.String**| Sending time in Y-m-d H:i:s format (e.g. 2016-05-27 13:02:33). This time is relative to sendingTimezone | 
 **sendingTimezone** | **optional.String**| ID or ISO-name of timezone used for sending when sendingDateTime parameter is set. E.g. if you specify sendingDateTime &#x3D; \\\&quot;2016-05-27 13:02:33\\\&quot; and sendingTimezone &#x3D; \\\&quot;America/Buenos_Aires\\\&quot;, your message will be sent at May 27, 2016 13:02:33 Buenos Aires time, or 16:02:33 UTC. Default is account timezone | 
 **contacts** | **optional.String**| Comma separated array of contact resources id message will be sent to | 
 **lists** | **optional.String**| Comma separated array of list resources id message will be sent to | 
 **phones** | **optional.String**| Comma separated array of E.164 phone numbers message will be sent to | 
 **cutExtra** | **optional.Int32**| Should sending method cut extra characters which not fit supplied partsCount or return 400 Bad request response instead. Default is 0 | [default to 0]
 **partsCount** | **optional.Int32**| Maximum message parts count (TextMagic allows sending 1 to 6 message parts). Default is 6 | [default to 6]
 **referenceId** | **optional.Int32**| Custom message reference id which can be used in your application infrastructure | 
 **from** | **optional.String**| One of allowed Sender ID (phone number or alphanumeric sender ID). If specified Sender ID is not allowed for some destinations, a fallback default Sender ID will be used to ensure delivery | 
 **rule** | **optional.String**| iCal RRULE parameter to create recurrent scheduled messages. When used, sendingTime is mandatory as start point of sending. See https://www.textmagic.com/free-tools/rrule-generator for format details | 
 **createChat** | **optional.Int32**| Should sending method try to create new Chat(if not exist) with specified recipients. Default is 0 | [default to 0]
 **tts** | **optional.Int32**| Send Text to Speech message. Default is 0 | [default to 0]
 **local** | **optional.Int32**| Treat phone numbers passed in \\&#39;phones\\&#39; field as local. Default is 0 | [default to 0]
 **localCountry** | **optional.String**| 2-letter ISO country code for local phone numbers, used when \\&#39;local\\&#39; is set to true. Default is account country | 

### Return type

[**GetMessagePriceResponse**](GetMessagePriceResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagePrices**
> GetMessagePricesResponse GetMessagePrices(ctx, )
Get pricing

Get message prices for all countries.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetMessagePricesResponse**](GetMessagePricesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOutboundMessage**
> MessageOut GetOutboundMessage(ctx, id)
Get a single message

Get a single outgoing message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**MessageOut**](MessageOut.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOutboundMessagesHistory**
> map[string]interface{} GetOutboundMessagesHistory(ctx, optional)
Get history

Get outbound messages history.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetOutboundMessagesHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetOutboundMessagesHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **lastId** | **optional.Int32**| Filter results by ID, selecting all values lesser than the specified ID. | 
 **query** | **optional.String**| Find message by specified search query | 
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

# **SearchOutboundMessages**
> map[string]interface{} SearchOutboundMessages(ctx, optional)
Find messages

Find outbound messages by given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchOutboundMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchOutboundMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **lastId** | **optional.Int32**| Filter results by ID, selecting all values lesser than the specified ID. Note that \\&#39;page\\&#39; parameter is ignored when \\&#39;lastId\\&#39; is specified | 
 **ids** | **optional.String**| Find message by ID(s) | 
 **sessionId** | **optional.Int32**| Find messages by session ID | 
 **statuses** | **optional.String**| Find messages by status | 
 **includeDeleted** | **optional.Int32**| Search also in deleted messages | [default to 0]
 **query** | **optional.String**| Find messages by specified search query | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendMessage**
> SendMessageResponse SendMessage(ctx, sendMessageInputObject, optional)
Send message

The main entrypoint to send messages. See examples above for the reference.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sendMessageInputObject** | [**SendMessageInputObject**](SendMessageInputObject.md)|  | 
 **optional** | ***SendMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SendMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIgnoreNullValues** | **optional.Bool**|  | [default to true]

### Return type

[**SendMessageResponse**](SendMessageResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadMessageAttachment**
> UploadMessageAttachmentResponse UploadMessageAttachment(ctx, file)
Upload message attachment

Upload a new file to insert it as a link.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File**| Attachment. Supports .jpg, .gif, .png, .pdf, .txt, .csv, .doc, .docx, .xls, .xlsx, .ppt, .pptx &amp; .vcf file formats | 

### Return type

[**UploadMessageAttachmentResponse**](UploadMessageAttachmentResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

