package models

import (
    "encoding/json"
    "fmt"
)

// VpnPathTrafficShaping represents a VpnPathTrafficShaping struct.
type VpnPathTrafficShaping struct {
    // percentages for different class of traffic: high / medium / low / best-effort adding up to 100
    ClassPercentage      []int                  `json:"class_percentage,omitempty"`
    Enabled              *bool                  `json:"enabled,omitempty"`
    MaxTxKbps            Optional[int]          `json:"max_tx_kbps"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VpnPathTrafficShaping,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VpnPathTrafficShaping) String() string {
    return fmt.Sprintf(
    	"VpnPathTrafficShaping[ClassPercentage=%v, Enabled=%v, MaxTxKbps=%v, AdditionalProperties=%v]",
    	v.ClassPercentage, v.Enabled, v.MaxTxKbps, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VpnPathTrafficShaping.
// It customizes the JSON marshaling process for VpnPathTrafficShaping objects.
func (v VpnPathTrafficShaping) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "class_percentage", "enabled", "max_tx_kbps"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VpnPathTrafficShaping object to a map representation for JSON marshaling.
func (v VpnPathTrafficShaping) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.ClassPercentage != nil {
        structMap["class_percentage"] = v.ClassPercentage
    }
    if v.Enabled != nil {
        structMap["enabled"] = v.Enabled
    }
    if v.MaxTxKbps.IsValueSet() {
        if v.MaxTxKbps.Value() != nil {
            structMap["max_tx_kbps"] = v.MaxTxKbps.Value()
        } else {
            structMap["max_tx_kbps"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VpnPathTrafficShaping.
// It customizes the JSON unmarshaling process for VpnPathTrafficShaping objects.
func (v *VpnPathTrafficShaping) UnmarshalJSON(input []byte) error {
    var temp tempVpnPathTrafficShaping
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "class_percentage", "enabled", "max_tx_kbps")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.ClassPercentage = temp.ClassPercentage
    v.Enabled = temp.Enabled
    v.MaxTxKbps = temp.MaxTxKbps
    return nil
}

// tempVpnPathTrafficShaping is a temporary struct used for validating the fields of VpnPathTrafficShaping.
type tempVpnPathTrafficShaping  struct {
    ClassPercentage []int         `json:"class_percentage,omitempty"`
    Enabled         *bool         `json:"enabled,omitempty"`
    MaxTxKbps       Optional[int] `json:"max_tx_kbps"`
}
