package models

import (
    "encoding/json"
    "fmt"
)

// ApAntennaMode represents a ApAntennaMode struct.
type ApAntennaMode struct {
    // Antenna Mode for AP which supports selectable antennas. enum: `external`, `internal`
    AntMode              *AntModeEnum           `json:"ant_mode,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApAntennaMode,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApAntennaMode) String() string {
    return fmt.Sprintf(
    	"ApAntennaMode[AntMode=%v, AdditionalProperties=%v]",
    	a.AntMode, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApAntennaMode.
// It customizes the JSON marshaling process for ApAntennaMode objects.
func (a ApAntennaMode) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "ant_mode"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApAntennaMode object to a map representation for JSON marshaling.
func (a ApAntennaMode) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.AntMode != nil {
        structMap["ant_mode"] = a.AntMode
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApAntennaMode.
// It customizes the JSON unmarshaling process for ApAntennaMode objects.
func (a *ApAntennaMode) UnmarshalJSON(input []byte) error {
    var temp tempApAntennaMode
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ant_mode")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.AntMode = temp.AntMode
    return nil
}

// tempApAntennaMode is a temporary struct used for validating the fields of ApAntennaMode.
type tempApAntennaMode  struct {
    AntMode *AntModeEnum `json:"ant_mode,omitempty"`
}
