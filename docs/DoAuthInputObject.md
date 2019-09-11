# DoAuthInputObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Account username or email | [default to null]
**Password** | **string** | Account password | [default to null]
**AppName** | **string** | Application name | [optional] [default to null]
**AppVersion** | **string** | Application version | [optional] [default to null]
**DeviceId** | **string** | Device ID for push notification service | [optional] [default to null]
**PushServiceType** | **string** | required when deviceId provided. Push notification service type: a for Apple Push Notifications, g for Google Cloud Messaging | [optional] [default to null]
**TfaCode** | **string** | required when 2FA enabled on account. 2FA challenge response (SMS code or security question answer) | [optional] [default to null]
**TfaId** | **string** | required when 2FA enabled on account. 2FA challenge response (SMS code or security question answer) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


