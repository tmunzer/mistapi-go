package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SynthetictestDevice represents a SynthetictestDevice struct.
type SynthetictestDevice struct {
    // If `type`==`lan_connectivity`
    Host                 *string                          `json:"host,omitempty"`
    // If `type`==`dns`
    Hostname             *string                          `json:"hostname,omitempty"`
    // If `type`==`arp`
    Ip                   *string                          `json:"ip,omitempty"`
    // If `type`==`radius`
    Password             *string                          `json:"password,omitempty"`
    // If `type`==`lan_connectivity`
    PingCount            *int                             `json:"ping_count,omitempty"`
    // If `type`==`lan_connectivity`
    PingDetails          *bool                            `json:"ping_details,omitempty"`
    // If `type`==`lan_connectivity`
    PingSize             *int                             `json:"ping_size,omitempty"`
    // If `type`==`speedtest`, required for ssr
    PortId               *string                          `json:"port_id,omitempty"`
    // if `type`==`lan_connectivity`. enum: `ping`, `traceroute`, `ping+traceroute`
    Protocol             *SynthetictestDeviceProtocolEnum `json:"protocol,omitempty"`
    // If `type`==`lan_connectivity`
    Tenant               *string                          `json:"tenant,omitempty"`
    // SRX only, traceroute udp port
    TracerouteUdpPort    *int                             `json:"traceroute_udp_port,omitempty"`
    // enum: `arp`, `curl`, `dhcp`, `dhcp6`, `dns`, `lan_connectivity`, `radius`, `speedtest`
    Type                 SynthetictestTypeEnum            `json:"type"`
    // If `type`==`curl`
    Url                  *string                          `json:"url,omitempty"`
    // If `type`==`radius`
    Username             *string                          `json:"username,omitempty"`
    // Required for AP
    VlanId               *SynthetictestDeviceVlanId       `json:"vlan_id,omitempty"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for SynthetictestDevice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SynthetictestDevice) String() string {
    return fmt.Sprintf(
    	"SynthetictestDevice[Host=%v, Hostname=%v, Ip=%v, Password=%v, PingCount=%v, PingDetails=%v, PingSize=%v, PortId=%v, Protocol=%v, Tenant=%v, TracerouteUdpPort=%v, Type=%v, Url=%v, Username=%v, VlanId=%v, AdditionalProperties=%v]",
    	s.Host, s.Hostname, s.Ip, s.Password, s.PingCount, s.PingDetails, s.PingSize, s.PortId, s.Protocol, s.Tenant, s.TracerouteUdpPort, s.Type, s.Url, s.Username, s.VlanId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestDevice.
// It customizes the JSON marshaling process for SynthetictestDevice objects.
func (s SynthetictestDevice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "host", "hostname", "ip", "password", "ping_count", "ping_details", "ping_size", "port_id", "protocol", "tenant", "traceroute_udp_port", "type", "url", "username", "vlan_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SynthetictestDevice object to a map representation for JSON marshaling.
func (s SynthetictestDevice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "host", "hostname", "ip", "password", "ping_count", "ping_details", "ping_size", "port_id", "protocol", "tenant", "traceroute_udp_port", "type", "url", "username", "vlan_id")
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
