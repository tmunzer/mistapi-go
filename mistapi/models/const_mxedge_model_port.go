package models

import (
    "encoding/json"
    "fmt"
)

// ConstMxedgeModelPort represents a ConstMxedgeModelPort struct.
type ConstMxedgeModelPort struct {
    Display              *string                `json:"display,omitempty"`
    Speed                *int                   `json:"speed,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstMxedgeModelPort,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstMxedgeModelPort) String() string {
    return fmt.Sprintf(
    	"ConstMxedgeModelPort[Display=%v, Speed=%v, AdditionalProperties=%v]",
    	c.Display, c.Speed, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstMxedgeModelPort.
// It customizes the JSON marshaling process for ConstMxedgeModelPort objects.
func (c ConstMxedgeModelPort) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "display", "speed"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstMxedgeModelPort object to a map representation for JSON marshaling.
func (c ConstMxedgeModelPort) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Display != nil {
        structMap["display"] = c.Display
    }
    if c.Speed != nil {
        structMap["speed"] = c.Speed
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstMxedgeModelPort.
// It customizes the JSON unmarshaling process for ConstMxedgeModelPort objects.
func (c *ConstMxedgeModelPort) UnmarshalJSON(input []byte) error {
    var temp tempConstMxedgeModelPort
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "display", "speed")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Display = temp.Display
    c.Speed = temp.Speed
    return nil
}

// tempConstMxedgeModelPort is a temporary struct used for validating the fields of ConstMxedgeModelPort.
type tempConstMxedgeModelPort  struct {
    Display *string `json:"display,omitempty"`
    Speed   *int    `json:"speed,omitempty"`
}
