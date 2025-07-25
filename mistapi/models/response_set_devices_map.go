// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ResponseSetDevicesMap represents a ResponseSetDevicesMap struct.
type ResponseSetDevicesMap struct {
    Locked               []string               `json:"locked,omitempty"`
    Moved                []string               `json:"moved,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSetDevicesMap,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSetDevicesMap) String() string {
    return fmt.Sprintf(
    	"ResponseSetDevicesMap[Locked=%v, Moved=%v, AdditionalProperties=%v]",
    	r.Locked, r.Moved, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSetDevicesMap.
// It customizes the JSON marshaling process for ResponseSetDevicesMap objects.
func (r ResponseSetDevicesMap) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "locked", "moved"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSetDevicesMap object to a map representation for JSON marshaling.
func (r ResponseSetDevicesMap) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Locked != nil {
        structMap["locked"] = r.Locked
    }
    if r.Moved != nil {
        structMap["moved"] = r.Moved
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSetDevicesMap.
// It customizes the JSON unmarshaling process for ResponseSetDevicesMap objects.
func (r *ResponseSetDevicesMap) UnmarshalJSON(input []byte) error {
    var temp tempResponseSetDevicesMap
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "locked", "moved")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Locked = temp.Locked
    r.Moved = temp.Moved
    return nil
}

// tempResponseSetDevicesMap is a temporary struct used for validating the fields of ResponseSetDevicesMap.
type tempResponseSetDevicesMap  struct {
    Locked []string `json:"locked,omitempty"`
    Moved  []string `json:"moved,omitempty"`
}
