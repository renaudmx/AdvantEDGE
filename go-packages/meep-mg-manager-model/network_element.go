/*
 * MEEP Mobility Group Manager Model
 *
 * Copyright (c) 2019 InterDigital Communications, Inc. All rights reserved. The information provided herein is the proprietary and confidential information of InterDigital Communications, Inc. 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Network element service mappings
type NetworkElement struct {

	// Network element name
	Name string `json:"name,omitempty"`

	ServiceMaps []MobilityGroupServiceMap `json:"serviceMaps,omitempty"`
}
