
# Network Vpn Access Config

VPN access settings for a network and VPN pair

## Structure

`NetworkVpnAccessConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdvertisedSubnet` | `*string` | Optional | If `routed`==`true`, whether to advertise an aggregated subnet toward HUB this is useful when there are multiple networks on SPOKE's side |
| `AllowPing` | `*bool` | Optional | Whether to allow ping from vpn into this routed network |
| `DestinationNat` | [`map[string]models.NetworkVpnAccessDestinationNatProperty`](../../doc/models/network-vpn-access-destination-nat-property.md) | Optional | Property key can be an External IP (i.e. "63.16.0.3"), an External IP:Port (i.e. "63.16.0.3:443"), an External Port (i.e. ":443"), an External CIDR (i.e. "63.16.0.0/30"), an External CIDR:Port (i.e. "63.16.0.0/30:443") or a Variable (i.e. "{{myvar}}"). At least one of the `internal_ip` or `port` must be defined |
| `NatPool` | `*string` | Optional | If `routed`==`false` (usually at Spoke), but some hosts needs to be reachable from Hub, a subnet is required to create and advertise the route to Hub |
| `NoReadvertiseToLanBgp` | `*bool` | Optional | toward LAN-side BGP peers<br><br>**Default**: `false` |
| `NoReadvertiseToLanOspf` | `*bool` | Optional | toward LAN-side OSPF peers<br><br>**Default**: `false` |
| `NoReadvertiseToOverlay` | `*bool` | Optional | toward overlay, how HUB should deal with routes it received from Spokes |
| `OtherVrfs` | `[]string` | Optional | By default, the routes are only readvertised toward the same vrf on spoke. To allow it to be leaked to other vrfs |
| `Routed` | `*bool` | Optional | Whether this network is routable |
| `SourceNat` | [`*models.NetworkSourceNat`](../../doc/models/network-source-nat.md) | Optional | If `routed`==`false` (usually at Spoke), but some hosts needs to be reachable from Hub |
| `StaticNat` | [`map[string]models.NetworkVpnAccessStaticNatProperty`](../../doc/models/network-vpn-access-static-nat-property.md) | Optional | Property key may be an External IP address (i.e. "63.16.0.3"), a CIDR (i.e. "63.16.0.12/20") or a Variable (i.e. "{{myvar}}") |
| `SummarizedSubnet` | `*string` | Optional | toward overlay, how HUB should deal with routes it received from Spokes |
| `SummarizedSubnetToLanBgp` | `*string` | Optional | toward LAN-side BGP peers |
| `SummarizedSubnetToLanOspf` | `*string` | Optional | toward LAN-side OSPF peers |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    networkVpnAccessConfig := models.NetworkVpnAccessConfig{
        AdvertisedSubnet:          models.ToPointer("172.16.0.0/24"),
        AllowPing:                 models.ToPointer(false),
        DestinationNat:            map[string]models.NetworkVpnAccessDestinationNatProperty{
            "key0": models.NetworkVpnAccessDestinationNatProperty{
                InternalIp:           models.ToPointer("internal_ip0"),
                Name:                 models.ToPointer("name4"),
                Port:                 models.ToPointer("port4"),
            },
            "key1": models.NetworkVpnAccessDestinationNatProperty{
                InternalIp:           models.ToPointer("internal_ip0"),
                Name:                 models.ToPointer("name4"),
                Port:                 models.ToPointer("port4"),
            },
            "key2": models.NetworkVpnAccessDestinationNatProperty{
                InternalIp:           models.ToPointer("internal_ip0"),
                Name:                 models.ToPointer("name4"),
                Port:                 models.ToPointer("port4"),
            },
        },
        NatPool:                   models.ToPointer("172.16.0.0/26"),
        NoReadvertiseToLanBgp:     models.ToPointer(false),
        NoReadvertiseToLanOspf:    models.ToPointer(false),
        SummarizedSubnet:          models.ToPointer("172.16.0.0/16"),
        SummarizedSubnetToLanBgp:  models.ToPointer("172.16.0.0/16"),
        SummarizedSubnetToLanOspf: models.ToPointer("172.16.0.0/16"),
    }

}
```

