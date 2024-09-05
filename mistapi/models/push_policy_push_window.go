package models

import (
    "encoding/json"
)

// PushPolicyPushWindow represents a PushPolicyPushWindow struct.
// if enabled, new config will only be pushed to device within the specified time window
type PushPolicyPushWindow struct {
    Enabled              *bool          `json:"enabled,omitempty"`
    // hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun).
    // **Note**: If the dow is not defined then it\u2019\ s treated as 00:00-23:59.
    Hours                *Hours         `json:"hours,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PushPolicyPushWindow.
// It customizes the JSON marshaling process for PushPolicyPushWindow objects.
func (p PushPolicyPushWindow) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PushPolicyPushWindow object to a map representation for JSON marshaling.
func (p PushPolicyPushWindow) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "hours")
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
