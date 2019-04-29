/*
 * MEEP TC controller API
 *
 * Copyright (c) 2019 InterDigital Communications, Inc. All rights reserved. The information provided herein is the proprietary and confidential information of InterDigital Communications, Inc. 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Client basic information object
type ClientBasicInfo struct {

	// Unique pod identifier
	PodId string `json:"podId,omitempty"`

	// IP address of the pod (client)
	Ip string `json:"ip,omitempty"`
}
