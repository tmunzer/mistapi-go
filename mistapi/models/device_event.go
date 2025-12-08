// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// DeviceEvent represents a DeviceEvent struct.
type DeviceEvent struct {
	// (will be deprecated soon; please use mac instead) ap mac
	Ap *string `json:"ap,omitempty"`
	// (will be deprecated soon; please use device_name instead) ap name
	ApName *string `json:"ap_name,omitempty"`
	Apfw   *string `json:"apfw,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	AuditId    *uuid.UUID `json:"audit_id,omitempty"`
	Bandwidth  *int       `json:"bandwidth,omitempty"`
	Channel    *int       `json:"channel,omitempty"`
	ChassisMac *string    `json:"chassis_mac,omitempty"`
	Count      *int       `json:"count,omitempty"`
	// Device name
	DeviceName *string `json:"device_name,omitempty"`
	// enum: `ap`, `gateway`, `switch`
	DeviceType *DeviceTypeEnum `json:"device_type,omitempty"`
	// (optional) event advisory. enum: `notice`, `warn`
	EvType *WebhookDeviceEventsEventEvTypeEnum `json:"ev_type,omitempty"`
	ExtIp  *string                             `json:"ext_ip,omitempty"`
	// Device mac
	Mac          *string   `json:"mac,omitempty"`
	Model        *string   `json:"model,omitempty"`
	Node         *string   `json:"node,omitempty"`
	OrgId        uuid.UUID `json:"org_id"`
	PortId       *string   `json:"port_id,omitempty"`
	Power        *int      `json:"power,omitempty"`
	PreBandwidth *int      `json:"pre_bandwidth,omitempty"`
	PreChannel   *int      `json:"pre_channel,omitempty"`
	PrePower     *int      `json:"pre_power,omitempty"`
	PreUsage     *int      `json:"pre_usage,omitempty"`
	// (optional) event reason
	Reason *string    `json:"reason,omitempty"`
	SiteId *uuid.UUID `json:"site_id,omitempty"`
	// Site name
	SiteName *string `json:"site_name,omitempty"`
	// (optional) event description
	Text *string `json:"text,omitempty"`
	// Epoch (seconds)
	Timestamp float64 `json:"timestamp"`
	// Event type
	Type                 string                 `json:"type"`
	Usage                *int                   `json:"usage,omitempty"`
	Version              *string                `json:"version,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DeviceEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DeviceEvent) String() string {
	return fmt.Sprintf(
		"DeviceEvent[Ap=%v, ApName=%v, Apfw=%v, AuditId=%v, Bandwidth=%v, Channel=%v, ChassisMac=%v, Count=%v, DeviceName=%v, DeviceType=%v, EvType=%v, ExtIp=%v, Mac=%v, Model=%v, Node=%v, OrgId=%v, PortId=%v, Power=%v, PreBandwidth=%v, PreChannel=%v, PrePower=%v, PreUsage=%v, Reason=%v, SiteId=%v, SiteName=%v, Text=%v, Timestamp=%v, Type=%v, Usage=%v, Version=%v, AdditionalProperties=%v]",
		d.Ap, d.ApName, d.Apfw, d.AuditId, d.Bandwidth, d.Channel, d.ChassisMac, d.Count, d.DeviceName, d.DeviceType, d.EvType, d.ExtIp, d.Mac, d.Model, d.Node, d.OrgId, d.PortId, d.Power, d.PreBandwidth, d.PreChannel, d.PrePower, d.PreUsage, d.Reason, d.SiteId, d.SiteName, d.Text, d.Timestamp, d.Type, d.Usage, d.Version, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DeviceEvent.
// It customizes the JSON marshaling process for DeviceEvent objects.
func (d DeviceEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(d.AdditionalProperties,
		"ap", "ap_name", "apfw", "audit_id", "bandwidth", "channel", "chassis_mac", "count", "device_name", "device_type", "ev_type", "ext_ip", "mac", "model", "node", "org_id", "port_id", "power", "pre_bandwidth", "pre_channel", "pre_power", "pre_usage", "reason", "site_id", "site_name", "text", "timestamp", "type", "usage", "version"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(d.toMap())
}

// toMap converts the DeviceEvent object to a map representation for JSON marshaling.
func (d DeviceEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, d.AdditionalProperties)
	if d.Ap != nil {
		structMap["ap"] = d.Ap
	}
	if d.ApName != nil {
		structMap["ap_name"] = d.ApName
	}
	if d.Apfw != nil {
		structMap["apfw"] = d.Apfw
	}
	if d.AuditId != nil {
		structMap["audit_id"] = d.AuditId
	}
	if d.Bandwidth != nil {
		structMap["bandwidth"] = d.Bandwidth
	}
	if d.Channel != nil {
		structMap["channel"] = d.Channel
	}
	if d.ChassisMac != nil {
		structMap["chassis_mac"] = d.ChassisMac
	}
	if d.Count != nil {
		structMap["count"] = d.Count
	}
	if d.DeviceName != nil {
		structMap["device_name"] = d.DeviceName
	}
	if d.DeviceType != nil {
		structMap["device_type"] = d.DeviceType
	}
	if d.EvType != nil {
		structMap["ev_type"] = d.EvType
	}
	if d.ExtIp != nil {
		structMap["ext_ip"] = d.ExtIp
	}
	if d.Mac != nil {
		structMap["mac"] = d.Mac
	}
	if d.Model != nil {
		structMap["model"] = d.Model
	}
	if d.Node != nil {
		structMap["node"] = d.Node
	}
	structMap["org_id"] = d.OrgId
	if d.PortId != nil {
		structMap["port_id"] = d.PortId
	}
	if d.Power != nil {
		structMap["power"] = d.Power
	}
	if d.PreBandwidth != nil {
		structMap["pre_bandwidth"] = d.PreBandwidth
	}
	if d.PreChannel != nil {
		structMap["pre_channel"] = d.PreChannel
	}
	if d.PrePower != nil {
		structMap["pre_power"] = d.PrePower
	}
	if d.PreUsage != nil {
		structMap["pre_usage"] = d.PreUsage
	}
	if d.Reason != nil {
		structMap["reason"] = d.Reason
	}
	if d.SiteId != nil {
		structMap["site_id"] = d.SiteId
	}
	if d.SiteName != nil {
		structMap["site_name"] = d.SiteName
	}
	if d.Text != nil {
		structMap["text"] = d.Text
	}
	structMap["timestamp"] = d.Timestamp
	structMap["type"] = d.Type
	if d.Usage != nil {
		structMap["usage"] = d.Usage
	}
	if d.Version != nil {
		structMap["version"] = d.Version
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceEvent.
// It customizes the JSON unmarshaling process for DeviceEvent objects.
func (d *DeviceEvent) UnmarshalJSON(input []byte) error {
	var temp tempDeviceEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "ap_name", "apfw", "audit_id", "bandwidth", "channel", "chassis_mac", "count", "device_name", "device_type", "ev_type", "ext_ip", "mac", "model", "node", "org_id", "port_id", "power", "pre_bandwidth", "pre_channel", "pre_power", "pre_usage", "reason", "site_id", "site_name", "text", "timestamp", "type", "usage", "version")
	if err != nil {
		return err
	}
	d.AdditionalProperties = additionalProperties

	d.Ap = temp.Ap
	d.ApName = temp.ApName
	d.Apfw = temp.Apfw
	d.AuditId = temp.AuditId
	d.Bandwidth = temp.Bandwidth
	d.Channel = temp.Channel
	d.ChassisMac = temp.ChassisMac
	d.Count = temp.Count
	d.DeviceName = temp.DeviceName
	d.DeviceType = temp.DeviceType
	d.EvType = temp.EvType
	d.ExtIp = temp.ExtIp
	d.Mac = temp.Mac
	d.Model = temp.Model
	d.Node = temp.Node
	d.OrgId = *temp.OrgId
	d.PortId = temp.PortId
	d.Power = temp.Power
	d.PreBandwidth = temp.PreBandwidth
	d.PreChannel = temp.PreChannel
	d.PrePower = temp.PrePower
	d.PreUsage = temp.PreUsage
	d.Reason = temp.Reason
	d.SiteId = temp.SiteId
	d.SiteName = temp.SiteName
	d.Text = temp.Text
	d.Timestamp = *temp.Timestamp
	d.Type = *temp.Type
	d.Usage = temp.Usage
	d.Version = temp.Version
	return nil
}

// tempDeviceEvent is a temporary struct used for validating the fields of DeviceEvent.
type tempDeviceEvent struct {
	Ap           *string                             `json:"ap,omitempty"`
	ApName       *string                             `json:"ap_name,omitempty"`
	Apfw         *string                             `json:"apfw,omitempty"`
	AuditId      *uuid.UUID                          `json:"audit_id,omitempty"`
	Bandwidth    *int                                `json:"bandwidth,omitempty"`
	Channel      *int                                `json:"channel,omitempty"`
	ChassisMac   *string                             `json:"chassis_mac,omitempty"`
	Count        *int                                `json:"count,omitempty"`
	DeviceName   *string                             `json:"device_name,omitempty"`
	DeviceType   *DeviceTypeEnum                     `json:"device_type,omitempty"`
	EvType       *WebhookDeviceEventsEventEvTypeEnum `json:"ev_type,omitempty"`
	ExtIp        *string                             `json:"ext_ip,omitempty"`
	Mac          *string                             `json:"mac,omitempty"`
	Model        *string                             `json:"model,omitempty"`
	Node         *string                             `json:"node,omitempty"`
	OrgId        *uuid.UUID                          `json:"org_id"`
	PortId       *string                             `json:"port_id,omitempty"`
	Power        *int                                `json:"power,omitempty"`
	PreBandwidth *int                                `json:"pre_bandwidth,omitempty"`
	PreChannel   *int                                `json:"pre_channel,omitempty"`
	PrePower     *int                                `json:"pre_power,omitempty"`
	PreUsage     *int                                `json:"pre_usage,omitempty"`
	Reason       *string                             `json:"reason,omitempty"`
	SiteId       *uuid.UUID                          `json:"site_id,omitempty"`
	SiteName     *string                             `json:"site_name,omitempty"`
	Text         *string                             `json:"text,omitempty"`
	Timestamp    *float64                            `json:"timestamp"`
	Type         *string                             `json:"type"`
	Usage        *int                                `json:"usage,omitempty"`
	Version      *string                             `json:"version,omitempty"`
}

func (d *tempDeviceEvent) validate() error {
	var errs []string
	if d.OrgId == nil {
		errs = append(errs, "required field `org_id` is missing for type `device_event`")
	}
	if d.Timestamp == nil {
		errs = append(errs, "required field `timestamp` is missing for type `device_event`")
	}
	if d.Type == nil {
		errs = append(errs, "required field `type` is missing for type `device_event`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
