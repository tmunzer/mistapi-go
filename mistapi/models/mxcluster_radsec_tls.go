package models

import (
    "encoding/json"
)

// MxclusterRadsecTls represents a MxclusterRadsecTls struct.
type MxclusterRadsecTls struct {
    Keypair              *string        `json:"keypair,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxclusterRadsecTls.
// It customizes the JSON marshaling process for MxclusterRadsecTls objects.
func (m MxclusterRadsecTls) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxclusterRadsecTls object to a map representation for JSON marshaling.
func (m MxclusterRadsecTls) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Keypair != nil {
        structMap["keypair"] = m.Keypair
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxclusterRadsecTls.
// It customizes the JSON unmarshaling process for MxclusterRadsecTls objects.
func (m *MxclusterRadsecTls) UnmarshalJSON(input []byte) error {
    var temp mxclusterRadsecTls
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "keypair")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Keypair = temp.Keypair
    return nil
}

// mxclusterRadsecTls is a temporary struct used for validating the fields of MxclusterRadsecTls.
type mxclusterRadsecTls  struct {
    Keypair *string `json:"keypair,omitempty"`
}
