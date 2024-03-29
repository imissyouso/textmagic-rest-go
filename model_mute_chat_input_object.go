/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type MuteChatInputObject struct {
	// Chat ID.
	Id int32 `json:"id"`
	// Mute notifications sound.
	Mute bool `json:"mute"`
	// Mute for N hours.
	For_ int32 `json:"for,omitempty"`
}
