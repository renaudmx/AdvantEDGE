/*
 * Location API
 *
 * The ETSI MEC ISG MEC012 Location API described using OpenAPI. The API is based on the Open Mobile Alliance's specification RESTful Network API for Zonal Presence
 *
 * API version: 1.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

// A type containing access point information.
type AccessPointInfo struct {
	AccessPointId string `json:"accessPointId"`

	LocationInfo *LocationInfo `json:"locationInfo,omitempty"`

	ConnectionType *ConnectionType `json:"connectionType"`

	OperationStatus *OperationStatus `json:"operationStatus"`

	NumberOfUsers uint32 `json:"numberOfUsers"`

	Timezone string `json:"timezone,omitempty"`

	InterestRealm string `json:"interestRealm,omitempty"`

	ResourceURL string `json:"resourceURL"`
}
