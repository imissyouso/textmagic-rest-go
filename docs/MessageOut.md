# MessageOut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Message ID. | [default to null]
**Sender** | **string** | Message sender (phone number or alphanumeric Sender ID). | [optional] [default to null]
**Receiver** | **string** | Recipient phone number. | [optional] [default to null]
**Text** | **string** |  | [default to null]
**Status** | **string** | Delivery status of the message. @TODO: Please see the table below to see different delivery statuses.  | [default to null]
**ContactId** | **int32** |  | [default to null]
**SessionId** | **int32** |  | [default to null]
**MessageTime** | [**time.Time**](time.Time.md) | Sending time. | [default to null]
**Avatar** | **string** |  | [default to null]
**Deleted** | **bool** |  | [optional] [default to null]
**Charset** | **string** | Message charset. Could be: *   **ISO-8859-1** for plaintext SMS *   **UTF-16BE** for Unicode SMS  | [default to null]
**CharsetLabel** | **string** |  | [default to null]
**FirstName** | **string** | @TODO: Contact first name. Could be substituted from your [Contacts](/docs/api/contacts/) (even if you submitted phone number instead of contact ID).  | [default to null]
**LastName** | **string** | Contact last name. | [default to null]
**Country** | **string** | Two-letter ISO country code of the recipient phone number.  | [default to null]
**Phone** | **string** |  | [optional] [default to null]
**Price** | **float32** |  | [optional] [default to null]
**PartsCount** | **int32** | Message parts (multiples of 160 characters) count. | [default to null]
**FromEmail** | **string** |  | [optional] [default to null]
**FromNumber** | **string** |  | [optional] [default to null]
**SmscId** | **string** |  | [optional] [default to null]
**Contact** | **string** |  | [optional] [default to null]
**Source** | **string** |  | [optional] [default to null]
**DeliveredCount** | **int32** |  | [optional] [default to null]
**NumbersCount** | **int32** |  | [optional] [default to null]
**UserId** | **int32** |  | [optional] [default to null]
**CreditsPrice** | **string** |  | [optional] [default to null]
**Chars** | **int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


