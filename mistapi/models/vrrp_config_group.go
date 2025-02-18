package models

import (
    "encoding/json"
    "fmt"
)

// VrrpConfigGroup represents a VrrpConfigGroup struct.
type VrrpConfigGroup struct {
    Priority             *int                   `json:"priority,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VrrpConfigGroup,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VrrpConfigGroup) String() string {
    return fmt.Sprintf(
    	"VrrpConfigGroup[Priority=%v, AdditionalProperties=%v]",
    	v.Priority, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VrrpConfigGroup.
// It customizes the JSON marshaling process for VrrpConfigGroup objects.
func (v VrrpConfigGroup) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "priority"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VrrpConfigGroup object to a map representation for JSON marshaling.
func (v VrrpConfigGroup) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Priority != nil {
        structMap["priority"] = v.Priority
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VrrpConfigGroup.
// It customizes the JSON unmarshaling process for VrrpConfigGroup objects.
func (v *VrrpConfigGroup) UnmarshalJSON(input []byte) error {
    var temp tempVrrpConfigGroup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "priority")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.Priority = temp.Priority
    return nil
}

// tempVrrpConfigGroup is a temporary struct used for validating the fields of VrrpConfigGroup.
type tempVrrpConfigGroup  struct {
    Priority *int `json:"priority,omitempty"`
}
