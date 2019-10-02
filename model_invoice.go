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

type Invoice struct {
	// The invoice ID.
	Id int32 `json:"id"`
	// Top up amount.
	Bundle int32 `json:"bundle"`
	// Top up currency.
	Currency string `json:"currency"`
	// VAT charged (if any).
	Vat float32 `json:"vat"`
	// Payment method description.
	PaymentMethod string `json:"paymentMethod"`
}
