/*
 * TextMagic API Documentation
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Contact: support@textmagi.biz
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type SearchTemplatesResponse struct {
	Page int32 `json:"page"`
	PageCount int32 `json:"pageCount"`
	Limit int32 `json:"limit"`
	Resources []MessageTemplate `json:"resources"`
}
