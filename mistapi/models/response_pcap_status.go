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

// ResponsePcapStatus represents a ResponsePcapStatus struct.
// Current status of a packet capture session
type ResponsePcapStatus struct {
	// AP MAC address targeted by the packet capture, or null when no single AP filter is set
	ApMac Optional[string] `json:"ap_mac"`
	// List of target APs to capture packets
	Aps []string `json:"aps,omitempty"`
	// Client MAC address filter applied to the packet capture, or null when no client filter is used
	ClientMac Optional[string] `json:"client_mac"`
	// Configured packet capture duration, in seconds
	Duration *int `json:"duration,omitempty"`
	// Whether the packet capture session is currently enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Epoch timestamp, in seconds, when the capture session expires
	Expiry *float64 `json:"expiry,omitempty"`
	// List of APs where configuration attempt failed
	Failed []string `json:"failed,omitempty"`
	// PCAP format. enum:
	// * `stream`: to Mist cloud
	// * `tzsp`: stream packets (over UDP as TZSP packets) to a remote host (typically running Wireshark)
	Format *CaptureMxedgeFormatEnum `json:"format,omitempty"`
	// Information on gateways to capture packets on if a gateway capture type is specified
	Gateways []string `json:"gateways,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id uuid.UUID `json:"id"`
	// Whether multicast traffic is included in the packet capture
	IncludesMcast *bool `json:"includes_mcast,omitempty"`
	// Map of Mist Edge IDs that could not be configured for capture
	InvalidMxedges *interface{} `json:"invalid_mxedges,omitempty"`
	// Max number of packets configured by user
	MaxNumPackets *int `json:"max_num_packets,omitempty"`
	// Maximum number of bytes captured from each packet
	MaxPktLen *int `json:"max_pkt_len,omitempty"`
	// Number of Mist Edges in the capture session
	MxedgeCount *int `json:"mxedge_count,omitempty"`
	// Dict of Mist Edges to capture on, property key is the Mist Edge ID
	Mxedges map[string]ResponsePcapStatusMxedgesItem `json:"mxedges,omitempty"`
	// total number of packets captured by all AP, not applicable for type [client, new_assoc]
	NumPackets *int `json:"num_packets,omitempty"`
	// List of target APs successfully configured to capture packets
	Ok []string `json:"ok,omitempty"`
	// Unique identifier of a Mist organization
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// Per-AP radio capture settings keyed by AP MAC address
	PcapAps map[string]ResponsePcapAp `json:"pcap_aps,omitempty"`
	// When `type`==`radiotap`, radiotap_tcpdump_expression expression provided by the user
	RadiotapTcpdumpExpression *string `json:"radiotap_tcpdump_expression,omitempty"`
	// Whether raw packet data is included in the capture output
	Raw *bool `json:"raw,omitempty"`
	// When `type`==`scan`, scan_tcpdump_expression provided by the user
	ScanTcpdumpExpression *string `json:"scan_tcpdump_expression,omitempty"`
	// Unique identifier of a Mist site
	SiteId *uuid.UUID `json:"site_id,omitempty"`
	// Wireless network SSID filter applied to the packet capture, or null when no SSID filter is used
	Ssid Optional[string] `json:"ssid"`
	// Epoch timestamp, in seconds, when the capture session started
	StartedTime *int `json:"started_time,omitempty"`
	// Information on switches to capture packets on if a switch capture type is specified. irb port interface is automatically added to capture as needed to ensure all desired packets are captured.
	Switches []string `json:"switches,omitempty"`
	// tcpdump expression provided by the user (common)
	TcpdumpExpression *string `json:"tcpdump_expression,omitempty"`
	// Epoch timestamp, in seconds
	Timestamp *float64 `json:"timestamp,omitempty"`
	// enum: `client`, `gateway`, `new_assoc`, `radiotap`, `radiotap,wired`, `wired`, `wireless`
	Type PcapTypeEnum `json:"type"`
	// Required if `format`==`tzsp`. Remote host accessible to Mist Edges over the network for receiving the captured packets.
	TzspHost *string `json:"tzsp_host,omitempty"`
	// If `format`==`tzsp`. Port on remote host for receiving the captured packets
	TzspPort *int `json:"tzsp_port,omitempty"`
	// When `type`==`wired`, wired_tcpdump_expression provided by the user
	WiredTcpdumpExpression *string `json:"wired_tcpdump_expression,omitempty"`
	// When `type`==`‘wireless’`, wireless_tcpdump_expression provided by the user
	WirelessTcpdumpExpression *string                `json:"wireless_tcpdump_expression,omitempty"`
	AdditionalProperties      map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponsePcapStatus,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponsePcapStatus) String() string {
	return fmt.Sprintf(
		"ResponsePcapStatus[ApMac=%v, Aps=%v, ClientMac=%v, Duration=%v, Enabled=%v, Expiry=%v, Failed=%v, Format=%v, Gateways=%v, Id=%v, IncludesMcast=%v, InvalidMxedges=%v, MaxNumPackets=%v, MaxPktLen=%v, MxedgeCount=%v, Mxedges=%v, NumPackets=%v, Ok=%v, OrgId=%v, PcapAps=%v, RadiotapTcpdumpExpression=%v, Raw=%v, ScanTcpdumpExpression=%v, SiteId=%v, Ssid=%v, StartedTime=%v, Switches=%v, TcpdumpExpression=%v, Timestamp=%v, Type=%v, TzspHost=%v, TzspPort=%v, WiredTcpdumpExpression=%v, WirelessTcpdumpExpression=%v, AdditionalProperties=%v]",
		r.ApMac, r.Aps, r.ClientMac, r.Duration, r.Enabled, r.Expiry, r.Failed, r.Format, r.Gateways, r.Id, r.IncludesMcast, r.InvalidMxedges, r.MaxNumPackets, r.MaxPktLen, r.MxedgeCount, r.Mxedges, r.NumPackets, r.Ok, r.OrgId, r.PcapAps, r.RadiotapTcpdumpExpression, r.Raw, r.ScanTcpdumpExpression, r.SiteId, r.Ssid, r.StartedTime, r.Switches, r.TcpdumpExpression, r.Timestamp, r.Type, r.TzspHost, r.TzspPort, r.WiredTcpdumpExpression, r.WirelessTcpdumpExpression, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponsePcapStatus.
// It customizes the JSON marshaling process for ResponsePcapStatus objects.
func (r ResponsePcapStatus) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"ap_mac", "aps", "client_mac", "duration", "enabled", "expiry", "failed", "format", "gateways", "id", "includes_mcast", "invalid_mxedges", "max_num_packets", "max_pkt_len", "mxedge_count", "mxedges", "num_packets", "ok", "org_id", "pcap_aps", "radiotap_tcpdump_expression", "raw", "scan_tcpdump_expression", "site_id", "ssid", "started_time", "switches", "tcpdump_expression", "timestamp", "type", "tzsp_host", "tzsp_port", "wired_tcpdump_expression", "wireless_tcpdump_expression"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponsePcapStatus object to a map representation for JSON marshaling.
func (r ResponsePcapStatus) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.ApMac.IsValueSet() {
		if r.ApMac.Value() != nil {
			structMap["ap_mac"] = r.ApMac.Value()
		} else {
			structMap["ap_mac"] = nil
		}
	}
	if r.Aps != nil {
		structMap["aps"] = r.Aps
	}
	if r.ClientMac.IsValueSet() {
		if r.ClientMac.Value() != nil {
			structMap["client_mac"] = r.ClientMac.Value()
		} else {
			structMap["client_mac"] = nil
		}
	}
	if r.Duration != nil {
		structMap["duration"] = r.Duration
	}
	if r.Enabled != nil {
		structMap["enabled"] = r.Enabled
	}
	if r.Expiry != nil {
		structMap["expiry"] = r.Expiry
	}
	if r.Failed != nil {
		structMap["failed"] = r.Failed
	}
	if r.Format != nil {
		structMap["format"] = r.Format
	}
	if r.Gateways != nil {
		structMap["gateways"] = r.Gateways
	}
	structMap["id"] = r.Id
	if r.IncludesMcast != nil {
		structMap["includes_mcast"] = r.IncludesMcast
	}
	if r.InvalidMxedges != nil {
		structMap["invalid_mxedges"] = r.InvalidMxedges
	}
	if r.MaxNumPackets != nil {
		structMap["max_num_packets"] = r.MaxNumPackets
	}
	if r.MaxPktLen != nil {
		structMap["max_pkt_len"] = r.MaxPktLen
	}
	if r.MxedgeCount != nil {
		structMap["mxedge_count"] = r.MxedgeCount
	}
	if r.Mxedges != nil {
		structMap["mxedges"] = r.Mxedges
	}
	if r.NumPackets != nil {
		structMap["num_packets"] = r.NumPackets
	}
	if r.Ok != nil {
		structMap["ok"] = r.Ok
	}
	if r.OrgId != nil {
		structMap["org_id"] = r.OrgId
	}
	if r.PcapAps != nil {
		structMap["pcap_aps"] = r.PcapAps
	}
	if r.RadiotapTcpdumpExpression != nil {
		structMap["radiotap_tcpdump_expression"] = r.RadiotapTcpdumpExpression
	}
	if r.Raw != nil {
		structMap["raw"] = r.Raw
	}
	if r.ScanTcpdumpExpression != nil {
		structMap["scan_tcpdump_expression"] = r.ScanTcpdumpExpression
	}
	if r.SiteId != nil {
		structMap["site_id"] = r.SiteId
	}
	if r.Ssid.IsValueSet() {
		if r.Ssid.Value() != nil {
			structMap["ssid"] = r.Ssid.Value()
		} else {
			structMap["ssid"] = nil
		}
	}
	if r.StartedTime != nil {
		structMap["started_time"] = r.StartedTime
	}
	if r.Switches != nil {
		structMap["switches"] = r.Switches
	}
	if r.TcpdumpExpression != nil {
		structMap["tcpdump_expression"] = r.TcpdumpExpression
	}
	if r.Timestamp != nil {
		structMap["timestamp"] = r.Timestamp
	}
	structMap["type"] = r.Type
	if r.TzspHost != nil {
		structMap["tzsp_host"] = r.TzspHost
	}
	if r.TzspPort != nil {
		structMap["tzsp_port"] = r.TzspPort
	}
	if r.WiredTcpdumpExpression != nil {
		structMap["wired_tcpdump_expression"] = r.WiredTcpdumpExpression
	}
	if r.WirelessTcpdumpExpression != nil {
		structMap["wireless_tcpdump_expression"] = r.WirelessTcpdumpExpression
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponsePcapStatus.
// It customizes the JSON unmarshaling process for ResponsePcapStatus objects.
func (r *ResponsePcapStatus) UnmarshalJSON(input []byte) error {
	var temp tempResponsePcapStatus
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_mac", "aps", "client_mac", "duration", "enabled", "expiry", "failed", "format", "gateways", "id", "includes_mcast", "invalid_mxedges", "max_num_packets", "max_pkt_len", "mxedge_count", "mxedges", "num_packets", "ok", "org_id", "pcap_aps", "radiotap_tcpdump_expression", "raw", "scan_tcpdump_expression", "site_id", "ssid", "started_time", "switches", "tcpdump_expression", "timestamp", "type", "tzsp_host", "tzsp_port", "wired_tcpdump_expression", "wireless_tcpdump_expression")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.ApMac = temp.ApMac
	r.Aps = temp.Aps
	r.ClientMac = temp.ClientMac
	r.Duration = temp.Duration
	r.Enabled = temp.Enabled
	r.Expiry = temp.Expiry
	r.Failed = temp.Failed
	r.Format = temp.Format
	r.Gateways = temp.Gateways
	r.Id = *temp.Id
	r.IncludesMcast = temp.IncludesMcast
	r.InvalidMxedges = temp.InvalidMxedges
	r.MaxNumPackets = temp.MaxNumPackets
	r.MaxPktLen = temp.MaxPktLen
	r.MxedgeCount = temp.MxedgeCount
	r.Mxedges = temp.Mxedges
	r.NumPackets = temp.NumPackets
	r.Ok = temp.Ok
	r.OrgId = temp.OrgId
	r.PcapAps = temp.PcapAps
	r.RadiotapTcpdumpExpression = temp.RadiotapTcpdumpExpression
	r.Raw = temp.Raw
	r.ScanTcpdumpExpression = temp.ScanTcpdumpExpression
	r.SiteId = temp.SiteId
	r.Ssid = temp.Ssid
	r.StartedTime = temp.StartedTime
	r.Switches = temp.Switches
	r.TcpdumpExpression = temp.TcpdumpExpression
	r.Timestamp = temp.Timestamp
	r.Type = *temp.Type
	r.TzspHost = temp.TzspHost
	r.TzspPort = temp.TzspPort
	r.WiredTcpdumpExpression = temp.WiredTcpdumpExpression
	r.WirelessTcpdumpExpression = temp.WirelessTcpdumpExpression
	return nil
}

// tempResponsePcapStatus is a temporary struct used for validating the fields of ResponsePcapStatus.
type tempResponsePcapStatus struct {
	ApMac                     Optional[string]                         `json:"ap_mac"`
	Aps                       []string                                 `json:"aps,omitempty"`
	ClientMac                 Optional[string]                         `json:"client_mac"`
	Duration                  *int                                     `json:"duration,omitempty"`
	Enabled                   *bool                                    `json:"enabled,omitempty"`
	Expiry                    *float64                                 `json:"expiry,omitempty"`
	Failed                    []string                                 `json:"failed,omitempty"`
	Format                    *CaptureMxedgeFormatEnum                 `json:"format,omitempty"`
	Gateways                  []string                                 `json:"gateways,omitempty"`
	Id                        *uuid.UUID                               `json:"id"`
	IncludesMcast             *bool                                    `json:"includes_mcast,omitempty"`
	InvalidMxedges            *interface{}                             `json:"invalid_mxedges,omitempty"`
	MaxNumPackets             *int                                     `json:"max_num_packets,omitempty"`
	MaxPktLen                 *int                                     `json:"max_pkt_len,omitempty"`
	MxedgeCount               *int                                     `json:"mxedge_count,omitempty"`
	Mxedges                   map[string]ResponsePcapStatusMxedgesItem `json:"mxedges,omitempty"`
	NumPackets                *int                                     `json:"num_packets,omitempty"`
	Ok                        []string                                 `json:"ok,omitempty"`
	OrgId                     *uuid.UUID                               `json:"org_id,omitempty"`
	PcapAps                   map[string]ResponsePcapAp                `json:"pcap_aps,omitempty"`
	RadiotapTcpdumpExpression *string                                  `json:"radiotap_tcpdump_expression,omitempty"`
	Raw                       *bool                                    `json:"raw,omitempty"`
	ScanTcpdumpExpression     *string                                  `json:"scan_tcpdump_expression,omitempty"`
	SiteId                    *uuid.UUID                               `json:"site_id,omitempty"`
	Ssid                      Optional[string]                         `json:"ssid"`
	StartedTime               *int                                     `json:"started_time,omitempty"`
	Switches                  []string                                 `json:"switches,omitempty"`
	TcpdumpExpression         *string                                  `json:"tcpdump_expression,omitempty"`
	Timestamp                 *float64                                 `json:"timestamp,omitempty"`
	Type                      *PcapTypeEnum                            `json:"type"`
	TzspHost                  *string                                  `json:"tzsp_host,omitempty"`
	TzspPort                  *int                                     `json:"tzsp_port,omitempty"`
	WiredTcpdumpExpression    *string                                  `json:"wired_tcpdump_expression,omitempty"`
	WirelessTcpdumpExpression *string                                  `json:"wireless_tcpdump_expression,omitempty"`
}

func (r *tempResponsePcapStatus) validate() error {
	var errs []string
	if r.Id == nil {
		errs = append(errs, "required field `id` is missing for type `response_pcap_status`")
	}
	if r.Type == nil {
		errs = append(errs, "required field `type` is missing for type `response_pcap_status`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
