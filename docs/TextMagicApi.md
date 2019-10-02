# TextMagic\TextMagicApi

All URIs are relative to *http://my.textmagic.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignContactsToList**](TextMagicApi.md#AssignContactsToList) | **Put** /api/v2/lists/{id}/contacts | Assign contacts to the specified list.
[**BlockContact**](TextMagicApi.md#BlockContact) | **Post** /api/v2/contacts/block | Block contact from inbound and outbound communication by phone number.
[**BuyDedicatedNumber**](TextMagicApi.md#BuyDedicatedNumber) | **Post** /api/v2/numbers | Buy a dedicated number
[**CancelSurvey**](TextMagicApi.md#CancelSurvey) | **Put** /api/v2/surveys/{id}/cancel | Cancel a survey.
[**CancelVerification**](TextMagicApi.md#CancelVerification) | **Delete** /api/v2/verify/{verifyId} | Cancel verification process
[**CheckPhoneVerificationCode**](TextMagicApi.md#CheckPhoneVerificationCode) | **Put** /api/v2/user/phone/verification | Check user phone verification code
[**CheckPhoneVerificationCode_0**](TextMagicApi.md#CheckPhoneVerificationCode_0) | **Put** /api/v2/verify | Step 2: Check the verification code 
[**ClearAndAssignContactsToList**](TextMagicApi.md#ClearAndAssignContactsToList) | **Post** /api/v2/lists/{id}/contacts | Reset list members to the specified contacts.
[**CloseChatsBulk**](TextMagicApi.md#CloseChatsBulk) | **Post** /api/v2/chats/close/bulk | Close chats (bulk)
[**CloseReadChats**](TextMagicApi.md#CloseReadChats) | **Post** /api/v2/chats/close/read | Close read chats
[**CloseSubaccount**](TextMagicApi.md#CloseSubaccount) | **Delete** /api/v2/subaccounts/{id} | Close subaccount.
[**CreateContact**](TextMagicApi.md#CreateContact) | **Post** /api/v2/contacts | Create a new contact from the submitted data.
[**CreateContactNote**](TextMagicApi.md#CreateContactNote) | **Post** /api/v2/contacts/{id}/notes | Create a new contact note.
[**CreateCustomField**](TextMagicApi.md#CreateCustomField) | **Post** /api/v2/customfields | Create a new custom field from the submitted data.
[**CreateList**](TextMagicApi.md#CreateList) | **Post** /api/v2/lists | Create a new list from the submitted data.
[**CreatePushToken**](TextMagicApi.md#CreatePushToken) | **Post** /api/v2/push/tokens | Add or update a device token.
[**CreateSurvey**](TextMagicApi.md#CreateSurvey) | **Post** /api/v2/surveys | Create a new survey from the submitted data.
[**CreateSurveyNode**](TextMagicApi.md#CreateSurveyNode) | **Post** /api/v2/surveys/{id}/nodes | Create a new node from the submitted data.
[**CreateTemplate**](TextMagicApi.md#CreateTemplate) | **Post** /api/v2/templates | Create a template
[**DeleteAllContacts**](TextMagicApi.md#DeleteAllContacts) | **Delete** /api/v2/contact/all | Delete all contacts.
[**DeleteAllOutboundMessages**](TextMagicApi.md#DeleteAllOutboundMessages) | **Delete** /api/v2/message/all | Delete all messages
[**DeleteAvatar**](TextMagicApi.md#DeleteAvatar) | **Delete** /api/v2/user/avatar | Delete an avatar for the current user.\\
[**DeleteChatMessages**](TextMagicApi.md#DeleteChatMessages) | **Post** /api/v2/chats/{id}/messages/delete | Delete chat messages by ID(s)
[**DeleteChatsBulk**](TextMagicApi.md#DeleteChatsBulk) | **Post** /api/v2/chats/delete | Delete chats (bulk)
[**DeleteContact**](TextMagicApi.md#DeleteContact) | **Delete** /api/v2/contacts/{id} | Delete a single contact.
[**DeleteContactAvatar**](TextMagicApi.md#DeleteContactAvatar) | **Delete** /api/v2/contacts/{id}/avatar | Delete an avatar for the contact.
[**DeleteContactNote**](TextMagicApi.md#DeleteContactNote) | **Delete** /api/v2/notes/{id} | Delete a single contact note.
[**DeleteContactNotesBulk**](TextMagicApi.md#DeleteContactNotesBulk) | **Post** /api/v2/contacts/{id}/notes/delete | Delete contact note by given ID(s) or delete all contact notes.
[**DeleteContactsByIds**](TextMagicApi.md#DeleteContactsByIds) | **Post** /api/v2/contacts/delete | Delete contact by given ID(s) or delete all contacts.
[**DeleteContactsFromList**](TextMagicApi.md#DeleteContactsFromList) | **Delete** /api/v2/lists/{id}/contacts | Unassign contacts from the specified list.
[**DeleteCustomField**](TextMagicApi.md#DeleteCustomField) | **Delete** /api/v2/customfields/{id} | Delete a single custom field.
[**DeleteDedicatedNumber**](TextMagicApi.md#DeleteDedicatedNumber) | **Delete** /api/v2/numbers/{id} | Cancel dedicated number subscription
[**DeleteInboundMessage**](TextMagicApi.md#DeleteInboundMessage) | **Delete** /api/v2/replies/{id} | Delete a single inbound message
[**DeleteInboundMessagesBulk**](TextMagicApi.md#DeleteInboundMessagesBulk) | **Post** /api/v2/replies/delete | Delete inbound messages (bulk)
[**DeleteList**](TextMagicApi.md#DeleteList) | **Delete** /api/v2/lists/{id} | Delete a single list.
[**DeleteListAvatar**](TextMagicApi.md#DeleteListAvatar) | **Delete** /api/v2/lists/{id}/avatar | Delete an avatar for the list.
[**DeleteListContactsBulk**](TextMagicApi.md#DeleteListContactsBulk) | **Post** /api/v2/lists/{id}/contacts/delete | Delete contact from list by given ID(s) or all contacts from list.
[**DeleteListsBulk**](TextMagicApi.md#DeleteListsBulk) | **Post** /api/v2/lists/delete | Delete list by given ID(s) or delete all lists.
[**DeleteMessageSession**](TextMagicApi.md#DeleteMessageSession) | **Delete** /api/v2/sessions/{id} | Delete a session
[**DeleteMessageSessionsBulk**](TextMagicApi.md#DeleteMessageSessionsBulk) | **Post** /api/v2/sessions/delete | Delete sessions (bulk)
[**DeleteOutboundMessage**](TextMagicApi.md#DeleteOutboundMessage) | **Delete** /api/v2/messages/{id} | Delete message
[**DeleteOutboundMessagesBulk**](TextMagicApi.md#DeleteOutboundMessagesBulk) | **Post** /api/v2/messages/delete | Delete messages (bulk)
[**DeletePushToken**](TextMagicApi.md#DeletePushToken) | **Delete** /api/v2/push/tokens/{type}/{deviceId} | Delete a push notification device token.
[**DeleteScheduledMessage**](TextMagicApi.md#DeleteScheduledMessage) | **Delete** /api/v2/schedules/{id} | Delete a single scheduled message
[**DeleteScheduledMessagesBulk**](TextMagicApi.md#DeleteScheduledMessagesBulk) | **Post** /api/v2/schedules/delete | Delete scheduled messages (bulk)
[**DeleteSenderId**](TextMagicApi.md#DeleteSenderId) | **Delete** /api/v2/senderids/{id} | Delete a Sender ID
[**DeleteSurvey**](TextMagicApi.md#DeleteSurvey) | **Delete** /api/v2/surveys/{id} | Delete a survey.
[**DeleteSurveyNode**](TextMagicApi.md#DeleteSurveyNode) | **Delete** /api/v2/surveys/nodes/{id} | Delete a node.
[**DeleteTemplate**](TextMagicApi.md#DeleteTemplate) | **Delete** /api/v2/templates/{id} | Delete a template
[**DeleteTemplatesBulk**](TextMagicApi.md#DeleteTemplatesBulk) | **Post** /api/v2/templates/delete | Delete templates (bulk)
[**DoAuth**](TextMagicApi.md#DoAuth) | **Post** /api/v2/auth | Authenticate user by given username and password.
[**DoCarrierLookup**](TextMagicApi.md#DoCarrierLookup) | **Get** /api/v2/lookups/{phone} | Carrier Lookup
[**DoEmailLookup**](TextMagicApi.md#DoEmailLookup) | **Get** /api/v2/email-lookups/{email} | Validate Email address using Email Lookup tool
[**DuplicateSurvey**](TextMagicApi.md#DuplicateSurvey) | **Put** /api/v2/surveys/{id}/duplicate | Duplicate a survey.
[**GetAllBulkSessions**](TextMagicApi.md#GetAllBulkSessions) | **Get** /api/v2/bulks | Get all bulk sending sessions.
[**GetAllChats**](TextMagicApi.md#GetAllChats) | **Get** /api/v2/chats | Get all chats
[**GetAllInboundMessages**](TextMagicApi.md#GetAllInboundMessages) | **Get** /api/v2/replies | Get all inbound messages
[**GetAllMessageSessions**](TextMagicApi.md#GetAllMessageSessions) | **Get** /api/v2/sessions | Get all sessions
[**GetAllOutboundMessages**](TextMagicApi.md#GetAllOutboundMessages) | **Get** /api/v2/messages | Get all messages
[**GetAllScheduledMessages**](TextMagicApi.md#GetAllScheduledMessages) | **Get** /api/v2/schedules | Get all scheduled messages
[**GetAllTemplates**](TextMagicApi.md#GetAllTemplates) | **Get** /api/v2/templates | Get all templates
[**GetAvailableDedicatedNumbers**](TextMagicApi.md#GetAvailableDedicatedNumbers) | **Get** /api/v2/numbers/available | Find dedicated numbers available for purchase
[**GetAvailableSenderSettingOptions**](TextMagicApi.md#GetAvailableSenderSettingOptions) | **Get** /api/v2/sources | Get available sender settings
[**GetBalanceNotificationOptions**](TextMagicApi.md#GetBalanceNotificationOptions) | **Get** /api/v2/user/notification/balance/bundles | Returns the list of available balance options which can be used as a bound to determine when to send email to user with low balance notification. See https://my.textmagic.com/online/account/notifications/balance
[**GetBalanceNotificationSettings**](TextMagicApi.md#GetBalanceNotificationSettings) | **Get** /api/v2/user/notification/balance | Get balance notification settings
[**GetBlockedContacts**](TextMagicApi.md#GetBlockedContacts) | **Get** /api/v2/contacts/block/list | Get blocked contacts.
[**GetBulkSession**](TextMagicApi.md#GetBulkSession) | **Get** /api/v2/bulks/{id} | Get bulk message session status.
[**GetCallbackSettings**](TextMagicApi.md#GetCallbackSettings) | **Get** /api/v2/callback/settings | Fetch callback URL settings
[**GetCallsPrices**](TextMagicApi.md#GetCallsPrices) | **Get** /api/v2/calls/price | Check pricing for a inbound/outbound call.
[**GetChat**](TextMagicApi.md#GetChat) | **Get** /api/v2/chats/{id} | Get a single chat
[**GetChatByPhone**](TextMagicApi.md#GetChatByPhone) | **Get** /api/v2/chats/{phone}/by/phone | Find chats by phone
[**GetChatMessages**](TextMagicApi.md#GetChatMessages) | **Get** /api/v2/chats/{id}/message | Get chat messages
[**GetContact**](TextMagicApi.md#GetContact) | **Get** /api/v2/contacts/{id} | Get a single contact.
[**GetContactByPhone**](TextMagicApi.md#GetContactByPhone) | **Get** /api/v2/contacts/phone/{phone} | Get a single contact by phone number.
[**GetContactIfBlocked**](TextMagicApi.md#GetContactIfBlocked) | **Get** /api/v2/contacts/block/phone | Check is that phone number blocked
[**GetContactImportSessionProgress**](TextMagicApi.md#GetContactImportSessionProgress) | **Get** /api/v2/contacts/import/progress/{id} | Get contact import session progress.
[**GetContactNote**](TextMagicApi.md#GetContactNote) | **Get** /api/v2/notes/{id} | Get a single contact note.
[**GetContactNotes**](TextMagicApi.md#GetContactNotes) | **Get** /api/v2/contacts/{id}/notes | Fetch notes assigned to the given contact.
[**GetContacts**](TextMagicApi.md#GetContacts) | **Get** /api/v2/contacts | Get all user contacts.
[**GetContactsAutocomplete**](TextMagicApi.md#GetContactsAutocomplete) | **Get** /api/v2/contacts/autocomplete | Get contacts autocomplete suggestions by given search term.
[**GetContactsByListId**](TextMagicApi.md#GetContactsByListId) | **Get** /api/v2/lists/{id}/contacts | Fetch user contacts by given group id.
[**GetCountries**](TextMagicApi.md#GetCountries) | **Get** /api/v2/countries | Return list of countries.
[**GetCurrentUser**](TextMagicApi.md#GetCurrentUser) | **Get** /api/v2/user | Get current user info.
[**GetCustomField**](TextMagicApi.md#GetCustomField) | **Get** /api/v2/customfields/{id} | Get a single custom field.
[**GetCustomFields**](TextMagicApi.md#GetCustomFields) | **Get** /api/v2/customfields | Get all contact custom fields.
[**GetDedicatedNumber**](TextMagicApi.md#GetDedicatedNumber) | **Get** /api/v2/numbers/{id} | Get the details of a specific dedicated number
[**GetDisallowedRules**](TextMagicApi.md#GetDisallowedRules) | **Get** /api/v2/user/disallowed-rules | Get an array of all rules that are disallowed to the current account.
[**GetFavourites**](TextMagicApi.md#GetFavourites) | **Get** /api/v2/contacts/favorite | Get favorite contacts and lists.
[**GetInboundMessage**](TextMagicApi.md#GetInboundMessage) | **Get** /api/v2/replies/{id} | Get a single inbound message
[**GetInboundMessagesNotificationSettings**](TextMagicApi.md#GetInboundMessagesNotificationSettings) | **Get** /api/v2/user/notification/inbound | Get inbound messages notification settings
[**GetInvoices**](TextMagicApi.md#GetInvoices) | **Get** /api/v2/invoices | Return account invoices.
[**GetList**](TextMagicApi.md#GetList) | **Get** /api/v2/lists/{id} | Get a single list.
[**GetListContactsIds**](TextMagicApi.md#GetListContactsIds) | **Get** /api/v2/lists/{id}/contacts/ids | Fetch all contacts IDs belonging to the list with ID.
[**GetListsOfContact**](TextMagicApi.md#GetListsOfContact) | **Get** /api/v2/contacts/{id}/lists | Return lists which contact belongs to.
[**GetMessagePreview**](TextMagicApi.md#GetMessagePreview) | **Get** /api/v2/messages/preview | Preview message
[**GetMessagePrice**](TextMagicApi.md#GetMessagePrice) | **Get** /api/v2/messages/price | Check price
[**GetMessagePrices**](TextMagicApi.md#GetMessagePrices) | **Get** /api/v2/messages/prices | Get pricing
[**GetMessageSession**](TextMagicApi.md#GetMessageSession) | **Get** /api/v2/sessions/{id} | Get a session details
[**GetMessageSessionStat**](TextMagicApi.md#GetMessageSessionStat) | **Get** /api/v2/sessions/{id}/stat | Get a session statistics
[**GetMessagesBySessionId**](TextMagicApi.md#GetMessagesBySessionId) | **Get** /api/v2/sessions/{id}/messages | Get a session messages
[**GetMessagingCounters**](TextMagicApi.md#GetMessagingCounters) | **Get** /api/v2/stats/messaging/data | Return counters for messaging data views.
[**GetMessagingStat**](TextMagicApi.md#GetMessagingStat) | **Get** /api/v2/stats/messaging | Return messaging statistics.
[**GetOutboundMessage**](TextMagicApi.md#GetOutboundMessage) | **Get** /api/v2/messages/{id} | Get a single message
[**GetOutboundMessagesHistory**](TextMagicApi.md#GetOutboundMessagesHistory) | **Get** /api/v2/history | Get history
[**GetPushTokens**](TextMagicApi.md#GetPushTokens) | **Get** /api/v2/push/tokens | Get all device tokens assigned to the current account
[**GetScheduledMessage**](TextMagicApi.md#GetScheduledMessage) | **Get** /api/v2/schedules/{id} | Get a single scheduled message
[**GetSenderId**](TextMagicApi.md#GetSenderId) | **Get** /api/v2/senderids/{id} | Get the details of a specific Sender ID
[**GetSenderIds**](TextMagicApi.md#GetSenderIds) | **Get** /api/v2/senderids | Get all your approved Sender IDs
[**GetSenderSettings**](TextMagicApi.md#GetSenderSettings) | **Get** /api/v2/sender/settings | Get current sender settings
[**GetSpendingStat**](TextMagicApi.md#GetSpendingStat) | **Get** /api/v2/stats/spending | Return account spending statistics.
[**GetState**](TextMagicApi.md#GetState) | **Get** /api/v2/state | Get current entities state
[**GetSubaccount**](TextMagicApi.md#GetSubaccount) | **Get** /api/v2/subaccounts/{id} | Get a single subaccount.
[**GetSubaccounts**](TextMagicApi.md#GetSubaccounts) | **Get** /api/v2/subaccounts | Get all subaccounts of current user.
[**GetSubaccountsWithTokens**](TextMagicApi.md#GetSubaccountsWithTokens) | **Post** /api/v2/subaccounts/tokens/list | Get all subaccounts with their REST API tokens associated with specified app name.
[**GetSurvey**](TextMagicApi.md#GetSurvey) | **Get** /api/v2/surveys/{id} | Get a survey by id.
[**GetSurveyNode**](TextMagicApi.md#GetSurveyNode) | **Get** /api/v2/surveys/nodes/{id} | Get a node by id.
[**GetSurveyNodes**](TextMagicApi.md#GetSurveyNodes) | **Get** /api/v2/surveys/{id}/nodes | Fetch nodes by given survey id.
[**GetSurveys**](TextMagicApi.md#GetSurveys) | **Get** /api/v2/surveys | Get all user surveys.
[**GetTemplate**](TextMagicApi.md#GetTemplate) | **Get** /api/v2/templates/{id} | Get a template details
[**GetTimezones**](TextMagicApi.md#GetTimezones) | **Get** /api/v2/timezones | Return all available timezone IDs.
[**GetUnreadMessagesTotal**](TextMagicApi.md#GetUnreadMessagesTotal) | **Get** /api/v2/chats/unread/count | Get unread messages number
[**GetUnsubscribedContact**](TextMagicApi.md#GetUnsubscribedContact) | **Get** /api/v2/unsubscribers/{id} | Get a single unsubscribed contact.
[**GetUnsubscribers**](TextMagicApi.md#GetUnsubscribers) | **Get** /api/v2/unsubscribers | Get all contact have unsubscribed from your communication.
[**GetUserDedicatedNumbers**](TextMagicApi.md#GetUserDedicatedNumbers) | **Get** /api/v2/numbers | Get all your dedicated numbers
[**GetUserLists**](TextMagicApi.md#GetUserLists) | **Get** /api/v2/lists | Get all user lists.
[**GetVersions**](TextMagicApi.md#GetVersions) | **Get** /api/v2/versions | Get minimal valid apps versions
[**InviteSubaccount**](TextMagicApi.md#InviteSubaccount) | **Post** /api/v2/subaccounts | Invite new subaccount.
[**MarkChatsReadBulk**](TextMagicApi.md#MarkChatsReadBulk) | **Post** /api/v2/chats/read/bulk | Mark chats as read (bulk)
[**MarkChatsUnreadBulk**](TextMagicApi.md#MarkChatsUnreadBulk) | **Post** /api/v2/chats/unread/bulk | Mark chats as unread (bulk)
[**MergeSurveyNodes**](TextMagicApi.md#MergeSurveyNodes) | **Post** /api/v2/surveys/nodes/merge | Merge two question nodes.
[**MuteChat**](TextMagicApi.md#MuteChat) | **Post** /api/v2/chats/mute | Mute chat sounds
[**MuteChatsBulk**](TextMagicApi.md#MuteChatsBulk) | **Post** /api/v2/chats/mute/bulk | Mute chats (bulk)
[**Ping**](TextMagicApi.md#Ping) | **Get** /api/v2/ping | Just does a pong.
[**ReopenChatsBulk**](TextMagicApi.md#ReopenChatsBulk) | **Post** /api/v2/chats/reopen/bulk | Reopen chats (bulk)
[**RequestNewSubaccountToken**](TextMagicApi.md#RequestNewSubaccountToken) | **Post** /api/v2/subaccounts/tokens | Request a new REST API token for subaccount.
[**RequestSenderId**](TextMagicApi.md#RequestSenderId) | **Post** /api/v2/senderids | Apply for a new Sender ID
[**ResetSurvey**](TextMagicApi.md#ResetSurvey) | **Put** /api/v2/surveys/{id}/reset | Reset a survey flow.
[**SearchChats**](TextMagicApi.md#SearchChats) | **Get** /api/v2/chats/search | Find chats by message text
[**SearchChatsByIds**](TextMagicApi.md#SearchChatsByIds) | **Get** /api/v2/chats/search/ids | Find chats (bulk)
[**SearchChatsByReceipent**](TextMagicApi.md#SearchChatsByReceipent) | **Get** /api/v2/chats/search/recipients | Find chats by recipient
[**SearchContacts**](TextMagicApi.md#SearchContacts) | **Get** /api/v2/contacts/search | Find user contacts by given parameters.
[**SearchInboundMessages**](TextMagicApi.md#SearchInboundMessages) | **Get** /api/v2/replies/search | Find inbound messages
[**SearchLists**](TextMagicApi.md#SearchLists) | **Get** /api/v2/lists/search | Find contact lists by given parameters.
[**SearchOutboundMessages**](TextMagicApi.md#SearchOutboundMessages) | **Get** /api/v2/messages/search | Find messages
[**SearchScheduledMessages**](TextMagicApi.md#SearchScheduledMessages) | **Get** /api/v2/schedules/search | Find scheduled messages
[**SearchTemplates**](TextMagicApi.md#SearchTemplates) | **Get** /api/v2/templates/search | Find templates by criteria
[**SendEmailVerificationCode**](TextMagicApi.md#SendEmailVerificationCode) | **Get** /api/v2/user/email/verification | Send user email verification
[**SendMessage**](TextMagicApi.md#SendMessage) | **Post** /api/v2/messages | Send message
[**SendPhoneVerificationCode**](TextMagicApi.md#SendPhoneVerificationCode) | **Get** /api/v2/user/phone/verification | Send user phone verification
[**SendPhoneVerificationCode_0**](TextMagicApi.md#SendPhoneVerificationCode_0) | **Post** /api/v2/verify | Step 1: Send a verification code 
[**SetChatStatus**](TextMagicApi.md#SetChatStatus) | **Post** /api/v2/chats/status | Change chat status
[**StartSurvey**](TextMagicApi.md#StartSurvey) | **Put** /api/v2/surveys/{id}/start | Start a survey.
[**UnblockContact**](TextMagicApi.md#UnblockContact) | **Post** /api/v2/contacts/unblock | Unblock contact by phone number.
[**UnblockContactsBulk**](TextMagicApi.md#UnblockContactsBulk) | **Post** /api/v2/contacts/unblock/bulk | Unblock several contacts by blocked contact ids or unblock all contacts
[**UnmuteChatsBulk**](TextMagicApi.md#UnmuteChatsBulk) | **Post** /api/v2/chats/unmute/bulk | Unmute chats (bulk)
[**UnsubscribeContact**](TextMagicApi.md#UnsubscribeContact) | **Post** /api/v2/unsubscribers | Unsubscribe contact from your communication by phone number.
[**UpdateBalanceNotificationSettings**](TextMagicApi.md#UpdateBalanceNotificationSettings) | **Put** /api/v2/user/notification/balance | Update balance notification settings
[**UpdateCallbackSettings**](TextMagicApi.md#UpdateCallbackSettings) | **Put** /api/v2/callback/settings | Update callback URL settings
[**UpdateChatDesktopNotificationSettings**](TextMagicApi.md#UpdateChatDesktopNotificationSettings) | **Put** /api/v2/user/desktop/notification | Update chat desktop notification settings
[**UpdateContact**](TextMagicApi.md#UpdateContact) | **Put** /api/v2/contacts/{id} | Update existing contact.
[**UpdateContactNote**](TextMagicApi.md#UpdateContactNote) | **Put** /api/v2/notes/{id} | Update existing contact note.
[**UpdateCurrentUser**](TextMagicApi.md#UpdateCurrentUser) | **Put** /api/v2/user | Update current user info.
[**UpdateCustomField**](TextMagicApi.md#UpdateCustomField) | **Put** /api/v2/customfields/{id} | Update existing custom field.
[**UpdateCustomFieldValue**](TextMagicApi.md#UpdateCustomFieldValue) | **Put** /api/v2/customfields/{id}/update | Update contact&#39;s custom field value.
[**UpdateInboundMessagesNotificationSettings**](TextMagicApi.md#UpdateInboundMessagesNotificationSettings) | **Put** /api/v2/user/notification/inbound | Update inbound messages notification settings
[**UpdateList**](TextMagicApi.md#UpdateList) | **Put** /api/v2/lists/{id} | Update existing list.
[**UpdatePassword**](TextMagicApi.md#UpdatePassword) | **Put** /api/v2/user/password/change | Change user password.
[**UpdateSenderSetting**](TextMagicApi.md#UpdateSenderSetting) | **Put** /api/v2/sender/settings | Change sender settings
[**UpdateSurvey**](TextMagicApi.md#UpdateSurvey) | **Put** /api/v2/surveys/{id} | Update existing survey.
[**UpdateSurveyNode**](TextMagicApi.md#UpdateSurveyNode) | **Put** /api/v2/surveys/nodes/{id} | Update existing node.
[**UpdateTemplate**](TextMagicApi.md#UpdateTemplate) | **Put** /api/v2/templates/{id} | Update a template


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

# **BlockContact**
> ResourceLinkResponse BlockContact(ctx, blockContactInputObject)
Block contact from inbound and outbound communication by phone number.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockContactInputObject** | [**BlockContactInputObject**](BlockContactInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BuyDedicatedNumber**
> BuyDedicatedNumber(ctx, buyDedicatedNumberInputObject)
Buy a dedicated number

To buy a dedicated number, you first need to find an available number matching your criteria using the `/api/v2/numbers/available` command described above.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buyDedicatedNumberInputObject** | [**BuyDedicatedNumberInputObject**](BuyDedicatedNumberInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelSurvey**
> ResourceLinkResponse CancelSurvey(ctx, id)
Cancel a survey.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelVerification**
> CancelVerification(ctx, verifyId)
Cancel verification process

You can cancel the verification not earlier than 30 seconds after the initial request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **verifyId** | **string**| the verifyId that you received in Step 1. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckPhoneVerificationCode**
> CheckPhoneVerificationCode(ctx, checkPhoneVerificationCodeInputObject)
Check user phone verification code



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **checkPhoneVerificationCodeInputObject** | [**CheckPhoneVerificationCodeInputObject**](CheckPhoneVerificationCodeInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckPhoneVerificationCode_0**
> CheckPhoneVerificationCode_0(ctx, checkPhoneVerificationCodeInputObject)
Step 2: Check the verification code 

Check received code from user with the code which was actually sent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **checkPhoneVerificationCodeInputObject** | [**CheckPhoneVerificationCodeInputObject1**](CheckPhoneVerificationCodeInputObject1.md)|  | 

### Return type

 (empty response body)

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

# **CloseChatsBulk**
> CloseChatsBulk(ctx, closeChatsBulkInputObject)
Close chats (bulk)

Close chats by chat ids or close all chats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **closeChatsBulkInputObject** | [**CloseChatsBulkInputObject**](CloseChatsBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloseReadChats**
> CloseReadChats(ctx, )
Close read chats

Close all chats that have no unread messages.

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

# **CloseSubaccount**
> CloseSubaccount(ctx, id)
Close subaccount.



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

# **CreateContact**
> ResourceLinkResponse CreateContact(ctx, createContactInputObject)
Create a new contact from the submitted data.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createContactInputObject** | [**CreateContactInputObject**](CreateContactInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateContactNote**
> ResourceLinkResponse CreateContactNote(ctx, createContactNoteInputObject, id)
Create a new contact note.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createContactNoteInputObject** | [**CreateContactNoteInputObject**](CreateContactNoteInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCustomField**
> ResourceLinkResponse CreateCustomField(ctx, createCustomFieldInputObject)
Create a new custom field from the submitted data.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createCustomFieldInputObject** | [**CreateCustomFieldInputObject**](CreateCustomFieldInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateList**
> ResourceLinkResponse CreateList(ctx, createListInputObject)
Create a new list from the submitted data.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createListInputObject** | [**CreateListInputObject**](CreateListInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePushToken**
> CreatePushToken(ctx, createPushTokenInputObject)
Add or update a device token.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createPushTokenInputObject** | [**CreatePushTokenInputObject**](CreatePushTokenInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSurvey**
> ResourceLinkResponse CreateSurvey(ctx, createSurveyInputObject)
Create a new survey from the submitted data.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSurveyInputObject** | [**CreateSurveyInputObject**](CreateSurveyInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSurveyNode**
> ResourceLinkResponse CreateSurveyNode(ctx, createSurveyNodeInputObject, id)
Create a new node from the submitted data.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSurveyNodeInputObject** | [**CreateSurveyNodeInputObject**](CreateSurveyNodeInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTemplate**
> ResourceLinkResponse CreateTemplate(ctx, createTemplateInputObject)
Create a template

There are times when creating a new template makes sense (such as when targeting specific clients or improving your business strategies).  You can create new SMS templates for marketing purposes or SMS templates for business campaigns. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createTemplateInputObject** | [**CreateTemplateInputObject**](CreateTemplateInputObject.md)|  | 

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

# **DeleteAvatar**
> DeleteAvatar(ctx, )
Delete an avatar for the current user.\\



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

# **DeleteChatMessages**
> DeleteChatMessages(ctx, deleteChatMessagesBulkInputObject, id)
Delete chat messages by ID(s)

Delete messages from chat by given messages ID(s).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteChatMessagesBulkInputObject** | [**DeleteChatMessagesBulkInputObject**](DeleteChatMessagesBulkInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteChatsBulk**
> DeleteChatsBulk(ctx, deleteChatsBulkInputObject)
Delete chats (bulk)

Delete chats by given ID(s) or delete all chats.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteChatsBulkInputObject** | [**DeleteChatsBulkInputObject**](DeleteChatsBulkInputObject.md)|  | 

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
> DeleteContactNotesBulk(ctx, id, deleteContactNotesBulkInputObject)
Delete contact note by given ID(s) or delete all contact notes.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **deleteContactNotesBulkInputObject** | [**DeleteContactNotesBulkInputObject**](DeleteContactNotesBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContactsByIds**
> DeleteContactsByIds(ctx, deleteContactsByIdsInputObject)
Delete contact by given ID(s) or delete all contacts.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteContactsByIdsInputObject** | [**DeleteContactsByIdsInputObject**](DeleteContactsByIdsInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

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

# **DeleteCustomField**
> DeleteCustomField(ctx, id)
Delete a single custom field.



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

# **DeleteDedicatedNumber**
> DeleteDedicatedNumber(ctx, id)
Cancel dedicated number subscription



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

# **DeleteInboundMessage**
> DeleteInboundMessage(ctx, id)
Delete a single inbound message

> Note, deleted inbound message will disappear from TextMagic Online, chats, and any other place they are referenced.  So, be careful! 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The unique numeric ID for the inbound message. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInboundMessagesBulk**
> DeleteInboundMessagesBulk(ctx, deleteInboundMessagesBulkInputObject)
Delete inbound messages (bulk)

> Note, deleted inbound message will disappear from TextMagic Online, chats, and any other place they are referenced.  So, be careful! 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteInboundMessagesBulkInputObject** | [**DeleteInboundMessagesBulkInputObject**](DeleteInboundMessagesBulkInputObject.md)|  | 

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
> DeleteListContactsBulk(ctx, deleteListContactsBulkInputObject, id)
Delete contact from list by given ID(s) or all contacts from list.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteListContactsBulkInputObject** | [**DeleteListContactsBulkInputObject**](DeleteListContactsBulkInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListsBulk**
> DeleteListsBulk(ctx, deleteListsBulkInputObject)
Delete list by given ID(s) or delete all lists.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteListsBulkInputObject** | [**DeleteListsBulkInputObject**](DeleteListsBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMessageSession**
> DeleteMessageSession(ctx, id)
Delete a session



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

# **DeleteMessageSessionsBulk**
> DeleteMessageSessionsBulk(ctx, deleteMessageSessionsBulkInputObject)
Delete sessions (bulk)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteMessageSessionsBulkInputObject** | [**DeleteMessageSessionsBulkInputObject**](DeleteMessageSessionsBulkInputObject.md)|  | 

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
> DeleteOutboundMessagesBulk(ctx, deleteOutboundMessagesBulkInputObject)
Delete messages (bulk)

Delete outbound messages by given ID(s) or delete all outbound messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteOutboundMessagesBulkInputObject** | [**DeleteOutboundMessagesBulkInputObject**](DeleteOutboundMessagesBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePushToken**
> DeletePushToken(ctx, type_, deviceId)
Delete a push notification device token.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 
  **deviceId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteScheduledMessage**
> DeleteScheduledMessage(ctx, id)
Delete a single scheduled message



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

# **DeleteScheduledMessagesBulk**
> DeleteScheduledMessagesBulk(ctx, deleteScheduledMessagesBulkInputObject)
Delete scheduled messages (bulk)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteScheduledMessagesBulkInputObject** | [**DeleteScheduledMessagesBulkInputObject**](DeleteScheduledMessagesBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSenderId**
> DeleteSenderId(ctx, id)
Delete a Sender ID



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

# **DeleteSurvey**
> DeleteSurvey(ctx, id)
Delete a survey.



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

# **DeleteSurveyNode**
> DeleteSurveyNode(ctx, id)
Delete a node.



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

# **DeleteTemplate**
> DeleteTemplate(ctx, id)
Delete a template



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

# **DeleteTemplatesBulk**
> DeleteTemplatesBulk(ctx, deleteTemplatesBulkInputObject)
Delete templates (bulk)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteTemplatesBulkInputObject** | [**DeleteTemplatesBulkInputObject**](DeleteTemplatesBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoAuth**
> DoAuthResponse DoAuth(ctx, doAuthInputObject)
Authenticate user by given username and password.

Returning a username and token that you should pass to the all requests (in X-TM-Username and X-TM-Key, respectively)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **doAuthInputObject** | [**DoAuthInputObject**](DoAuthInputObject.md)|  | 

### Return type

[**DoAuthResponse**](DoAuthResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **DuplicateSurvey**
> ResourceLinkResponse DuplicateSurvey(ctx, id)
Duplicate a survey.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllBulkSessions**
> GetAllBulkSessionsPaginatedResponse GetAllBulkSessions(ctx, optional)
Get all bulk sending sessions.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllBulkSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllBulkSessionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetAllBulkSessionsPaginatedResponse**](GetAllBulkSessionsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllChats**
> GetAllChatsPaginatedResponse GetAllChats(ctx, optional)
Get all chats



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllChatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllChatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.String**| Fetch only (a)ctive, (c)losed or (d)eleted chats | 
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **voice** | **optional.Int32**| Fetch results with voice calls | [default to 0]
 **flat** | **optional.Int32**| Should additional contact info be included | [default to 0]

### Return type

[**GetAllChatsPaginatedResponse**](GetAllChatsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllInboundMessages**
> GetAllInboundMessagesPaginatedResponse GetAllInboundMessages(ctx, optional)
Get all inbound messages



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllInboundMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllInboundMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**GetAllInboundMessagesPaginatedResponse**](GetAllInboundMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllMessageSessions**
> GetAllMessageSessionsPaginatedResponse GetAllMessageSessions(ctx, optional)
Get all sessions



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllMessageSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllMessageSessionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetAllMessageSessionsPaginatedResponse**](GetAllMessageSessionsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllOutboundMessages**
> GetAllOutboundMessagesPaginatedResponse GetAllOutboundMessages(ctx, optional)
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
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **lastId** | **optional.Int32**| Filter results by ID, selecting all values lesser than the specified ID. Note that \\&#39;page\\&#39; parameter is ignored when \\&#39;lastId\\&#39; is specified | 

### Return type

[**GetAllOutboundMessagesPaginatedResponse**](GetAllOutboundMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllScheduledMessages**
> GetAllScheduledMessagesPaginatedResponse GetAllScheduledMessages(ctx, optional)
Get all scheduled messages



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllScheduledMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllScheduledMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **status** | **optional.String**| Fetch schedules with the specific status: a - actual, c - completed, x - all | [default to x]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**GetAllScheduledMessagesPaginatedResponse**](GetAllScheduledMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTemplates**
> GetAllTemplatesPaginatedResponse GetAllTemplates(ctx, optional)
Get all templates



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | 
 **limit** | **optional.Int32**| The number of results per page. | 

### Return type

[**GetAllTemplatesPaginatedResponse**](GetAllTemplatesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvailableDedicatedNumbers**
> GetAvailableDedicatedNumbersResponse GetAvailableDedicatedNumbers(ctx, country, optional)
Find dedicated numbers available for purchase



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Two-letter dedicated number country ISO code. | 
 **optional** | ***GetAvailableDedicatedNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAvailableDedicatedNumbersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefix** | **optional.Int32**| Desired number prefix. Should include country code (i.e. 447 for UK phone number format). Leave blank to get all the available numbers for the specified country. | [default to 1]
 **tollfree** | **optional.Int32**| Should we show only tollfree numbers (tollfree available only for US). | [default to 0]

### Return type

[**GetAvailableDedicatedNumbersResponse**](GetAvailableDedicatedNumbersResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvailableSenderSettingOptions**
> GetAvailableSenderSettingOptionsResponse GetAvailableSenderSettingOptions(ctx, optional)
Get available sender settings

Get all available sender setting options which could be used in \"from\" parameter of POST messages method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAvailableSenderSettingOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAvailableSenderSettingOptionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **optional.String**| Two-letter ISO country ID. If not specified, it returns all the available sender settings. | 

### Return type

[**GetAvailableSenderSettingOptionsResponse**](GetAvailableSenderSettingOptionsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBalanceNotificationOptions**
> GetBalanceNotificationOptionsResponse GetBalanceNotificationOptions(ctx, )
Returns the list of available balance options which can be used as a bound to determine when to send email to user with low balance notification. See https://my.textmagic.com/online/account/notifications/balance



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetBalanceNotificationOptionsResponse**](GetBalanceNotificationOptionsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBalanceNotificationSettings**
> GetBalanceNotificationSettingsResponse GetBalanceNotificationSettings(ctx, )
Get balance notification settings



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetBalanceNotificationSettingsResponse**](GetBalanceNotificationSettingsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockedContacts**
> GetBlockedContactsPaginatedResponse GetBlockedContacts(ctx, optional)
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
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **query** | **optional.String**| Find blocked contacts by specified search query | 
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**GetBlockedContactsPaginatedResponse**](GetBlockedContactsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBulkSession**
> BulkSession GetBulkSession(ctx, id)
Get bulk message session status.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**BulkSession**](BulkSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCallbackSettings**
> GetCallbackSettingsResponse GetCallbackSettings(ctx, )
Fetch callback URL settings



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetCallbackSettingsResponse**](GetCallbackSettingsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCallsPrices**
> GetCallsPricesResponse GetCallsPrices(ctx, )
Check pricing for a inbound/outbound call.



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetCallsPricesResponse**](GetCallsPricesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChat**
> Chat GetChat(ctx, id)
Get a single chat



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**Chat**](Chat.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChatByPhone**
> Chat GetChatByPhone(ctx, phone, optional)
Find chats by phone



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **phone** | **string**|  | 
 **optional** | ***GetChatByPhoneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetChatByPhoneOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsert** | **optional.Int32**| Create a new chat if not found | [default to 0]
 **reopen** | **optional.Int32**| Reopen chat if found or do not change status | [default to 0]

### Return type

[**Chat**](Chat.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChatMessages**
> GetChatMessagesPaginatedResponse GetChatMessages(ctx, id, optional)
Get chat messages



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetChatMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetChatMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **query** | **optional.String**| Find messages by specified search query | 
 **start** | **optional.Int32**| Return messages since specified timestamp only | 
 **end** | **optional.Int32**| Return messages up to specified timestamp only | 
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]
 **voice** | **optional.Int32**| Fetch results with voice calls | [default to 0]

### Return type

[**GetChatMessagesPaginatedResponse**](GetChatMessagesPaginatedResponse.md)

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
> GetContactNotesPaginatedResponse GetContactNotes(ctx, id, optional)
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

 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetContactNotesPaginatedResponse**](GetContactNotesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContacts**
> GetContactsPaginatedResponse GetContacts(ctx, optional)
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
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **shared** | **optional.Int32**| Should shared contacts to be included | [default to 0]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**GetContactsPaginatedResponse**](GetContactsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactsAutocomplete**
> GetContactsAutocompleteResponse GetContactsAutocomplete(ctx, query, optional)
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

 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **lists** | **optional.Int32**| Should lists be returned or not | [default to 0]

### Return type

[**GetContactsAutocompleteResponse**](GetContactsAutocompleteResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactsByListId**
> GetContactsByListIdPaginatedResponse GetContactsByListId(ctx, id, optional)
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

 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**GetContactsByListIdPaginatedResponse**](GetContactsByListIdPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCountries**
> GetCountriesResponse GetCountries(ctx, )
Return list of countries.



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetCountriesResponse**](GetCountriesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentUser**
> User GetCurrentUser(ctx, )
Get current user info.



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomField**
> UserCustomField GetCustomField(ctx, id)
Get a single custom field.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**UserCustomField**](UserCustomField.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomFields**
> GetCustomFieldsPaginatedResponse GetCustomFields(ctx, optional)
Get all contact custom fields.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCustomFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCustomFieldsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetCustomFieldsPaginatedResponse**](GetCustomFieldsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDedicatedNumber**
> UsersInbound GetDedicatedNumber(ctx, id)
Get the details of a specific dedicated number



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**UsersInbound**](UsersInbound.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDisallowedRules**
> GetDisallowedRulesResponse GetDisallowedRules(ctx, )
Get an array of all rules that are disallowed to the current account.



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetDisallowedRulesResponse**](GetDisallowedRulesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFavourites**
> GetFavouritesPaginatedResponse GetFavourites(ctx, optional)
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
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **query** | **optional.String**| Find contacts or lists by specified search query | [default to A]

### Return type

[**GetFavouritesPaginatedResponse**](GetFavouritesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInboundMessage**
> MessageIn GetInboundMessage(ctx, id)
Get a single inbound message



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The unique numeric ID for the inbound message. | 

### Return type

[**MessageIn**](MessageIn.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInboundMessagesNotificationSettings**
> GetInboundMessagesNotificationSettingsResponse GetInboundMessagesNotificationSettings(ctx, )
Get inbound messages notification settings



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetInboundMessagesNotificationSettingsResponse**](GetInboundMessagesNotificationSettingsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoices**
> GetInvoicesPaginatedResponse GetInvoices(ctx, optional)
Return account invoices.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetInvoicesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetInvoicesPaginatedResponse**](GetInvoicesPaginatedResponse.md)

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
> GetListsOfContactPaginatedResponse GetListsOfContact(ctx, id, optional)
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

 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetListsOfContactPaginatedResponse**](GetListsOfContactPaginatedResponse.md)

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

# **GetMessageSession**
> MessageSession GetMessageSession(ctx, id)
Get a session details



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| a session ID | 

### Return type

[**MessageSession**](MessageSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessageSessionStat**
> GetMessageSessionStatResponse GetMessageSessionStat(ctx, id, optional)
Get a session statistics



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetMessageSessionStatOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMessageSessionStatOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDeleted** | **optional.Int32**| Search also in deleted messages | [default to 0]

### Return type

[**GetMessageSessionStatResponse**](GetMessageSessionStatResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagesBySessionId**
> GetMessagesBySessionIdPaginatedResponse GetMessagesBySessionId(ctx, id, optional)
Get a session messages

A useful synonym for \"messages/search\" command with provided \"sessionId\" parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetMessagesBySessionIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMessagesBySessionIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **statuses** | **optional.String**| Find messages by status | 
 **includeDeleted** | **optional.Int32**| Search also in deleted messages | [default to 0]

### Return type

[**GetMessagesBySessionIdPaginatedResponse**](GetMessagesBySessionIdPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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
> GetOutboundMessagesHistoryPaginatedResponse GetOutboundMessagesHistory(ctx, optional)
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
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **lastId** | **optional.Int32**| Filter results by ID, selecting all values lesser than the specified ID. | 
 **query** | **optional.String**| Find message by specified search query | 
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**GetOutboundMessagesHistoryPaginatedResponse**](GetOutboundMessagesHistoryPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPushTokens**
> GetPushTokensResponse GetPushTokens(ctx, )
Get all device tokens assigned to the current account



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetPushTokensResponse**](GetPushTokensResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScheduledMessage**
> MessagesIcs GetScheduledMessage(ctx, id)
Get a single scheduled message



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**MessagesIcs**](MessagesIcs.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSenderId**
> SenderId GetSenderId(ctx, id)
Get the details of a specific Sender ID



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**SenderId**](SenderId.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSenderIds**
> GetSenderIdsPaginatedResponse GetSenderIds(ctx, optional)
Get all your approved Sender IDs



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSenderIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSenderIdsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetSenderIdsPaginatedResponse**](GetSenderIdsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSenderSettings**
> GetSenderSettingsResponse GetSenderSettings(ctx, optional)
Get current sender settings

@TODO

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSenderSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSenderSettingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **optional.String**| Return sender settings enabled for sending to specified country. Two upper case characters | 

### Return type

[**GetSenderSettingsResponse**](GetSenderSettingsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpendingStat**
> GetSpendingStatPaginatedResponse GetSpendingStat(ctx, optional)
Return account spending statistics.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSpendingStatOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSpendingStatOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **start** | **optional.Int32**| Optional. Start date in unix timestamp format. Default is 7 days ago | 
 **end** | **optional.Int32**| Optional. End date in unix timestamp format. Default is now | 

### Return type

[**GetSpendingStatPaginatedResponse**](GetSpendingStatPaginatedResponse.md)

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

# **GetSubaccount**
> User GetSubaccount(ctx, id)
Get a single subaccount.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**User**](User.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubaccounts**
> User GetSubaccounts(ctx, optional)
Get all subaccounts of current user.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSubaccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSubaccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**User**](User.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubaccountsWithTokens**
> GetSubaccountsWithTokensResponse GetSubaccountsWithTokens(ctx, getSubaccountsWithTokensInputObject, optional)
Get all subaccounts with their REST API tokens associated with specified app name.

When more than one token related to app name, last key will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **getSubaccountsWithTokensInputObject** | [**GetSubaccountsWithTokensInputObject**](GetSubaccountsWithTokensInputObject.md)|  | 
 **optional** | ***GetSubaccountsWithTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSubaccountsWithTokensOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetSubaccountsWithTokensResponse**](GetSubaccountsWithTokensResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSurvey**
> Survey GetSurvey(ctx, id)
Get a survey by id.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**Survey**](Survey.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSurveyNode**
> SurveyNode GetSurveyNode(ctx, id)
Get a node by id.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**SurveyNode**](SurveyNode.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSurveyNodes**
> GetSurveyNodesResponse GetSurveyNodes(ctx, id)
Fetch nodes by given survey id.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**GetSurveyNodesResponse**](GetSurveyNodesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSurveys**
> GetSurveysPaginatedResponse GetSurveys(ctx, optional)
Get all user surveys.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSurveysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSurveysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetSurveysPaginatedResponse**](GetSurveysPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplate**
> MessageTemplate GetTemplate(ctx, id)
Get a template details



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**MessageTemplate**](MessageTemplate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimezones**
> GetTimezonesResponse GetTimezones(ctx, optional)
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

[**GetTimezonesResponse**](GetTimezonesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUnreadMessagesTotal**
> GetUnreadMessagesTotalResponse GetUnreadMessagesTotal(ctx, )
Get unread messages number

Get total amount of unread messages in the current user chats.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetUnreadMessagesTotalResponse**](GetUnreadMessagesTotalResponse.md)

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
> GetUnsubscribersPaginatedResponse GetUnsubscribers(ctx, optional)
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
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetUnsubscribersPaginatedResponse**](GetUnsubscribersPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserDedicatedNumbers**
> GetUserDedicatedNumbersPaginatedResponse GetUserDedicatedNumbers(ctx, optional)
Get all your dedicated numbers



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUserDedicatedNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUserDedicatedNumbersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **surveyId** | **optional.Int32**| Fetch only that numbers which are ready for the survey | 

### Return type

[**GetUserDedicatedNumbersPaginatedResponse**](GetUserDedicatedNumbersPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserLists**
> GetUserListsPaginatedResponse GetUserLists(ctx, optional)
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
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]
 **favoriteOnly** | **optional.Int32**| Return only favorite lists | [default to 0]
 **onlyMine** | **optional.Int32**| Return only current user lists | [default to 0]

### Return type

[**GetUserListsPaginatedResponse**](GetUserListsPaginatedResponse.md)

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

# **InviteSubaccount**
> InviteSubaccount(ctx, inviteSubaccountInputObject)
Invite new subaccount.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inviteSubaccountInputObject** | [**InviteSubaccountInputObject**](InviteSubaccountInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarkChatsReadBulk**
> MarkChatsReadBulk(ctx, markChatsReadBulkInputObject)
Mark chats as read (bulk)

Mark several chats as read by chat ids or mark all chats as read

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **markChatsReadBulkInputObject** | [**MarkChatsReadBulkInputObject**](MarkChatsReadBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarkChatsUnreadBulk**
> MarkChatsUnreadBulk(ctx, markChatsUnreadBulkInputObject)
Mark chats as unread (bulk)

Mark several chats as UNread by chat ids or mark all chats as UNread

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **markChatsUnreadBulkInputObject** | [**MarkChatsUnreadBulkInputObject**](MarkChatsUnreadBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MergeSurveyNodes**
> MergeSurveyNodes(ctx, mergeSurveyNodesInputObject)
Merge two question nodes.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mergeSurveyNodesInputObject** | [**MergeSurveyNodesInputObject**](MergeSurveyNodesInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MuteChat**
> ResourceLinkResponse MuteChat(ctx, muteChatInputObject)
Mute chat sounds



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **muteChatInputObject** | [**MuteChatInputObject**](MuteChatInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MuteChatsBulk**
> MuteChatsBulk(ctx, muteChatsBulkInputObject)
Mute chats (bulk)

Mute several chats by chat ids or mute all chats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **muteChatsBulkInputObject** | [**MuteChatsBulkInputObject**](MuteChatsBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

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

# **ReopenChatsBulk**
> ReopenChatsBulk(ctx, reopenChatsBulkInputObject)
Reopen chats (bulk)

Reopen chats by chat ids or reopen all chats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reopenChatsBulkInputObject** | [**ReopenChatsBulkInputObject**](ReopenChatsBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestNewSubaccountToken**
> User RequestNewSubaccountToken(ctx, requestNewSubaccountTokenInputObject)
Request a new REST API token for subaccount.

Returning user object, key and app name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestNewSubaccountTokenInputObject** | [**RequestNewSubaccountTokenInputObject**](RequestNewSubaccountTokenInputObject.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestSenderId**
> ResourceLinkResponse RequestSenderId(ctx, requestSenderIdInputObject)
Apply for a new Sender ID

> Sender IDs are shared between all of your sub-accounts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestSenderIdInputObject** | [**RequestSenderIdInputObject**](RequestSenderIdInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetSurvey**
> ResourceLinkResponse ResetSurvey(ctx, id)
Reset a survey flow.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchChats**
> SearchChatsPaginatedResponse SearchChats(ctx, optional)
Find chats by message text



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchChatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchChatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **query** | **optional.String**| Find chats by specified search query | 

### Return type

[**SearchChatsPaginatedResponse**](SearchChatsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchChatsByIds**
> SearchChatsByIdsPaginatedResponse SearchChatsByIds(ctx, optional)
Find chats (bulk)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchChatsByIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchChatsByIdsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **ids** | **optional.String**| Find chats by ID(s) | 

### Return type

[**SearchChatsByIdsPaginatedResponse**](SearchChatsByIdsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchChatsByReceipent**
> SearchChatsByReceipentPaginatedResponse SearchChatsByReceipent(ctx, optional)
Find chats by recipient

Find chats by recipient (contact, list name or phone number).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchChatsByReceipentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchChatsByReceipentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **query** | **optional.String**| Find chats by specified search query | 
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]

### Return type

[**SearchChatsByReceipentPaginatedResponse**](SearchChatsByReceipentPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchContacts**
> SearchContactsPaginatedResponse SearchContacts(ctx, optional)
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
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
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

[**SearchContactsPaginatedResponse**](SearchContactsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchInboundMessages**
> SearchInboundMessagesPaginatedResponse SearchInboundMessages(ctx, optional)
Find inbound messages

Find inbound messages by given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchInboundMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchInboundMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **ids** | **optional.String**| Find message by ID(s) | 
 **query** | **optional.String**| Find recipients by specified search query | 
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]
 **expand** | **optional.Int32**| Expand by adding firstName, lastName and contactId | [default to 0]

### Return type

[**SearchInboundMessagesPaginatedResponse**](SearchInboundMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchLists**
> SearchListsPaginatedResponse SearchLists(ctx, optional)
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
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **ids** | **optional.String**| Find lists by ID(s) | 
 **query** | **optional.String**| Find lists by specified search query | 
 **onlyMine** | **optional.Int32**| Return only current user lists | [default to 0]
 **onlyDefault** | **optional.Int32**| Return only default lists | [default to 0]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**SearchListsPaginatedResponse**](SearchListsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOutboundMessages**
> SearchOutboundMessagesPaginatedResponse SearchOutboundMessages(ctx, optional)
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
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **lastId** | **optional.Int32**| Filter results by ID, selecting all values lesser than the specified ID. Note that \\&#39;page\\&#39; parameter is ignored when \\&#39;lastId\\&#39; is specified | 
 **ids** | **optional.String**| Find message by ID(s) | 
 **sessionId** | **optional.Int32**| Find messages by session ID | 
 **statuses** | **optional.String**| Find messages by status | 
 **includeDeleted** | **optional.Int32**| Search also in deleted messages | [default to 0]
 **query** | **optional.String**| Find messages by specified search query | 

### Return type

[**SearchOutboundMessagesPaginatedResponse**](SearchOutboundMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchScheduledMessages**
> SearchScheduledMessagesPaginatedResponse SearchScheduledMessages(ctx, optional)
Find scheduled messages



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchScheduledMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchScheduledMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **query** | **optional.String**| Find messages by specified search query | 
 **ids** | **optional.String**| Find schedules by ID(s) | 
 **status** | **optional.String**| Fetch schedules with the specific status: a - actual, c - completed, x - all | [default to x]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**SearchScheduledMessagesPaginatedResponse**](SearchScheduledMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTemplates**
> SearchTemplatesPaginatedResponse SearchTemplates(ctx, optional)
Find templates by criteria



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **ids** | **optional.String**| Find template by ID(s) | 
 **name** | **optional.String**| Find template by name | 
 **content** | **optional.String**| Find template by content | 

### Return type

[**SearchTemplatesPaginatedResponse**](SearchTemplatesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendEmailVerificationCode**
> SendEmailVerificationCode(ctx, )
Send user email verification



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendMessage**
> SendMessageResponse SendMessage(ctx, sendMessageInputObject)
Send message

The main entrypoint to send messages. See examples above for the reference.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sendMessageInputObject** | [**SendMessageInputObject**](SendMessageInputObject.md)|  | 

### Return type

[**SendMessageResponse**](SendMessageResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendPhoneVerificationCode**
> SendPhoneVerificationCode(ctx, )
Send user phone verification



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendPhoneVerificationCode_0**
> SendPhoneVerificationCodeResponse SendPhoneVerificationCode_0(ctx, sendPhoneVerificationCodeInputObject)
Step 1: Send a verification code 

Sends verification code to specified phone number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sendPhoneVerificationCodeInputObject** | [**SendPhoneVerificationCodeInputObject**](SendPhoneVerificationCodeInputObject.md)|  | 

### Return type

[**SendPhoneVerificationCodeResponse**](SendPhoneVerificationCodeResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetChatStatus**
> ResourceLinkResponse SetChatStatus(ctx, setChatStatusInputObject)
Change chat status

Set status of the chat given by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **setChatStatusInputObject** | [**SetChatStatusInputObject**](SetChatStatusInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartSurvey**
> ResourceLinkResponse StartSurvey(ctx, id)
Start a survey.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnblockContact**
> UnblockContact(ctx, unblockContactInputObject)
Unblock contact by phone number.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unblockContactInputObject** | [**UnblockContactInputObject**](UnblockContactInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnblockContactsBulk**
> UnblockContactsBulk(ctx, unblockContactsBulkInputObject)
Unblock several contacts by blocked contact ids or unblock all contacts



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unblockContactsBulkInputObject** | [**UnblockContactsBulkInputObject**](UnblockContactsBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnmuteChatsBulk**
> UnmuteChatsBulk(ctx, unmuteChatsBulkInputObject)
Unmute chats (bulk)

Unmute several chats by chat ids or unmute all chats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unmuteChatsBulkInputObject** | [**UnmuteChatsBulkInputObject**](UnmuteChatsBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsubscribeContact**
> ResourceLinkResponse UnsubscribeContact(ctx, unsubscribeContactInputObject)
Unsubscribe contact from your communication by phone number.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unsubscribeContactInputObject** | [**UnsubscribeContactInputObject**](UnsubscribeContactInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBalanceNotificationSettings**
> UpdateBalanceNotificationSettings(ctx, updateBalanceNotificationSettingsInputObject)
Update balance notification settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateBalanceNotificationSettingsInputObject** | [**UpdateBalanceNotificationSettingsInputObject**](UpdateBalanceNotificationSettingsInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCallbackSettings**
> UpdateCallbackSettings(ctx, updateCallbackSettingsInputObject)
Update callback URL settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateCallbackSettingsInputObject** | [**UpdateCallbackSettingsInputObject**](UpdateCallbackSettingsInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChatDesktopNotificationSettings**
> UpdateChatDesktopNotificationSettings(ctx, updateChatDesktopNotificationSettingsInputObject)
Update chat desktop notification settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateChatDesktopNotificationSettingsInputObject** | [**UpdateChatDesktopNotificationSettingsInputObject**](UpdateChatDesktopNotificationSettingsInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContact**
> ResourceLinkResponse UpdateContact(ctx, updateContactInputObject, id)
Update existing contact.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateContactInputObject** | [**UpdateContactInputObject**](UpdateContactInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContactNote**
> ResourceLinkResponse UpdateContactNote(ctx, updateContactNoteInputObject, id)
Update existing contact note.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateContactNoteInputObject** | [**UpdateContactNoteInputObject**](UpdateContactNoteInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCurrentUser**
> UpdateCurrentUserResponse UpdateCurrentUser(ctx, updateCurrentUserInputObject)
Update current user info.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateCurrentUserInputObject** | [**UpdateCurrentUserInputObject**](UpdateCurrentUserInputObject.md)|  | 

### Return type

[**UpdateCurrentUserResponse**](UpdateCurrentUserResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomField**
> ResourceLinkResponse UpdateCustomField(ctx, updateCustomFieldInputObject, id)
Update existing custom field.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateCustomFieldInputObject** | [**UpdateCustomFieldInputObject**](UpdateCustomFieldInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomFieldValue**
> ResourceLinkResponse UpdateCustomFieldValue(ctx, updateCustomFieldValueInputObject, id)
Update contact's custom field value.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateCustomFieldValueInputObject** | [**UpdateCustomFieldValueInputObject**](UpdateCustomFieldValueInputObject.md)|  | 
  **id** | **string**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInboundMessagesNotificationSettings**
> UpdateInboundMessagesNotificationSettings(ctx, updateInboundMessagesNotificationSettingsInputObject)
Update inbound messages notification settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateInboundMessagesNotificationSettingsInputObject** | [**UpdateInboundMessagesNotificationSettingsInputObject**](UpdateInboundMessagesNotificationSettingsInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

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

# **UpdatePassword**
> UpdatePassword(ctx, updatePasswordInputObject)
Change user password.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatePasswordInputObject** | [**UpdatePasswordInputObject**](UpdatePasswordInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSenderSetting**
> UpdateSenderSetting(ctx, updateSenderSettingInputObject)
Change sender settings

@TODO

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateSenderSettingInputObject** | [**UpdateSenderSettingInputObject**](UpdateSenderSettingInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSurvey**
> ResourceLinkResponse UpdateSurvey(ctx, updateSurveyInputObject, id)
Update existing survey.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateSurveyInputObject** | [**UpdateSurveyInputObject**](UpdateSurveyInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSurveyNode**
> ResourceLinkResponse UpdateSurveyNode(ctx, updateSurveyNodeInputObject, id)
Update existing node.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateSurveyNodeInputObject** | [**UpdateSurveyNodeInputObject**](UpdateSurveyNodeInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTemplate**
> ResourceLinkResponse UpdateTemplate(ctx, updateTemplateInputObject, id)
Update a template



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateTemplateInputObject** | [**UpdateTemplateInputObject**](UpdateTemplateInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

