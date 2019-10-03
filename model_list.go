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
	Id int32 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Favorited bool `json:"favorited"`
	MembersCount int32 `json:"membersCount"`
	User *User `json:"user"`
	Service bool `json:"service"`
	Shared bool `json:"shared"`
	Avatar *ListImage `json:"avatar"`
	IsDefault bool `json:"isDefault"`
}