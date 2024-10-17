package models

import (
	"encoding/json"
)

// TuntermPortConfig represents a TuntermPortConfig struct.
// ethernet port configurations
type TuntermPortConfig struct {
	// list of ports to be used for downstream (to AP) purpose
	DownstreamPorts []string `json:"downstream_ports,omitempty"`
	// weather to separate upstream / downstream ports. default is false where all ports will be used.
	SeparateUpstreamDownstream *bool `json:"separate_upstream_downstream,omitempty"`
	// native VLAN id for upstream ports
	UpstreamPortVlanId *int `json:"upstream_port_vlan_id,omitempty"`
	// list of ports to be used for upstrea purpose (to LAN)
	UpstreamPorts        []string       `json:"upstream_ports,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TuntermPortConfig.
// It customizes the JSON marshaling process for TuntermPortConfig objects.
func (t TuntermPortConfig) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TuntermPortConfig object to a map representation for JSON marshaling.
func (t TuntermPortConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, t.AdditionalProperties)
	if t.DownstreamPorts != nil {
		structMap["downstream_ports"] = t.DownstreamPorts
	}
	if t.SeparateUpstreamDownstream != nil {
		structMap["separate_upstream_downstream"] = t.SeparateUpstreamDownstream
	}
	if t.UpstreamPortVlanId != nil {
		structMap["upstream_port_vlan_id"] = t.UpstreamPortVlanId
	}
	if t.UpstreamPorts != nil {
		structMap["upstream_ports"] = t.UpstreamPorts
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TuntermPortConfig.
// It customizes the JSON unmarshaling process for TuntermPortConfig objects.
func (t *TuntermPortConfig) UnmarshalJSON(input []byte) error {
	var temp tempTuntermPortConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "downstream_ports", "separate_upstream_downstream", "upstream_port_vlan_id", "upstream_ports")
	if err != nil {
		return err
	}

	t.AdditionalProperties = additionalProperties
	t.DownstreamPorts = temp.DownstreamPorts
	t.SeparateUpstreamDownstream = temp.SeparateUpstreamDownstream
	t.UpstreamPortVlanId = temp.UpstreamPortVlanId
	t.UpstreamPorts = temp.UpstreamPorts
	return nil
}

// tempTuntermPortConfig is a temporary struct used for validating the fields of TuntermPortConfig.
type tempTuntermPortConfig struct {
	DownstreamPorts            []string `json:"downstream_ports,omitempty"`
	SeparateUpstreamDownstream *bool    `json:"separate_upstream_downstream,omitempty"`
	UpstreamPortVlanId         *int     `json:"upstream_port_vlan_id,omitempty"`
	UpstreamPorts              []string `json:"upstream_ports,omitempty"`
}
