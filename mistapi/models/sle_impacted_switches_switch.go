package models

import (
    "encoding/json"
)

// SleImpactedSwitchesSwitch represents a SleImpactedSwitchesSwitch struct.
type SleImpactedSwitchesSwitch struct {
    Degraded             *float64               `json:"degraded,omitempty"`
    Duration             *float64               `json:"duration,omitempty"`
    Interface            []string               `json:"interface,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    SwitchMac            *string                `json:"switch_mac,omitempty"`
    SwitchModel          *string                `json:"switch_model,omitempty"`
    SwitchVersion        *string                `json:"switch_version,omitempty"`
    Total                *float64               `json:"total,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedSwitchesSwitch.
// It customizes the JSON marshaling process for SleImpactedSwitchesSwitch objects.
func (s SleImpactedSwitchesSwitch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "degraded", "duration", "interface", "name", "switch_mac", "switch_model", "switch_version", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedSwitchesSwitch object to a map representation for JSON marshaling.
func (s SleImpactedSwitchesSwitch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Degraded != nil {
        structMap["degraded"] = s.Degraded
    }
    if s.Duration != nil {
        structMap["duration"] = s.Duration
    }
    if s.Interface != nil {
        structMap["interface"] = s.Interface
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.SwitchMac != nil {
        structMap["switch_mac"] = s.SwitchMac
    }
    if s.SwitchModel != nil {
        structMap["switch_model"] = s.SwitchModel
    }
    if s.SwitchVersion != nil {
        structMap["switch_version"] = s.SwitchVersion
    }
    if s.Total != nil {
        structMap["total"] = s.Total
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedSwitchesSwitch.
// It customizes the JSON unmarshaling process for SleImpactedSwitchesSwitch objects.
func (s *SleImpactedSwitchesSwitch) UnmarshalJSON(input []byte) error {
    var temp tempSleImpactedSwitchesSwitch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "degraded", "duration", "interface", "name", "switch_mac", "switch_model", "switch_version", "total")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Degraded = temp.Degraded
    s.Duration = temp.Duration
    s.Interface = temp.Interface
    s.Name = temp.Name
    s.SwitchMac = temp.SwitchMac
    s.SwitchModel = temp.SwitchModel
    s.SwitchVersion = temp.SwitchVersion
    s.Total = temp.Total
    return nil
}

// tempSleImpactedSwitchesSwitch is a temporary struct used for validating the fields of SleImpactedSwitchesSwitch.
type tempSleImpactedSwitchesSwitch  struct {
    Degraded      *float64 `json:"degraded,omitempty"`
    Duration      *float64 `json:"duration,omitempty"`
    Interface     []string `json:"interface,omitempty"`
    Name          *string  `json:"name,omitempty"`
    SwitchMac     *string  `json:"switch_mac,omitempty"`
    SwitchModel   *string  `json:"switch_model,omitempty"`
    SwitchVersion *string  `json:"switch_version,omitempty"`
    Total         *float64 `json:"total,omitempty"`
}
