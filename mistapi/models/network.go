package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// Network represents a Network struct.
// Networks are usually subnets that have cross-site significance. `networks`in Org Settings will got merged into `networks`in Site Setting. For gateways, they can be used to define Service Routes.
type Network struct {
    // when the object has been created, in epoch
    CreatedTime          *float64                          `json:"created_time,omitempty"`
    // whether to disallow Mist Devices in the network
    DisallowMistServices *bool                             `json:"disallow_mist_services,omitempty"`
    Gateway              *string                           `json:"gateway,omitempty"`
    Gateway6             *string                           `json:"gateway6,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID                        `json:"id,omitempty"`
    InternalAccess       *NetworkInternalAccess            `json:"internal_access,omitempty"`
    // whether this network has direct internet access
    InternetAccess       *NetworkInternetAccess            `json:"internet_access,omitempty"`
    // whether to allow clients in the network to talk to each other
    Isolation            *bool                             `json:"isolation,omitempty"`
    // when the object has been modified for the last time, in epoch
    ModifiedTime         *float64                          `json:"modified_time,omitempty"`
    // whether to enable multicast support (only PIM-sparse mode is supported)
    Multicast            *NetworkMulticast                 `json:"multicast,omitempty"`
    Name                 string                            `json:"name"`
    OrgId                *uuid.UUID                        `json:"org_id,omitempty"`
    // for a Network (usually LAN), it can be routable to other networks (e.g. OSPF)
    RoutedForNetworks    []string                          `json:"routed_for_networks,omitempty"`
    Subnet               *string                           `json:"subnet,omitempty"`
    Subnet6              *string                           `json:"subnet6,omitempty"`
    Tenants              map[string]NetworkTenant          `json:"tenants,omitempty"`
    VlanId               *VlanIdWithVariable               `json:"vlan_id,omitempty"`
    // Property key is the VPN name. Whether this network can be accessed from vpn
    VpnAccess            map[string]NetworkVpnAccessConfig `json:"vpn_access,omitempty"`
    AdditionalProperties map[string]interface{}            `json:"_"`
}

// String implements the fmt.Stringer interface for Network,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n Network) String() string {
    return fmt.Sprintf(
    	"Network[CreatedTime=%v, DisallowMistServices=%v, Gateway=%v, Gateway6=%v, Id=%v, InternalAccess=%v, InternetAccess=%v, Isolation=%v, ModifiedTime=%v, Multicast=%v, Name=%v, OrgId=%v, RoutedForNetworks=%v, Subnet=%v, Subnet6=%v, Tenants=%v, VlanId=%v, VpnAccess=%v, AdditionalProperties=%v]",
    	n.CreatedTime, n.DisallowMistServices, n.Gateway, n.Gateway6, n.Id, n.InternalAccess, n.InternetAccess, n.Isolation, n.ModifiedTime, n.Multicast, n.Name, n.OrgId, n.RoutedForNetworks, n.Subnet, n.Subnet6, n.Tenants, n.VlanId, n.VpnAccess, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Network.
// It customizes the JSON marshaling process for Network objects.
func (n Network) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(n.AdditionalProperties,
        "created_time", "disallow_mist_services", "gateway", "gateway6", "id", "internal_access", "internet_access", "isolation", "modified_time", "multicast", "name", "org_id", "routed_for_networks", "subnet", "subnet6", "tenants", "vlan_id", "vpn_access"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(n.toMap())
}

// toMap converts the Network object to a map representation for JSON marshaling.
func (n Network) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, n.AdditionalProperties)
    if n.CreatedTime != nil {
        structMap["created_time"] = n.CreatedTime
    }
    if n.DisallowMistServices != nil {
        structMap["disallow_mist_services"] = n.DisallowMistServices
    }
    if n.Gateway != nil {
        structMap["gateway"] = n.Gateway
    }
    if n.Gateway6 != nil {
        structMap["gateway6"] = n.Gateway6
    }
    if n.Id != nil {
        structMap["id"] = n.Id
    }
    if n.InternalAccess != nil {
        structMap["internal_access"] = n.InternalAccess.toMap()
    }
    if n.InternetAccess != nil {
        structMap["internet_access"] = n.InternetAccess.toMap()
    }
    if n.Isolation != nil {
        structMap["isolation"] = n.Isolation
    }
    if n.ModifiedTime != nil {
        structMap["modified_time"] = n.ModifiedTime
    }
    if n.Multicast != nil {
        structMap["multicast"] = n.Multicast.toMap()
    }
    structMap["name"] = n.Name
    if n.OrgId != nil {
        structMap["org_id"] = n.OrgId
    }
    if n.RoutedForNetworks != nil {
        structMap["routed_for_networks"] = n.RoutedForNetworks
    }
    if n.Subnet != nil {
        structMap["subnet"] = n.Subnet
    }
    if n.Subnet6 != nil {
        structMap["subnet6"] = n.Subnet6
    }
    if n.Tenants != nil {
        structMap["tenants"] = n.Tenants
    }
    if n.VlanId != nil {
        structMap["vlan_id"] = n.VlanId.toMap()
    }
    if n.VpnAccess != nil {
        structMap["vpn_access"] = n.VpnAccess
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Network.
// It customizes the JSON unmarshaling process for Network objects.
func (n *Network) UnmarshalJSON(input []byte) error {
    var temp tempNetwork
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "disallow_mist_services", "gateway", "gateway6", "id", "internal_access", "internet_access", "isolation", "modified_time", "multicast", "name", "org_id", "routed_for_networks", "subnet", "subnet6", "tenants", "vlan_id", "vpn_access")
    if err != nil {
    	return err
    }
    n.AdditionalProperties = additionalProperties
    
    n.CreatedTime = temp.CreatedTime
    n.DisallowMistServices = temp.DisallowMistServices
    n.Gateway = temp.Gateway
    n.Gateway6 = temp.Gateway6
    n.Id = temp.Id
    n.InternalAccess = temp.InternalAccess
    n.InternetAccess = temp.InternetAccess
    n.Isolation = temp.Isolation
    n.ModifiedTime = temp.ModifiedTime
    n.Multicast = temp.Multicast
    n.Name = *temp.Name
    n.OrgId = temp.OrgId
    n.RoutedForNetworks = temp.RoutedForNetworks
    n.Subnet = temp.Subnet
    n.Subnet6 = temp.Subnet6
    n.Tenants = temp.Tenants
    n.VlanId = temp.VlanId
    n.VpnAccess = temp.VpnAccess
    return nil
}

// tempNetwork is a temporary struct used for validating the fields of Network.
type tempNetwork  struct {
    CreatedTime          *float64                          `json:"created_time,omitempty"`
    DisallowMistServices *bool                             `json:"disallow_mist_services,omitempty"`
    Gateway              *string                           `json:"gateway,omitempty"`
    Gateway6             *string                           `json:"gateway6,omitempty"`
    Id                   *uuid.UUID                        `json:"id,omitempty"`
    InternalAccess       *NetworkInternalAccess            `json:"internal_access,omitempty"`
    InternetAccess       *NetworkInternetAccess            `json:"internet_access,omitempty"`
    Isolation            *bool                             `json:"isolation,omitempty"`
    ModifiedTime         *float64                          `json:"modified_time,omitempty"`
    Multicast            *NetworkMulticast                 `json:"multicast,omitempty"`
    Name                 *string                           `json:"name"`
    OrgId                *uuid.UUID                        `json:"org_id,omitempty"`
    RoutedForNetworks    []string                          `json:"routed_for_networks,omitempty"`
    Subnet               *string                           `json:"subnet,omitempty"`
    Subnet6              *string                           `json:"subnet6,omitempty"`
    Tenants              map[string]NetworkTenant          `json:"tenants,omitempty"`
    VlanId               *VlanIdWithVariable               `json:"vlan_id,omitempty"`
    VpnAccess            map[string]NetworkVpnAccessConfig `json:"vpn_access,omitempty"`
}

func (n *tempNetwork) validate() error {
    var errs []string
    if n.Name == nil {
        errs = append(errs, "required field `name` is missing for type `network`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
