# TextMagic\TextMagicApi

All URIs are relative to *http://my.textmagic.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignContactsToList**](TextMagicApi.md#AssignContactsToList) | **Put** /api/v2/lists/{id}/contacts | Assign contacts to the specified list.
[**BlockContact**](TextMagicApi.md#BlockContact) | **Post** /api/v2/contacts/block | Block contact from inbound and outbound communication by phone number.
[**BuyDedicatedNumber**](TextMagicApi.md#BuyDedicatedNumber) | **Post** /api/v2/numbers | Buy a dedicated number and assign it to the specified account.
[**CancelSurvey**](TextMagicApi.md#CancelSurvey) | **Put** /api/v2/surveys/{id}/cancel | Cancel a survey.
[**CheckPhoneVerificationCode**](TextMagicApi.md#CheckPhoneVerificationCode) | **Put** /api/v2/user/phone/verification | Check user phone verification code
[**ClearAndAssignContactsToList**](TextMagicApi.md#ClearAndAssignContactsToList) | **Post** /api/v2/lists/{id}/contacts | Reset list members to the specified contacts.
[**CloseChatsBulk**](TextMagicApi.md#CloseChatsBulk) | **Post** /api/v2/chats/close/bulk | Close chats by chat ids or close all chats
[**CloseReadChats**](TextMagicApi.md#CloseReadChats) | **Post** /api/v2/chats/close/read | Close all chats that have no unread messages.
[**CloseSubaccount**](TextMagicApi.md#CloseSubaccount) | **Delete** /api/v2/subaccounts/{id} | Close subaccount.
[**CreateContact**](TextMagicApi.md#CreateContact) | **Post** /api/v2/contacts | Create a new contact from the submitted data.
[**CreateContactNote**](TextMagicApi.md#CreateContactNote) | **Post** /api/v2/contacts/{id}/notes | Create a new contact note.
[**CreateCustomField**](TextMagicApi.md#CreateCustomField) | **Post** /api/v2/customfields | Create a new custom field from the submitted data.
[**CreateList**](TextMagicApi.md#CreateList) | **Post** /api/v2/lists | Create a new list from the submitted data.
[**CreatePushToken**](TextMagicApi.md#CreatePushToken) | **Post** /api/v2/push/tokens | Add or update a device token.
[**CreateSurvey**](TextMagicApi.md#CreateSurvey) | **Post** /api/v2/surveys | Create a new survey from the submitted data.
[**CreateSurveyNode**](TextMagicApi.md#CreateSurveyNode) | **Post** /api/v2/surveys/{id}/nodes | Create a new node from the submitted data.
[**CreateTemplate**](TextMagicApi.md#CreateTemplate) | **Post** /api/v2/templates | Create a new template from the submitted data.
[**DeleteAllContacts**](TextMagicApi.md#DeleteAllContacts) | **Delete** /api/v2/contact/all | Delete all contacts.
[**DeleteAllOutboundMessages**](TextMagicApi.md#DeleteAllOutboundMessages) | **Delete** /api/v2/message/all | Delete all messages
[**DeleteAvatar**](TextMagicApi.md#DeleteAvatar) | **Delete** /api/v2/user/avatar | Delete an avatar for the current user.\\
[**DeleteChatMessages**](TextMagicApi.md#DeleteChatMessages) | **Post** /api/v2/chats/{id}/messages/delete | Delete messages from chat by given messages ID(s).
[**DeleteChatsBulk**](TextMagicApi.md#DeleteChatsBulk) | **Post** /api/v2/chats/delete | Delete chats by given ID(s) or delete all chats.
[**DeleteContact**](TextMagicApi.md#DeleteContact) | **Delete** /api/v2/contacts/{id} | Delete a single contact.
[**DeleteContactAvatar**](TextMagicApi.md#DeleteContactAvatar) | **Delete** /api/v2/contacts/{id}/avatar | Delete an avatar for the contact.
[**DeleteContactNote**](TextMagicApi.md#DeleteContactNote) | **Delete** /api/v2/notes/{id} | Delete a single contact note.
[**DeleteContactNotesBulk**](TextMagicApi.md#DeleteContactNotesBulk) | **Post** /api/v2/contacts/{id}/notes/delete | Delete contact note by given ID(s) or delete all contact notes.
[**DeleteContactsByIds**](TextMagicApi.md#DeleteContactsByIds) | **Post** /api/v2/contacts/delete | Delete contact by given ID(s) or delete all contacts.
[**DeleteContactsFromList**](TextMagicApi.md#DeleteContactsFromList) | **Delete** /api/v2/lists/{id}/contacts | Unassign contacts from the specified list.
[**DeleteCustomField**](TextMagicApi.md#DeleteCustomField) | **Delete** /api/v2/customfields/{id} | Delete a single custom field.
[**DeleteDedicatedNumber**](TextMagicApi.md#DeleteDedicatedNumber) | **Delete** /api/v2/numbers/{id} | Cancel dedicated number subscription.
[**DeleteInboundMessage**](TextMagicApi.md#DeleteInboundMessage) | **Delete** /api/v2/replies/{id} | Delete the incoming message.
[**DeleteInboundMessagesBulk**](TextMagicApi.md#DeleteInboundMessagesBulk) | **Post** /api/v2/replies/delete | Delete inbound messages by given ID(s) or delete all inbound messages.
[**DeleteList**](TextMagicApi.md#DeleteList) | **Delete** /api/v2/lists/{id} | Delete a single list.
[**DeleteListAvatar**](TextMagicApi.md#DeleteListAvatar) | **Delete** /api/v2/lists/{id}/avatar | Delete an avatar for the list.
[**DeleteListContactsBulk**](TextMagicApi.md#DeleteListContactsBulk) | **Post** /api/v2/lists/{id}/contacts/delete | Delete contact from list by given ID(s) or all contacts from list.
[**DeleteListsBulk**](TextMagicApi.md#DeleteListsBulk) | **Post** /api/v2/lists/delete | Delete list by given ID(s) or delete all lists.
[**DeleteMessageSession**](TextMagicApi.md#DeleteMessageSession) | **Delete** /api/v2/sessions/{id} | Delete a message session, together with all nested messages.
[**DeleteMessageSessionsBulk**](TextMagicApi.md#DeleteMessageSessionsBulk) | **Post** /api/v2/sessions/delete | Delete messages sessions, together with all nested messages, by given ID(s) or delete all messages sessions.
[**DeleteOutboundMessage**](TextMagicApi.md#DeleteOutboundMessage) | **Delete** /api/v2/messages/{id} | Delete message
[**DeleteOutboundMessagesBulk**](TextMagicApi.md#DeleteOutboundMessagesBulk) | **Post** /api/v2/messages/delete | Delete messages by IDs
[**DeletePushToken**](TextMagicApi.md#DeletePushToken) | **Delete** /api/v2/push/tokens/{type}/{deviceId} | Delete a push notification device token.
[**DeleteScheduledMessage**](TextMagicApi.md#DeleteScheduledMessage) | **Delete** /api/v2/schedules/{id} | Delete a message session, together with all nested messages.
[**DeleteScheduledMessagesBulk**](TextMagicApi.md#DeleteScheduledMessagesBulk) | **Post** /api/v2/schedules/delete | Delete scheduled messages by given ID(s) or delete all scheduled messages.
[**DeleteSenderId**](TextMagicApi.md#DeleteSenderId) | **Delete** /api/v2/senderids/{id} | Delete a Sender ID.
[**DeleteSurvey**](TextMagicApi.md#DeleteSurvey) | **Delete** /api/v2/surveys/{id} | Delete a survey.
[**DeleteSurveyNode**](TextMagicApi.md#DeleteSurveyNode) | **Delete** /api/v2/surveys/nodes/{id} | Delete a node.
[**DeleteTemplate**](TextMagicApi.md#DeleteTemplate) | **Delete** /api/v2/templates/{id} | Delete a single template.
[**DeleteTemplatesBulk**](TextMagicApi.md#DeleteTemplatesBulk) | **Post** /api/v2/templates/delete | Delete template by given ID(s) or delete all templates.
[**DoAuth**](TextMagicApi.md#DoAuth) | **Post** /api/v2/auth | Authenticate user by given username and password.
[**DoCarrierLookup**](TextMagicApi.md#DoCarrierLookup) | **Get** /api/v2/lookups/{phone} | Carrier Lookup
[**DoEmailLookup**](TextMagicApi.md#DoEmailLookup) | **Get** /api/v2/email-lookups/{email} | Validate Email address using Email Lookup tool
[**DuplicateSurvey**](TextMagicApi.md#DuplicateSurvey) | **Put** /api/v2/surveys/{id}/duplicate | Duplicate a survey.
[**GetAllBulkSessions**](TextMagicApi.md#GetAllBulkSessions) | **Get** /api/v2/bulks | Get all bulk sending sessions.
[**GetAllChats**](TextMagicApi.md#GetAllChats) | **Get** /api/v2/chats | Get all user chats.
[**GetAllInboundMessages**](TextMagicApi.md#GetAllInboundMessages) | **Get** /api/v2/replies | Get all inbox messages.
[**GetAllMessageSessions**](TextMagicApi.md#GetAllMessageSessions) | **Get** /api/v2/sessions | Get all message sending sessions.
[**GetAllOutboundMessages**](TextMagicApi.md#GetAllOutboundMessages) | **Get** /api/v2/messages | Get all messages
[**GetAllScheduledMessages**](TextMagicApi.md#GetAllScheduledMessages) | **Get** /api/v2/schedules | Get all scheduled messages.
[**GetAllTemplates**](TextMagicApi.md#GetAllTemplates) | **Get** /api/v2/templates | Get all user templates.
[**GetAvailableDedicatedNumbers**](TextMagicApi.md#GetAvailableDedicatedNumbers) | **Get** /api/v2/numbers/available | Find available dedicated numbers to buy.
[**GetAvailableSenderSettingOptions**](TextMagicApi.md#GetAvailableSenderSettingOptions) | **Get** /api/v2/sources | Get all available sender setting options which could be used in \&quot;from\&quot; parameter of POST messages method.
[**GetBalanceNotificationOptions**](TextMagicApi.md#GetBalanceNotificationOptions) | **Get** /api/v2/user/notification/balance/bundles | Returns the list of available balance options which can be used as a bound to determine when to send email to user with low balance notification. See https://my.textmagic.com/online/account/notifications/balance
[**GetBalanceNotificationSettings**](TextMagicApi.md#GetBalanceNotificationSettings) | **Get** /api/v2/user/notification/balance | Get balance notification settings
[**GetBlockedContacts**](TextMagicApi.md#GetBlockedContacts) | **Get** /api/v2/contacts/block/list | Get blocked contacts.
[**GetBulkSession**](TextMagicApi.md#GetBulkSession) | **Get** /api/v2/bulks/{id} | Get bulk message session status.
[**GetCallbackSettings**](TextMagicApi.md#GetCallbackSettings) | **Get** /api/v2/callback/settings | Fetch callback URL settings
[**GetCallsPrices**](TextMagicApi.md#GetCallsPrices) | **Get** /api/v2/calls/price | Check pricing for a inbound/outbound call.
[**GetChat**](TextMagicApi.md#GetChat) | **Get** /api/v2/chats/{id} | Get a single chat.
[**GetChatByPhone**](TextMagicApi.md#GetChatByPhone) | **Get** /api/v2/chats/{phone}/by/phone | Find chats by phone.
[**GetChatMessages**](TextMagicApi.md#GetChatMessages) | **Get** /api/v2/chats/{id}/message | Fetch messages from chat with specified chat id.
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
[**GetDedicatedNumber**](TextMagicApi.md#GetDedicatedNumber) | **Get** /api/v2/numbers/{id} | Get a single dedicated number.
[**GetDisallowedRules**](TextMagicApi.md#GetDisallowedRules) | **Get** /api/v2/user/disallowed-rules | Get an array of all rules that are disallowed to the current account.
[**GetFavourites**](TextMagicApi.md#GetFavourites) | **Get** /api/v2/contacts/favorite | Get favorite contacts and lists.
[**GetForwardedCalls**](TextMagicApi.md#GetForwardedCalls) | **Get** /api/v2/calls | Get all forwarded calls.
[**GetInboundMessage**](TextMagicApi.md#GetInboundMessage) | **Get** /api/v2/replies/{id} | Get a single inbox message.
[**GetInboundMessagesNotificationSettings**](TextMagicApi.md#GetInboundMessagesNotificationSettings) | **Get** /api/v2/user/notification/inbound | Get inbound messages notification settings
[**GetInvoices**](TextMagicApi.md#GetInvoices) | **Get** /api/v2/invoices | Return account invoices.
[**GetList**](TextMagicApi.md#GetList) | **Get** /api/v2/lists/{id} | Get a single list.
[**GetListContactsIds**](TextMagicApi.md#GetListContactsIds) | **Get** /api/v2/lists/{id}/contacts/ids | Fetch all contacts IDs belonging to the list with ID.
[**GetListsOfContact**](TextMagicApi.md#GetListsOfContact) | **Get** /api/v2/contacts/{id}/lists | Return lists which contact belongs to.
[**GetMessagePreview**](TextMagicApi.md#GetMessagePreview) | **Get** /api/v2/messages/preview | Preview message
[**GetMessagePrice**](TextMagicApi.md#GetMessagePrice) | **Get** /api/v2/messages/price | Check price
[**GetMessagePrices**](TextMagicApi.md#GetMessagePrices) | **Get** /api/v2/messages/prices | Get pricing
[**GetMessageSession**](TextMagicApi.md#GetMessageSession) | **Get** /api/v2/sessions/{id} | Get a message session.
[**GetMessageSessionStat**](TextMagicApi.md#GetMessageSessionStat) | **Get** /api/v2/sessions/{id}/stat | Get sending session statistics.
[**GetMessagesBySessionId**](TextMagicApi.md#GetMessagesBySessionId) | **Get** /api/v2/sessions/{id}/messages | Fetch messages by given session id.
[**GetMessagingCounters**](TextMagicApi.md#GetMessagingCounters) | **Get** /api/v2/stats/messaging/data | Return counters for messaging data views.
[**GetMessagingStat**](TextMagicApi.md#GetMessagingStat) | **Get** /api/v2/stats/messaging | Return messaging statistics.
[**GetOutboundMessage**](TextMagicApi.md#GetOutboundMessage) | **Get** /api/v2/messages/{id} | Get a single message
[**GetOutboundMessagesHistory**](TextMagicApi.md#GetOutboundMessagesHistory) | **Get** /api/v2/history | Get history
[**GetPushTokens**](TextMagicApi.md#GetPushTokens) | **Get** /api/v2/push/tokens | Get all device tokens assigned to the current account
[**GetScheduledMessage**](TextMagicApi.md#GetScheduledMessage) | **Get** /api/v2/schedules/{id} | Get message schedule.
[**GetSenderId**](TextMagicApi.md#GetSenderId) | **Get** /api/v2/senderids/{id} | Get a single Sender ID.
[**GetSenderIds**](TextMagicApi.md#GetSenderIds) | **Get** /api/v2/senderids | Get all sender IDs of current user.
[**GetSenderSettings**](TextMagicApi.md#GetSenderSettings) | **Get** /api/v2/sender/settings | Get current user sender settings.
[**GetSpendingStat**](TextMagicApi.md#GetSpendingStat) | **Get** /api/v2/stats/spending | Return account spending statistics.
[**GetState**](TextMagicApi.md#GetState) | **Get** /api/v2/state | Get current entities state
[**GetSubaccount**](TextMagicApi.md#GetSubaccount) | **Get** /api/v2/subaccounts/{id} | Get a single subaccount.
[**GetSubaccounts**](TextMagicApi.md#GetSubaccounts) | **Get** /api/v2/subaccounts | Get all subaccounts of current user.
[**GetSubaccountsWithTokens**](TextMagicApi.md#GetSubaccountsWithTokens) | **Post** /api/v2/subaccounts/tokens/list | Get all subaccounts with their REST API tokens associated with specified app name.
[**GetSurvey**](TextMagicApi.md#GetSurvey) | **Get** /api/v2/surveys/{id} | Get a survey by id.
[**GetSurveyNode**](TextMagicApi.md#GetSurveyNode) | **Get** /api/v2/surveys/nodes/{id} | Get a node by id.
[**GetSurveyNodes**](TextMagicApi.md#GetSurveyNodes) | **Get** /api/v2/surveys/{id}/nodes | Fetch nodes by given survey id.
[**GetSurveys**](TextMagicApi.md#GetSurveys) | **Get** /api/v2/surveys | Get all user surveys.
[**GetTemplate**](TextMagicApi.md#GetTemplate) | **Get** /api/v2/templates/{id} | Get a single template.
[**GetTimezones**](TextMagicApi.md#GetTimezones) | **Get** /api/v2/timezones | Return all available timezone IDs.
[**GetUnreadMessagesTotal**](TextMagicApi.md#GetUnreadMessagesTotal) | **Get** /api/v2/chats/unread/count | Get total amount of unread messages in the current user chats.
[**GetUnsubscribedContact**](TextMagicApi.md#GetUnsubscribedContact) | **Get** /api/v2/unsubscribers/{id} | Get a single unsubscribed contact.
[**GetUnsubscribers**](TextMagicApi.md#GetUnsubscribers) | **Get** /api/v2/unsubscribers | Get all contact have unsubscribed from your communication.
[**GetUserDedicatedNumbers**](TextMagicApi.md#GetUserDedicatedNumbers) | **Get** /api/v2/numbers | Get user&#39;s dedicated numbers.
[**GetUserLists**](TextMagicApi.md#GetUserLists) | **Get** /api/v2/lists | Get all user lists.
[**GetVersions**](TextMagicApi.md#GetVersions) | **Get** /api/v2/versions | Get minimal valid apps versions
[**InviteSubaccount**](TextMagicApi.md#InviteSubaccount) | **Post** /api/v2/subaccounts | Invite new subaccount.
[**MarkChatsReadBulk**](TextMagicApi.md#MarkChatsReadBulk) | **Post** /api/v2/chats/read/bulk | Mark several chats as read by chat ids or mark all chats as read
[**MarkChatsUnreadBulk**](TextMagicApi.md#MarkChatsUnreadBulk) | **Post** /api/v2/chats/unread/bulk | Mark several chats as UNread by chat ids or mark all chats as UNread
[**MergeSurveyNodes**](TextMagicApi.md#MergeSurveyNodes) | **Post** /api/v2/surveys/nodes/merge | Merge two question nodes.
[**MuteChat**](TextMagicApi.md#MuteChat) | **Post** /api/v2/chats/mute | Set mute mode.
[**MuteChatsBulk**](TextMagicApi.md#MuteChatsBulk) | **Post** /api/v2/chats/mute/bulk | Mute several chats by chat ids or mute all chats
[**Ping**](TextMagicApi.md#Ping) | **Get** /api/v2/ping | Just does a pong.
[**ReopenChatsBulk**](TextMagicApi.md#ReopenChatsBulk) | **Post** /api/v2/chats/reopen/bulk | Reopen chats by chat ids or reopen all chats
[**RequestNewSubaccountToken**](TextMagicApi.md#RequestNewSubaccountToken) | **Post** /api/v2/subaccounts/tokens | Request a new REST API token for subaccount.
[**RequestSenderId**](TextMagicApi.md#RequestSenderId) | **Post** /api/v2/senderids | Request for a new Sender ID.
[**ResetSurvey**](TextMagicApi.md#ResetSurvey) | **Put** /api/v2/surveys/{id}/reset | Reset a survey flow.
[**SearchChats**](TextMagicApi.md#SearchChats) | **Get** /api/v2/chats/search | Find chats by inbound or outbound messages text.
[**SearchChatsByIds**](TextMagicApi.md#SearchChatsByIds) | **Get** /api/v2/chats/search/ids | Find chats by IDs.
[**SearchChatsByReceipent**](TextMagicApi.md#SearchChatsByReceipent) | **Get** /api/v2/chats/search/recipients | Find chats by recipient (contact, list name or phone number).
[**SearchContacts**](TextMagicApi.md#SearchContacts) | **Get** /api/v2/contacts/search | Find user contacts by given parameters.
[**SearchInboundMessages**](TextMagicApi.md#SearchInboundMessages) | **Get** /api/v2/replies/search | Find inbound messages by given parameters.
[**SearchLists**](TextMagicApi.md#SearchLists) | **Get** /api/v2/lists/search | Find contact lists by given parameters.
[**SearchOutboundMessages**](TextMagicApi.md#SearchOutboundMessages) | **Get** /api/v2/messages/search | Find messages
[**SearchScheduledMessages**](TextMagicApi.md#SearchScheduledMessages) | **Get** /api/v2/schedules/search | Find scheduled messages by given parameters.
[**SearchTemplates**](TextMagicApi.md#SearchTemplates) | **Get** /api/v2/templates/search | Find user templates by given parameters.
[**SendEmailVerificationCode**](TextMagicApi.md#SendEmailVerificationCode) | **Get** /api/v2/user/email/verification | Send user email verification
[**SendMessage**](TextMagicApi.md#SendMessage) | **Post** /api/v2/messages | Send message
[**SendPhoneVerificationCode**](TextMagicApi.md#SendPhoneVerificationCode) | **Get** /api/v2/user/phone/verification | Send user phone verification
[**SetChatStatus**](TextMagicApi.md#SetChatStatus) | **Post** /api/v2/chats/status | Set status of the chat given by ID.
[**StartSurvey**](TextMagicApi.md#StartSurvey) | **Put** /api/v2/surveys/{id}/start | Start a survey.
[**UnblockContact**](TextMagicApi.md#UnblockContact) | **Post** /api/v2/contacts/unblock | Unblock contact by phone number.
[**UnblockContactsBulk**](TextMagicApi.md#UnblockContactsBulk) | **Post** /api/v2/contacts/unblock/bulk | Unblock several contacts by blocked contact ids or unblock all contacts
[**UnmuteChatsBulk**](TextMagicApi.md#UnmuteChatsBulk) | **Post** /api/v2/chats/unmute/bulk | Unmute several chats by chat ids or unmute all chats
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
[**UpdateSenderSetting**](TextMagicApi.md#UpdateSenderSetting) | **Put** /api/v2/sender/settings | Change sender settings for specified country.
[**UpdateSurvey**](TextMagicApi.md#UpdateSurvey) | **Put** /api/v2/surveys/{id} | Update existing survey.
[**UpdateSurveyNode**](TextMagicApi.md#UpdateSurveyNode) | **Put** /api/v2/surveys/nodes/{id} | Update existing node.
[**UpdateTemplate**](TextMagicApi.md#UpdateTemplate) | **Put** /api/v2/templates/{id} | Update existing template.


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

# **BuyDedicatedNumber**
> BuyDedicatedNumber(ctx, buyDedicatedNumberInputObject, optional)
Buy a dedicated number and assign it to the specified account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buyDedicatedNumberInputObject** | [**BuyDedicatedNumberInputObject**](BuyDedicatedNumberInputObject.md)|  | 
 **optional** | ***BuyDedicatedNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BuyDedicatedNumberOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIgnoreNullValues** | **optional.Bool**|  | [default to true]

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

# **CheckPhoneVerificationCode**
> CheckPhoneVerificationCode(ctx, checkPhoneVerificationCodeInputObject, optional)
Check user phone verification code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **checkPhoneVerificationCodeInputObject** | [**CheckPhoneVerificationCodeInputObject**](CheckPhoneVerificationCodeInputObject.md)|  | 
 **optional** | ***CheckPhoneVerificationCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CheckPhoneVerificationCodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIgnoreNullValues** | **optional.Bool**|  | [default to true]

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
> CloseChatsBulk(ctx, closeChatsBulkInputObject, optional)
Close chats by chat ids or close all chats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **closeChatsBulkInputObject** | [**CloseChatsBulkInputObject**](CloseChatsBulkInputObject.md)|  | 
 **optional** | ***CloseChatsBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloseChatsBulkOpts struct

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

# **CloseReadChats**
> CloseReadChats(ctx, )
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

# **CreateCustomField**
> ResourceLinkResponse CreateCustomField(ctx, createCustomFieldInputObject, optional)
Create a new custom field from the submitted data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createCustomFieldInputObject** | [**CreateCustomFieldInputObject**](CreateCustomFieldInputObject.md)|  | 
 **optional** | ***CreateCustomFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateCustomFieldOpts struct

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

# **CreatePushToken**
> CreatePushToken(ctx, createPushTokenInputObject, optional)
Add or update a device token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createPushTokenInputObject** | [**CreatePushTokenInputObject**](CreatePushTokenInputObject.md)|  | 
 **optional** | ***CreatePushTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreatePushTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIgnoreNullValues** | **optional.Bool**|  | [default to true]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSurvey**
> ResourceLinkResponse CreateSurvey(ctx, createSurveyInputObject, optional)
Create a new survey from the submitted data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSurveyInputObject** | [**CreateSurveyInputObject**](CreateSurveyInputObject.md)|  | 
 **optional** | ***CreateSurveyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateSurveyOpts struct

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

# **CreateSurveyNode**
> ResourceLinkResponse CreateSurveyNode(ctx, createSurveyNodeInputObject, id, optional)
Create a new node from the submitted data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSurveyNodeInputObject** | [**CreateSurveyNodeInputObject**](CreateSurveyNodeInputObject.md)|  | 
  **id** | **int32**|  | 
 **optional** | ***CreateSurveyNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateSurveyNodeOpts struct

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

# **CreateTemplate**
> ResourceLinkResponse CreateTemplate(ctx, createTemplateInputObject, optional)
Create a new template from the submitted data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createTemplateInputObject** | [**CreateTemplateInputObject**](CreateTemplateInputObject.md)|  | 
 **optional** | ***CreateTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateTemplateOpts struct

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
> DeleteChatMessages(ctx, deleteChatMessagesBulkInputObject, id, optional)
Delete messages from chat by given messages ID(s).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteChatMessagesBulkInputObject** | [**DeleteChatMessagesBulkInputObject**](DeleteChatMessagesBulkInputObject.md)|  | 
  **id** | **int32**|  | 
 **optional** | ***DeleteChatMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteChatMessagesOpts struct

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

# **DeleteChatsBulk**
> DeleteChatsBulk(ctx, deleteChatsBulkInputObject, optional)
Delete chats by given ID(s) or delete all chats.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteChatsBulkInputObject** | [**DeleteChatsBulkInputObject**](DeleteChatsBulkInputObject.md)|  | 
 **optional** | ***DeleteChatsBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteChatsBulkOpts struct

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
Cancel dedicated number subscription.

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
Delete the incoming message.

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

# **DeleteInboundMessagesBulk**
> DeleteInboundMessagesBulk(ctx, deleteInboundMessagesBulkInputObject, optional)
Delete inbound messages by given ID(s) or delete all inbound messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteInboundMessagesBulkInputObject** | [**DeleteInboundMessagesBulkInputObject**](DeleteInboundMessagesBulkInputObject.md)|  | 
 **optional** | ***DeleteInboundMessagesBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteInboundMessagesBulkOpts struct

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

# **DeleteMessageSession**
> DeleteMessageSession(ctx, id)
Delete a message session, together with all nested messages.

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
> DeleteMessageSessionsBulk(ctx, deleteMessageSessionsBulkInputObject, optional)
Delete messages sessions, together with all nested messages, by given ID(s) or delete all messages sessions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteMessageSessionsBulkInputObject** | [**DeleteMessageSessionsBulkInputObject**](DeleteMessageSessionsBulkInputObject.md)|  | 
 **optional** | ***DeleteMessageSessionsBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteMessageSessionsBulkOpts struct

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
Delete a message session, together with all nested messages.

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
> DeleteScheduledMessagesBulk(ctx, deleteScheduledMessagesBulkInputObject, optional)
Delete scheduled messages by given ID(s) or delete all scheduled messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteScheduledMessagesBulkInputObject** | [**DeleteScheduledMessagesBulkInputObject**](DeleteScheduledMessagesBulkInputObject.md)|  | 
 **optional** | ***DeleteScheduledMessagesBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteScheduledMessagesBulkOpts struct

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

# **DeleteSenderId**
> DeleteSenderId(ctx, id)
Delete a Sender ID.

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
Delete a single template.

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
> DeleteTemplatesBulk(ctx, deleteTemplatesBulkInputObject, optional)
Delete template by given ID(s) or delete all templates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteTemplatesBulkInputObject** | [**DeleteTemplatesBulkInputObject**](DeleteTemplatesBulkInputObject.md)|  | 
 **optional** | ***DeleteTemplatesBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteTemplatesBulkOpts struct

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

# **DoAuth**
> DoAuthResponse DoAuth(ctx, doAuthInputObject, optional)
Authenticate user by given username and password.

Returning a username and token that you should pass to the all requests (in X-TM-Username and X-TM-Key, respectively)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **doAuthInputObject** | [**DoAuthInputObject**](DoAuthInputObject.md)|  | 
 **optional** | ***DoAuthOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DoAuthOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIgnoreNullValues** | **optional.Bool**|  | [default to true]

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
> GetAllBulkSessionsResponse GetAllBulkSessions(ctx, optional)
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
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]

### Return type

[**GetAllBulkSessionsResponse**](GetAllBulkSessionsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllChats**
> GetAllChatsResponse GetAllChats(ctx, optional)
Get all user chats.

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
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **voice** | **optional.Int32**| Fetch results with voice calls | [default to 0]
 **flat** | **optional.Int32**| Should additional contact info be included | [default to 0]

### Return type

[**GetAllChatsResponse**](GetAllChatsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllInboundMessages**
> GetAllInboundMessagesResponse GetAllInboundMessages(ctx, optional)
Get all inbox messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllInboundMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllInboundMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**GetAllInboundMessagesResponse**](GetAllInboundMessagesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllMessageSessions**
> GetAllMessageSessionsResponse GetAllMessageSessions(ctx, optional)
Get all message sending sessions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllMessageSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllMessageSessionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]

### Return type

[**GetAllMessageSessionsResponse**](GetAllMessageSessionsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllOutboundMessages**
> GetAllOutboundMessagesResponse GetAllOutboundMessages(ctx, optional)
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

[**GetAllOutboundMessagesResponse**](GetAllOutboundMessagesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllScheduledMessages**
> GetAllScheduledMessagesResponse GetAllScheduledMessages(ctx, optional)
Get all scheduled messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllScheduledMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllScheduledMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **status** | **optional.String**| Fetch schedules with the specific status: a - actual, c - completed, x - all | [default to x]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**GetAllScheduledMessagesResponse**](GetAllScheduledMessagesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTemplates**
> GetAllTemplatesResponse GetAllTemplates(ctx, optional)
Get all user templates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | 
 **limit** | **optional.Int32**| How many results to return | 

### Return type

[**GetAllTemplatesResponse**](GetAllTemplatesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvailableDedicatedNumbers**
> GetAvailableDedicatedNumbersResponse GetAvailableDedicatedNumbers(ctx, country, optional)
Find available dedicated numbers to buy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Dedicated number country. Two letters in upper case | 
 **optional** | ***GetAvailableDedicatedNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAvailableDedicatedNumbersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefix** | **optional.Int32**| Desired number prefix. Should include country code (i.e. 447 for GB). In case provide tollfree &#x3D; 1 parameter and there are available tollfree numbers, this parameter will be ignore. | [default to 1]
 **tollfree** | **optional.Int32**| Should we show only tollfree numbers (tollfree available only for US). Default is false. | [default to 0]

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
 **country** | **optional.String**| Return sender setting options available in specific country only. Two upper case characters | 

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
> GetBlockedContactsResponse GetBlockedContacts(ctx, optional)
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

[**GetBlockedContactsResponse**](GetBlockedContactsResponse.md)

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
> map[string]interface{} GetCallsPrices(ctx, )
Check pricing for a inbound/outbound call.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChat**
> Chat GetChat(ctx, id)
Get a single chat.

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
Find chats by phone.

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
> GetChatMessagesResponse GetChatMessages(ctx, id, optional)
Fetch messages from chat with specified chat id.

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

 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **query** | **optional.String**| Find messages by specified search query | 
 **start** | **optional.Int32**| Return messages since specified timestamp only | 
 **end** | **optional.Int32**| Return messages up to specified timestamp only | 
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]
 **voice** | **optional.Int32**| Fetch results with voice calls | [default to 0]

### Return type

[**GetChatMessagesResponse**](GetChatMessagesResponse.md)

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
> GetContactNotesResponse GetContactNotes(ctx, id, optional)
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

[**GetContactNotesResponse**](GetContactNotesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContacts**
> GetContactsResponse GetContacts(ctx, optional)
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

[**GetContactsResponse**](GetContactsResponse.md)

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

# **GetContactsByListId**
> GetContactsByListIdResponse GetContactsByListId(ctx, id, optional)
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

[**GetContactsByListIdResponse**](GetContactsByListIdResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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
> GetCustomFieldsResponse GetCustomFields(ctx, optional)
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
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]

### Return type

[**GetCustomFieldsResponse**](GetCustomFieldsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDedicatedNumber**
> UsersInbound GetDedicatedNumber(ctx, id)
Get a single dedicated number.

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
> []string GetDisallowedRules(ctx, )
Get an array of all rules that are disallowed to the current account.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFavourites**
> GetFavouritesResponse GetFavourites(ctx, optional)
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

[**GetFavouritesResponse**](GetFavouritesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetForwardedCalls**
> GetForwardedCallsResponse GetForwardedCalls(ctx, optional)
Get all forwarded calls.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetForwardedCallsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetForwardedCallsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]

### Return type

[**GetForwardedCallsResponse**](GetForwardedCallsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInboundMessage**
> MessageIn GetInboundMessage(ctx, id)
Get a single inbox message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

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
> GetInvoicesResponse GetInvoices(ctx, optional)
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
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]

### Return type

[**GetInvoicesResponse**](GetInvoicesResponse.md)

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
> GetListsOfContactResponse GetListsOfContact(ctx, id, optional)
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

[**GetListsOfContactResponse**](GetListsOfContactResponse.md)

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
Get a message session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

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
Get sending session statistics.

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
> GetMessagesBySessionIdResponse GetMessagesBySessionId(ctx, id, optional)
Fetch messages by given session id.

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

 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **statuses** | **optional.String**| Find messages by status | 
 **includeDeleted** | **optional.Int32**| Search also in deleted messages | [default to 0]

### Return type

[**GetMessagesBySessionIdResponse**](GetMessagesBySessionIdResponse.md)

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
> GetOutboundMessagesHistoryResponse GetOutboundMessagesHistory(ctx, optional)
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

[**GetOutboundMessagesHistoryResponse**](GetOutboundMessagesHistoryResponse.md)

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
Get message schedule.

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
Get a single Sender ID.

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
> GetSenderIdsResponse GetSenderIds(ctx, optional)
Get all sender IDs of current user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSenderIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSenderIdsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]

### Return type

[**GetSenderIdsResponse**](GetSenderIdsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSenderSettings**
> GetSenderSettingsResponse GetSenderSettings(ctx, optional)
Get current user sender settings.

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
> GetSpendingStatResponse GetSpendingStat(ctx, optional)
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
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **start** | **optional.Int32**| Optional. Start date in unix timestamp format. Default is 7 days ago | 
 **end** | **optional.Int32**| Optional. End date in unix timestamp format. Default is now | 

### Return type

[**GetSpendingStatResponse**](GetSpendingStatResponse.md)

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
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]

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

 **page** | **optional.Float32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **xIgnoreNullValues** | **optional.Bool**|  | [default to true]

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
> GetSurveysResponse GetSurveys(ctx, optional)
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
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]

### Return type

[**GetSurveysResponse**](GetSurveysResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplate**
> MessageTemplate GetTemplate(ctx, id)
Get a single template.

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

# **GetUnreadMessagesTotal**
> GetUnreadMessagesTotalResponse GetUnreadMessagesTotal(ctx, )
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
> GetUnsubscribersResponse GetUnsubscribers(ctx, optional)
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

[**GetUnsubscribersResponse**](GetUnsubscribersResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserDedicatedNumbers**
> GetUserDedicatedNumbersResponse GetUserDedicatedNumbers(ctx, optional)
Get user's dedicated numbers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUserDedicatedNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUserDedicatedNumbersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **surveyId** | **optional.Int32**| Fetch only that numbers which are ready for the survey | 

### Return type

[**GetUserDedicatedNumbersResponse**](GetUserDedicatedNumbersResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserLists**
> GetUserListsResponse GetUserLists(ctx, optional)
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

[**GetUserListsResponse**](GetUserListsResponse.md)

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
> InviteSubaccount(ctx, inviteSubaccountInputObject, optional)
Invite new subaccount.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inviteSubaccountInputObject** | [**InviteSubaccountInputObject**](InviteSubaccountInputObject.md)|  | 
 **optional** | ***InviteSubaccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InviteSubaccountOpts struct

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

# **MarkChatsReadBulk**
> MarkChatsReadBulk(ctx, markChatsReadBulkInputObject, optional)
Mark several chats as read by chat ids or mark all chats as read

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **markChatsReadBulkInputObject** | [**MarkChatsReadBulkInputObject**](MarkChatsReadBulkInputObject.md)|  | 
 **optional** | ***MarkChatsReadBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarkChatsReadBulkOpts struct

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

# **MarkChatsUnreadBulk**
> MarkChatsUnreadBulk(ctx, markChatsUnreadBulkInputObject, optional)
Mark several chats as UNread by chat ids or mark all chats as UNread

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **markChatsUnreadBulkInputObject** | [**MarkChatsUnreadBulkInputObject**](MarkChatsUnreadBulkInputObject.md)|  | 
 **optional** | ***MarkChatsUnreadBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarkChatsUnreadBulkOpts struct

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

# **MergeSurveyNodes**
> MergeSurveyNodes(ctx, mergeSurveyNodesInputObject, optional)
Merge two question nodes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mergeSurveyNodesInputObject** | [**MergeSurveyNodesInputObject**](MergeSurveyNodesInputObject.md)|  | 
 **optional** | ***MergeSurveyNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeSurveyNodesOpts struct

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

# **MuteChat**
> ResourceLinkResponse MuteChat(ctx, muteChatInputObject, optional)
Set mute mode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **muteChatInputObject** | [**MuteChatInputObject**](MuteChatInputObject.md)|  | 
 **optional** | ***MuteChatOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MuteChatOpts struct

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

# **MuteChatsBulk**
> MuteChatsBulk(ctx, muteChatsBulkInputObject, optional)
Mute several chats by chat ids or mute all chats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **muteChatsBulkInputObject** | [**MuteChatsBulkInputObject**](MuteChatsBulkInputObject.md)|  | 
 **optional** | ***MuteChatsBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MuteChatsBulkOpts struct

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
> ReopenChatsBulk(ctx, reopenChatsBulkInputObject, optional)
Reopen chats by chat ids or reopen all chats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reopenChatsBulkInputObject** | [**ReopenChatsBulkInputObject**](ReopenChatsBulkInputObject.md)|  | 
 **optional** | ***ReopenChatsBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReopenChatsBulkOpts struct

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

# **RequestNewSubaccountToken**
> User RequestNewSubaccountToken(ctx, requestNewSubaccountTokenInputObject, optional)
Request a new REST API token for subaccount.

Returning user object, key and app name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestNewSubaccountTokenInputObject** | [**RequestNewSubaccountTokenInputObject**](RequestNewSubaccountTokenInputObject.md)|  | 
 **optional** | ***RequestNewSubaccountTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RequestNewSubaccountTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIgnoreNullValues** | **optional.Bool**|  | [default to true]

### Return type

[**User**](User.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestSenderId**
> ResourceLinkResponse RequestSenderId(ctx, requestSenderIdInputObject, optional)
Request for a new Sender ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestSenderIdInputObject** | [**RequestSenderIdInputObject**](RequestSenderIdInputObject.md)|  | 
 **optional** | ***RequestSenderIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RequestSenderIdOpts struct

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
> SearchChatsResponse SearchChats(ctx, optional)
Find chats by inbound or outbound messages text.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchChatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchChatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **query** | **optional.String**| Find chats by specified search query | 

### Return type

[**SearchChatsResponse**](SearchChatsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchChatsByIds**
> SearchChatsByIdsResponse SearchChatsByIds(ctx, optional)
Find chats by IDs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchChatsByIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchChatsByIdsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **ids** | **optional.String**| Find chats by ID(s) | 

### Return type

[**SearchChatsByIdsResponse**](SearchChatsByIdsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchChatsByReceipent**
> SearchChatsByReceipentResponse SearchChatsByReceipent(ctx, optional)
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
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **query** | **optional.String**| Find chats by specified search query | 
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]

### Return type

[**SearchChatsByReceipentResponse**](SearchChatsByReceipentResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchContacts**
> SearchContactsResponse SearchContacts(ctx, optional)
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

[**SearchContactsResponse**](SearchContactsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchInboundMessages**
> SearchInboundMessagesResponse SearchInboundMessages(ctx, optional)
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
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **ids** | **optional.String**| Find message by ID(s) | 
 **query** | **optional.String**| Find recipients by specified search query | 
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]
 **expand** | **optional.Int32**| Expand by adding firstName, lastName and contactId | [default to 0]

### Return type

[**SearchInboundMessagesResponse**](SearchInboundMessagesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchLists**
> SearchListsResponse SearchLists(ctx, optional)
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

[**SearchListsResponse**](SearchListsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOutboundMessages**
> SearchOutboundMessagesResponse SearchOutboundMessages(ctx, optional)
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

[**SearchOutboundMessagesResponse**](SearchOutboundMessagesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchScheduledMessages**
> SearchScheduledMessagesResponse SearchScheduledMessages(ctx, optional)
Find scheduled messages by given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchScheduledMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchScheduledMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **query** | **optional.String**| Find messages by specified search query | 
 **ids** | **optional.String**| Find schedules by ID(s) | 
 **status** | **optional.String**| Fetch schedules with the specific status: a - actual, c - completed, x - all | [default to x]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**SearchScheduledMessagesResponse**](SearchScheduledMessagesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTemplates**
> SearchTemplatesResponse SearchTemplates(ctx, optional)
Find user templates by given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page | [default to 1]
 **limit** | **optional.Int32**| How many results to return | [default to 10]
 **ids** | **optional.String**| Find template by ID(s) | 
 **name** | **optional.String**| Find template by name | 
 **content** | **optional.String**| Find template by content | 

### Return type

[**SearchTemplatesResponse**](SearchTemplatesResponse.md)

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

# **SetChatStatus**
> ResourceLinkResponse SetChatStatus(ctx, setChatStatusInputObject, optional)
Set status of the chat given by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **setChatStatusInputObject** | [**SetChatStatusInputObject**](SetChatStatusInputObject.md)|  | 
 **optional** | ***SetChatStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetChatStatusOpts struct

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

# **UnmuteChatsBulk**
> UnmuteChatsBulk(ctx, unmuteChatsBulkInputObject, optional)
Unmute several chats by chat ids or unmute all chats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unmuteChatsBulkInputObject** | [**UnmuteChatsBulkInputObject**](UnmuteChatsBulkInputObject.md)|  | 
 **optional** | ***UnmuteChatsBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UnmuteChatsBulkOpts struct

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

# **UpdateBalanceNotificationSettings**
> UpdateBalanceNotificationSettings(ctx, updateBalanceNotificationSettingsInputObject, optional)
Update balance notification settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateBalanceNotificationSettingsInputObject** | [**UpdateBalanceNotificationSettingsInputObject**](UpdateBalanceNotificationSettingsInputObject.md)|  | 
 **optional** | ***UpdateBalanceNotificationSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateBalanceNotificationSettingsOpts struct

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

# **UpdateCallbackSettings**
> UpdateCallbackSettings(ctx, updateCallbackSettingsInputObject, optional)
Update callback URL settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateCallbackSettingsInputObject** | [**UpdateCallbackSettingsInputObject**](UpdateCallbackSettingsInputObject.md)|  | 
 **optional** | ***UpdateCallbackSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateCallbackSettingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIgnoreNullValues** | **optional.Bool**|  | [default to true]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChatDesktopNotificationSettings**
> UpdateChatDesktopNotificationSettings(ctx, updateChatDesktopNotificationSettingsInputObject, optional)
Update chat desktop notification settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateChatDesktopNotificationSettingsInputObject** | [**UpdateChatDesktopNotificationSettingsInputObject**](UpdateChatDesktopNotificationSettingsInputObject.md)|  | 
 **optional** | ***UpdateChatDesktopNotificationSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateChatDesktopNotificationSettingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIgnoreNullValues** | **optional.Bool**|  | [default to true]

### Return type

 (empty response body)

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

# **UpdateCurrentUser**
> UpdateCurrentUserResponse UpdateCurrentUser(ctx, updateCurrentUserInputObject, optional)
Update current user info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateCurrentUserInputObject** | [**UpdateCurrentUserInputObject**](UpdateCurrentUserInputObject.md)|  | 
 **optional** | ***UpdateCurrentUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateCurrentUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIgnoreNullValues** | **optional.Bool**|  | [default to true]

### Return type

[**UpdateCurrentUserResponse**](UpdateCurrentUserResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomField**
> ResourceLinkResponse UpdateCustomField(ctx, updateCustomFieldInputObject, id, optional)
Update existing custom field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateCustomFieldInputObject** | [**UpdateCustomFieldInputObject**](UpdateCustomFieldInputObject.md)|  | 
  **id** | **int32**|  | 
 **optional** | ***UpdateCustomFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateCustomFieldOpts struct

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

# **UpdateCustomFieldValue**
> ResourceLinkResponse UpdateCustomFieldValue(ctx, updateCustomFieldValueInputObject, id, optional)
Update contact's custom field value.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateCustomFieldValueInputObject** | [**UpdateCustomFieldValueInputObject**](UpdateCustomFieldValueInputObject.md)|  | 
  **id** | **string**|  | 
 **optional** | ***UpdateCustomFieldValueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateCustomFieldValueOpts struct

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

# **UpdateInboundMessagesNotificationSettings**
> UpdateInboundMessagesNotificationSettings(ctx, updateInboundMessagesNotificationSettingsInputObject, optional)
Update inbound messages notification settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateInboundMessagesNotificationSettingsInputObject** | [**UpdateInboundMessagesNotificationSettingsInputObject**](UpdateInboundMessagesNotificationSettingsInputObject.md)|  | 
 **optional** | ***UpdateInboundMessagesNotificationSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateInboundMessagesNotificationSettingsOpts struct

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
> UpdatePassword(ctx, updatePasswordInputObject, optional)
Change user password.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatePasswordInputObject** | [**UpdatePasswordInputObject**](UpdatePasswordInputObject.md)|  | 
 **optional** | ***UpdatePasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdatePasswordOpts struct

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

# **UpdateSenderSetting**
> UpdateSenderSetting(ctx, updateSenderSettingInputObject, optional)
Change sender settings for specified country.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateSenderSettingInputObject** | [**UpdateSenderSettingInputObject**](UpdateSenderSettingInputObject.md)|  | 
 **optional** | ***UpdateSenderSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateSenderSettingOpts struct

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

# **UpdateSurvey**
> ResourceLinkResponse UpdateSurvey(ctx, updateSurveyInputObject, id, optional)
Update existing survey.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateSurveyInputObject** | [**UpdateSurveyInputObject**](UpdateSurveyInputObject.md)|  | 
  **id** | **int32**|  | 
 **optional** | ***UpdateSurveyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateSurveyOpts struct

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

# **UpdateSurveyNode**
> ResourceLinkResponse UpdateSurveyNode(ctx, updateSurveyNodeInputObject, id, optional)
Update existing node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateSurveyNodeInputObject** | [**UpdateSurveyNodeInputObject**](UpdateSurveyNodeInputObject.md)|  | 
  **id** | **int32**|  | 
 **optional** | ***UpdateSurveyNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateSurveyNodeOpts struct

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

# **UpdateTemplate**
> ResourceLinkResponse UpdateTemplate(ctx, updateTemplateInputObject, id, optional)
Update existing template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateTemplateInputObject** | [**UpdateTemplateInputObject**](UpdateTemplateInputObject.md)|  | 
  **id** | **int32**|  | 
 **optional** | ***UpdateTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateTemplateOpts struct

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

