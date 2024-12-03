package models

import (
    "encoding/json"
)

// MxedgeDas represents a MxedgeDas struct.
// configure cloud-assisted dynamic authorization service on this cluster of mist edges
type MxedgeDas struct {
    // dynamic authorization clients configured to send CoA|DM to mist edges on port 3799
    CoaServers           []MxedgeDasCoaServer   `json:"coa_servers,omitempty"`
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeDas.
// It customizes the JSON marshaling process for MxedgeDas objects.
func (m MxedgeDas) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "coa_servers", "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeDas object to a map representation for JSON marshaling.
func (m MxedgeDas) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.CoaServers != nil {
        structMap["coa_servers"] = m.CoaServers
    }
    if m.Enabled != nil {
        structMap["enabled"] = m.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeDas.
// It customizes the JSON unmarshaling process for MxedgeDas objects.
func (m *MxedgeDas) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeDas
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "coa_servers", "enabled")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.CoaServers = temp.CoaServers
    m.Enabled = temp.Enabled
    return nil
}

// tempMxedgeDas is a temporary struct used for validating the fields of MxedgeDas.
type tempMxedgeDas  struct {
    CoaServers []MxedgeDasCoaServer `json:"coa_servers,omitempty"`
    Enabled    *bool                `json:"enabled,omitempty"`
}
