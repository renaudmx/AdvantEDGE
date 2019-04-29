/*
 * MEEP TC controller API
 *
 * Copyright (c) 2019 InterDigital Communications, Inc. All rights reserved. The information provided herein is the proprietary and confidential information of InterDigital Communications, Inc. 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Service object
type ServiceConfig struct {

	// Unique service name
	Name string `json:"name,omitempty"`

	// Multi-Edge service name, if any
	MeSvcName string `json:"meSvcName,omitempty"`

	Ports []ServicePort `json:"ports,omitempty"`
}
