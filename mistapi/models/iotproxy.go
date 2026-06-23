// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// Iotproxy represents a Iotproxy struct.
// IoT proxy configuration for the site
type Iotproxy struct {
	// Whether the site IoT proxy is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Visionline integration settings for IoT proxy
	Visionline           *IotproxyVisionline    `json:"visionline,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Iotproxy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i Iotproxy) String() string {
	return fmt.Sprintf(
		"Iotproxy[Enabled=%v, Visionline=%v, AdditionalProperties=%v]",
		i.Enabled, i.Visionline, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Iotproxy.
// It customizes the JSON marshaling process for Iotproxy objects.
func (i Iotproxy) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(i.AdditionalProperties,
		"enabled", "visionline"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(i.toMap())
}

// toMap converts the Iotproxy object to a map representation for JSON marshaling.
func (i Iotproxy) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, i.AdditionalProperties)
	if i.Enabled != nil {
		structMap["enabled"] = i.Enabled
	}
	if i.Visionline != nil {
		structMap["visionline"] = i.Visionline.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Iotproxy.
// It customizes the JSON unmarshaling process for Iotproxy objects.
func (i *Iotproxy) UnmarshalJSON(input []byte) error {
	var temp tempIotproxy
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "visionline")
	if err != nil {
		return err
	}
	i.AdditionalProperties = additionalProperties

	i.Enabled = temp.Enabled
	i.Visionline = temp.Visionline
	return nil
}

// tempIotproxy is a temporary struct used for validating the fields of Iotproxy.
type tempIotproxy struct {
	Enabled    *bool               `json:"enabled,omitempty"`
	Visionline *IotproxyVisionline `json:"visionline,omitempty"`
}
