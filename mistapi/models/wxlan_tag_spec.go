package models

import (
	"encoding/json"
)

// WxlanTagSpec represents a WxlanTagSpec struct.
type WxlanTagSpec struct {
	// matched destination port, "0" means any
	PortRange *string `json:"port_range,omitempty"`
	// tcp / udp / icmp / gre / any / ":protocol_number", `protocol_number` is between 1-254
	Protocol *string `json:"protocol,omitempty"`
	// matched destination subnets and/or IP Addresses
	Subnets              []string       `json:"subnets,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WxlanTagSpec.
// It customizes the JSON marshaling process for WxlanTagSpec objects.
func (w WxlanTagSpec) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(w.toMap())
}

// toMap converts the WxlanTagSpec object to a map representation for JSON marshaling.
func (w WxlanTagSpec) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, w.AdditionalProperties)
	if w.PortRange != nil {
		structMap["port_range"] = w.PortRange
	}
	if w.Protocol != nil {
		structMap["protocol"] = w.Protocol
	}
	if w.Subnets != nil {
		structMap["subnets"] = w.Subnets
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WxlanTagSpec.
// It customizes the JSON unmarshaling process for WxlanTagSpec objects.
func (w *WxlanTagSpec) UnmarshalJSON(input []byte) error {
	var temp tempWxlanTagSpec
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "port_range", "protocol", "subnets")
	if err != nil {
		return err
	}

	w.AdditionalProperties = additionalProperties
	w.PortRange = temp.PortRange
	w.Protocol = temp.Protocol
	w.Subnets = temp.Subnets
	return nil
}

// tempWxlanTagSpec is a temporary struct used for validating the fields of WxlanTagSpec.
type tempWxlanTagSpec struct {
	PortRange *string  `json:"port_range,omitempty"`
	Protocol  *string  `json:"protocol,omitempty"`
	Subnets   []string `json:"subnets,omitempty"`
}
