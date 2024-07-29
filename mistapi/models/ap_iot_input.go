package models

import (
    "encoding/json"
)

// ApIotInput represents a ApIotInput struct.
// IoT Input AP settings
type ApIotInput struct {
    // whether to enable a pin
    Enabled              *bool            `json:"enabled,omitempty"`
    // optional; descriptive pin name
    Name                 *string          `json:"name,omitempty"`
    // the type of pull-up the pin uses. enum: `external`, `internal`, `none`
    Pullup               *ApIotPullupEnum `json:"pullup,omitempty"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApIotInput.
// It customizes the JSON marshaling process for ApIotInput objects.
func (a ApIotInput) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApIotInput object to a map representation for JSON marshaling.
func (a ApIotInput) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Enabled != nil {
        structMap["enabled"] = a.Enabled
    }
    if a.Name != nil {
        structMap["name"] = a.Name
    }
    if a.Pullup != nil {
        structMap["pullup"] = a.Pullup
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApIotInput.
// It customizes the JSON unmarshaling process for ApIotInput objects.
func (a *ApIotInput) UnmarshalJSON(input []byte) error {
    var temp apIotInput
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "name", "pullup")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Enabled = temp.Enabled
    a.Name = temp.Name
    a.Pullup = temp.Pullup
    return nil
}

// apIotInput is a temporary struct used for validating the fields of ApIotInput.
type apIotInput  struct {
    Enabled *bool            `json:"enabled,omitempty"`
    Name    *string          `json:"name,omitempty"`
    Pullup  *ApIotPullupEnum `json:"pullup,omitempty"`
}
