# MessageSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Session ID. | [default to null]
**StartTime** | **string** | Session creation time. | [default to null]
**Text** | **string** | Session text. If a template was used for the session text (see [Messages: Send](#tag/Outbound-Messages) for details), it may contain template tags.  | [default to null]
**Source** | **string** | *   **O** for TextMagic Online *   **A** for API *   **M** for TextMagic Messenger *   **E** for [Email to SMS](/docs/api/send-email-to-sms/) *   **X** for [Distribution lists](/docs/api/distribution-lists/)  | [default to null]
**ReferenceId** | **string** | Custom reference ID (see [Messages: Send](/docs/api/send-sms/) for details).  | [default to null]
**Price** | **float32** | Session cost (in account currency). | [default to null]
**NumbersCount** | **int32** | Session recipient count. | [default to null]
**Destination** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


