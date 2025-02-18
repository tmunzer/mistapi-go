package models

import (
    "encoding/json"
    "fmt"
)

// StatsMxedgeTuntermPortConfig represents a StatsMxedgeTuntermPortConfig struct.
type StatsMxedgeTuntermPortConfig struct {
    DownstreamPorts            []string               `json:"downstream_ports,omitempty"`
    SeparateUpstreamDownstream *bool                  `json:"separate_upstream_downstream,omitempty"`
    UpstreamPorts              []string               `json:"upstream_ports,omitempty"`
    AdditionalProperties       map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMxedgeTuntermPortConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMxedgeTuntermPortConfig) String() string {
    return fmt.Sprintf(
    	"StatsMxedgeTuntermPortConfig[DownstreamPorts=%v, SeparateUpstreamDownstream=%v, UpstreamPorts=%v, AdditionalProperties=%v]",
    	s.DownstreamPorts, s.SeparateUpstreamDownstream, s.UpstreamPorts, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgeTuntermPortConfig.
// It customizes the JSON marshaling process for StatsMxedgeTuntermPortConfig objects.
func (s StatsMxedgeTuntermPortConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "downstream_ports", "separate_upstream_downstream", "upstream_ports"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgeTuntermPortConfig object to a map representation for JSON marshaling.
func (s StatsMxedgeTuntermPortConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "downstream_ports", "separate_upstream_downstream", "upstream_ports")
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
type tempStatsMxedgeTuntermPortConfig  struct {
    DownstreamPorts            []string `json:"downstream_ports,omitempty"`
    SeparateUpstreamDownstream *bool    `json:"separate_upstream_downstream,omitempty"`
    UpstreamPorts              []string `json:"upstream_ports,omitempty"`
}
