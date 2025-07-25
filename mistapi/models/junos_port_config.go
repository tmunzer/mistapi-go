// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// JunosPortConfig represents a JunosPortConfig struct.
// Switch port config
type JunosPortConfig struct {
    // To disable LACP support for the AE interface
    AeDisableLacp        *bool                      `json:"ae_disable_lacp,omitempty"`
    // Users could force to use the designated AE name
    AeIdx                *int                       `json:"ae_idx,omitempty"`
    // To use fast timeout
    AeLacpSlow           *bool                      `json:"ae_lacp_slow,omitempty"`
    Aggregated           *bool                      `json:"aggregated,omitempty"`
    // To generate port up/down alarm
    Critical             *bool                      `json:"critical,omitempty"`
    Description          *string                    `json:"description,omitempty"`
    // If `speed` and `duplex` are specified, whether to disable autonegotiation
    DisableAutoneg       *bool                      `json:"disable_autoneg,omitempty"`
    // enum: `auto`, `full`, `half`
    Duplex               *JunosPortConfigDuplexEnum `json:"duplex,omitempty"`
    // Enable dynamic usage for this port. Set to `dynamic` to enable.
    DynamicUsage         Optional[string]           `json:"dynamic_usage"`
    Esilag               *bool                      `json:"esilag,omitempty"`
    // Media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation
    Mtu                  *int                       `json:"mtu,omitempty"`
    // Prevent helpdesk to override the port config
    NoLocalOverwrite     *bool                      `json:"no_local_overwrite,omitempty"`
    PoeDisabled          *bool                      `json:"poe_disabled,omitempty"`
    // Required if `usage`==`vlan_tunnel`. Q-in-Q tunneling using All-in-one bundling. This also enables standard L2PT for interfaces that are not encapsulation tunnel interfaces and uses MAC rewrite operation. [View more information](https://www.juniper.net/documentation/us/en/software/junos/multicast-l2/topics/topic-map/q-in-q.html#id-understanding-qinq-tunneling-and-vlan-translation)
    PortNetwork          *string                    `json:"port_network,omitempty"`
    // enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `10g`, `25g`, `40g`, `100g`,`auto`
    Speed                *JunosPortConfigSpeedEnum  `json:"speed,omitempty"`
    // Port usage name. For Q-in-Q, use `vlan_tunnel`. If EVPN is used, use `evpn_uplink`or `evpn_downlink`
    Usage                string                     `json:"usage"`
    AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for JunosPortConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (j JunosPortConfig) String() string {
    return fmt.Sprintf(
    	"JunosPortConfig[AeDisableLacp=%v, AeIdx=%v, AeLacpSlow=%v, Aggregated=%v, Critical=%v, Description=%v, DisableAutoneg=%v, Duplex=%v, DynamicUsage=%v, Esilag=%v, Mtu=%v, NoLocalOverwrite=%v, PoeDisabled=%v, PortNetwork=%v, Speed=%v, Usage=%v, AdditionalProperties=%v]",
    	j.AeDisableLacp, j.AeIdx, j.AeLacpSlow, j.Aggregated, j.Critical, j.Description, j.DisableAutoneg, j.Duplex, j.DynamicUsage, j.Esilag, j.Mtu, j.NoLocalOverwrite, j.PoeDisabled, j.PortNetwork, j.Speed, j.Usage, j.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for JunosPortConfig.
// It customizes the JSON marshaling process for JunosPortConfig objects.
func (j JunosPortConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(j.AdditionalProperties,
        "ae_disable_lacp", "ae_idx", "ae_lacp_slow", "aggregated", "critical", "description", "disable_autoneg", "duplex", "dynamic_usage", "esilag", "mtu", "no_local_overwrite", "poe_disabled", "port_network", "speed", "usage"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(j.toMap())
}

// toMap converts the JunosPortConfig object to a map representation for JSON marshaling.
func (j JunosPortConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, j.AdditionalProperties)
    if j.AeDisableLacp != nil {
        structMap["ae_disable_lacp"] = j.AeDisableLacp
    }
    if j.AeIdx != nil {
        structMap["ae_idx"] = j.AeIdx
    }
    if j.AeLacpSlow != nil {
        structMap["ae_lacp_slow"] = j.AeLacpSlow
    }
    if j.Aggregated != nil {
        structMap["aggregated"] = j.Aggregated
    }
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
    if j.DynamicUsage.IsValueSet() {
        if j.DynamicUsage.Value() != nil {
            structMap["dynamic_usage"] = j.DynamicUsage.Value()
        } else {
            structMap["dynamic_usage"] = nil
        }
    }
    if j.Esilag != nil {
        structMap["esilag"] = j.Esilag
    }
    if j.Mtu != nil {
        structMap["mtu"] = j.Mtu
    }
    if j.NoLocalOverwrite != nil {
        structMap["no_local_overwrite"] = j.NoLocalOverwrite
    }
    if j.PoeDisabled != nil {
        structMap["poe_disabled"] = j.PoeDisabled
    }
    if j.PortNetwork != nil {
        structMap["port_network"] = j.PortNetwork
    }
    if j.Speed != nil {
        structMap["speed"] = j.Speed
    }
    structMap["usage"] = j.Usage
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JunosPortConfig.
// It customizes the JSON unmarshaling process for JunosPortConfig objects.
func (j *JunosPortConfig) UnmarshalJSON(input []byte) error {
    var temp tempJunosPortConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ae_disable_lacp", "ae_idx", "ae_lacp_slow", "aggregated", "critical", "description", "disable_autoneg", "duplex", "dynamic_usage", "esilag", "mtu", "no_local_overwrite", "poe_disabled", "port_network", "speed", "usage")
    if err != nil {
    	return err
    }
    j.AdditionalProperties = additionalProperties
    
    j.AeDisableLacp = temp.AeDisableLacp
    j.AeIdx = temp.AeIdx
    j.AeLacpSlow = temp.AeLacpSlow
    j.Aggregated = temp.Aggregated
    j.Critical = temp.Critical
    j.Description = temp.Description
    j.DisableAutoneg = temp.DisableAutoneg
    j.Duplex = temp.Duplex
    j.DynamicUsage = temp.DynamicUsage
    j.Esilag = temp.Esilag
    j.Mtu = temp.Mtu
    j.NoLocalOverwrite = temp.NoLocalOverwrite
    j.PoeDisabled = temp.PoeDisabled
    j.PortNetwork = temp.PortNetwork
    j.Speed = temp.Speed
    j.Usage = *temp.Usage
    return nil
}

// tempJunosPortConfig is a temporary struct used for validating the fields of JunosPortConfig.
type tempJunosPortConfig  struct {
    AeDisableLacp    *bool                      `json:"ae_disable_lacp,omitempty"`
    AeIdx            *int                       `json:"ae_idx,omitempty"`
    AeLacpSlow       *bool                      `json:"ae_lacp_slow,omitempty"`
    Aggregated       *bool                      `json:"aggregated,omitempty"`
    Critical         *bool                      `json:"critical,omitempty"`
    Description      *string                    `json:"description,omitempty"`
    DisableAutoneg   *bool                      `json:"disable_autoneg,omitempty"`
    Duplex           *JunosPortConfigDuplexEnum `json:"duplex,omitempty"`
    DynamicUsage     Optional[string]           `json:"dynamic_usage"`
    Esilag           *bool                      `json:"esilag,omitempty"`
    Mtu              *int                       `json:"mtu,omitempty"`
    NoLocalOverwrite *bool                      `json:"no_local_overwrite,omitempty"`
    PoeDisabled      *bool                      `json:"poe_disabled,omitempty"`
    PortNetwork      *string                    `json:"port_network,omitempty"`
    Speed            *JunosPortConfigSpeedEnum  `json:"speed,omitempty"`
    Usage            *string                    `json:"usage"`
}

func (j *tempJunosPortConfig) validate() error {
    var errs []string
    if j.Usage == nil {
        errs = append(errs, "required field `usage` is missing for type `junos_port_config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
