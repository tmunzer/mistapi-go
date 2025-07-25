// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// CaptureMxedgeMxedges represents a CaptureMxedgeMxedges struct.
// Property key is the Mx Edge ID, currently limited to one mxedge per org capture session
type CaptureMxedgeMxedges struct {
    Interfaces           map[string]CaptureMxedgeMxedgesInterfaces `json:"interfaces,omitempty"`
    AdditionalProperties map[string]interface{}                    `json:"_"`
}

// String implements the fmt.Stringer interface for CaptureMxedgeMxedges,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CaptureMxedgeMxedges) String() string {
    return fmt.Sprintf(
    	"CaptureMxedgeMxedges[Interfaces=%v, AdditionalProperties=%v]",
    	c.Interfaces, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CaptureMxedgeMxedges.
// It customizes the JSON marshaling process for CaptureMxedgeMxedges objects.
func (c CaptureMxedgeMxedges) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "interfaces"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureMxedgeMxedges object to a map representation for JSON marshaling.
func (c CaptureMxedgeMxedges) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Interfaces != nil {
        structMap["interfaces"] = c.Interfaces
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureMxedgeMxedges.
// It customizes the JSON unmarshaling process for CaptureMxedgeMxedges objects.
func (c *CaptureMxedgeMxedges) UnmarshalJSON(input []byte) error {
    var temp tempCaptureMxedgeMxedges
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "interfaces")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Interfaces = temp.Interfaces
    return nil
}

// tempCaptureMxedgeMxedges is a temporary struct used for validating the fields of CaptureMxedgeMxedges.
type tempCaptureMxedgeMxedges  struct {
    Interfaces map[string]CaptureMxedgeMxedgesInterfaces `json:"interfaces,omitempty"`
}
