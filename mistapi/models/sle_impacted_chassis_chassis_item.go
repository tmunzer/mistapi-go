package models

import (
    "encoding/json"
)

// SleImpactedChassisChassisItem represents a SleImpactedChassisChassisItem struct.
type SleImpactedChassisChassisItem struct {
    Chassis              *string                `json:"chassis,omitempty"`
    Degraded             *float64               `json:"degraded,omitempty"`
    Duration             *float64               `json:"duration,omitempty"`
    Role                 *string                `json:"role,omitempty"`
    SwitchMac            *string                `json:"switch_mac,omitempty"`
    SwitchName           *string                `json:"switch_name,omitempty"`
    Total                *float64               `json:"total,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedChassisChassisItem.
// It customizes the JSON marshaling process for SleImpactedChassisChassisItem objects.
func (s SleImpactedChassisChassisItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "chassis", "degraded", "duration", "role", "switch_mac", "switch_name", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedChassisChassisItem object to a map representation for JSON marshaling.
func (s SleImpactedChassisChassisItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Chassis != nil {
        structMap["chassis"] = s.Chassis
    }
    if s.Degraded != nil {
        structMap["degraded"] = s.Degraded
    }
    if s.Duration != nil {
        structMap["duration"] = s.Duration
    }
    if s.Role != nil {
        structMap["role"] = s.Role
    }
    if s.SwitchMac != nil {
        structMap["switch_mac"] = s.SwitchMac
    }
    if s.SwitchName != nil {
        structMap["switch_name"] = s.SwitchName
    }
    if s.Total != nil {
        structMap["total"] = s.Total
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedChassisChassisItem.
// It customizes the JSON unmarshaling process for SleImpactedChassisChassisItem objects.
func (s *SleImpactedChassisChassisItem) UnmarshalJSON(input []byte) error {
    var temp tempSleImpactedChassisChassisItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "chassis", "degraded", "duration", "role", "switch_mac", "switch_name", "total")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Chassis = temp.Chassis
    s.Degraded = temp.Degraded
    s.Duration = temp.Duration
    s.Role = temp.Role
    s.SwitchMac = temp.SwitchMac
    s.SwitchName = temp.SwitchName
    s.Total = temp.Total
    return nil
}

// tempSleImpactedChassisChassisItem is a temporary struct used for validating the fields of SleImpactedChassisChassisItem.
type tempSleImpactedChassisChassisItem  struct {
    Chassis    *string  `json:"chassis,omitempty"`
    Degraded   *float64 `json:"degraded,omitempty"`
    Duration   *float64 `json:"duration,omitempty"`
    Role       *string  `json:"role,omitempty"`
    SwitchMac  *string  `json:"switch_mac,omitempty"`
    SwitchName *string  `json:"switch_name,omitempty"`
    Total      *float64 `json:"total,omitempty"`
}
