package models

import (
    "encoding/json"
)

// UtilsClearBpdu represents a UtilsClearBpdu struct.
type UtilsClearBpdu struct {
    // the port on which to clear the detected BPDU error, or `all` for all ports
    Port                 *string                `json:"port,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsClearBpdu.
// It customizes the JSON marshaling process for UtilsClearBpdu objects.
func (u UtilsClearBpdu) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "port"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsClearBpdu object to a map representation for JSON marshaling.
func (u UtilsClearBpdu) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Port != nil {
        structMap["port"] = u.Port
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsClearBpdu.
// It customizes the JSON unmarshaling process for UtilsClearBpdu objects.
func (u *UtilsClearBpdu) UnmarshalJSON(input []byte) error {
    var temp tempUtilsClearBpdu
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "port")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Port = temp.Port
    return nil
}

// tempUtilsClearBpdu is a temporary struct used for validating the fields of UtilsClearBpdu.
type tempUtilsClearBpdu  struct {
    Port *string `json:"port,omitempty"`
}
