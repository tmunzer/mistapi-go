
# Network Vpn Access Config

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
| `StaticNat` | [`map[string]models.NetworkVpnAccessStaticNatProperty`](../../doc/models/network-vpn-access-static-nat-property.md) | Optional | Property key may be an External IP Address (i.e. "63.16.0.3"), a CIDR (i.e. "63.16.0.12/20") or a Variable (i.e. "{{myvar}}") |
| `SummarizedSubnet` | `*string` | Optional | toward overlay, how HUB should deal with routes it received from Spokes |
| `SummarizedSubnetToLanBgp` | `*string` | Optional | toward LAN-side BGP peers |
| `SummarizedSubnetToLanOspf` | `*string` | Optional | toward LAN-side OSPF peers |

## Example (as JSON)

```json
{
  "advertised_subnet": "172.16.0.0/24",
  "nat_pool": "172.16.0.0/26",
  "no_readvertise_to_lan_bgp": false,
  "no_readvertise_to_lan_ospf": false,
  "summarized_subnet": "172.16.0.0/16",
  "summarized_subnet_to_lan_bgp": "172.16.0.0/16",
  "summarized_subnet_to_lan_ospf": "172.16.0.0/16",
  "allow_ping": false,
  "destination_nat": {
    "key0": {
      "internal_ip": "internal_ip0",
      "name": "name4",
      "port": "port4"
    },
    "key1": {
      "internal_ip": "internal_ip0",
      "name": "name4",
      "port": "port4"
    }
  }
}
```

