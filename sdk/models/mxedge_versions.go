package models

import (
    "encoding/json"
)

// MxedgeVersions represents a MxedgeVersions struct.
type MxedgeVersions struct {
    Mxagent              *string        `json:"mxagent,omitempty"`
    Tuntnerm             *string        `json:"tuntnerm,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeVersions.
// It customizes the JSON marshaling process for MxedgeVersions objects.
func (m MxedgeVersions) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeVersions object to a map representation for JSON marshaling.
func (m MxedgeVersions) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Mxagent != nil {
        structMap["mxagent"] = m.Mxagent
    }
    if m.Tuntnerm != nil {
        structMap["tuntnerm"] = m.Tuntnerm
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeVersions.
// It customizes the JSON unmarshaling process for MxedgeVersions objects.
func (m *MxedgeVersions) UnmarshalJSON(input []byte) error {
    var temp mxedgeVersions
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mxagent", "tuntnerm")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Mxagent = temp.Mxagent
    m.Tuntnerm = temp.Tuntnerm
    return nil
}

// mxedgeVersions is a temporary struct used for validating the fields of MxedgeVersions.
type mxedgeVersions  struct {
    Mxagent  *string `json:"mxagent,omitempty"`
    Tuntnerm *string `json:"tuntnerm,omitempty"`
}
