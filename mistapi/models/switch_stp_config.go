package models

import (
    "encoding/json"
    "fmt"
)

// SwitchStpConfig represents a SwitchStpConfig struct.
type SwitchStpConfig struct {
    // Switch STP priority: from `0k` to `15k`
    BridgePriority       *string                `json:"bridge_priority,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchStpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchStpConfig) String() string {
    return fmt.Sprintf(
    	"SwitchStpConfig[BridgePriority=%v, AdditionalProperties=%v]",
    	s.BridgePriority, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchStpConfig.
// It customizes the JSON marshaling process for SwitchStpConfig objects.
func (s SwitchStpConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "bridge_priority"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchStpConfig object to a map representation for JSON marshaling.
func (s SwitchStpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.BridgePriority != nil {
        structMap["bridge_priority"] = s.BridgePriority
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchStpConfig.
// It customizes the JSON unmarshaling process for SwitchStpConfig objects.
func (s *SwitchStpConfig) UnmarshalJSON(input []byte) error {
    var temp tempSwitchStpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bridge_priority")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.BridgePriority = temp.BridgePriority
    return nil
}

// tempSwitchStpConfig is a temporary struct used for validating the fields of SwitchStpConfig.
type tempSwitchStpConfig  struct {
    BridgePriority *string `json:"bridge_priority,omitempty"`
}
