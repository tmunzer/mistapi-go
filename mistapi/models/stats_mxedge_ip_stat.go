package models

import (
    "encoding/json"
    "fmt"
)

// StatsMxedgeIpStat represents a StatsMxedgeIpStat struct.
// IP stats
type StatsMxedgeIpStat struct {
    Ip                   *string                `json:"ip,omitempty"`
    Ip6                  *string                `json:"ip6,omitempty"`
    // Property key is the interface name. IPs for each net interface
    Ips                  map[string]string      `json:"ips,omitempty"`
    // Property key is the interface name. MAC for each net interface
    Macs                 map[string]string      `json:"macs,omitempty"`
    Netmask              *string                `json:"netmask,omitempty"`
    Netmask6             *string                `json:"netmask6,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMxedgeIpStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMxedgeIpStat) String() string {
    return fmt.Sprintf(
    	"StatsMxedgeIpStat[Ip=%v, Ip6=%v, Ips=%v, Macs=%v, Netmask=%v, Netmask6=%v, AdditionalProperties=%v]",
    	s.Ip, s.Ip6, s.Ips, s.Macs, s.Netmask, s.Netmask6, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgeIpStat.
// It customizes the JSON marshaling process for StatsMxedgeIpStat objects.
func (s StatsMxedgeIpStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "ip", "ip6", "ips", "macs", "netmask", "netmask6"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgeIpStat object to a map representation for JSON marshaling.
func (s StatsMxedgeIpStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Ip != nil {
        structMap["ip"] = s.Ip
    }
    if s.Ip6 != nil {
        structMap["ip6"] = s.Ip6
    }
    if s.Ips != nil {
        structMap["ips"] = s.Ips
    }
    if s.Macs != nil {
        structMap["macs"] = s.Macs
    }
    if s.Netmask != nil {
        structMap["netmask"] = s.Netmask
    }
    if s.Netmask6 != nil {
        structMap["netmask6"] = s.Netmask6
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgeIpStat.
// It customizes the JSON unmarshaling process for StatsMxedgeIpStat objects.
func (s *StatsMxedgeIpStat) UnmarshalJSON(input []byte) error {
    var temp tempStatsMxedgeIpStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ip", "ip6", "ips", "macs", "netmask", "netmask6")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Ip = temp.Ip
    s.Ip6 = temp.Ip6
    s.Ips = temp.Ips
    s.Macs = temp.Macs
    s.Netmask = temp.Netmask
    s.Netmask6 = temp.Netmask6
    return nil
}

// tempStatsMxedgeIpStat is a temporary struct used for validating the fields of StatsMxedgeIpStat.
type tempStatsMxedgeIpStat  struct {
    Ip       *string           `json:"ip,omitempty"`
    Ip6      *string           `json:"ip6,omitempty"`
    Ips      map[string]string `json:"ips,omitempty"`
    Macs     map[string]string `json:"macs,omitempty"`
    Netmask  *string           `json:"netmask,omitempty"`
    Netmask6 *string           `json:"netmask6,omitempty"`
}
