/*
 * MEEP Location Service - Notification Callback
 *
 * This is MEEP Location Services Notification Callback API
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// A type containing zone information.
type ZoneInfo struct {
	ZoneId string `json:"zoneId"`
	NumberOfAccessPoints uint32 `json:"numberOfAccessPoints"`
	NumberOfUsers uint32 `json:"numberOfUsers"`
}
