/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type UploadMessageAttachmentResponse struct {
	// `href` field characters count. 
	Chars int32 `json:"chars"`
	// This is a relative link to your file. To construct a full link, just add “[https://my.textmagic.com/”](https://my.textmagic.com/%E2%80%9D) to the beginning (like this: [https://my.textmagic.com/click/Zwcj9](https://my.textmagic.com/click/Zwcj9)). For most modern devices, you can omit “https://” part and write just [my.textmagic.com/click/Zwcj9](https://my.textmagic.com/click/Zwcj9), it will save you 8 characters. 
	Href string `json:"href"`
	// File name of uploaded file. 
	Name string `json:"name"`
	// Attachment size in bytes.
	Size int32 `json:"size"`
}
