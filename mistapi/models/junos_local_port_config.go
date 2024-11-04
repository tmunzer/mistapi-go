package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// JunosLocalPortConfig represents a JunosLocalPortConfig struct.
// Switch port config
type JunosLocalPortConfig struct {
    // if want to generate port up/down alarm
    Critical             *bool                      `json:"critical,omitempty"`
    Description          *string                    `json:"description,omitempty"`
    // if `speed` and `duplex` are specified, whether to disable autonegotiation
    DisableAutoneg       *bool                      `json:"disable_autoneg,omitempty"`
    // enum: `auto`, `full`, `half`
    Duplex               *JunosPortConfigDuplexEnum `json:"duplex,omitempty"`
    // media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation
    Mtu                  *int                       `json:"mtu,omitempty"`
    PoeDisabled          *bool                      `json:"poe_disabled,omitempty"`
    // enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `auto`
    Speed                *JunosPortConfigSpeedEnum  `json:"speed,omitempty"`
    // port usage name.
    // If EVPN is used, use `evpn_uplink`or `evpn_downlink`
    Usage                string                     `json:"usage"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for JunosLocalPortConfig.
// It customizes the JSON marshaling process for JunosLocalPortConfig objects.
func (j JunosLocalPortConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(j.toMap())
}

// toMap converts the JunosLocalPortConfig object to a map representation for JSON marshaling.
func (j JunosLocalPortConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, j.AdditionalProperties)
    if j.Critical != nil {
        structMap["critical"] = j.Critical
    }
    if j.Description != nil {
        structMap["description"] = j.Description
    }
    if j.DisableAutoneg != nil {
        structMap["disable_autoneg"] = j.DisableAutoneg
    }
    if j.Duplex != nil {
        structMap["duplex"] = j.Duplex
    }
    if j.Mtu != nil {
        structMap["mtu"] = j.Mtu
    }
    if j.PoeDisabled != nil {
        structMap["poe_disabled"] = j.PoeDisabled
    }
    if j.Speed != nil {
        structMap["speed"] = j.Speed
    }
    structMap["usage"] = j.Usage
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JunosLocalPortConfig.
// It customizes the JSON unmarshaling process for JunosLocalPortConfig objects.
func (j *JunosLocalPortConfig) UnmarshalJSON(input []byte) error {
    var temp tempJunosLocalPortConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "critical", "description", "disable_autoneg", "duplex", "mtu", "poe_disabled", "speed", "usage")
    if err != nil {
    	return err
    }
    
    j.AdditionalProperties = additionalProperties
    j.Critical = temp.Critical
    j.Description = temp.Description
    j.DisableAutoneg = temp.DisableAutoneg
    j.Duplex = temp.Duplex
    j.Mtu = temp.Mtu
    j.PoeDisabled = temp.PoeDisabled
    j.Speed = temp.Speed
    j.Usage = *temp.Usage
    return nil
}

// tempJunosLocalPortConfig is a temporary struct used for validating the fields of JunosLocalPortConfig.
type tempJunosLocalPortConfig  struct {
    Critical       *bool                      `json:"critical,omitempty"`
    Description    *string                    `json:"description,omitempty"`
    DisableAutoneg *bool                      `json:"disable_autoneg,omitempty"`
    Duplex         *JunosPortConfigDuplexEnum `json:"duplex,omitempty"`
    Mtu            *int                       `json:"mtu,omitempty"`
    PoeDisabled    *bool                      `json:"poe_disabled,omitempty"`
    Speed          *JunosPortConfigSpeedEnum  `json:"speed,omitempty"`
    Usage          *string                    `json:"usage"`
}

func (j *tempJunosLocalPortConfig) validate() error {
    var errs []string
    if j.Usage == nil {
        errs = append(errs, "required field `usage` is missing for type `junos_local_port_config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
