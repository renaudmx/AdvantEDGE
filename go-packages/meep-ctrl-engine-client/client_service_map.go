/*
 * MEEP Controller REST API
 *
 * Copyright (c) 2019 InterDigital Communications, Inc. All rights reserved. The information provided herein is the proprietary and confidential information of InterDigital Communications, Inc. 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Client-specific list of mappings of exposed port to internal service
type ClientServiceMap struct {

	// Unique external client identifier
	Client string `json:"client,omitempty"`

	ServiceMap []ServiceMap `json:"serviceMap,omitempty"`
}
