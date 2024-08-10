package models

import (
    "encoding/json"
)

// AppProbingCustomApp represents a AppProbingCustomApp struct.
type AppProbingCustomApp struct {
    // if `protocol`==`icmp`
    Address              *string                          `json:"address,omitempty"`
    AppType              *string                          `json:"app_type,omitempty"`
    // if `protocol`==`http`
    Hostname             []string                         `json:"hostname,omitempty"`
    Name                 *string                          `json:"name,omitempty"`
    Network              *string                          `json:"network,omitempty"`
    // enum: `http`, `udp`
    Protocol             *AppProbingCustomAppProtocolEnum `json:"protocol,omitempty"`
    // if `protocol`==`http`
    Url                  *string                          `json:"url,omitempty"`
    Vrf                  *string                          `json:"vrf,omitempty"`
    AdditionalProperties map[string]any                   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AppProbingCustomApp.
// It customizes the JSON marshaling process for AppProbingCustomApp objects.
func (a AppProbingCustomApp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AppProbingCustomApp object to a map representation for JSON marshaling.
func (a AppProbingCustomApp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Address != nil {
        structMap["address"] = a.Address
    }
    if a.AppType != nil {
        structMap["app_type"] = a.AppType
    }
    if a.Hostname != nil {
        structMap["hostname"] = a.Hostname
    }
    if a.Name != nil {
        structMap["name"] = a.Name
    }
    if a.Network != nil {
        structMap["network"] = a.Network
    }
    if a.Protocol != nil {
        structMap["protocol"] = a.Protocol
    }
    if a.Url != nil {
        structMap["url"] = a.Url
    }
    if a.Vrf != nil {
        structMap["vrf"] = a.Vrf
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AppProbingCustomApp.
// It customizes the JSON unmarshaling process for AppProbingCustomApp objects.
func (a *AppProbingCustomApp) UnmarshalJSON(input []byte) error {
    var temp tempAppProbingCustomApp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "address", "app_type", "hostname", "name", "network", "protocol", "url", "vrf")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Address = temp.Address
    a.AppType = temp.AppType
    a.Hostname = temp.Hostname
    a.Name = temp.Name
    a.Network = temp.Network
    a.Protocol = temp.Protocol
    a.Url = temp.Url
    a.Vrf = temp.Vrf
    return nil
}

// tempAppProbingCustomApp is a temporary struct used for validating the fields of AppProbingCustomApp.
type tempAppProbingCustomApp  struct {
    Address  *string                          `json:"address,omitempty"`
    AppType  *string                          `json:"app_type,omitempty"`
    Hostname []string                         `json:"hostname,omitempty"`
    Name     *string                          `json:"name,omitempty"`
    Network  *string                          `json:"network,omitempty"`
    Protocol *AppProbingCustomAppProtocolEnum `json:"protocol,omitempty"`
    Url      *string                          `json:"url,omitempty"`
    Vrf      *string                          `json:"vrf,omitempty"`
}
