// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// MxclusterTuntermExtraRoute represents a MxclusterTuntermExtraRoute struct.
type MxclusterTuntermExtraRoute struct {
    Via                  *string                `json:"via,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxclusterTuntermExtraRoute,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxclusterTuntermExtraRoute) String() string {
    return fmt.Sprintf(
    	"MxclusterTuntermExtraRoute[Via=%v, AdditionalProperties=%v]",
    	m.Via, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxclusterTuntermExtraRoute.
// It customizes the JSON marshaling process for MxclusterTuntermExtraRoute objects.
func (m MxclusterTuntermExtraRoute) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "via"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxclusterTuntermExtraRoute object to a map representation for JSON marshaling.
func (m MxclusterTuntermExtraRoute) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Via != nil {
        structMap["via"] = m.Via
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxclusterTuntermExtraRoute.
// It customizes the JSON unmarshaling process for MxclusterTuntermExtraRoute objects.
func (m *MxclusterTuntermExtraRoute) UnmarshalJSON(input []byte) error {
    var temp tempMxclusterTuntermExtraRoute
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "via")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Via = temp.Via
    return nil
}

// tempMxclusterTuntermExtraRoute is a temporary struct used for validating the fields of MxclusterTuntermExtraRoute.
type tempMxclusterTuntermExtraRoute  struct {
    Via *string `json:"via,omitempty"`
}
