package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// UsermacsId represents a UsermacsId struct.
type UsermacsId struct {
    UsermacIds           []uuid.UUID            `json:"usermac_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UsermacsId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UsermacsId) String() string {
    return fmt.Sprintf(
    	"UsermacsId[UsermacIds=%v, AdditionalProperties=%v]",
    	u.UsermacIds, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UsermacsId.
// It customizes the JSON marshaling process for UsermacsId objects.
func (u UsermacsId) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "usermac_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UsermacsId object to a map representation for JSON marshaling.
func (u UsermacsId) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.UsermacIds != nil {
        structMap["usermac_ids"] = u.UsermacIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UsermacsId.
// It customizes the JSON unmarshaling process for UsermacsId objects.
func (u *UsermacsId) UnmarshalJSON(input []byte) error {
    var temp tempUsermacsId
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "usermac_ids")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.UsermacIds = temp.UsermacIds
    return nil
}

// tempUsermacsId is a temporary struct used for validating the fields of UsermacsId.
type tempUsermacsId  struct {
    UsermacIds []uuid.UUID `json:"usermac_ids,omitempty"`
}
