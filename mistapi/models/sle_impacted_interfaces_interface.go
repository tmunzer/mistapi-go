package models

import (
    "encoding/json"
    "fmt"
)

// SleImpactedInterfacesInterface represents a SleImpactedInterfacesInterface struct.
type SleImpactedInterfacesInterface struct {
    Degraded             *float64               `json:"degraded,omitempty"`
    Duration             *float64               `json:"duration,omitempty"`
    InterfaceName        *string                `json:"interface_name,omitempty"`
    SwitchMac            *string                `json:"switch_mac,omitempty"`
    SwitchName           *string                `json:"switch_name,omitempty"`
    Total                *float64               `json:"total,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactedInterfacesInterface,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactedInterfacesInterface) String() string {
    return fmt.Sprintf(
    	"SleImpactedInterfacesInterface[Degraded=%v, Duration=%v, InterfaceName=%v, SwitchMac=%v, SwitchName=%v, Total=%v, AdditionalProperties=%v]",
    	s.Degraded, s.Duration, s.InterfaceName, s.SwitchMac, s.SwitchName, s.Total, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedInterfacesInterface.
// It customizes the JSON marshaling process for SleImpactedInterfacesInterface objects.
func (s SleImpactedInterfacesInterface) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "degraded", "duration", "interface_name", "switch_mac", "switch_name", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedInterfacesInterface object to a map representation for JSON marshaling.
func (s SleImpactedInterfacesInterface) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Degraded != nil {
        structMap["degraded"] = s.Degraded
    }
    if s.Duration != nil {
        structMap["duration"] = s.Duration
    }
    if s.InterfaceName != nil {
        structMap["interface_name"] = s.InterfaceName
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

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedInterfacesInterface.
// It customizes the JSON unmarshaling process for SleImpactedInterfacesInterface objects.
func (s *SleImpactedInterfacesInterface) UnmarshalJSON(input []byte) error {
    var temp tempSleImpactedInterfacesInterface
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "degraded", "duration", "interface_name", "switch_mac", "switch_name", "total")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Degraded = temp.Degraded
    s.Duration = temp.Duration
    s.InterfaceName = temp.InterfaceName
    s.SwitchMac = temp.SwitchMac
    s.SwitchName = temp.SwitchName
    s.Total = temp.Total
    return nil
}

// tempSleImpactedInterfacesInterface is a temporary struct used for validating the fields of SleImpactedInterfacesInterface.
type tempSleImpactedInterfacesInterface  struct {
    Degraded      *float64 `json:"degraded,omitempty"`
    Duration      *float64 `json:"duration,omitempty"`
    InterfaceName *string  `json:"interface_name,omitempty"`
    SwitchMac     *string  `json:"switch_mac,omitempty"`
    SwitchName    *string  `json:"switch_name,omitempty"`
    Total         *float64 `json:"total,omitempty"`
}
