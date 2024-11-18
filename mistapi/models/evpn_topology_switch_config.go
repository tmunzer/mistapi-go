package models

import (
    "encoding/json"
)

// EvpnTopologySwitchConfig represents a EvpnTopologySwitchConfig struct.
type EvpnTopologySwitchConfig struct {
    // Property key is the port name or range (e.g. "ge-0/0/0-10")
    PortConfig           map[string]JunosPortConfig `json:"port_config,omitempty"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EvpnTopologySwitchConfig.
// It customizes the JSON marshaling process for EvpnTopologySwitchConfig objects.
func (e EvpnTopologySwitchConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnTopologySwitchConfig object to a map representation for JSON marshaling.
func (e EvpnTopologySwitchConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    if e.PortConfig != nil {
        structMap["port_config"] = e.PortConfig
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EvpnTopologySwitchConfig.
// It customizes the JSON unmarshaling process for EvpnTopologySwitchConfig objects.
func (e *EvpnTopologySwitchConfig) UnmarshalJSON(input []byte) error {
    var temp tempEvpnTopologySwitchConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "port_config")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.PortConfig = temp.PortConfig
    return nil
}

// tempEvpnTopologySwitchConfig is a temporary struct used for validating the fields of EvpnTopologySwitchConfig.
type tempEvpnTopologySwitchConfig  struct {
    PortConfig map[string]JunosPortConfig `json:"port_config,omitempty"`
}
