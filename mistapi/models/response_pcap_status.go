package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// ResponsePcapStatus represents a ResponsePcapStatus struct.
type ResponsePcapStatus struct {
    ApMac                     Optional[string]          `json:"ap_mac"`
    // List of target APs to capture packets
    Aps                       []string                  `json:"aps,omitempty"`
    ClientMac                 Optional[string]          `json:"client_mac"`
    Duration                  *int                      `json:"duration,omitempty"`
    // List of APs where configuration attempt failed
    Failed                    []string                  `json:"failed,omitempty"`
    // PCAP format. enum:
    // * `stream`: to Mist cloud
    // * `tzsp`: tream packets (over UDP as TZSP packets) to a remote host (typically running Wireshark)
    Format                    *CaptureMxedgeFormatEnum  `json:"format,omitempty"`
    // Information on gateways to capture packets on if a gateway capture type is specified
    Gateways                  []string                  `json:"gateways,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                        uuid.UUID                 `json:"id"`
    IncludesMcast             *bool                     `json:"includes_mcast,omitempty"`
    // Max number of packets configured by user
    MaxNumPackets             *int                      `json:"max_num_packets,omitempty"`
    MaxPktLen                 *int                      `json:"max_pkt_len,omitempty"`
    // nformation on mxedges to capture packets on if a mxedge capture type is specified
    Mxedges                   []string                  `json:"mxedges,omitempty"`
    // total number of packets captured by all AP, not applicable for type [client, new_assoc]
    NumPackets                *int                      `json:"num_packets,omitempty"`
    // List of target APs successfully configured to capture packets
    Ok                        []string                  `json:"ok,omitempty"`
    PcapAps                   map[string]ResponsePcapAp `json:"pcap_aps,omitempty"`
    // When `type`==`radiotap`, radiotap_tcpdump_expression expression provided by the user
    RadiotapTcpdumpExpression *string                   `json:"radiotap_tcpdump_expression,omitempty"`
    // When `type`==`scan`, scan_tcpdump_expression provided by the user
    ScanTcpdumpExpression     *string                   `json:"scan_tcpdump_expression,omitempty"`
    Ssid                      Optional[string]          `json:"ssid"`
    StartedTime               *int                      `json:"started_time,omitempty"`
    // Information on switches to capture packets on if a switch capture type is specified. irb port interface is automatically added to capture as needed to ensure all desired packets are captured.
    Switches                  []string                  `json:"switches,omitempty"`
    // tcpdump expression provided by the user (common)
    TcpdumpExpression         *string                   `json:"tcpdump_expression,omitempty"`
    // enum: `client`, `gateway`, `new_assoc`, `radiotap`, `radiotap,wired`, `wired`, `wireless`
    Type                      PcapTypeEnum              `json:"type"`
    // Required if `format`==`tzsp`. Remote host accessible to mxedges over the network for receiving the captured packets.
    TzspHost                  *string                   `json:"tzsp_host,omitempty"`
    // If `format`==`tzsp`. Port on remote host for receiving the captured packets
    TzspPort                  *int                      `json:"tzsp_port,omitempty"`
    // When `type`==`wired`, wired_tcpdump_expression provided by the user
    WiredTcpdumpExpression    *string                   `json:"wired_tcpdump_expression,omitempty"`
    // When `type`==`‘wireless’`, wireless_tcpdump_expression provided by the user
    WirelessTcpdumpExpression *string                   `json:"wireless_tcpdump_expression,omitempty"`
    AdditionalProperties      map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for ResponsePcapStatus,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponsePcapStatus) String() string {
    return fmt.Sprintf(
    	"ResponsePcapStatus[ApMac=%v, Aps=%v, ClientMac=%v, Duration=%v, Failed=%v, Format=%v, Gateways=%v, Id=%v, IncludesMcast=%v, MaxNumPackets=%v, MaxPktLen=%v, Mxedges=%v, NumPackets=%v, Ok=%v, PcapAps=%v, RadiotapTcpdumpExpression=%v, ScanTcpdumpExpression=%v, Ssid=%v, StartedTime=%v, Switches=%v, TcpdumpExpression=%v, Type=%v, TzspHost=%v, TzspPort=%v, WiredTcpdumpExpression=%v, WirelessTcpdumpExpression=%v, AdditionalProperties=%v]",
    	r.ApMac, r.Aps, r.ClientMac, r.Duration, r.Failed, r.Format, r.Gateways, r.Id, r.IncludesMcast, r.MaxNumPackets, r.MaxPktLen, r.Mxedges, r.NumPackets, r.Ok, r.PcapAps, r.RadiotapTcpdumpExpression, r.ScanTcpdumpExpression, r.Ssid, r.StartedTime, r.Switches, r.TcpdumpExpression, r.Type, r.TzspHost, r.TzspPort, r.WiredTcpdumpExpression, r.WirelessTcpdumpExpression, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponsePcapStatus.
// It customizes the JSON marshaling process for ResponsePcapStatus objects.
func (r ResponsePcapStatus) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "ap_mac", "aps", "client_mac", "duration", "failed", "format", "gateways", "id", "includes_mcast", "max_num_packets", "max_pkt_len", "mxedges", "num_packets", "ok", "pcap_aps", "radiotap_tcpdump_expression", "scan_tcpdump_expression", "ssid", "started_time", "switches", "tcpdump_expression", "type", "tzsp_host", "tzsp_port", "wired_tcpdump_expression", "wireless_tcpdump_expression"); err != nil {
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
    if r.MaxNumPackets != nil {
        structMap["max_num_packets"] = r.MaxNumPackets
    }
    if r.MaxPktLen != nil {
        structMap["max_pkt_len"] = r.MaxPktLen
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
    if r.PcapAps != nil {
        structMap["pcap_aps"] = r.PcapAps
    }
    if r.RadiotapTcpdumpExpression != nil {
        structMap["radiotap_tcpdump_expression"] = r.RadiotapTcpdumpExpression
    }
    if r.ScanTcpdumpExpression != nil {
        structMap["scan_tcpdump_expression"] = r.ScanTcpdumpExpression
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_mac", "aps", "client_mac", "duration", "failed", "format", "gateways", "id", "includes_mcast", "max_num_packets", "max_pkt_len", "mxedges", "num_packets", "ok", "pcap_aps", "radiotap_tcpdump_expression", "scan_tcpdump_expression", "ssid", "started_time", "switches", "tcpdump_expression", "type", "tzsp_host", "tzsp_port", "wired_tcpdump_expression", "wireless_tcpdump_expression")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.ApMac = temp.ApMac
    r.Aps = temp.Aps
    r.ClientMac = temp.ClientMac
    r.Duration = temp.Duration
    r.Failed = temp.Failed
    r.Format = temp.Format
    r.Gateways = temp.Gateways
    r.Id = *temp.Id
    r.IncludesMcast = temp.IncludesMcast
    r.MaxNumPackets = temp.MaxNumPackets
    r.MaxPktLen = temp.MaxPktLen
    r.Mxedges = temp.Mxedges
    r.NumPackets = temp.NumPackets
    r.Ok = temp.Ok
    r.PcapAps = temp.PcapAps
    r.RadiotapTcpdumpExpression = temp.RadiotapTcpdumpExpression
    r.ScanTcpdumpExpression = temp.ScanTcpdumpExpression
    r.Ssid = temp.Ssid
    r.StartedTime = temp.StartedTime
    r.Switches = temp.Switches
    r.TcpdumpExpression = temp.TcpdumpExpression
    r.Type = *temp.Type
    r.TzspHost = temp.TzspHost
    r.TzspPort = temp.TzspPort
    r.WiredTcpdumpExpression = temp.WiredTcpdumpExpression
    r.WirelessTcpdumpExpression = temp.WirelessTcpdumpExpression
    return nil
}

// tempResponsePcapStatus is a temporary struct used for validating the fields of ResponsePcapStatus.
type tempResponsePcapStatus  struct {
    ApMac                     Optional[string]          `json:"ap_mac"`
    Aps                       []string                  `json:"aps,omitempty"`
    ClientMac                 Optional[string]          `json:"client_mac"`
    Duration                  *int                      `json:"duration,omitempty"`
    Failed                    []string                  `json:"failed,omitempty"`
    Format                    *CaptureMxedgeFormatEnum  `json:"format,omitempty"`
    Gateways                  []string                  `json:"gateways,omitempty"`
    Id                        *uuid.UUID                `json:"id"`
    IncludesMcast             *bool                     `json:"includes_mcast,omitempty"`
    MaxNumPackets             *int                      `json:"max_num_packets,omitempty"`
    MaxPktLen                 *int                      `json:"max_pkt_len,omitempty"`
    Mxedges                   []string                  `json:"mxedges,omitempty"`
    NumPackets                *int                      `json:"num_packets,omitempty"`
    Ok                        []string                  `json:"ok,omitempty"`
    PcapAps                   map[string]ResponsePcapAp `json:"pcap_aps,omitempty"`
    RadiotapTcpdumpExpression *string                   `json:"radiotap_tcpdump_expression,omitempty"`
    ScanTcpdumpExpression     *string                   `json:"scan_tcpdump_expression,omitempty"`
    Ssid                      Optional[string]          `json:"ssid"`
    StartedTime               *int                      `json:"started_time,omitempty"`
    Switches                  []string                  `json:"switches,omitempty"`
    TcpdumpExpression         *string                   `json:"tcpdump_expression,omitempty"`
    Type                      *PcapTypeEnum             `json:"type"`
    TzspHost                  *string                   `json:"tzsp_host,omitempty"`
    TzspPort                  *int                      `json:"tzsp_port,omitempty"`
    WiredTcpdumpExpression    *string                   `json:"wired_tcpdump_expression,omitempty"`
    WirelessTcpdumpExpression *string                   `json:"wireless_tcpdump_expression,omitempty"`
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
    return errors.New(strings.Join (errs, "\n"))
}
