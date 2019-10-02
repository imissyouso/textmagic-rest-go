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

type SenderId struct {
	// Numeric sender ID.
	Id int32 `json:"id"`
	// Alphanumeric ID.
	SenderId string `json:"senderId"`
	User *User `json:"user"`
	// *   **P** for Pending. This Sender ID is being reviewed by our support team. *   **R** for Rejected. Our support team rejected your application for this Sender ID. *   **A** for Active. 
	Status string `json:"status"`
}
