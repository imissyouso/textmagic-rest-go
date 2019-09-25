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

type UpdateListObject struct {
	// List name
	Name string `json:"name"`
	// Should this list be shared with sub-accounts
	Shared bool `json:"shared"`
	// Is list favorited. Default is false
	Favorited bool `json:"favorited,omitempty"`
	// Is list default for new contacts (web only).
	IsDefault bool `json:"isDefault,omitempty"`
}
