package models

import (
    "encoding/json"
    "fmt"
)

// StatsSwitchApRedundancyModule represents a StatsSwitchApRedundancyModule struct.
type StatsSwitchApRedundancyModule struct {
    NumAps                     *int                   `json:"num_aps,omitempty"`
    NumApsWithSwitchRedundancy *int                   `json:"num_aps_with_switch_redundancy,omitempty"`
    AdditionalProperties       map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsSwitchApRedundancyModule,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsSwitchApRedundancyModule) String() string {
    return fmt.Sprintf(
    	"StatsSwitchApRedundancyModule[NumAps=%v, NumApsWithSwitchRedundancy=%v, AdditionalProperties=%v]",
    	s.NumAps, s.NumApsWithSwitchRedundancy, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsSwitchApRedundancyModule.
// It customizes the JSON marshaling process for StatsSwitchApRedundancyModule objects.
func (s StatsSwitchApRedundancyModule) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "num_aps", "num_aps_with_switch_redundancy"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsSwitchApRedundancyModule object to a map representation for JSON marshaling.
func (s StatsSwitchApRedundancyModule) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.NumAps != nil {
        structMap["num_aps"] = s.NumAps
    }
    if s.NumApsWithSwitchRedundancy != nil {
        structMap["num_aps_with_switch_redundancy"] = s.NumApsWithSwitchRedundancy
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsSwitchApRedundancyModule.
// It customizes the JSON unmarshaling process for StatsSwitchApRedundancyModule objects.
func (s *StatsSwitchApRedundancyModule) UnmarshalJSON(input []byte) error {
    var temp tempStatsSwitchApRedundancyModule
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "num_aps", "num_aps_with_switch_redundancy")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.NumAps = temp.NumAps
    s.NumApsWithSwitchRedundancy = temp.NumApsWithSwitchRedundancy
    return nil
}

// tempStatsSwitchApRedundancyModule is a temporary struct used for validating the fields of StatsSwitchApRedundancyModule.
type tempStatsSwitchApRedundancyModule  struct {
    NumAps                     *int `json:"num_aps,omitempty"`
    NumApsWithSwitchRedundancy *int `json:"num_aps_with_switch_redundancy,omitempty"`
}
