package models

import (
    "encoding/json"
    "fmt"
)

// UseAutoApValues represents a UseAutoApValues struct.
type UseAutoApValues struct {
    // If accept is true, accepts placement for devices in list otherwise. If false, reject for devices in list.
    Accept               *bool                   `json:"accept,omitempty"`
    // The selector to choose auto placement or auto orientation. enum: `orientation`, `placement`
    For                  *UseAutoApValuesForEnum `json:"for,omitempty"`
    // A list of macs to accept/reject. If a list is not provided the API will accept/reject for the full map.
    Macs                 []string                `json:"macs,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for UseAutoApValues,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UseAutoApValues) String() string {
    return fmt.Sprintf(
    	"UseAutoApValues[Accept=%v, For=%v, Macs=%v, AdditionalProperties=%v]",
    	u.Accept, u.For, u.Macs, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UseAutoApValues.
// It customizes the JSON marshaling process for UseAutoApValues objects.
func (u UseAutoApValues) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "accept", "for", "macs"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UseAutoApValues object to a map representation for JSON marshaling.
func (u UseAutoApValues) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Accept != nil {
        structMap["accept"] = u.Accept
    }
    if u.For != nil {
        structMap["for"] = u.For
    }
    if u.Macs != nil {
        structMap["macs"] = u.Macs
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UseAutoApValues.
// It customizes the JSON unmarshaling process for UseAutoApValues objects.
func (u *UseAutoApValues) UnmarshalJSON(input []byte) error {
    var temp tempUseAutoApValues
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "accept", "for", "macs")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Accept = temp.Accept
    u.For = temp.For
    u.Macs = temp.Macs
    return nil
}

// tempUseAutoApValues is a temporary struct used for validating the fields of UseAutoApValues.
type tempUseAutoApValues  struct {
    Accept *bool                   `json:"accept,omitempty"`
    For    *UseAutoApValuesForEnum `json:"for,omitempty"`
    Macs   []string                `json:"macs,omitempty"`
}
