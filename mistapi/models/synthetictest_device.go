package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SynthetictestDevice represents a SynthetictestDevice struct.
type SynthetictestDevice struct {
    // if `type`==`dns`
    Hostname             *string                    `json:"hostname,omitempty"`
    // if `type`==`arp`
    Ip                   *string                    `json:"ip,omitempty"`
    // if `type`==`radius`
    Password             *string                    `json:"password,omitempty"`
    // if `type`==`ssr`
    PortId               *string                    `json:"port_id,omitempty"`
    Type                 SynthetictestTypeEnum      `json:"type"`
    // if `type`==`curl`
    Url                  *string                    `json:"url,omitempty"`
    // if `type`==`radius`
    Username             *string                    `json:"username,omitempty"`
    // required for AP
    VlanId               *SynthetictestDeviceVlanId `json:"vlan_id,omitempty"`
    AdditionalProperties map[string]any             `json:"_"`
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
    if s.Hostname != nil {
        structMap["hostname"] = s.Hostname
    }
    if s.Ip != nil {
        structMap["ip"] = s.Ip
    }
    if s.Password != nil {
        structMap["password"] = s.Password
    }
    if s.PortId != nil {
        structMap["port_id"] = s.PortId
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
    var temp synthetictestDevice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "hostname", "ip", "password", "port_id", "type", "url", "username", "vlan_id")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Hostname = temp.Hostname
    s.Ip = temp.Ip
    s.Password = temp.Password
    s.PortId = temp.PortId
    s.Type = *temp.Type
    s.Url = temp.Url
    s.Username = temp.Username
    s.VlanId = temp.VlanId
    return nil
}

// synthetictestDevice is a temporary struct used for validating the fields of SynthetictestDevice.
type synthetictestDevice  struct {
    Hostname *string                    `json:"hostname,omitempty"`
    Ip       *string                    `json:"ip,omitempty"`
    Password *string                    `json:"password,omitempty"`
    PortId   *string                    `json:"port_id,omitempty"`
    Type     *SynthetictestTypeEnum     `json:"type"`
    Url      *string                    `json:"url,omitempty"`
    Username *string                    `json:"username,omitempty"`
    VlanId   *SynthetictestDeviceVlanId `json:"vlan_id,omitempty"`
}

func (s *synthetictestDevice) validate() error {
    var errs []string
    if s.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Synthetictest_Device`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
