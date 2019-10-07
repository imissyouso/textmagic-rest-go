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

type DoEmailLookupResponse struct {
	// The email address passed to the call.
	Address string `json:"address"`
	// The email is `valid` or `invalid`.
	Status string `json:"status"`
	// The delivery status of the email address is`deliverable`, `undeliverable`  or `unknown`.
	Deliverability string `json:"deliverability"`
	// The reason why the checked email is invalid/undeliverable.
	Reason string `json:"reason"`
	// The risk score of the email is`high`, `medium`, `low` or `null`.
	Risk string `json:"risk"`
	// The email address type (domain) is `free` or `corporate`.
	AddressType string `json:"addressType"`
	// This is be `true` if the domain is in the list of disposable email addresses, otherwise returns as `false`.
	IsDisposableAddress bool `json:"isDisposableAddress"`
	// Null if nothing is suggested, however, if there is a potential typo in the email address, the closest suggestion is provided.
	Suggestion string `json:"suggestion"`
	// Checks the mailbox part of the email whether it matches a specific role type (‘admin’, ‘sales’, ‘webmaster’)
	EmailRole string `json:"emailRole"`
	// The local part of the email address.
	LocalPart string `json:"localPart"`
	// The domain part of the email address.
	DomainPart string `json:"domainPart"`
	// Email exchange server domain name (MX record value).
	Exchange string `json:"exchange"`
	// MX record preference.
	Preference int32 `json:"preference"`
	// `true` if the email address exists in TextMagic whitelist. 
	IsInWhiteList bool `json:"isInWhiteList"`
	// `true` if the email address exists in TextMagic blacklist. 
	IsInBlackList bool `json:"isInBlackList"`
	// `true` if the email address domain has an MX record. 
	HasMx bool `json:"hasMx"`
	HasAa bool `json:"hasAa"`
	// `true` if the email address domain has an AAAA record (IPv6). 
	HasAaaa bool `json:"hasAaaa"`
}
