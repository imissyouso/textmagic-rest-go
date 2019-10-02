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

type CheckPhoneVerificationCodeInputObject1 struct {
	// Verification code that was received by the user and entered into the form field.
	Code int32 `json:"code"`
	// VerifyId from Step 1 to match both requests together.
	VerifyId string `json:"verifyId"`
}