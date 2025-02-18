package models

import (
    "encoding/json"
    "fmt"
)

// UtilsShowArp represents a UtilsShowArp struct.
type UtilsShowArp struct {
    // Duration in sec for which refresh is enabled. Should be set only if interval is configured to non-zero value.
    Duration             *int                   `json:"duration,omitempty"`
    // Rate at which output will refresh
    Interval             *int                   `json:"interval,omitempty"`
    // IP Address
    Ip                   *string                `json:"ip,omitempty"`
    // Device Port ID
    PortId               *string                `json:"port_id,omitempty"`
    // VRF Name
    Vrf                  *string                `json:"vrf,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsShowArp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsShowArp) String() string {
    return fmt.Sprintf(
    	"UtilsShowArp[Duration=%v, Interval=%v, Ip=%v, PortId=%v, Vrf=%v, AdditionalProperties=%v]",
    	u.Duration, u.Interval, u.Ip, u.PortId, u.Vrf, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowArp.
// It customizes the JSON marshaling process for UtilsShowArp objects.
func (u UtilsShowArp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "duration", "interval", "ip", "port_id", "vrf"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowArp object to a map representation for JSON marshaling.
func (u UtilsShowArp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Duration != nil {
        structMap["duration"] = u.Duration
    }
    if u.Interval != nil {
        structMap["interval"] = u.Interval
    }
    if u.Ip != nil {
        structMap["ip"] = u.Ip
    }
    if u.PortId != nil {
        structMap["port_id"] = u.PortId
    }
    if u.Vrf != nil {
        structMap["vrf"] = u.Vrf
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowArp.
// It customizes the JSON unmarshaling process for UtilsShowArp objects.
func (u *UtilsShowArp) UnmarshalJSON(input []byte) error {
    var temp tempUtilsShowArp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "duration", "interval", "ip", "port_id", "vrf")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Duration = temp.Duration
    u.Interval = temp.Interval
    u.Ip = temp.Ip
    u.PortId = temp.PortId
    u.Vrf = temp.Vrf
    return nil
}

// tempUtilsShowArp is a temporary struct used for validating the fields of UtilsShowArp.
type tempUtilsShowArp  struct {
    Duration *int    `json:"duration,omitempty"`
    Interval *int    `json:"interval,omitempty"`
    Ip       *string `json:"ip,omitempty"`
    PortId   *string `json:"port_id,omitempty"`
    Vrf      *string `json:"vrf,omitempty"`
}
