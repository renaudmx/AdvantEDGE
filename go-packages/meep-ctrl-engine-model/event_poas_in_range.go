/*
 * MEEP Model
 *
 * Copyright (c) 2019 InterDigital Communications, Inc. All rights reserved. The information provided herein is the proprietary and confidential information of InterDigital Communications, Inc. 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// POAs In Range Event object
type EventPoasInRange struct {

	// UE identifier
	Ue string `json:"ue,omitempty"`

	PoasInRange []string `json:"poasInRange,omitempty"`
}
