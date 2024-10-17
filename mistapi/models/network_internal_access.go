package models

import (
	"encoding/json"
)

// NetworkInternalAccess represents a NetworkInternalAccess struct.
type NetworkInternalAccess struct {
	Enabled              *bool          `json:"enabled,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NetworkInternalAccess.
// It customizes the JSON marshaling process for NetworkInternalAccess objects.
func (n NetworkInternalAccess) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(n.toMap())
}

// toMap converts the NetworkInternalAccess object to a map representation for JSON marshaling.
func (n NetworkInternalAccess) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, n.AdditionalProperties)
	if n.Enabled != nil {
		structMap["enabled"] = n.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkInternalAccess.
// It customizes the JSON unmarshaling process for NetworkInternalAccess objects.
func (n *NetworkInternalAccess) UnmarshalJSON(input []byte) error {
	var temp tempNetworkInternalAccess
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled")
	if err != nil {
		return err
	}

	n.AdditionalProperties = additionalProperties
	n.Enabled = temp.Enabled
	return nil
}

// tempNetworkInternalAccess is a temporary struct used for validating the fields of NetworkInternalAccess.
type tempNetworkInternalAccess struct {
	Enabled *bool `json:"enabled,omitempty"`
}
