package models

import (
    "encoding/json"
)

// MxedgeVersions represents a MxedgeVersions struct.
type MxedgeVersions struct {
    Mxagent              *string                `json:"mxagent,omitempty"`
    Tuntnerm             *string                `json:"tuntnerm,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeVersions.
// It customizes the JSON marshaling process for MxedgeVersions objects.
func (m MxedgeVersions) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "mxagent", "tuntnerm"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeVersions object to a map representation for JSON marshaling.
func (m MxedgeVersions) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
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
    var temp tempMxedgeVersions
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mxagent", "tuntnerm")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Mxagent = temp.Mxagent
    m.Tuntnerm = temp.Tuntnerm
    return nil
}

// tempMxedgeVersions is a temporary struct used for validating the fields of MxedgeVersions.
type tempMxedgeVersions  struct {
    Mxagent  *string `json:"mxagent,omitempty"`
    Tuntnerm *string `json:"tuntnerm,omitempty"`
}
