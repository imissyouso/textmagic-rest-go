/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type GetMessagePriceResponseCountriesItem struct {
	// Two-letter ISO country code
	Country string `json:"country"`
	// Country name
	CountryName string `json:"country_name"`
	// Is allow to use dedicated number
	AllowDedicated bool `json:"allow_dedicated"`
	// Parts count to send
	Count float32 `json:"count"`
	// Maximum parts to send
	Max float32 `json:"max"`
	// Total price to send
	Sum string `json:"sum"`
	// Is this landline number?
	Landline float32 `json:"landline"`
}
