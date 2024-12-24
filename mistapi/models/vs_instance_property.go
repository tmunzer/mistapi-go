package models

import (
    "encoding/json"
    "fmt"
)

// VsInstanceProperty represents a VsInstanceProperty struct.
type VsInstanceProperty struct {
    Networks             []string               `json:"networks,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VsInstanceProperty,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VsInstanceProperty) String() string {
    return fmt.Sprintf(
    	"VsInstanceProperty[Networks=%v, AdditionalProperties=%v]",
    	v.Networks, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VsInstanceProperty.
// It customizes the JSON marshaling process for VsInstanceProperty objects.
func (v VsInstanceProperty) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "networks"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VsInstanceProperty object to a map representation for JSON marshaling.
func (v VsInstanceProperty) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Networks != nil {
        structMap["networks"] = v.Networks
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VsInstanceProperty.
// It customizes the JSON unmarshaling process for VsInstanceProperty objects.
func (v *VsInstanceProperty) UnmarshalJSON(input []byte) error {
    var temp tempVsInstanceProperty
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "networks")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.Networks = temp.Networks
    return nil
}

// tempVsInstanceProperty is a temporary struct used for validating the fields of VsInstanceProperty.
type tempVsInstanceProperty  struct {
    Networks []string `json:"networks,omitempty"`
}
