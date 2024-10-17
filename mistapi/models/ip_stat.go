package models

import (
	"encoding/json"
)

// IpStat represents a IpStat struct.
type IpStat struct {
	DhcpServer           Optional[string]  `json:"dhcp_server"`
	Dns                  []string          `json:"dns,omitempty"`
	DnsSuffix            []string          `json:"dns_suffix,omitempty"`
	Gateway              Optional[string]  `json:"gateway"`
	Gateway6             Optional[string]  `json:"gateway6"`
	Ip                   Optional[string]  `json:"ip"`
	Ip6                  Optional[string]  `json:"ip6"`
	Ips                  map[string]string `json:"ips,omitempty"`
	Netmask              Optional[string]  `json:"netmask"`
	Netmask6             Optional[string]  `json:"netmask6"`
	AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for IpStat.
// It customizes the JSON marshaling process for IpStat objects.
func (i IpStat) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the IpStat object to a map representation for JSON marshaling.
func (i IpStat) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, i.AdditionalProperties)
	if i.DhcpServer.IsValueSet() {
		if i.DhcpServer.Value() != nil {
			structMap["dhcp_server"] = i.DhcpServer.Value()
		} else {
			structMap["dhcp_server"] = nil
		}
	}
	if i.Dns != nil {
		structMap["dns"] = i.Dns
	}
	if i.DnsSuffix != nil {
		structMap["dns_suffix"] = i.DnsSuffix
	}
	if i.Gateway.IsValueSet() {
		if i.Gateway.Value() != nil {
			structMap["gateway"] = i.Gateway.Value()
		} else {
			structMap["gateway"] = nil
		}
	}
	if i.Gateway6.IsValueSet() {
		if i.Gateway6.Value() != nil {
			structMap["gateway6"] = i.Gateway6.Value()
		} else {
			structMap["gateway6"] = nil
		}
	}
	if i.Ip.IsValueSet() {
		if i.Ip.Value() != nil {
			structMap["ip"] = i.Ip.Value()
		} else {
			structMap["ip"] = nil
		}
	}
	if i.Ip6.IsValueSet() {
		if i.Ip6.Value() != nil {
			structMap["ip6"] = i.Ip6.Value()
		} else {
			structMap["ip6"] = nil
		}
	}
	if i.Ips != nil {
		structMap["ips"] = i.Ips
	}
	if i.Netmask.IsValueSet() {
		if i.Netmask.Value() != nil {
			structMap["netmask"] = i.Netmask.Value()
		} else {
			structMap["netmask"] = nil
		}
	}
	if i.Netmask6.IsValueSet() {
		if i.Netmask6.Value() != nil {
			structMap["netmask6"] = i.Netmask6.Value()
		} else {
			structMap["netmask6"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IpStat.
// It customizes the JSON unmarshaling process for IpStat objects.
func (i *IpStat) UnmarshalJSON(input []byte) error {
	var temp tempIpStat
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "dhcp_server", "dns", "dns_suffix", "gateway", "gateway6", "ip", "ip6", "ips", "netmask", "netmask6")
	if err != nil {
		return err
	}

	i.AdditionalProperties = additionalProperties
	i.DhcpServer = temp.DhcpServer
	i.Dns = temp.Dns
	i.DnsSuffix = temp.DnsSuffix
	i.Gateway = temp.Gateway
	i.Gateway6 = temp.Gateway6
	i.Ip = temp.Ip
	i.Ip6 = temp.Ip6
	i.Ips = temp.Ips
	i.Netmask = temp.Netmask
	i.Netmask6 = temp.Netmask6
	return nil
}

// tempIpStat is a temporary struct used for validating the fields of IpStat.
type tempIpStat struct {
	DhcpServer Optional[string]  `json:"dhcp_server"`
	Dns        []string          `json:"dns,omitempty"`
	DnsSuffix  []string          `json:"dns_suffix,omitempty"`
	Gateway    Optional[string]  `json:"gateway"`
	Gateway6   Optional[string]  `json:"gateway6"`
	Ip         Optional[string]  `json:"ip"`
	Ip6        Optional[string]  `json:"ip6"`
	Ips        map[string]string `json:"ips,omitempty"`
	Netmask    Optional[string]  `json:"netmask"`
	Netmask6   Optional[string]  `json:"netmask6"`
}
