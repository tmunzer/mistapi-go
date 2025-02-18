package models

import (
    "encoding/json"
    "fmt"
)

// ApIotOutput represents a ApIotOutput struct.
// IoT output AP settings
type ApIotOutput struct {
    // Whether to enable a pin
    Enabled              *bool                  `json:"enabled,omitempty"`
    // Optional; descriptive pin name
    Name                 *string                `json:"name,omitempty"`
    // Whether the pin is configured as an output. DO and A1-A4 can be repurposed by changing
    Output               *bool                  `json:"output,omitempty"`
    // the type of pull-up the pin uses. enum: `external`, `internal`, `none`
    Pullup               *ApIotPullupEnum       `json:"pullup,omitempty"`
    // Output pin signal level, default 0
    Value                *int                   `json:"value,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApIotOutput,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApIotOutput) String() string {
    return fmt.Sprintf(
    	"ApIotOutput[Enabled=%v, Name=%v, Output=%v, Pullup=%v, Value=%v, AdditionalProperties=%v]",
    	a.Enabled, a.Name, a.Output, a.Pullup, a.Value, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApIotOutput.
// It customizes the JSON marshaling process for ApIotOutput objects.
func (a ApIotOutput) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "enabled", "name", "output", "pullup", "value"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApIotOutput object to a map representation for JSON marshaling.
func (a ApIotOutput) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Enabled != nil {
        structMap["enabled"] = a.Enabled
    }
    if a.Name != nil {
        structMap["name"] = a.Name
    }
    if a.Output != nil {
        structMap["output"] = a.Output
    }
    if a.Pullup != nil {
        structMap["pullup"] = a.Pullup
    }
    if a.Value != nil {
        structMap["value"] = a.Value
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApIotOutput.
// It customizes the JSON unmarshaling process for ApIotOutput objects.
func (a *ApIotOutput) UnmarshalJSON(input []byte) error {
    var temp tempApIotOutput
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "name", "output", "pullup", "value")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Enabled = temp.Enabled
    a.Name = temp.Name
    a.Output = temp.Output
    a.Pullup = temp.Pullup
    a.Value = temp.Value
    return nil
}

// tempApIotOutput is a temporary struct used for validating the fields of ApIotOutput.
type tempApIotOutput  struct {
    Enabled *bool            `json:"enabled,omitempty"`
    Name    *string          `json:"name,omitempty"`
    Output  *bool            `json:"output,omitempty"`
    Pullup  *ApIotPullupEnum `json:"pullup,omitempty"`
    Value   *int             `json:"value,omitempty"`
}
