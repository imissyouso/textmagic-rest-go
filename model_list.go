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

type List struct {
	// List ID.
	Id int32 `json:"id"`
	// List name.
	Name string `json:"name"`
	// List description.
	Description string `json:"description"`
	Favorited bool `json:"favorited"`
	// List members count.
	MembersCount int32 `json:"membersCount"`
	User *User `json:"user"`
	Service bool `json:"service"`
	// Is the list **shared** among all sub-accounts?
	Shared bool `json:"shared"`
	Avatar *ListImage `json:"avatar"`
	IsDefault bool `json:"isDefault"`
}
