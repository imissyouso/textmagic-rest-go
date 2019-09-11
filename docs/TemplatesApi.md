# TextMagic\TemplatesApi

All URIs are relative to *http://my.textmagic.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplate**](TemplatesApi.md#CreateTemplate) | **Post** /api/v2/templates | Create a new template from the submitted data.
[**DeleteTemplate**](TemplatesApi.md#DeleteTemplate) | **Delete** /api/v2/templates/{id} | Delete a single template.
[**DeleteTemplatesBulk**](TemplatesApi.md#DeleteTemplatesBulk) | **Post** /api/v2/templates/delete | Delete template by given ID(s) or delete all templates.
[**GetAllTemplates**](TemplatesApi.md#GetAllTemplates) | **Get** /api/v2/templates | Get all user templates.
[**GetTemplate**](TemplatesApi.md#GetTemplate) | **Get** /api/v2/templates/{id} | Get a single template.
[**SearchTemplates**](TemplatesApi.md#SearchTemplates) | **Get** /api/v2/templates/search | Find user templates by given parameters.
[**UpdateTemplate**](TemplatesApi.md#UpdateTemplate) | **Put** /api/v2/templates/{id} | Update existing template.


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

# **GetAllTemplates**
> map[string]interface{} GetAllTemplates(ctx, optional)
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

[**map[string]interface{}**](interface{}.md)

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

# **SearchTemplates**
> map[string]interface{} SearchTemplates(ctx, optional)
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

[**map[string]interface{}**](interface{}.md)

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

