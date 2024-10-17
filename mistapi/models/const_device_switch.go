package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ConstDeviceSwitch represents a ConstDeviceSwitch struct.
type ConstDeviceSwitch struct {
	Alias                *string                   `json:"alias,omitempty"`
	Defaults             *ConstDeviceSwitchDefault `json:"defaults,omitempty"`
	Description          *string                   `json:"description,omitempty"`
	Display              *string                   `json:"display,omitempty"`
	EvolvedOs            *bool                     `json:"evolved_os,omitempty"`
	EvpnRiType           *string                   `json:"evpn_ri_type,omitempty"`
	Experimental         *bool                     `json:"experimental,omitempty"`
	FansPluggable        *bool                     `json:"fans_pluggable,omitempty"`
	HasBgp               *bool                     `json:"has_bgp,omitempty"`
	HasEts               *bool                     `json:"has_ets,omitempty"`
	HasEvpn              *bool                     `json:"has_evpn,omitempty"`
	HasIrb               *bool                     `json:"has_irb,omitempty"`
	HasPoeOut            *bool                     `json:"has_poe_out,omitempty"`
	HasSnapshot          *bool                     `json:"has_snapshot,omitempty"`
	HasVc                *bool                     `json:"has_vc,omitempty"`
	Model                *string                   `json:"model,omitempty"`
	Modular              *bool                     `json:"modular,omitempty"`
	NoShapingRate        *bool                     `json:"no_shaping_rate,omitempty"`
	NumberFans           *int                      `json:"number_fans,omitempty"`
	OcDevice             *bool                     `json:"oc_device,omitempty"`
	OobInterface         *string                   `json:"oob_interface,omitempty"`
	PacketActionDropOnly *bool                     `json:"packet_action_drop_only,omitempty"`
	// Object Key is the PIC number
	Pic         map[string]string `json:"pic,omitempty"`
	SubRequired *string           `json:"sub_required,omitempty"`
	// Device Type. enum: `switch`
	Type                 string         `json:"type"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstDeviceSwitch.
// It customizes the JSON marshaling process for ConstDeviceSwitch objects.
func (c ConstDeviceSwitch) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ConstDeviceSwitch object to a map representation for JSON marshaling.
func (c ConstDeviceSwitch) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, c.AdditionalProperties)
	if c.Alias != nil {
		structMap["alias"] = c.Alias
	}
	if c.Defaults != nil {
		structMap["defaults"] = c.Defaults.toMap()
	}
	if c.Description != nil {
		structMap["description"] = c.Description
	}
	if c.Display != nil {
		structMap["display"] = c.Display
	}
	if c.EvolvedOs != nil {
		structMap["evolved_os"] = c.EvolvedOs
	}
	if c.EvpnRiType != nil {
		structMap["evpn_ri_type"] = c.EvpnRiType
	}
	if c.Experimental != nil {
		structMap["experimental"] = c.Experimental
	}
	if c.FansPluggable != nil {
		structMap["fans_pluggable"] = c.FansPluggable
	}
	if c.HasBgp != nil {
		structMap["has_bgp"] = c.HasBgp
	}
	if c.HasEts != nil {
		structMap["has_ets"] = c.HasEts
	}
	if c.HasEvpn != nil {
		structMap["has_evpn"] = c.HasEvpn
	}
	if c.HasIrb != nil {
		structMap["has_irb"] = c.HasIrb
	}
	if c.HasPoeOut != nil {
		structMap["has_poe_out"] = c.HasPoeOut
	}
	if c.HasSnapshot != nil {
		structMap["has_snapshot"] = c.HasSnapshot
	}
	if c.HasVc != nil {
		structMap["has_vc"] = c.HasVc
	}
	if c.Model != nil {
		structMap["model"] = c.Model
	}
	if c.Modular != nil {
		structMap["modular"] = c.Modular
	}
	if c.NoShapingRate != nil {
		structMap["no_shaping_rate"] = c.NoShapingRate
	}
	if c.NumberFans != nil {
		structMap["number_fans"] = c.NumberFans
	}
	if c.OcDevice != nil {
		structMap["oc_device"] = c.OcDevice
	}
	if c.OobInterface != nil {
		structMap["oob_interface"] = c.OobInterface
	}
	if c.PacketActionDropOnly != nil {
		structMap["packet_action_drop_only"] = c.PacketActionDropOnly
	}
	if c.Pic != nil {
		structMap["pic"] = c.Pic
	}
	if c.SubRequired != nil {
		structMap["sub_required"] = c.SubRequired
	}
	structMap["type"] = c.Type
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstDeviceSwitch.
// It customizes the JSON unmarshaling process for ConstDeviceSwitch objects.
func (c *ConstDeviceSwitch) UnmarshalJSON(input []byte) error {
	var temp tempConstDeviceSwitch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "alias", "defaults", "description", "display", "evolved_os", "evpn_ri_type", "experimental", "fans_pluggable", "has_bgp", "has_ets", "has_evpn", "has_irb", "has_poe_out", "has_snapshot", "has_vc", "model", "modular", "no_shaping_rate", "number_fans", "oc_device", "oob_interface", "packet_action_drop_only", "pic", "sub_required", "type")
	if err != nil {
		return err
	}

	c.AdditionalProperties = additionalProperties
	c.Alias = temp.Alias
	c.Defaults = temp.Defaults
	c.Description = temp.Description
	c.Display = temp.Display
	c.EvolvedOs = temp.EvolvedOs
	c.EvpnRiType = temp.EvpnRiType
	c.Experimental = temp.Experimental
	c.FansPluggable = temp.FansPluggable
	c.HasBgp = temp.HasBgp
	c.HasEts = temp.HasEts
	c.HasEvpn = temp.HasEvpn
	c.HasIrb = temp.HasIrb
	c.HasPoeOut = temp.HasPoeOut
	c.HasSnapshot = temp.HasSnapshot
	c.HasVc = temp.HasVc
	c.Model = temp.Model
	c.Modular = temp.Modular
	c.NoShapingRate = temp.NoShapingRate
	c.NumberFans = temp.NumberFans
	c.OcDevice = temp.OcDevice
	c.OobInterface = temp.OobInterface
	c.PacketActionDropOnly = temp.PacketActionDropOnly
	c.Pic = temp.Pic
	c.SubRequired = temp.SubRequired
	c.Type = *temp.Type
	return nil
}

// tempConstDeviceSwitch is a temporary struct used for validating the fields of ConstDeviceSwitch.
type tempConstDeviceSwitch struct {
	Alias                *string                   `json:"alias,omitempty"`
	Defaults             *ConstDeviceSwitchDefault `json:"defaults,omitempty"`
	Description          *string                   `json:"description,omitempty"`
	Display              *string                   `json:"display,omitempty"`
	EvolvedOs            *bool                     `json:"evolved_os,omitempty"`
	EvpnRiType           *string                   `json:"evpn_ri_type,omitempty"`
	Experimental         *bool                     `json:"experimental,omitempty"`
	FansPluggable        *bool                     `json:"fans_pluggable,omitempty"`
	HasBgp               *bool                     `json:"has_bgp,omitempty"`
	HasEts               *bool                     `json:"has_ets,omitempty"`
	HasEvpn              *bool                     `json:"has_evpn,omitempty"`
	HasIrb               *bool                     `json:"has_irb,omitempty"`
	HasPoeOut            *bool                     `json:"has_poe_out,omitempty"`
	HasSnapshot          *bool                     `json:"has_snapshot,omitempty"`
	HasVc                *bool                     `json:"has_vc,omitempty"`
	Model                *string                   `json:"model,omitempty"`
	Modular              *bool                     `json:"modular,omitempty"`
	NoShapingRate        *bool                     `json:"no_shaping_rate,omitempty"`
	NumberFans           *int                      `json:"number_fans,omitempty"`
	OcDevice             *bool                     `json:"oc_device,omitempty"`
	OobInterface         *string                   `json:"oob_interface,omitempty"`
	PacketActionDropOnly *bool                     `json:"packet_action_drop_only,omitempty"`
	Pic                  map[string]string         `json:"pic,omitempty"`
	SubRequired          *string                   `json:"sub_required,omitempty"`
	Type                 *string                   `json:"type"`
}

func (c *tempConstDeviceSwitch) validate() error {
	var errs []string
	if c.Type == nil {
		errs = append(errs, "required field `type` is missing for type `const_device_switch`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
