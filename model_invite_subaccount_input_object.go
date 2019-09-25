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

type InviteSubaccountInputObject struct {
	// Subaccount email
	Email string `json:"email"`
	// Subaccount role: A for administrator or U for regular user
	Role string `json:"role"`
}
