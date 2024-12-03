package models

import (
    "encoding/json"
)

// SecintelProfileProfile represents a SecintelProfileProfile struct.
type SecintelProfileProfile struct {
    // enum: `default`, `standard`, `strict`
    Action               *SecintelProfileProfileActionEnum   `json:"action,omitempty"`
    // enum: `CC`, `IH` (Infected Host), `DNS`
    Category             *SecintelProfileProfileCategoryEnum `json:"category,omitempty"`
    AdditionalProperties map[string]interface{}              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SecintelProfileProfile.
// It customizes the JSON marshaling process for SecintelProfileProfile objects.
func (s SecintelProfileProfile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "action", "category"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SecintelProfileProfile object to a map representation for JSON marshaling.
func (s SecintelProfileProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Action != nil {
        structMap["action"] = s.Action
    }
    if s.Category != nil {
        structMap["category"] = s.Category
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SecintelProfileProfile.
// It customizes the JSON unmarshaling process for SecintelProfileProfile objects.
func (s *SecintelProfileProfile) UnmarshalJSON(input []byte) error {
    var temp tempSecintelProfileProfile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "action", "category")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Action = temp.Action
    s.Category = temp.Category
    return nil
}

// tempSecintelProfileProfile is a temporary struct used for validating the fields of SecintelProfileProfile.
type tempSecintelProfileProfile  struct {
    Action   *SecintelProfileProfileActionEnum   `json:"action,omitempty"`
    Category *SecintelProfileProfileCategoryEnum `json:"category,omitempty"`
}
