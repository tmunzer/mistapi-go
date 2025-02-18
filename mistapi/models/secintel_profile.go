package models

import (
    "encoding/json"
    "fmt"
)

// SecintelProfile represents a SecintelProfile struct.
type SecintelProfile struct {
    Name                 *string                  `json:"name,omitempty"`
    Profiles             []SecintelProfileProfile `json:"profiles,omitempty"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for SecintelProfile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SecintelProfile) String() string {
    return fmt.Sprintf(
    	"SecintelProfile[Name=%v, Profiles=%v, AdditionalProperties=%v]",
    	s.Name, s.Profiles, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SecintelProfile.
// It customizes the JSON marshaling process for SecintelProfile objects.
func (s SecintelProfile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "name", "profiles"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SecintelProfile object to a map representation for JSON marshaling.
func (s SecintelProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "profiles")
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
