package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ConstDeviceGateway represents a ConstDeviceGateway struct.
type ConstDeviceGateway struct {
    // Object Key is the interface type name (e.g. "lan_ports", "wan_ports", ...)
    Defaults             map[string]string        `json:"defaults,omitempty"`
    Description          *string                  `json:"description,omitempty"`
    Experimental         *bool                    `json:"experimental,omitempty"`
    FansPluggable        *bool                    `json:"fans_pluggable,omitempty"`
    HaNode0Fpc           *int                     `json:"ha_node0_fpc,omitempty"`
    HaNode1Fpc           *int                     `json:"ha_node1_fpc,omitempty"`
    HasBgp               *bool                    `json:"has_bgp,omitempty"`
    HasFxp0              *bool                    `json:"has_fxp0,omitempty"`
    HasHaControl         *bool                    `json:"has_ha_control,omitempty"`
    HasHaData            *bool                    `json:"has_ha_data,omitempty"`
    HasIrb               *bool                    `json:"has_irb,omitempty"`
    HasPoeOut            *bool                    `json:"has_poe_out,omitempty"`
    HasSnapshot          *bool                    `json:"has_snapshot,omitempty"`
    IrbDisabledByDefault *bool                    `json:"irb_disabled_by_default,omitempty"`
    Model                *string                  `json:"model,omitempty"`
    NumberFans           *int                     `json:"number_fans,omitempty"`
    OcDevice             *bool                    `json:"oc_device,omitempty"`
    // Object Key is the PIC number
    Pic                  map[string]string        `json:"pic,omitempty"`
    // Object Key is the interface name (e.g. "ge-0/0/1", ...)
    Ports                *ConstDeviceGatewayPorts `json:"ports,omitempty"`
    SubRequired          *string                  `json:"sub_required,omitempty"`
    T128Device           *bool                    `json:"t128_device,omitempty"`
    // Device Type. enum: `gateway`
    Type                 string                   `json:"type"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for ConstDeviceGateway,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstDeviceGateway) String() string {
    return fmt.Sprintf(
    	"ConstDeviceGateway[Defaults=%v, Description=%v, Experimental=%v, FansPluggable=%v, HaNode0Fpc=%v, HaNode1Fpc=%v, HasBgp=%v, HasFxp0=%v, HasHaControl=%v, HasHaData=%v, HasIrb=%v, HasPoeOut=%v, HasSnapshot=%v, IrbDisabledByDefault=%v, Model=%v, NumberFans=%v, OcDevice=%v, Pic=%v, Ports=%v, SubRequired=%v, T128Device=%v, Type=%v, AdditionalProperties=%v]",
    	c.Defaults, c.Description, c.Experimental, c.FansPluggable, c.HaNode0Fpc, c.HaNode1Fpc, c.HasBgp, c.HasFxp0, c.HasHaControl, c.HasHaData, c.HasIrb, c.HasPoeOut, c.HasSnapshot, c.IrbDisabledByDefault, c.Model, c.NumberFans, c.OcDevice, c.Pic, c.Ports, c.SubRequired, c.T128Device, c.Type, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstDeviceGateway.
// It customizes the JSON marshaling process for ConstDeviceGateway objects.
func (c ConstDeviceGateway) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "defaults", "description", "experimental", "fans_pluggable", "ha_node0_fpc", "ha_node1_fpc", "has_bgp", "has_fxp0", "has_ha_control", "has_ha_data", "has_irb", "has_poe_out", "has_snapshot", "irb_disabled_by_default", "model", "number_fans", "oc_device", "pic", "ports", "sub_required", "t128_device", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstDeviceGateway object to a map representation for JSON marshaling.
func (c ConstDeviceGateway) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Defaults != nil {
        structMap["defaults"] = c.Defaults
    }
    if c.Description != nil {
        structMap["description"] = c.Description
    }
    if c.Experimental != nil {
        structMap["experimental"] = c.Experimental
    }
    if c.FansPluggable != nil {
        structMap["fans_pluggable"] = c.FansPluggable
    }
    if c.HaNode0Fpc != nil {
        structMap["ha_node0_fpc"] = c.HaNode0Fpc
    }
    if c.HaNode1Fpc != nil {
        structMap["ha_node1_fpc"] = c.HaNode1Fpc
    }
    if c.HasBgp != nil {
        structMap["has_bgp"] = c.HasBgp
    }
    if c.HasFxp0 != nil {
        structMap["has_fxp0"] = c.HasFxp0
    }
    if c.HasHaControl != nil {
        structMap["has_ha_control"] = c.HasHaControl
    }
    if c.HasHaData != nil {
        structMap["has_ha_data"] = c.HasHaData
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
    if c.IrbDisabledByDefault != nil {
        structMap["irb_disabled_by_default"] = c.IrbDisabledByDefault
    }
    if c.Model != nil {
        structMap["model"] = c.Model
    }
    if c.NumberFans != nil {
        structMap["number_fans"] = c.NumberFans
    }
    if c.OcDevice != nil {
        structMap["oc_device"] = c.OcDevice
    }
    if c.Pic != nil {
        structMap["pic"] = c.Pic
    }
    if c.Ports != nil {
        structMap["ports"] = c.Ports.toMap()
    }
    if c.SubRequired != nil {
        structMap["sub_required"] = c.SubRequired
    }
    if c.T128Device != nil {
        structMap["t128_device"] = c.T128Device
    }
    structMap["type"] = c.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstDeviceGateway.
// It customizes the JSON unmarshaling process for ConstDeviceGateway objects.
func (c *ConstDeviceGateway) UnmarshalJSON(input []byte) error {
    var temp tempConstDeviceGateway
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "defaults", "description", "experimental", "fans_pluggable", "ha_node0_fpc", "ha_node1_fpc", "has_bgp", "has_fxp0", "has_ha_control", "has_ha_data", "has_irb", "has_poe_out", "has_snapshot", "irb_disabled_by_default", "model", "number_fans", "oc_device", "pic", "ports", "sub_required", "t128_device", "type")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Defaults = temp.Defaults
    c.Description = temp.Description
    c.Experimental = temp.Experimental
    c.FansPluggable = temp.FansPluggable
    c.HaNode0Fpc = temp.HaNode0Fpc
    c.HaNode1Fpc = temp.HaNode1Fpc
    c.HasBgp = temp.HasBgp
    c.HasFxp0 = temp.HasFxp0
    c.HasHaControl = temp.HasHaControl
    c.HasHaData = temp.HasHaData
    c.HasIrb = temp.HasIrb
    c.HasPoeOut = temp.HasPoeOut
    c.HasSnapshot = temp.HasSnapshot
    c.IrbDisabledByDefault = temp.IrbDisabledByDefault
    c.Model = temp.Model
    c.NumberFans = temp.NumberFans
    c.OcDevice = temp.OcDevice
    c.Pic = temp.Pic
    c.Ports = temp.Ports
    c.SubRequired = temp.SubRequired
    c.T128Device = temp.T128Device
    c.Type = *temp.Type
    return nil
}

// tempConstDeviceGateway is a temporary struct used for validating the fields of ConstDeviceGateway.
type tempConstDeviceGateway  struct {
    Defaults             map[string]string        `json:"defaults,omitempty"`
    Description          *string                  `json:"description,omitempty"`
    Experimental         *bool                    `json:"experimental,omitempty"`
    FansPluggable        *bool                    `json:"fans_pluggable,omitempty"`
    HaNode0Fpc           *int                     `json:"ha_node0_fpc,omitempty"`
    HaNode1Fpc           *int                     `json:"ha_node1_fpc,omitempty"`
    HasBgp               *bool                    `json:"has_bgp,omitempty"`
    HasFxp0              *bool                    `json:"has_fxp0,omitempty"`
    HasHaControl         *bool                    `json:"has_ha_control,omitempty"`
    HasHaData            *bool                    `json:"has_ha_data,omitempty"`
    HasIrb               *bool                    `json:"has_irb,omitempty"`
    HasPoeOut            *bool                    `json:"has_poe_out,omitempty"`
    HasSnapshot          *bool                    `json:"has_snapshot,omitempty"`
    IrbDisabledByDefault *bool                    `json:"irb_disabled_by_default,omitempty"`
    Model                *string                  `json:"model,omitempty"`
    NumberFans           *int                     `json:"number_fans,omitempty"`
    OcDevice             *bool                    `json:"oc_device,omitempty"`
    Pic                  map[string]string        `json:"pic,omitempty"`
    Ports                *ConstDeviceGatewayPorts `json:"ports,omitempty"`
    SubRequired          *string                  `json:"sub_required,omitempty"`
    T128Device           *bool                    `json:"t128_device,omitempty"`
    Type                 *string                  `json:"type"`
}

func (c *tempConstDeviceGateway) validate() error {
    var errs []string
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `const_device_gateway`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
