
# Network

Organization-level Layer 3 network definition that can be merged into site settings and used for service routes. Networks are used to define the service routes in the Gateway settings or templates

*This model accepts additional fields of type interface{}.*

## Structure

`Network`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DisallowMistServices` | `*bool` | Optional | Whether to disallow Mist Devices in the network<br><br>**Default**: `false` |
| `Gateway` | `*string` | Optional | IPv4 gateway address for this network |
| `Gateway6` | `*string` | Optional | IPv6 gateway address for this network |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `InternalAccess` | [`*models.NetworkInternalAccess`](../../doc/models/network-internal-access.md) | Optional | Internal access settings for an organization network |
| `InternetAccess` | [`*models.NetworkInternetAccess`](../../doc/models/network-internet-access.md) | Optional | Direct internet access settings for an organization network |
| `Isolation` | `*bool` | Optional | Whether to allow clients in the network to talk to each other |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Multicast` | [`*models.NetworkMulticast`](../../doc/models/network-multicast.md) | Optional | Whether to enable multicast support (only PIM-sparse mode is supported) |
| `Name` | `string` | Required | Display name of the organization network |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `RoutedForNetworks` | `[]string` | Optional | For a Network (usually LAN), it can be routable to other networks (e.g. OSPF) |
| `Subnet` | `*string` | Optional | IPv4 subnet CIDR for this network |
| `Subnet6` | `*string` | Optional | IPv6 subnet CIDR for this network |
| `Tenants` | [`map[string]models.NetworkTenant`](../../doc/models/network-tenant.md) | Optional | Property key must be the user/tenant name (i.e. "printer-1") or a Variable (i.e. "{{myvar}}") |
| `VlanId` | [`*models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Optional | VLAN ID, either numeric or expressed as a template variable string |
| `VpnAccess` | [`map[string]models.NetworkVpnAccessConfig`](../../doc/models/network-vpn-access-config.md) | Optional | Property key is the VPN name. Whether this network can be accessed from vpn |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    network := models.Network{
        CreatedTime:          models.ToPointer(float64(103.24)),
        DisallowMistServices: models.ToPointer(false),
        Gateway:              models.ToPointer("192.168.70.1"),
        Gateway6:             models.ToPointer("fdad:b0bc:f29e::1"),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Name:                 "name4",
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Subnet:               models.ToPointer("192.168.70.0/24"),
        Subnet6:              models.ToPointer("fdad:b0bc:f29e::/32"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

