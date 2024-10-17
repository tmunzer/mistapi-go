package models

import (
	"encoding/json"
)

// TunnelProviderOptionsJse represents a TunnelProviderOptionsJse struct.
// for jse-ipsec, this allow provisioning of adequate resource on JSE. Make sure adequate licenses are added
type TunnelProviderOptionsJse struct {
	Name                 *string        `json:"name,omitempty"`
	NumUsers             *int           `json:"num_users,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TunnelProviderOptionsJse.
// It customizes the JSON marshaling process for TunnelProviderOptionsJse objects.
func (t TunnelProviderOptionsJse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TunnelProviderOptionsJse object to a map representation for JSON marshaling.
func (t TunnelProviderOptionsJse) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, t.AdditionalProperties)
	if t.Name != nil {
		structMap["name"] = t.Name
	}
	if t.NumUsers != nil {
		structMap["num_users"] = t.NumUsers
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelProviderOptionsJse.
// It customizes the JSON unmarshaling process for TunnelProviderOptionsJse objects.
func (t *TunnelProviderOptionsJse) UnmarshalJSON(input []byte) error {
	var temp tempTunnelProviderOptionsJse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "name", "num_users")
	if err != nil {
		return err
	}

	t.AdditionalProperties = additionalProperties
	t.Name = temp.Name
	t.NumUsers = temp.NumUsers
	return nil
}

// tempTunnelProviderOptionsJse is a temporary struct used for validating the fields of TunnelProviderOptionsJse.
type tempTunnelProviderOptionsJse struct {
	Name     *string `json:"name,omitempty"`
	NumUsers *int    `json:"num_users,omitempty"`
}
