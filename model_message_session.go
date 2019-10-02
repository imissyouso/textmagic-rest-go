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

type MessageSession struct {
	// Session ID.
	Id int32 `json:"id"`
	// Session creation time.
	StartTime string `json:"startTime"`
	// Session text. If a template was used for the session text (see [Messages: Send](#tag/Outbound-Messages) for details), it may contain template tags. 
	Text string `json:"text"`
	// *   **O** for TextMagic Online *   **A** for API *   **M** for TextMagic Messenger *   **E** for [Email to SMS](/docs/api/send-email-to-sms/) *   **X** for [Distribution lists](/docs/api/distribution-lists/) 
	Source string `json:"source"`
	// Custom reference ID (see [Messages: Send](/docs/api/send-sms/) for details). 
	ReferenceId string `json:"referenceId"`
	// Session cost (in account currency).
	Price float32 `json:"price"`
	// Session recipient count.
	NumbersCount int32 `json:"numbersCount"`
	Destination string `json:"destination"`
}
