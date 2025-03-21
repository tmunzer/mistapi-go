package models

import (
    "encoding/json"
    "fmt"
)

// SynthetictestConfig represents a SynthetictestConfig struct.
type SynthetictestConfig struct {
    Disabled             *bool                            `json:"disabled,omitempty"`
    Vlans                []SynthetictestProperties        `json:"vlans,omitempty"`
    WanSpeedtest         *SynthetictestConfigWanSpeedtest `json:"wan_speedtest,omitempty"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for SynthetictestConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SynthetictestConfig) String() string {
    return fmt.Sprintf(
    	"SynthetictestConfig[Disabled=%v, Vlans=%v, WanSpeedtest=%v, AdditionalProperties=%v]",
    	s.Disabled, s.Vlans, s.WanSpeedtest, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestConfig.
// It customizes the JSON marshaling process for SynthetictestConfig objects.
func (s SynthetictestConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "disabled", "vlans", "wan_speedtest"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SynthetictestConfig object to a map representation for JSON marshaling.
func (s SynthetictestConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Disabled != nil {
        structMap["disabled"] = s.Disabled
    }
    if s.Vlans != nil {
        structMap["vlans"] = s.Vlans
    }
    if s.WanSpeedtest != nil {
        structMap["wan_speedtest"] = s.WanSpeedtest.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SynthetictestConfig.
// It customizes the JSON unmarshaling process for SynthetictestConfig objects.
func (s *SynthetictestConfig) UnmarshalJSON(input []byte) error {
    var temp tempSynthetictestConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "disabled", "vlans", "wan_speedtest")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Disabled = temp.Disabled
    s.Vlans = temp.Vlans
    s.WanSpeedtest = temp.WanSpeedtest
    return nil
}

// tempSynthetictestConfig is a temporary struct used for validating the fields of SynthetictestConfig.
type tempSynthetictestConfig  struct {
    Disabled     *bool                            `json:"disabled,omitempty"`
    Vlans        []SynthetictestProperties        `json:"vlans,omitempty"`
    WanSpeedtest *SynthetictestConfigWanSpeedtest `json:"wan_speedtest,omitempty"`
}
