package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// ResponsePcapSearchItem represents a ResponsePcapSearchItem struct.
type ResponsePcapSearchItem struct {
    ApMacs               []string                                     `json:"ap_macs,omitempty"`
    Aps                  []string                                     `json:"aps,omitempty"`
    Duration             *float64                                     `json:"duration,omitempty"`
    Format               *string                                      `json:"format,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID                                   `json:"id,omitempty"`
    MaxNumPackets        *float64                                     `json:"max_num_packets,omitempty"`
    OrgId                *uuid.UUID                                   `json:"org_id,omitempty"`
    PcapAps              map[string]ResponsePcapSearchItemPcapApsItem `json:"pcap_aps,omitempty"`
    PcapUrl              *string                                      `json:"pcap_url,omitempty"`
    SiteId               *uuid.UUID                                   `json:"site_id,omitempty"`
    TerminationReason    *string                                      `json:"termination_reason,omitempty"`
    Timestamp            float64                                      `json:"timestamp"`
    Type                 string                                       `json:"type"`
    Url                  string                                       `json:"url"`
    AdditionalProperties map[string]interface{}                       `json:"_"`
}

// String implements the fmt.Stringer interface for ResponsePcapSearchItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponsePcapSearchItem) String() string {
    return fmt.Sprintf(
    	"ResponsePcapSearchItem[ApMacs=%v, Aps=%v, Duration=%v, Format=%v, Id=%v, MaxNumPackets=%v, OrgId=%v, PcapAps=%v, PcapUrl=%v, SiteId=%v, TerminationReason=%v, Timestamp=%v, Type=%v, Url=%v, AdditionalProperties=%v]",
    	r.ApMacs, r.Aps, r.Duration, r.Format, r.Id, r.MaxNumPackets, r.OrgId, r.PcapAps, r.PcapUrl, r.SiteId, r.TerminationReason, r.Timestamp, r.Type, r.Url, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponsePcapSearchItem.
// It customizes the JSON marshaling process for ResponsePcapSearchItem objects.
func (r ResponsePcapSearchItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "ap_macs", "aps", "duration", "format", "id", "max_num_packets", "org_id", "pcap_aps", "pcap_url", "site_id", "termination_reason", "timestamp", "type", "url"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponsePcapSearchItem object to a map representation for JSON marshaling.
func (r ResponsePcapSearchItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.ApMacs != nil {
        structMap["ap_macs"] = r.ApMacs
    }
    if r.Aps != nil {
        structMap["aps"] = r.Aps
    }
    if r.Duration != nil {
        structMap["duration"] = r.Duration
    }
    if r.Format != nil {
        structMap["format"] = r.Format
    }
    if r.Id != nil {
        structMap["id"] = r.Id
    }
    if r.MaxNumPackets != nil {
        structMap["max_num_packets"] = r.MaxNumPackets
    }
    if r.OrgId != nil {
        structMap["org_id"] = r.OrgId
    }
    if r.PcapAps != nil {
        structMap["pcap_aps"] = r.PcapAps
    }
    if r.PcapUrl != nil {
        structMap["pcap_url"] = r.PcapUrl
    }
    if r.SiteId != nil {
        structMap["site_id"] = r.SiteId
    }
    if r.TerminationReason != nil {
        structMap["termination_reason"] = r.TerminationReason
    }
    structMap["timestamp"] = r.Timestamp
    structMap["type"] = r.Type
    structMap["url"] = r.Url
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponsePcapSearchItem.
// It customizes the JSON unmarshaling process for ResponsePcapSearchItem objects.
func (r *ResponsePcapSearchItem) UnmarshalJSON(input []byte) error {
    var temp tempResponsePcapSearchItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_macs", "aps", "duration", "format", "id", "max_num_packets", "org_id", "pcap_aps", "pcap_url", "site_id", "termination_reason", "timestamp", "type", "url")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.ApMacs = temp.ApMacs
    r.Aps = temp.Aps
    r.Duration = temp.Duration
    r.Format = temp.Format
    r.Id = temp.Id
    r.MaxNumPackets = temp.MaxNumPackets
    r.OrgId = temp.OrgId
    r.PcapAps = temp.PcapAps
    r.PcapUrl = temp.PcapUrl
    r.SiteId = temp.SiteId
    r.TerminationReason = temp.TerminationReason
    r.Timestamp = *temp.Timestamp
    r.Type = *temp.Type
    r.Url = *temp.Url
    return nil
}

// tempResponsePcapSearchItem is a temporary struct used for validating the fields of ResponsePcapSearchItem.
type tempResponsePcapSearchItem  struct {
    ApMacs            []string                                     `json:"ap_macs,omitempty"`
    Aps               []string                                     `json:"aps,omitempty"`
    Duration          *float64                                     `json:"duration,omitempty"`
    Format            *string                                      `json:"format,omitempty"`
    Id                *uuid.UUID                                   `json:"id,omitempty"`
    MaxNumPackets     *float64                                     `json:"max_num_packets,omitempty"`
    OrgId             *uuid.UUID                                   `json:"org_id,omitempty"`
    PcapAps           map[string]ResponsePcapSearchItemPcapApsItem `json:"pcap_aps,omitempty"`
    PcapUrl           *string                                      `json:"pcap_url,omitempty"`
    SiteId            *uuid.UUID                                   `json:"site_id,omitempty"`
    TerminationReason *string                                      `json:"termination_reason,omitempty"`
    Timestamp         *float64                                     `json:"timestamp"`
    Type              *string                                      `json:"type"`
    Url               *string                                      `json:"url"`
}

func (r *tempResponsePcapSearchItem) validate() error {
    var errs []string
    if r.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `response_pcap_search_item`")
    }
    if r.Type == nil {
        errs = append(errs, "required field `type` is missing for type `response_pcap_search_item`")
    }
    if r.Url == nil {
        errs = append(errs, "required field `url` is missing for type `response_pcap_search_item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
