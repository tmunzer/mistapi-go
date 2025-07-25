// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// VrfExtraRoute represents a VrfExtraRoute struct.
type VrfExtraRoute struct {
    // Next-hop address
    Via                  *string                `json:"via,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VrfExtraRoute,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VrfExtraRoute) String() string {
    return fmt.Sprintf(
    	"VrfExtraRoute[Via=%v, AdditionalProperties=%v]",
    	v.Via, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VrfExtraRoute.
// It customizes the JSON marshaling process for VrfExtraRoute objects.
func (v VrfExtraRoute) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "via"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VrfExtraRoute object to a map representation for JSON marshaling.
func (v VrfExtraRoute) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Via != nil {
        structMap["via"] = v.Via
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VrfExtraRoute.
// It customizes the JSON unmarshaling process for VrfExtraRoute objects.
func (v *VrfExtraRoute) UnmarshalJSON(input []byte) error {
    var temp tempVrfExtraRoute
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

// tempVrfExtraRoute is a temporary struct used for validating the fields of VrfExtraRoute.
type tempVrfExtraRoute  struct {
    Via *string `json:"via,omitempty"`
}
