// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ApIotInput represents a ApIotInput struct.
// IoT Input AP settings
type ApIotInput struct {
    // Whether to enable a pin
    Enabled              *bool                  `json:"enabled,omitempty"`
    // Optional; descriptive pin name
    Name                 *string                `json:"name,omitempty"`
    // the type of pull-up the pin uses. enum: `external`, `internal`, `none`
    Pullup               *ApIotPullupEnum       `json:"pullup,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApIotInput,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApIotInput) String() string {
    return fmt.Sprintf(
    	"ApIotInput[Enabled=%v, Name=%v, Pullup=%v, AdditionalProperties=%v]",
    	a.Enabled, a.Name, a.Pullup, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApIotInput.
// It customizes the JSON marshaling process for ApIotInput objects.
func (a ApIotInput) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "enabled", "name", "pullup"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApIotInput object to a map representation for JSON marshaling.
func (a ApIotInput) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Enabled != nil {
        structMap["enabled"] = a.Enabled
    }
    if a.Name != nil {
        structMap["name"] = a.Name
    }
    if a.Pullup != nil {
        structMap["pullup"] = a.Pullup
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApIotInput.
// It customizes the JSON unmarshaling process for ApIotInput objects.
func (a *ApIotInput) UnmarshalJSON(input []byte) error {
    var temp tempApIotInput
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "name", "pullup")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Enabled = temp.Enabled
    a.Name = temp.Name
    a.Pullup = temp.Pullup
    return nil
}

// tempApIotInput is a temporary struct used for validating the fields of ApIotInput.
type tempApIotInput  struct {
    Enabled *bool            `json:"enabled,omitempty"`
    Name    *string          `json:"name,omitempty"`
    Pullup  *ApIotPullupEnum `json:"pullup,omitempty"`
}
