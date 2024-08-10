package models

import (
    "encoding/json"
)

// WlanInjectDhcpOption82 represents a WlanInjectDhcpOption82 struct.
type WlanInjectDhcpOption82 struct {
    CircuitId            *string        `json:"circuit_id,omitempty"`
    // whether to inject option 82 when forwarding DHCP packets
    Enabled              *bool          `json:"enabled,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanInjectDhcpOption82.
// It customizes the JSON marshaling process for WlanInjectDhcpOption82 objects.
func (w WlanInjectDhcpOption82) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WlanInjectDhcpOption82 object to a map representation for JSON marshaling.
func (w WlanInjectDhcpOption82) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.CircuitId != nil {
        structMap["circuit_id"] = w.CircuitId
    }
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanInjectDhcpOption82.
// It customizes the JSON unmarshaling process for WlanInjectDhcpOption82 objects.
func (w *WlanInjectDhcpOption82) UnmarshalJSON(input []byte) error {
    var temp tempWlanInjectDhcpOption82
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "circuit_id", "enabled")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.CircuitId = temp.CircuitId
    w.Enabled = temp.Enabled
    return nil
}

// tempWlanInjectDhcpOption82 is a temporary struct used for validating the fields of WlanInjectDhcpOption82.
type tempWlanInjectDhcpOption82  struct {
    CircuitId *string `json:"circuit_id,omitempty"`
    Enabled   *bool   `json:"enabled,omitempty"`
}
