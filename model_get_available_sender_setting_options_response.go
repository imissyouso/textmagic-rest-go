/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type GetAvailableSenderSettingOptionsResponse struct {
	// Array of dedicated numbers strings.
	Dedicated []string `json:"dedicated"`
	// Array of verified account phone numbers (currently only one).
	User []string `json:"user"`
	// Array of shared numbers strings.
	Shared []string `json:"shared"`
	// Array of alphanumeric sender IDs.
	SenderIds []string `json:"senderIds"`
}
