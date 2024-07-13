package models

import (
    "encoding/json"
)

// UserMacImport represents a UserMacImport struct.
type UserMacImport struct {
    Added                []string       `json:"added,omitempty"`
    Errors               []string       `json:"errors,omitempty"`
    Updated              []string       `json:"updated,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UserMacImport.
// It customizes the JSON marshaling process for UserMacImport objects.
func (u UserMacImport) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UserMacImport object to a map representation for JSON marshaling.
func (u UserMacImport) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Added != nil {
        structMap["added"] = u.Added
    }
    if u.Errors != nil {
        structMap["errors"] = u.Errors
    }
    if u.Updated != nil {
        structMap["updated"] = u.Updated
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UserMacImport.
// It customizes the JSON unmarshaling process for UserMacImport objects.
func (u *UserMacImport) UnmarshalJSON(input []byte) error {
    var temp userMacImport
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "added", "errors", "updated")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Added = temp.Added
    u.Errors = temp.Errors
    u.Updated = temp.Updated
    return nil
}

// userMacImport is a temporary struct used for validating the fields of UserMacImport.
type userMacImport  struct {
    Added   []string `json:"added,omitempty"`
    Errors  []string `json:"errors,omitempty"`
    Updated []string `json:"updated,omitempty"`
}
