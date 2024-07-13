package models

import (
    "encoding/json"
)

// ApStatsAutoUpgrade represents a ApStatsAutoUpgrade struct.
type ApStatsAutoUpgrade struct {
    Lastcheck            Optional[int64] `json:"lastcheck"`
    AdditionalProperties map[string]any  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsAutoUpgrade.
// It customizes the JSON marshaling process for ApStatsAutoUpgrade objects.
func (a ApStatsAutoUpgrade) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsAutoUpgrade object to a map representation for JSON marshaling.
func (a ApStatsAutoUpgrade) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Lastcheck.IsValueSet() {
        if a.Lastcheck.Value() != nil {
            structMap["lastcheck"] = a.Lastcheck.Value()
        } else {
            structMap["lastcheck"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsAutoUpgrade.
// It customizes the JSON unmarshaling process for ApStatsAutoUpgrade objects.
func (a *ApStatsAutoUpgrade) UnmarshalJSON(input []byte) error {
    var temp apStatsAutoUpgrade
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "lastcheck")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Lastcheck = temp.Lastcheck
    return nil
}

// apStatsAutoUpgrade is a temporary struct used for validating the fields of ApStatsAutoUpgrade.
type apStatsAutoUpgrade  struct {
    Lastcheck Optional[int64] `json:"lastcheck"`
}
