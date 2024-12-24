package models

import (
    "encoding/json"
    "fmt"
)

// UserMacImport represents a UserMacImport struct.
type UserMacImport struct {
    Added                []string               `json:"added,omitempty"`
    Errors               []string               `json:"errors,omitempty"`
    Updated              []string               `json:"updated,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UserMacImport,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UserMacImport) String() string {
    return fmt.Sprintf(
    	"UserMacImport[Added=%v, Errors=%v, Updated=%v, AdditionalProperties=%v]",
    	u.Added, u.Errors, u.Updated, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UserMacImport.
// It customizes the JSON marshaling process for UserMacImport objects.
func (u UserMacImport) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "added", "errors", "updated"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UserMacImport object to a map representation for JSON marshaling.
func (u UserMacImport) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
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
    var temp tempUserMacImport
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "added", "errors", "updated")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Added = temp.Added
    u.Errors = temp.Errors
    u.Updated = temp.Updated
    return nil
}

// tempUserMacImport is a temporary struct used for validating the fields of UserMacImport.
type tempUserMacImport  struct {
    Added   []string `json:"added,omitempty"`
    Errors  []string `json:"errors,omitempty"`
    Updated []string `json:"updated,omitempty"`
}
