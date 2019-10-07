/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type UpdateSurveyInputObject struct {
	// Survey name
	Name string `json:"name"`
	// Array of contact resources id message will be sent to
	Contacts string `json:"contacts,omitempty"`
	// Array of list resources id message will be sent to
	Lists string `json:"lists,omitempty"`
	// Array of E.164 phone numbers message will be sent to
	Phones string `json:"phones,omitempty"`
	// Country values mapping (country => inbound phone id), e.g. country[GB] = \"123\", country[US] = \"123\"
	Country *interface{} `json:"country,omitempty"`
}
