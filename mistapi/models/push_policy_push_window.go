// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// PushPolicyPushWindow represents a PushPolicyPushWindow struct.
// If enabled, new config will only be pushed to device within the specified time window
type PushPolicyPushWindow struct {
    Enabled              *bool                  `json:"enabled,omitempty"`
    // Days/Hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun)
    Hours                *Hours                 `json:"hours,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PushPolicyPushWindow,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PushPolicyPushWindow) String() string {
    return fmt.Sprintf(
    	"PushPolicyPushWindow[Enabled=%v, Hours=%v, AdditionalProperties=%v]",
    	p.Enabled, p.Hours, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PushPolicyPushWindow.
// It customizes the JSON marshaling process for PushPolicyPushWindow objects.
func (p PushPolicyPushWindow) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "enabled", "hours"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PushPolicyPushWindow object to a map representation for JSON marshaling.
func (p PushPolicyPushWindow) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Enabled != nil {
        structMap["enabled"] = p.Enabled
    }
    if p.Hours != nil {
        structMap["hours"] = p.Hours.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PushPolicyPushWindow.
// It customizes the JSON unmarshaling process for PushPolicyPushWindow objects.
func (p *PushPolicyPushWindow) UnmarshalJSON(input []byte) error {
    var temp tempPushPolicyPushWindow
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "hours")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.Enabled = temp.Enabled
    p.Hours = temp.Hours
    return nil
}

// tempPushPolicyPushWindow is a temporary struct used for validating the fields of PushPolicyPushWindow.
type tempPushPolicyPushWindow  struct {
    Enabled *bool  `json:"enabled,omitempty"`
    Hours   *Hours `json:"hours,omitempty"`
}
