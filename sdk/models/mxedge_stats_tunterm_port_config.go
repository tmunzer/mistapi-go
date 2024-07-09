package models

import (
    "encoding/json"
)

// MxedgeStatsTuntermPortConfig represents a MxedgeStatsTuntermPortConfig struct.
type MxedgeStatsTuntermPortConfig struct {
    DownstreamPorts            []string       `json:"downstream_ports,omitempty"`
    SeparateUpstreamDownstream *bool          `json:"separate_upstream_downstream,omitempty"`
    UpstreamPorts              []string       `json:"upstream_ports,omitempty"`
    AdditionalProperties       map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeStatsTuntermPortConfig.
// It customizes the JSON marshaling process for MxedgeStatsTuntermPortConfig objects.
func (m MxedgeStatsTuntermPortConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeStatsTuntermPortConfig object to a map representation for JSON marshaling.
func (m MxedgeStatsTuntermPortConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.DownstreamPorts != nil {
        structMap["downstream_ports"] = m.DownstreamPorts
    }
    if m.SeparateUpstreamDownstream != nil {
        structMap["separate_upstream_downstream"] = m.SeparateUpstreamDownstream
    }
    if m.UpstreamPorts != nil {
        structMap["upstream_ports"] = m.UpstreamPorts
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeStatsTuntermPortConfig.
// It customizes the JSON unmarshaling process for MxedgeStatsTuntermPortConfig objects.
func (m *MxedgeStatsTuntermPortConfig) UnmarshalJSON(input []byte) error {
    var temp mxedgeStatsTuntermPortConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "downstream_ports", "separate_upstream_downstream", "upstream_ports")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.DownstreamPorts = temp.DownstreamPorts
    m.SeparateUpstreamDownstream = temp.SeparateUpstreamDownstream
    m.UpstreamPorts = temp.UpstreamPorts
    return nil
}

// mxedgeStatsTuntermPortConfig is a temporary struct used for validating the fields of MxedgeStatsTuntermPortConfig.
type mxedgeStatsTuntermPortConfig  struct {
    DownstreamPorts            []string `json:"downstream_ports,omitempty"`
    SeparateUpstreamDownstream *bool    `json:"separate_upstream_downstream,omitempty"`
    UpstreamPorts              []string `json:"upstream_ports,omitempty"`
}
