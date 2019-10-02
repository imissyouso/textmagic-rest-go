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

import (
	"time"
)

type Conversation struct {
	Id int32 `json:"id"`
	// Message type: inbound or outbound. 
	Direction string `json:"direction"`
	// Sender phone number.
	Sender string `json:"sender"`
	// Time when message arrived at TextMagic.
	MessageTime time.Time `json:"messageTime"`
	// Message text.
	Text string `json:"text"`
	// Receiver phone number.
	Receiver string `json:"receiver"`
	// Message status (for chats outbound only). See [message delivery statuses](/docs/api/sms-sessions/#message-delivery-statuses) for details.
	Status string `json:"status"`
	// Contact first name.
	FirstName string `json:"firstName"`
	// Contact last name.
	LastName string `json:"lastName"`
	SessionId int32 `json:"sessionId"`
}
