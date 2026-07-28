// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GatewayWanProbeOverrideHttp represents a GatewayWanProbeOverrideHttp struct.
// HTTP probe settings for a WAN probe override
type GatewayWanProbeOverrideHttp struct {
	// HTTP response status codes that indicate a successful probe. Defaults to 200 if not specified.
	AcceptedStatusCodes []int `json:"accepted_status_codes,omitempty"`
	// HTTP or HTTPS URLs to probe
	Urls                 []string               `json:"urls,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayWanProbeOverrideHttp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayWanProbeOverrideHttp) String() string {
	return fmt.Sprintf(
		"GatewayWanProbeOverrideHttp[AcceptedStatusCodes=%v, Urls=%v, AdditionalProperties=%v]",
		g.AcceptedStatusCodes, g.Urls, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayWanProbeOverrideHttp.
// It customizes the JSON marshaling process for GatewayWanProbeOverrideHttp objects.
func (g GatewayWanProbeOverrideHttp) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"accepted_status_codes", "urls"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GatewayWanProbeOverrideHttp object to a map representation for JSON marshaling.
func (g GatewayWanProbeOverrideHttp) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.AcceptedStatusCodes != nil {
		structMap["accepted_status_codes"] = g.AcceptedStatusCodes
	}
	if g.Urls != nil {
		structMap["urls"] = g.Urls
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayWanProbeOverrideHttp.
// It customizes the JSON unmarshaling process for GatewayWanProbeOverrideHttp objects.
func (g *GatewayWanProbeOverrideHttp) UnmarshalJSON(input []byte) error {
	var temp tempGatewayWanProbeOverrideHttp
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "accepted_status_codes", "urls")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.AcceptedStatusCodes = temp.AcceptedStatusCodes
	g.Urls = temp.Urls
	return nil
}

// tempGatewayWanProbeOverrideHttp is a temporary struct used for validating the fields of GatewayWanProbeOverrideHttp.
type tempGatewayWanProbeOverrideHttp struct {
	AcceptedStatusCodes []int    `json:"accepted_status_codes,omitempty"`
	Urls                []string `json:"urls,omitempty"`
}
