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

type MessagesIcsTextParameters struct {
	Cost float32 `json:"cost"`
	Parts int32 `json:"parts"`
	Chars int32 `json:"chars"`
	Encoding string `json:"encoding"`
	Countries []string `json:"countries"`
	CharsetLabel string `json:"charsetLabel"`
}
