package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// ResponsePcapStart represents a ResponsePcapStart struct.
type ResponsePcapStart struct {
    ApCount                 *int                   `json:"ap_count,omitempty"`
    Aps                     []string               `json:"aps,omitempty"`
    ClientMac               Optional[string]       `json:"client_mac"`
    Duration                *float64               `json:"duration,omitempty"`
    Enabled                 *bool                  `json:"enabled,omitempty"`
    Expiry                  *float64               `json:"expiry,omitempty"`
    Format                  *string                `json:"format,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                      uuid.UUID              `json:"id"`
    IncludeMcast            *bool                  `json:"include_mcast,omitempty"`
    MaxPktLen               *int                   `json:"max_pkt_len,omitempty"`
    NumPackets              *int                   `json:"num_packets,omitempty"`
    OrgId                   uuid.UUID              `json:"org_id"`
    Raw                     *bool                  `json:"raw,omitempty"`
    SiteId                  uuid.UUID              `json:"site_id"`
    Ssid                    Optional[string]       `json:"ssid"`
    TcpdumpParserExpression Optional[string]       `json:"tcpdump_parser_expression"`
    Timestamp               float64                `json:"timestamp"`
    Type                    string                 `json:"type"`
    AdditionalProperties    map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponsePcapStart.
// It customizes the JSON marshaling process for ResponsePcapStart objects.
func (r ResponsePcapStart) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "ap_count", "aps", "client_mac", "duration", "enabled", "expiry", "format", "id", "include_mcast", "max_pkt_len", "num_packets", "org_id", "raw", "site_id", "ssid", "tcpdump_parser_expression", "timestamp", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponsePcapStart object to a map representation for JSON marshaling.
func (r ResponsePcapStart) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.ApCount != nil {
        structMap["ap_count"] = r.ApCount
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
    if r.Format != nil {
        structMap["format"] = r.Format
    }
    structMap["id"] = r.Id
    if r.IncludeMcast != nil {
        structMap["include_mcast"] = r.IncludeMcast
    }
    if r.MaxPktLen != nil {
        structMap["max_pkt_len"] = r.MaxPktLen
    }
    if r.NumPackets != nil {
        structMap["num_packets"] = r.NumPackets
    }
    structMap["org_id"] = r.OrgId
    if r.Raw != nil {
        structMap["raw"] = r.Raw
    }
    structMap["site_id"] = r.SiteId
    if r.Ssid.IsValueSet() {
        if r.Ssid.Value() != nil {
            structMap["ssid"] = r.Ssid.Value()
        } else {
            structMap["ssid"] = nil
        }
    }
    if r.TcpdumpParserExpression.IsValueSet() {
        if r.TcpdumpParserExpression.Value() != nil {
            structMap["tcpdump_parser_expression"] = r.TcpdumpParserExpression.Value()
        } else {
            structMap["tcpdump_parser_expression"] = nil
        }
    }
    structMap["timestamp"] = r.Timestamp
    structMap["type"] = r.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponsePcapStart.
// It customizes the JSON unmarshaling process for ResponsePcapStart objects.
func (r *ResponsePcapStart) UnmarshalJSON(input []byte) error {
    var temp tempResponsePcapStart
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_count", "aps", "client_mac", "duration", "enabled", "expiry", "format", "id", "include_mcast", "max_pkt_len", "num_packets", "org_id", "raw", "site_id", "ssid", "tcpdump_parser_expression", "timestamp", "type")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.ApCount = temp.ApCount
    r.Aps = temp.Aps
    r.ClientMac = temp.ClientMac
    r.Duration = temp.Duration
    r.Enabled = temp.Enabled
    r.Expiry = temp.Expiry
    r.Format = temp.Format
    r.Id = *temp.Id
    r.IncludeMcast = temp.IncludeMcast
    r.MaxPktLen = temp.MaxPktLen
    r.NumPackets = temp.NumPackets
    r.OrgId = *temp.OrgId
    r.Raw = temp.Raw
    r.SiteId = *temp.SiteId
    r.Ssid = temp.Ssid
    r.TcpdumpParserExpression = temp.TcpdumpParserExpression
    r.Timestamp = *temp.Timestamp
    r.Type = *temp.Type
    return nil
}

// tempResponsePcapStart is a temporary struct used for validating the fields of ResponsePcapStart.
type tempResponsePcapStart  struct {
    ApCount                 *int             `json:"ap_count,omitempty"`
    Aps                     []string         `json:"aps,omitempty"`
    ClientMac               Optional[string] `json:"client_mac"`
    Duration                *float64         `json:"duration,omitempty"`
    Enabled                 *bool            `json:"enabled,omitempty"`
    Expiry                  *float64         `json:"expiry,omitempty"`
    Format                  *string          `json:"format,omitempty"`
    Id                      *uuid.UUID       `json:"id"`
    IncludeMcast            *bool            `json:"include_mcast,omitempty"`
    MaxPktLen               *int             `json:"max_pkt_len,omitempty"`
    NumPackets              *int             `json:"num_packets,omitempty"`
    OrgId                   *uuid.UUID       `json:"org_id"`
    Raw                     *bool            `json:"raw,omitempty"`
    SiteId                  *uuid.UUID       `json:"site_id"`
    Ssid                    Optional[string] `json:"ssid"`
    TcpdumpParserExpression Optional[string] `json:"tcpdump_parser_expression"`
    Timestamp               *float64         `json:"timestamp"`
    Type                    *string          `json:"type"`
}

func (r *tempResponsePcapStart) validate() error {
    var errs []string
    if r.Id == nil {
        errs = append(errs, "required field `id` is missing for type `response_pcap_start`")
    }
    if r.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `response_pcap_start`")
    }
    if r.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `response_pcap_start`")
    }
    if r.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `response_pcap_start`")
    }
    if r.Type == nil {
        errs = append(errs, "required field `type` is missing for type `response_pcap_start`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
