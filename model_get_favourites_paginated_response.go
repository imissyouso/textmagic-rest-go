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

type GetFavouritesPaginatedResponse struct {
	Page int32 `json:"page"`
	// The total number of pages.
	PageCount int32 `json:"pageCount"`
	// The number of results per page.
	Limit int32 `json:"limit"`
	Resources []FavoriteContact `json:"resources"`
}
