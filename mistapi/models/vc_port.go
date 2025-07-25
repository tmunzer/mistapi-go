// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// VcPort represents a VcPort struct.
type VcPort struct {
    // enum: `network`, `vcp-higig`, `vcp-hgoe` (only for EX4400-24X)
    Mode                 *VcPortModeEnum        `json:"mode,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VcPort,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VcPort) String() string {
    return fmt.Sprintf(
    	"VcPort[Mode=%v, AdditionalProperties=%v]",
    	v.Mode, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VcPort.
// It customizes the JSON marshaling process for VcPort objects.
func (v VcPort) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "mode"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VcPort object to a map representation for JSON marshaling.
func (v VcPort) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Mode != nil {
        structMap["mode"] = v.Mode
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VcPort.
// It customizes the JSON unmarshaling process for VcPort objects.
func (v *VcPort) UnmarshalJSON(input []byte) error {
    var temp tempVcPort
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mode")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.Mode = temp.Mode
    return nil
}

// tempVcPort is a temporary struct used for validating the fields of VcPort.
type tempVcPort  struct {
    Mode *VcPortModeEnum `json:"mode,omitempty"`
}
