package models

import (
	"encoding/json"
)

// StatsMxedgeTuntermPortConfig represents a StatsMxedgeTuntermPortConfig struct.
type StatsMxedgeTuntermPortConfig struct {
	DownstreamPorts            []string       `json:"downstream_ports,omitempty"`
	SeparateUpstreamDownstream *bool          `json:"separate_upstream_downstream,omitempty"`
	UpstreamPorts              []string       `json:"upstream_ports,omitempty"`
	AdditionalProperties       map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgeTuntermPortConfig.
// It customizes the JSON marshaling process for StatsMxedgeTuntermPortConfig objects.
func (s StatsMxedgeTuntermPortConfig) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgeTuntermPortConfig object to a map representation for JSON marshaling.
func (s StatsMxedgeTuntermPortConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.DownstreamPorts != nil {
		structMap["downstream_ports"] = s.DownstreamPorts
	}
	if s.SeparateUpstreamDownstream != nil {
		structMap["separate_upstream_downstream"] = s.SeparateUpstreamDownstream
	}
	if s.UpstreamPorts != nil {
		structMap["upstream_ports"] = s.UpstreamPorts
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgeTuntermPortConfig.
// It customizes the JSON unmarshaling process for StatsMxedgeTuntermPortConfig objects.
func (s *StatsMxedgeTuntermPortConfig) UnmarshalJSON(input []byte) error {
	var temp tempStatsMxedgeTuntermPortConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "downstream_ports", "separate_upstream_downstream", "upstream_ports")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.DownstreamPorts = temp.DownstreamPorts
	s.SeparateUpstreamDownstream = temp.SeparateUpstreamDownstream
	s.UpstreamPorts = temp.UpstreamPorts
	return nil
}

// tempStatsMxedgeTuntermPortConfig is a temporary struct used for validating the fields of StatsMxedgeTuntermPortConfig.
type tempStatsMxedgeTuntermPortConfig struct {
	DownstreamPorts            []string `json:"downstream_ports,omitempty"`
	SeparateUpstreamDownstream *bool    `json:"separate_upstream_downstream,omitempty"`
	UpstreamPorts              []string `json:"upstream_ports,omitempty"`
}
