package models

import (
    "encoding/json"
    "fmt"
)

// MxedgeVersions represents a MxedgeVersions struct.
type MxedgeVersions struct {
    Mxagent              *string                `json:"mxagent,omitempty"`
    Tunterm              *string                `json:"tunterm,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeVersions,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeVersions) String() string {
    return fmt.Sprintf(
    	"MxedgeVersions[Mxagent=%v, Tunterm=%v, AdditionalProperties=%v]",
    	m.Mxagent, m.Tunterm, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeVersions.
// It customizes the JSON marshaling process for MxedgeVersions objects.
func (m MxedgeVersions) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "mxagent", "tunterm"); err != nil {
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
    if m.Tunterm != nil {
        structMap["tunterm"] = m.Tunterm
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mxagent", "tunterm")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Mxagent = temp.Mxagent
    m.Tunterm = temp.Tunterm
    return nil
}

// tempMxedgeVersions is a temporary struct used for validating the fields of MxedgeVersions.
type tempMxedgeVersions  struct {
    Mxagent *string `json:"mxagent,omitempty"`
    Tunterm *string `json:"tunterm,omitempty"`
}
