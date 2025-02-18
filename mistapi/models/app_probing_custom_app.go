package models

import (
    "encoding/json"
    "fmt"
)

// AppProbingCustomApp represents a AppProbingCustomApp struct.
type AppProbingCustomApp struct {
    // Required if `protocol`==`icmp`
    Address              *string                          `json:"address,omitempty"`
    AppType              *string                          `json:"app_type,omitempty"`
    // If `protocol`==`http`
    Hostnames            []string                         `json:"hostnames,omitempty"`
    Key                  *string                          `json:"key,omitempty"`
    Name                 *string                          `json:"name,omitempty"`
    Network              *string                          `json:"network,omitempty"`
    // If `protocol`==`icmp`
    PacketSize           *int                             `json:"packetSize,omitempty"`
    // enum: `http`, `icmp`
    Protocol             *AppProbingCustomAppProtocolEnum `json:"protocol,omitempty"`
    // If `protocol`==`http`
    Url                  *string                          `json:"url,omitempty"`
    Vrf                  *string                          `json:"vrf,omitempty"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for AppProbingCustomApp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AppProbingCustomApp) String() string {
    return fmt.Sprintf(
    	"AppProbingCustomApp[Address=%v, AppType=%v, Hostnames=%v, Key=%v, Name=%v, Network=%v, PacketSize=%v, Protocol=%v, Url=%v, Vrf=%v, AdditionalProperties=%v]",
    	a.Address, a.AppType, a.Hostnames, a.Key, a.Name, a.Network, a.PacketSize, a.Protocol, a.Url, a.Vrf, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AppProbingCustomApp.
// It customizes the JSON marshaling process for AppProbingCustomApp objects.
func (a AppProbingCustomApp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "address", "app_type", "hostnames", "key", "name", "network", "packetSize", "protocol", "url", "vrf"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AppProbingCustomApp object to a map representation for JSON marshaling.
func (a AppProbingCustomApp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Address != nil {
        structMap["address"] = a.Address
    }
    if a.AppType != nil {
        structMap["app_type"] = a.AppType
    }
    if a.Hostnames != nil {
        structMap["hostnames"] = a.Hostnames
    }
    if a.Key != nil {
        structMap["key"] = a.Key
    }
    if a.Name != nil {
        structMap["name"] = a.Name
    }
    if a.Network != nil {
        structMap["network"] = a.Network
    }
    if a.PacketSize != nil {
        structMap["packetSize"] = a.PacketSize
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "address", "app_type", "hostnames", "key", "name", "network", "packetSize", "protocol", "url", "vrf")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Address = temp.Address
    a.AppType = temp.AppType
    a.Hostnames = temp.Hostnames
    a.Key = temp.Key
    a.Name = temp.Name
    a.Network = temp.Network
    a.PacketSize = temp.PacketSize
    a.Protocol = temp.Protocol
    a.Url = temp.Url
    a.Vrf = temp.Vrf
    return nil
}

// tempAppProbingCustomApp is a temporary struct used for validating the fields of AppProbingCustomApp.
type tempAppProbingCustomApp  struct {
    Address    *string                          `json:"address,omitempty"`
    AppType    *string                          `json:"app_type,omitempty"`
    Hostnames  []string                         `json:"hostnames,omitempty"`
    Key        *string                          `json:"key,omitempty"`
    Name       *string                          `json:"name,omitempty"`
    Network    *string                          `json:"network,omitempty"`
    PacketSize *int                             `json:"packetSize,omitempty"`
    Protocol   *AppProbingCustomAppProtocolEnum `json:"protocol,omitempty"`
    Url        *string                          `json:"url,omitempty"`
    Vrf        *string                          `json:"vrf,omitempty"`
}
