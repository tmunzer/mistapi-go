// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GatewayMgmtHostOutPolicySyslogServer represents a GatewayMgmtHostOutPolicySyslogServer struct.
// Allows to define the host_out_policy per Syslog Server. The Property key is the Syslog name
type GatewayMgmtHostOutPolicySyslogServer struct {
	Host                 *string                `json:"host,omitempty"`
	PathPreference       *string                `json:"path_preference,omitempty"`
	ServerName           *string                `json:"server_name,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayMgmtHostOutPolicySyslogServer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayMgmtHostOutPolicySyslogServer) String() string {
	return fmt.Sprintf(
		"GatewayMgmtHostOutPolicySyslogServer[Host=%v, PathPreference=%v, ServerName=%v, AdditionalProperties=%v]",
		g.Host, g.PathPreference, g.ServerName, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayMgmtHostOutPolicySyslogServer.
// It customizes the JSON marshaling process for GatewayMgmtHostOutPolicySyslogServer objects.
func (g GatewayMgmtHostOutPolicySyslogServer) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"host", "path_preference", "server_name"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GatewayMgmtHostOutPolicySyslogServer object to a map representation for JSON marshaling.
func (g GatewayMgmtHostOutPolicySyslogServer) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.Host != nil {
		structMap["host"] = g.Host
	}
	if g.PathPreference != nil {
		structMap["path_preference"] = g.PathPreference
	}
	if g.ServerName != nil {
		structMap["server_name"] = g.ServerName
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayMgmtHostOutPolicySyslogServer.
// It customizes the JSON unmarshaling process for GatewayMgmtHostOutPolicySyslogServer objects.
func (g *GatewayMgmtHostOutPolicySyslogServer) UnmarshalJSON(input []byte) error {
	var temp tempGatewayMgmtHostOutPolicySyslogServer
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "host", "path_preference", "server_name")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.Host = temp.Host
	g.PathPreference = temp.PathPreference
	g.ServerName = temp.ServerName
	return nil
}

// tempGatewayMgmtHostOutPolicySyslogServer is a temporary struct used for validating the fields of GatewayMgmtHostOutPolicySyslogServer.
type tempGatewayMgmtHostOutPolicySyslogServer struct {
	Host           *string `json:"host,omitempty"`
	PathPreference *string `json:"path_preference,omitempty"`
	ServerName     *string `json:"server_name,omitempty"`
}
