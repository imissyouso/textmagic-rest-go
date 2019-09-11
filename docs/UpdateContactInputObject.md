# UpdateContactInputObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | Contact first name | [optional] [default to null]
**LastName** | **string** | Contact last name | [optional] [default to null]
**Phone** | **string** | Contact phone number in E.164 (international) format without leading + or zeroes | [default to null]
**Email** | **string** | Contact email | [optional] [default to null]
**CompanyName** | **string** | Contact company name | [optional] [default to null]
**Lists** | **string** | Array of list resources id contact will be assigned to | [default to null]
**Favorited** | **bool** | Is contact favorited | [optional] [default to null]
**Blocked** | **bool** | Is contact blocked for outgoing and incoming messaging | [optional] [default to null]
**Type_** | **int32** | Force type of phone. Possible values: 0 - landline, 1 - mobile. Default is -1 (auto detection) | [optional] [default to null]
**CustomFieldValues** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Local** | **int32** | Treat phone number passed in request body as local | [optional] [default to null]
**Country** | **string** | 2-letter ISO country code for local phone numbers, used when local is  set to true. Default is account country | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


