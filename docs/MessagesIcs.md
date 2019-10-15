# MessagesIcs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Schedule ID. | [default to null]
**NextSend** | [**time.Time**](time.Time.md) | Next send date in [ISO 8601](https://en.wikipedia.org/?title&#x3D;ISO_8601) format.  | [default to null]
**Rrule** | **string** | [iCal RRULE](http://www.kanzaki.com/docs/ical/rrule.html) string.  | [default to null]
**Session** | [***MessageSession**](MessageSession.md) |  | [default to null]
**LastSent** | [**time.Time**](time.Time.md) | Date and time when last message has been sent. | [default to null]
**ContactName** | **string** |  | [default to null]
**Parameters** | [***MessagesIcsParameters**](MessagesIcs_parameters.md) |  | [default to null]
**Type_** | **string** |  | [default to null]
**Summary** | **string** |  | [default to null]
**TextParameters** | [***MessagesIcsTextParameters**](MessagesIcs_textParameters.md) |  | [default to null]
**FirstOccurrence** | [**time.Time**](time.Time.md) |  | [default to null]
**LastOccurrence** | [**time.Time**](time.Time.md) |  | [default to null]
**RecipientsCount** | **int32** | Amount of actual recipients. | [default to null]
**Timezone** | **string** | User-friendly timezone name (with spaces replaced by underscores). | [default to null]
**Completed** | **bool** | Indicates that schedling has been completed. | [default to null]
**Avatar** | **string** | TODO | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Scheduling creation time. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


