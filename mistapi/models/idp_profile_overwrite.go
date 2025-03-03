package models

import (
    "encoding/json"
    "fmt"
)

// IdpProfileOverwrite represents a IdpProfileOverwrite struct.
type IdpProfileOverwrite struct {
    // enum:
    // * alert (default)
    // * drop: silently dropping packets
    // * close: notify client/server to close connection
    Action               *IdpProfileActionEnum  `json:"action,omitempty"`
    Matching             *IdpProfileMatching    `json:"matching,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for IdpProfileOverwrite,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i IdpProfileOverwrite) String() string {
    return fmt.Sprintf(
    	"IdpProfileOverwrite[Action=%v, Matching=%v, Name=%v, AdditionalProperties=%v]",
    	i.Action, i.Matching, i.Name, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for IdpProfileOverwrite.
// It customizes the JSON marshaling process for IdpProfileOverwrite objects.
func (i IdpProfileOverwrite) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "action", "matching", "name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the IdpProfileOverwrite object to a map representation for JSON marshaling.
func (i IdpProfileOverwrite) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Action != nil {
        structMap["action"] = i.Action
    }
    if i.Matching != nil {
        structMap["matching"] = i.Matching.toMap()
    }
    if i.Name != nil {
        structMap["name"] = i.Name
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IdpProfileOverwrite.
// It customizes the JSON unmarshaling process for IdpProfileOverwrite objects.
func (i *IdpProfileOverwrite) UnmarshalJSON(input []byte) error {
    var temp tempIdpProfileOverwrite
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "action", "matching", "name")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.Action = temp.Action
    i.Matching = temp.Matching
    i.Name = temp.Name
    return nil
}

// tempIdpProfileOverwrite is a temporary struct used for validating the fields of IdpProfileOverwrite.
type tempIdpProfileOverwrite  struct {
    Action   *IdpProfileActionEnum `json:"action,omitempty"`
    Matching *IdpProfileMatching   `json:"matching,omitempty"`
    Name     *string               `json:"name,omitempty"`
}
