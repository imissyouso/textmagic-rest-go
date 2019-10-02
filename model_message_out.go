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

type MessageOut struct {
	// Message ID.
	Id int32 `json:"id"`
	// Message sender (phone number or alphanumeric Sender ID).
	Sender string `json:"sender,omitempty"`
	// Recipient phone number.
	Receiver string `json:"receiver,omitempty"`
	Text string `json:"text"`
	// Delivery status of the message. TODO: Please see the table below to see different delivery statuses. 
	Status string `json:"status"`
	ContactId int32 `json:"contactId"`
	SessionId int32 `json:"sessionId"`
	// Sending time.
	MessageTime time.Time `json:"messageTime"`
	Avatar string `json:"avatar"`
	Deleted bool `json:"deleted,omitempty"`
	// Message charset. Could be: *   **ISO-8859-1** for plaintext SMS *   **UTF-16BE** for Unicode SMS 
	Charset string `json:"charset"`
	CharsetLabel string `json:"charsetLabel"`
	// TODO: Contact first name. Could be substituted from your [Contacts](/docs/api/contacts/) (even if you submitted phone number instead of contact ID). 
	FirstName string `json:"firstName"`
	// Contact last name.
	LastName string `json:"lastName"`
	// Two-letter ISO country code of the recipient phone number. 
	Country string `json:"country"`
	Phone string `json:"phone,omitempty"`
	Price float32 `json:"price,omitempty"`
	// Message parts (multiples of 160 characters) count.
	PartsCount int32 `json:"partsCount"`
	FromEmail string `json:"fromEmail,omitempty"`
	FromNumber string `json:"fromNumber,omitempty"`
	SmscId string `json:"smscId,omitempty"`
	Contact string `json:"contact,omitempty"`
	Source string `json:"source,omitempty"`
	DeliveredCount int32 `json:"deliveredCount,omitempty"`
	NumbersCount int32 `json:"numbersCount,omitempty"`
	UserId int32 `json:"userId,omitempty"`
	CreditsPrice string `json:"creditsPrice,omitempty"`
	Chars int32 `json:"chars,omitempty"`
}
