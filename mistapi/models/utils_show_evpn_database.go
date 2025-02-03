package models

import (
    "encoding/json"
    "fmt"
)

// UtilsShowEvpnDatabase represents a UtilsShowEvpnDatabase struct.
type UtilsShowEvpnDatabase struct {
    // Duration in sec for which refresh is enabled. Should be set only if interval is configured to non-zero value.
    Duration             *int                   `json:"duration,omitempty"`
    // Rate at which output will refresh
    Interval             *int                   `json:"interval,omitempty"`
    // Xlient mac filter
    Mac                  *string                `json:"mac,omitempty"`
    // Interface name
    PortId               *string                `json:"port_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsShowEvpnDatabase,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsShowEvpnDatabase) String() string {
    return fmt.Sprintf(
    	"UtilsShowEvpnDatabase[Duration=%v, Interval=%v, Mac=%v, PortId=%v, AdditionalProperties=%v]",
    	u.Duration, u.Interval, u.Mac, u.PortId, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowEvpnDatabase.
// It customizes the JSON marshaling process for UtilsShowEvpnDatabase objects.
func (u UtilsShowEvpnDatabase) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "duration", "interval", "mac", "port_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowEvpnDatabase object to a map representation for JSON marshaling.
func (u UtilsShowEvpnDatabase) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Duration != nil {
        structMap["duration"] = u.Duration
    }
    if u.Interval != nil {
        structMap["interval"] = u.Interval
    }
    if u.Mac != nil {
        structMap["mac"] = u.Mac
    }
    if u.PortId != nil {
        structMap["port_id"] = u.PortId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowEvpnDatabase.
// It customizes the JSON unmarshaling process for UtilsShowEvpnDatabase objects.
func (u *UtilsShowEvpnDatabase) UnmarshalJSON(input []byte) error {
    var temp tempUtilsShowEvpnDatabase
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "duration", "interval", "mac", "port_id")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Duration = temp.Duration
    u.Interval = temp.Interval
    u.Mac = temp.Mac
    u.PortId = temp.PortId
    return nil
}

// tempUtilsShowEvpnDatabase is a temporary struct used for validating the fields of UtilsShowEvpnDatabase.
type tempUtilsShowEvpnDatabase  struct {
    Duration *int    `json:"duration,omitempty"`
    Interval *int    `json:"interval,omitempty"`
    Mac      *string `json:"mac,omitempty"`
    PortId   *string `json:"port_id,omitempty"`
}
