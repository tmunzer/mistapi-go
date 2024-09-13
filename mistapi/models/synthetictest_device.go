package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SynthetictestDevice represents a SynthetictestDevice struct.
type SynthetictestDevice struct {
    // if `type`==`lan_connectivity`
    Host                 *string                          `json:"host,omitempty"`
    // if `type`==`dns`
    Hostname             *string                          `json:"hostname,omitempty"`
    // if `type`==`arp`
    Ip                   *string                          `json:"ip,omitempty"`
    // if `type`==`radius`
    Password             *string                          `json:"password,omitempty"`
    // if `type`==`lan_connectivity`
    PingCount            *int                             `json:"ping_count,omitempty"`
    // if `type`==`lan_connectivity`
    PingDetails          *bool                            `json:"ping_details,omitempty"`
    // if `type`==`lan_connectivity`
    PingSize             *int                             `json:"ping_size,omitempty"`
    // if `type`==`speedtest`, required for ssr
    PortId               *string                          `json:"port_id,omitempty"`
    // if `type`==`lan_connectivity`. enum: `ping`, `traceroute`, `ping+traceroute`
    Protocol             *SynthetictestDeviceProtocolEnum `json:"protocol,omitempty"`
    // if `type`==`lan_connectivity`
    Tenant               *string                          `json:"tenant,omitempty"`
    // SRX only, traceroute udp port
    TracerouteUdpPort    *int                             `json:"traceroute_udp_port,omitempty"`
    // enum: `arp`, `curl`, `dhcp`, `dhcp6`, `dns`, `lan_connectivity`, `radius`, `speedtest`
    Type                 SynthetictestTypeEnum            `json:"type"`
    // if `type`==`curl`
    Url                  *string                          `json:"url,omitempty"`
    // if `type`==`radius`
    Username             *string                          `json:"username,omitempty"`
    // required for AP
    VlanId               *SynthetictestDeviceVlanId       `json:"vlan_id,omitempty"`
    AdditionalProperties map[string]any                   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestDevice.
// It customizes the JSON marshaling process for SynthetictestDevice objects.
func (s SynthetictestDevice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SynthetictestDevice object to a map representation for JSON marshaling.
func (s SynthetictestDevice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Host != nil {
        structMap["host"] = s.Host
    }
    if s.Hostname != nil {
        structMap["hostname"] = s.Hostname
    }
    if s.Ip != nil {
        structMap["ip"] = s.Ip
    }
    if s.Password != nil {
        structMap["password"] = s.Password
    }
    if s.PingCount != nil {
        structMap["ping_count"] = s.PingCount
    }
    if s.PingDetails != nil {
        structMap["ping_details"] = s.PingDetails
    }
    if s.PingSize != nil {
        structMap["ping_size"] = s.PingSize
    }
    if s.PortId != nil {
        structMap["port_id"] = s.PortId
    }
    if s.Protocol != nil {
        structMap["protocol"] = s.Protocol
    }
    if s.Tenant != nil {
        structMap["tenant"] = s.Tenant
    }
    if s.TracerouteUdpPort != nil {
        structMap["traceroute_udp_port"] = s.TracerouteUdpPort
    }
    structMap["type"] = s.Type
    if s.Url != nil {
        structMap["url"] = s.Url
    }
    if s.Username != nil {
        structMap["username"] = s.Username
    }
    if s.VlanId != nil {
        structMap["vlan_id"] = s.VlanId.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SynthetictestDevice.
// It customizes the JSON unmarshaling process for SynthetictestDevice objects.
func (s *SynthetictestDevice) UnmarshalJSON(input []byte) error {
    var temp tempSynthetictestDevice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "host", "hostname", "ip", "password", "ping_count", "ping_details", "ping_size", "port_id", "protocol", "tenant", "traceroute_udp_port", "type", "url", "username", "vlan_id")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Host = temp.Host
    s.Hostname = temp.Hostname
    s.Ip = temp.Ip
    s.Password = temp.Password
    s.PingCount = temp.PingCount
    s.PingDetails = temp.PingDetails
    s.PingSize = temp.PingSize
    s.PortId = temp.PortId
    s.Protocol = temp.Protocol
    s.Tenant = temp.Tenant
    s.TracerouteUdpPort = temp.TracerouteUdpPort
    s.Type = *temp.Type
    s.Url = temp.Url
    s.Username = temp.Username
    s.VlanId = temp.VlanId
    return nil
}

// tempSynthetictestDevice is a temporary struct used for validating the fields of SynthetictestDevice.
type tempSynthetictestDevice  struct {
    Host              *string                          `json:"host,omitempty"`
    Hostname          *string                          `json:"hostname,omitempty"`
    Ip                *string                          `json:"ip,omitempty"`
    Password          *string                          `json:"password,omitempty"`
    PingCount         *int                             `json:"ping_count,omitempty"`
    PingDetails       *bool                            `json:"ping_details,omitempty"`
    PingSize          *int                             `json:"ping_size,omitempty"`
    PortId            *string                          `json:"port_id,omitempty"`
    Protocol          *SynthetictestDeviceProtocolEnum `json:"protocol,omitempty"`
    Tenant            *string                          `json:"tenant,omitempty"`
    TracerouteUdpPort *int                             `json:"traceroute_udp_port,omitempty"`
    Type              *SynthetictestTypeEnum           `json:"type"`
    Url               *string                          `json:"url,omitempty"`
    Username          *string                          `json:"username,omitempty"`
    VlanId            *SynthetictestDeviceVlanId       `json:"vlan_id,omitempty"`
}

func (s *tempSynthetictestDevice) validate() error {
    var errs []string
    if s.Type == nil {
        errs = append(errs, "required field `type` is missing for type `synthetictest_device`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
