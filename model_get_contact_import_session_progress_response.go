/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type GetContactImportSessionProgressResponse struct {
	Status int32 `json:"status"`
	// How many contacts have been imported.
	Processed int32 `json:"processed"`
}
