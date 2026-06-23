// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// VrfExtraRoute6 represents a VrfExtraRoute6 struct.
// Additional IPv6 static route for a VRF instance
type VrfExtraRoute6 struct {
	// IPv6 next-hop address for this VRF extra route
	Via                  *string                `json:"via,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VrfExtraRoute6,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VrfExtraRoute6) String() string {
	return fmt.Sprintf(
		"VrfExtraRoute6[Via=%v, AdditionalProperties=%v]",
		v.Via, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VrfExtraRoute6.
// It customizes the JSON marshaling process for VrfExtraRoute6 objects.
func (v VrfExtraRoute6) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(v.AdditionalProperties,
		"via"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(v.toMap())
}

// toMap converts the VrfExtraRoute6 object to a map representation for JSON marshaling.
func (v VrfExtraRoute6) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, v.AdditionalProperties)
	if v.Via != nil {
		structMap["via"] = v.Via
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VrfExtraRoute6.
// It customizes the JSON unmarshaling process for VrfExtraRoute6 objects.
func (v *VrfExtraRoute6) UnmarshalJSON(input []byte) error {
	var temp tempVrfExtraRoute6
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "via")
	if err != nil {
		return err
	}
	v.AdditionalProperties = additionalProperties

	v.Via = temp.Via
	return nil
}

// tempVrfExtraRoute6 is a temporary struct used for validating the fields of VrfExtraRoute6.
type tempVrfExtraRoute6 struct {
	Via *string `json:"via,omitempty"`
}
