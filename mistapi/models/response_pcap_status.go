package models

import (
    "encoding/json"
    "errors"
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
    // List of target Gateways to capture packets if a gateway capture type is specified
    Gateways                  []string                  `json:"gateways,omitempty"`
    // unique id for the capture
    Id                        uuid.UUID                 `json:"id"`
    IncludesMcast             *bool                     `json:"includes_mcast,omitempty"`
    // max number of packets configured by user
    MaxNumPackets             *int                      `json:"max_num_packets,omitempty"`
    MaxPktLen                 *int                      `json:"max_pkt_len,omitempty"`
    // total number of packets captured by all AP, not applicable for type [client, new_assoc]
    NumPackets                *int                      `json:"num_packets,omitempty"`
    // List of target APs successfully configured to capture packets
    Ok                        []string                  `json:"ok,omitempty"`
    PcapAps                   map[string]ResponsePcapAp `json:"pcap_aps,omitempty"`
    // when `type`==`radiotap`, radiotap_tcpdump_expression expression provided by the user
    RadiotapTcpdumpExpression *string                   `json:"radiotap_tcpdump_expression,omitempty"`
    // when `type`==`scan`, scan_tcpdump_expression provided by the user
    ScanTcpdumpExpression     *string                   `json:"scan_tcpdump_expression,omitempty"`
    Ssid                      Optional[string]          `json:"ssid"`
    StartedTime               *int                      `json:"started_time,omitempty"`
    // List of target Switches to capture packets if a switch capture type is specified
    Switches                  []string                  `json:"switches,omitempty"`
    // tcpdump expression provided by the user (common)
    TcpdumpExpression         *string                   `json:"tcpdump_expression,omitempty"`
    // enum: `client`, `gateway`, `new_assoc`, `radiotap`, `radiotap,wired`, `wired`, `wireless`
    Type                      PcapTypeEnum              `json:"type"`
    // when `type`==`wired`, wired_tcpdump_expression provided by the user
    WiredTcpdumpExpression    *string                   `json:"wired_tcpdump_expression,omitempty"`
    // when `type`==`‘wireless’`, wireless_tcpdump_expression provided by the user
    WirelessTcpdumpExpression *string                   `json:"wireless_tcpdump_expression,omitempty"`
    AdditionalProperties      map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponsePcapStatus.
// It customizes the JSON marshaling process for ResponsePcapStatus objects.
func (r ResponsePcapStatus) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponsePcapStatus object to a map representation for JSON marshaling.
func (r ResponsePcapStatus) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
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
    var temp responsePcapStatus
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_mac", "aps", "client_mac", "duration", "failed", "gateways", "id", "includes_mcast", "max_num_packets", "max_pkt_len", "num_packets", "ok", "pcap_aps", "radiotap_tcpdump_expression", "scan_tcpdump_expression", "ssid", "started_time", "switches", "tcpdump_expression", "type", "wired_tcpdump_expression", "wireless_tcpdump_expression")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.ApMac = temp.ApMac
    r.Aps = temp.Aps
    r.ClientMac = temp.ClientMac
    r.Duration = temp.Duration
    r.Failed = temp.Failed
    r.Gateways = temp.Gateways
    r.Id = *temp.Id
    r.IncludesMcast = temp.IncludesMcast
    r.MaxNumPackets = temp.MaxNumPackets
    r.MaxPktLen = temp.MaxPktLen
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
    r.WiredTcpdumpExpression = temp.WiredTcpdumpExpression
    r.WirelessTcpdumpExpression = temp.WirelessTcpdumpExpression
    return nil
}

// responsePcapStatus is a temporary struct used for validating the fields of ResponsePcapStatus.
type responsePcapStatus  struct {
    ApMac                     Optional[string]          `json:"ap_mac"`
    Aps                       []string                  `json:"aps,omitempty"`
    ClientMac                 Optional[string]          `json:"client_mac"`
    Duration                  *int                      `json:"duration,omitempty"`
    Failed                    []string                  `json:"failed,omitempty"`
    Gateways                  []string                  `json:"gateways,omitempty"`
    Id                        *uuid.UUID                `json:"id"`
    IncludesMcast             *bool                     `json:"includes_mcast,omitempty"`
    MaxNumPackets             *int                      `json:"max_num_packets,omitempty"`
    MaxPktLen                 *int                      `json:"max_pkt_len,omitempty"`
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
    WiredTcpdumpExpression    *string                   `json:"wired_tcpdump_expression,omitempty"`
    WirelessTcpdumpExpression *string                   `json:"wireless_tcpdump_expression,omitempty"`
}

func (r *responsePcapStatus) validate() error {
    var errs []string
    if r.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Response_Pcap_Status`")
    }
    if r.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Response_Pcap_Status`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
