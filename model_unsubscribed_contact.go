/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

import (
	"time"
)

type UnsubscribedContact struct {
	// Unsubscribed contact ID.
	Id int32 `json:"id"`
	// Phone number in [E.164 format](https://en.wikipedia.org/wiki/E.164).
	Phone string `json:"phone"`
	// Time when contact has been opted-out.
	UnsubscribeTime time.Time `json:"unsubscribeTime"`
	// Unsubscribed contact first name.
	FirstName string `json:"firstName"`
	// Unsubscribed contact last name.
	LastName string `json:"lastName"`
}
