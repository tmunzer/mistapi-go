package models

import (
    "encoding/json"
)

// ApEslConfig represents a ApEslConfig struct.
type ApEslConfig struct {
    // Only if `type`==`imagotag` or `type`==`native`
    Cacert               *string        `json:"cacert,omitempty"`
    // Only if `type`==`imagotag` or `type`==`native`
    Channel              *int           `json:"channel,omitempty"`
    // usb_config is ignored if esl_config enabled
    Enabled              *bool          `json:"enabled,omitempty"`
    // Only if `type`==`imagotag` or `type`==`native`
    Host                 *string        `json:"host,omitempty"`
    // Only if `type`==`imagotag` or `type`==`native`
    Port                 *int           `json:"port,omitempty"`
    // note: ble_config will be ingored if esl_config is enabled and with native mode. enum: `hanshow`, `imagotag`, `native`, `solum`
    Type                 *ApEslTypeEnum `json:"type,omitempty"`
    // Only if `type`==`imagotag` or `type`==`native`
    VerifyCert           *bool          `json:"verify_cert,omitempty"`
    // Only if `type`==`solum` or `type`==`hanshow`
    VlanId               *int           `json:"vlan_id,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApEslConfig.
// It customizes the JSON marshaling process for ApEslConfig objects.
func (a ApEslConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApEslConfig object to a map representation for JSON marshaling.
func (a ApEslConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Cacert != nil {
        structMap["cacert"] = a.Cacert
    }
    if a.Channel != nil {
        structMap["channel"] = a.Channel
    }
    if a.Enabled != nil {
        structMap["enabled"] = a.Enabled
    }
    if a.Host != nil {
        structMap["host"] = a.Host
    }
    if a.Port != nil {
        structMap["port"] = a.Port
    }
    if a.Type != nil {
        structMap["type"] = a.Type
    }
    if a.VerifyCert != nil {
        structMap["verify_cert"] = a.VerifyCert
    }
    if a.VlanId != nil {
        structMap["vlan_id"] = a.VlanId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApEslConfig.
// It customizes the JSON unmarshaling process for ApEslConfig objects.
func (a *ApEslConfig) UnmarshalJSON(input []byte) error {
    var temp tempApEslConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "cacert", "channel", "enabled", "host", "port", "type", "verify_cert", "vlan_id")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Cacert = temp.Cacert
    a.Channel = temp.Channel
    a.Enabled = temp.Enabled
    a.Host = temp.Host
    a.Port = temp.Port
    a.Type = temp.Type
    a.VerifyCert = temp.VerifyCert
    a.VlanId = temp.VlanId
    return nil
}

// tempApEslConfig is a temporary struct used for validating the fields of ApEslConfig.
type tempApEslConfig  struct {
    Cacert     *string        `json:"cacert,omitempty"`
    Channel    *int           `json:"channel,omitempty"`
    Enabled    *bool          `json:"enabled,omitempty"`
    Host       *string        `json:"host,omitempty"`
    Port       *int           `json:"port,omitempty"`
    Type       *ApEslTypeEnum `json:"type,omitempty"`
    VerifyCert *bool          `json:"verify_cert,omitempty"`
    VlanId     *int           `json:"vlan_id,omitempty"`
}
