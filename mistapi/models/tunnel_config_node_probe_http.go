// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// TunnelConfigNodeProbeHttp represents a TunnelConfigNodeProbeHttp struct.
// HTTP probe settings for a custom IPsec tunnel node
type TunnelConfigNodeProbeHttp struct {
	// HTTP response status codes that indicate a successful probe. Defaults to 200 if not specified.
	AcceptedStatusCodes []int `json:"accepted_status_codes,omitempty"`
	// HTTP or HTTPS URLs to probe
	Urls                 []string               `json:"urls,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TunnelConfigNodeProbeHttp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TunnelConfigNodeProbeHttp) String() string {
	return fmt.Sprintf(
		"TunnelConfigNodeProbeHttp[AcceptedStatusCodes=%v, Urls=%v, AdditionalProperties=%v]",
		t.AcceptedStatusCodes, t.Urls, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TunnelConfigNodeProbeHttp.
// It customizes the JSON marshaling process for TunnelConfigNodeProbeHttp objects.
func (t TunnelConfigNodeProbeHttp) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(t.AdditionalProperties,
		"accepted_status_codes", "urls"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(t.toMap())
}

// toMap converts the TunnelConfigNodeProbeHttp object to a map representation for JSON marshaling.
func (t TunnelConfigNodeProbeHttp) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, t.AdditionalProperties)
	if t.AcceptedStatusCodes != nil {
		structMap["accepted_status_codes"] = t.AcceptedStatusCodes
	}
	if t.Urls != nil {
		structMap["urls"] = t.Urls
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelConfigNodeProbeHttp.
// It customizes the JSON unmarshaling process for TunnelConfigNodeProbeHttp objects.
func (t *TunnelConfigNodeProbeHttp) UnmarshalJSON(input []byte) error {
	var temp tempTunnelConfigNodeProbeHttp
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "accepted_status_codes", "urls")
	if err != nil {
		return err
	}
	t.AdditionalProperties = additionalProperties

	t.AcceptedStatusCodes = temp.AcceptedStatusCodes
	t.Urls = temp.Urls
	return nil
}

// tempTunnelConfigNodeProbeHttp is a temporary struct used for validating the fields of TunnelConfigNodeProbeHttp.
type tempTunnelConfigNodeProbeHttp struct {
	AcceptedStatusCodes []int    `json:"accepted_status_codes,omitempty"`
	Urls                []string `json:"urls,omitempty"`
}
