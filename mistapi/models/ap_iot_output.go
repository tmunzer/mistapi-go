package models

import (
    "encoding/json"
)

// ApIotOutput represents a ApIotOutput struct.
// IoT output AP settings
type ApIotOutput struct {
    // whether to enable a pin
    Enabled              *bool            `json:"enabled,omitempty"`
    // optional; descriptive pin name
    Name                 *string          `json:"name,omitempty"`
    // whether the pin is configured as an output. DO and A1-A4 can be repurposed by changing
    Output               *bool            `json:"output,omitempty"`
    // the type of pull-up the pin uses. enum: `external`, `internal`, `none`
    Pullup               *ApIotPullupEnum `json:"pullup,omitempty"`
    // output pin signal level, default 0
    Value                *int             `json:"value,omitempty"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApIotOutput.
// It customizes the JSON marshaling process for ApIotOutput objects.
func (a ApIotOutput) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApIotOutput object to a map representation for JSON marshaling.
func (a ApIotOutput) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
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
    var temp apIotOutput
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "name", "output", "pullup", "value")
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

// apIotOutput is a temporary struct used for validating the fields of ApIotOutput.
type apIotOutput  struct {
    Enabled *bool            `json:"enabled,omitempty"`
    Name    *string          `json:"name,omitempty"`
    Output  *bool            `json:"output,omitempty"`
    Pullup  *ApIotPullupEnum `json:"pullup,omitempty"`
    Value   *int             `json:"value,omitempty"`
}
