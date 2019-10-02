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

type GetOutboundMessagesHistoryPaginatedResponse struct {
	LastId int32 `json:"lastId"`
	NextLastId int32 `json:"nextLastId"`
	// The number of results per page.
	Limit int32 `json:"limit"`
	Resources []MessageOut `json:"resources"`
}