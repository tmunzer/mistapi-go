package models

import (
    "encoding/json"
    "fmt"
)

// VpnLinksSteering represents a VpnLinksSteering struct.
type VpnLinksSteering struct {
    Preference           *int                   `json:"preference,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VpnLinksSteering,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VpnLinksSteering) String() string {
    return fmt.Sprintf(
    	"VpnLinksSteering[Preference=%v, AdditionalProperties=%v]",
    	v.Preference, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VpnLinksSteering.
// It customizes the JSON marshaling process for VpnLinksSteering objects.
func (v VpnLinksSteering) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "preference"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VpnLinksSteering object to a map representation for JSON marshaling.
func (v VpnLinksSteering) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Preference != nil {
        structMap["preference"] = v.Preference
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VpnLinksSteering.
// It customizes the JSON unmarshaling process for VpnLinksSteering objects.
func (v *VpnLinksSteering) UnmarshalJSON(input []byte) error {
    var temp tempVpnLinksSteering
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "preference")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.Preference = temp.Preference
    return nil
}

// tempVpnLinksSteering is a temporary struct used for validating the fields of VpnLinksSteering.
type tempVpnLinksSteering  struct {
    Preference *int `json:"preference,omitempty"`
}
