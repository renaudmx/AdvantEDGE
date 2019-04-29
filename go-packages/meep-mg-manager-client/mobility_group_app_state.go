/*
 * MEEP Mobility Group Manager REST API
 *
 * Copyright (c) 2019 InterDigital Communications, Inc. All rights reserved. The information provided herein is the proprietary and confidential information of InterDigital Communications, Inc. 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Mobility Group Application State
type MobilityGroupAppState struct {

	// Mobility Group UE Identifier
	UeId string `json:"ueId,omitempty"`

	// Mobility Group Application State for provided UE
	UeState string `json:"ueState,omitempty"`
}
