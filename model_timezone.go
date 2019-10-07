/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type Timezone struct {
	// Internal timezone ID.
	Id int32 `json:"id"`
	// Timezone area.
	Area string `json:"area"`
	// Is Daylight saving time used in this timezone?
	Dst int32 `json:"dst"`
	// Offset from UTC time in seconds. In this example it is 21600/60/60=6 hours.
	Offset int32 `json:"offset"`
	// User-friendly timezone name (with spaces replaced by underscores).
	Timezone string `json:"timezone"`
}
