// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// EvpnOptionsVsInstance represents a EvpnOptionsVsInstance struct.
type EvpnOptionsVsInstance struct {
    Networks             []string               `json:"networks,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EvpnOptionsVsInstance,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EvpnOptionsVsInstance) String() string {
    return fmt.Sprintf(
    	"EvpnOptionsVsInstance[Networks=%v, AdditionalProperties=%v]",
    	e.Networks, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EvpnOptionsVsInstance.
// It customizes the JSON marshaling process for EvpnOptionsVsInstance objects.
func (e EvpnOptionsVsInstance) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "networks"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnOptionsVsInstance object to a map representation for JSON marshaling.
func (e EvpnOptionsVsInstance) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Networks != nil {
        structMap["networks"] = e.Networks
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EvpnOptionsVsInstance.
// It customizes the JSON unmarshaling process for EvpnOptionsVsInstance objects.
func (e *EvpnOptionsVsInstance) UnmarshalJSON(input []byte) error {
    var temp tempEvpnOptionsVsInstance
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "networks")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.Networks = temp.Networks
    return nil
}

// tempEvpnOptionsVsInstance is a temporary struct used for validating the fields of EvpnOptionsVsInstance.
type tempEvpnOptionsVsInstance  struct {
    Networks []string `json:"networks,omitempty"`
}
