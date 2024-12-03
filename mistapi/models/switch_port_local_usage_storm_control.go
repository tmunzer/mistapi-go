package models

import (
    "encoding/json"
)

// SwitchPortLocalUsageStormControl represents a SwitchPortLocalUsageStormControl struct.
// Switch storm control
type SwitchPortLocalUsageStormControl struct {
    // whether to disable storm control on broadcast traffic
    NoBroadcast           *bool                  `json:"no_broadcast,omitempty"`
    // whether to disable storm control on multicast traffic
    NoMulticast           *bool                  `json:"no_multicast,omitempty"`
    // whether to disable storm control on registered multicast traffic
    NoRegisteredMulticast *bool                  `json:"no_registered_multicast,omitempty"`
    // whether to disable storm control on unknown unicast traffic
    NoUnknownUnicast      *bool                  `json:"no_unknown_unicast,omitempty"`
    // bandwidth-percentage, configures the storm control level as a percentage of the available bandwidth
    Percentage            *int                   `json:"percentage,omitempty"`
    AdditionalProperties  map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchPortLocalUsageStormControl.
// It customizes the JSON marshaling process for SwitchPortLocalUsageStormControl objects.
func (s SwitchPortLocalUsageStormControl) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "no_broadcast", "no_multicast", "no_registered_multicast", "no_unknown_unicast", "percentage"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchPortLocalUsageStormControl object to a map representation for JSON marshaling.
func (s SwitchPortLocalUsageStormControl) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.NoBroadcast != nil {
        structMap["no_broadcast"] = s.NoBroadcast
    }
    if s.NoMulticast != nil {
        structMap["no_multicast"] = s.NoMulticast
    }
    if s.NoRegisteredMulticast != nil {
        structMap["no_registered_multicast"] = s.NoRegisteredMulticast
    }
    if s.NoUnknownUnicast != nil {
        structMap["no_unknown_unicast"] = s.NoUnknownUnicast
    }
    if s.Percentage != nil {
        structMap["percentage"] = s.Percentage
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchPortLocalUsageStormControl.
// It customizes the JSON unmarshaling process for SwitchPortLocalUsageStormControl objects.
func (s *SwitchPortLocalUsageStormControl) UnmarshalJSON(input []byte) error {
    var temp tempSwitchPortLocalUsageStormControl
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "no_broadcast", "no_multicast", "no_registered_multicast", "no_unknown_unicast", "percentage")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.NoBroadcast = temp.NoBroadcast
    s.NoMulticast = temp.NoMulticast
    s.NoRegisteredMulticast = temp.NoRegisteredMulticast
    s.NoUnknownUnicast = temp.NoUnknownUnicast
    s.Percentage = temp.Percentage
    return nil
}

// tempSwitchPortLocalUsageStormControl is a temporary struct used for validating the fields of SwitchPortLocalUsageStormControl.
type tempSwitchPortLocalUsageStormControl  struct {
    NoBroadcast           *bool `json:"no_broadcast,omitempty"`
    NoMulticast           *bool `json:"no_multicast,omitempty"`
    NoRegisteredMulticast *bool `json:"no_registered_multicast,omitempty"`
    NoUnknownUnicast      *bool `json:"no_unknown_unicast,omitempty"`
    Percentage            *int  `json:"percentage,omitempty"`
}
