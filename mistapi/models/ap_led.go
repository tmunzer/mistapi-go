package models

import (
    "encoding/json"
    "fmt"
)

// ApLed represents a ApLed struct.
// LED AP settings
type ApLed struct {
    Brightness           *int                   `json:"brightness,omitempty"`
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApLed,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApLed) String() string {
    return fmt.Sprintf(
    	"ApLed[Brightness=%v, Enabled=%v, AdditionalProperties=%v]",
    	a.Brightness, a.Enabled, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApLed.
// It customizes the JSON marshaling process for ApLed objects.
func (a ApLed) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "brightness", "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApLed object to a map representation for JSON marshaling.
func (a ApLed) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Brightness != nil {
        structMap["brightness"] = a.Brightness
    }
    if a.Enabled != nil {
        structMap["enabled"] = a.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApLed.
// It customizes the JSON unmarshaling process for ApLed objects.
func (a *ApLed) UnmarshalJSON(input []byte) error {
    var temp tempApLed
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "brightness", "enabled")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Brightness = temp.Brightness
    a.Enabled = temp.Enabled
    return nil
}

// tempApLed is a temporary struct used for validating the fields of ApLed.
type tempApLed  struct {
    Brightness *int  `json:"brightness,omitempty"`
    Enabled    *bool `json:"enabled,omitempty"`
}
