package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// ApSearchWlan represents a ApSearchWlan struct.
type ApSearchWlan struct {
    Id                   *uuid.UUID             `json:"id,omitempty"`
    Ssid                 *string                `json:"ssid,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApSearchWlan,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApSearchWlan) String() string {
    return fmt.Sprintf(
    	"ApSearchWlan[Id=%v, Ssid=%v, AdditionalProperties=%v]",
    	a.Id, a.Ssid, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApSearchWlan.
// It customizes the JSON marshaling process for ApSearchWlan objects.
func (a ApSearchWlan) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "id", "ssid"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApSearchWlan object to a map representation for JSON marshaling.
func (a ApSearchWlan) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Id != nil {
        structMap["id"] = a.Id
    }
    if a.Ssid != nil {
        structMap["ssid"] = a.Ssid
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApSearchWlan.
// It customizes the JSON unmarshaling process for ApSearchWlan objects.
func (a *ApSearchWlan) UnmarshalJSON(input []byte) error {
    var temp tempApSearchWlan
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "ssid")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Id = temp.Id
    a.Ssid = temp.Ssid
    return nil
}

// tempApSearchWlan is a temporary struct used for validating the fields of ApSearchWlan.
type tempApSearchWlan  struct {
    Id   *uuid.UUID `json:"id,omitempty"`
    Ssid *string    `json:"ssid,omitempty"`
}
