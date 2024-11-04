package models

import (
    "encoding/json"
)

// SecintelProfile represents a SecintelProfile struct.
type SecintelProfile struct {
    Name                 *string                  `json:"name,omitempty"`
    Profiles             []SecintelProfileProfile `json:"profiles,omitempty"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SecintelProfile.
// It customizes the JSON marshaling process for SecintelProfile objects.
func (s SecintelProfile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SecintelProfile object to a map representation for JSON marshaling.
func (s SecintelProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.Profiles != nil {
        structMap["profiles"] = s.Profiles
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SecintelProfile.
// It customizes the JSON unmarshaling process for SecintelProfile objects.
func (s *SecintelProfile) UnmarshalJSON(input []byte) error {
    var temp tempSecintelProfile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "name", "profiles")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Name = temp.Name
    s.Profiles = temp.Profiles
    return nil
}

// tempSecintelProfile is a temporary struct used for validating the fields of SecintelProfile.
type tempSecintelProfile  struct {
    Name     *string                  `json:"name,omitempty"`
    Profiles []SecintelProfileProfile `json:"profiles,omitempty"`
}
