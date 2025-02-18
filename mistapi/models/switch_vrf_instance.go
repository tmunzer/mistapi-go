package models

import (
    "encoding/json"
    "fmt"
)

// SwitchVrfInstance represents a SwitchVrfInstance struct.
type SwitchVrfInstance struct {
    // Property key is the destination CIDR (e.g. "10.0.0.0/8")
    ExtraRoutes          map[string]VrfExtraRoute `json:"extra_routes,omitempty"`
    Networks             []string                 `json:"networks,omitempty"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchVrfInstance,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchVrfInstance) String() string {
    return fmt.Sprintf(
    	"SwitchVrfInstance[ExtraRoutes=%v, Networks=%v, AdditionalProperties=%v]",
    	s.ExtraRoutes, s.Networks, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchVrfInstance.
// It customizes the JSON marshaling process for SwitchVrfInstance objects.
func (s SwitchVrfInstance) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "extra_routes", "networks"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchVrfInstance object to a map representation for JSON marshaling.
func (s SwitchVrfInstance) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ExtraRoutes != nil {
        structMap["extra_routes"] = s.ExtraRoutes
    }
    if s.Networks != nil {
        structMap["networks"] = s.Networks
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchVrfInstance.
// It customizes the JSON unmarshaling process for SwitchVrfInstance objects.
func (s *SwitchVrfInstance) UnmarshalJSON(input []byte) error {
    var temp tempSwitchVrfInstance
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "extra_routes", "networks")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ExtraRoutes = temp.ExtraRoutes
    s.Networks = temp.Networks
    return nil
}

// tempSwitchVrfInstance is a temporary struct used for validating the fields of SwitchVrfInstance.
type tempSwitchVrfInstance  struct {
    ExtraRoutes map[string]VrfExtraRoute `json:"extra_routes,omitempty"`
    Networks    []string                 `json:"networks,omitempty"`
}
