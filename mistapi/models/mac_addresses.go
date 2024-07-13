package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// MacAddresses represents a MacAddresses struct.
type MacAddresses struct {
    Macs                 []string       `json:"macs"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MacAddresses.
// It customizes the JSON marshaling process for MacAddresses objects.
func (m MacAddresses) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MacAddresses object to a map representation for JSON marshaling.
func (m MacAddresses) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["macs"] = m.Macs
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MacAddresses.
// It customizes the JSON unmarshaling process for MacAddresses objects.
func (m *MacAddresses) UnmarshalJSON(input []byte) error {
    var temp macAddresses
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "macs")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Macs = *temp.Macs
    return nil
}

// macAddresses is a temporary struct used for validating the fields of MacAddresses.
type macAddresses  struct {
    Macs *[]string `json:"macs"`
}

func (m *macAddresses) validate() error {
    var errs []string
    if m.Macs == nil {
        errs = append(errs, "required field `macs` is missing for type `Mac_Addresses`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
