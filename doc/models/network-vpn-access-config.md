
# Network Vpn Access Config

*This model accepts additional fields of type interface{}.*

## Structure

`NetworkVpnAccessConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdvertisedSubnet` | `*string` | Optional | if `routed`==`true`, whether to advertise an aggregated subnet toward HUB this is useful when there are multiple networks on SPOKE's side |
| `AllowPing` | `*bool` | Optional | whether to allow ping from vpn into this routed network |
| `DestinationNat` | [`map[string]models.NetworkDestinationNatProperty`](../../doc/models/network-destination-nat-property.md) | Optional | Property key may be an IP/Port (i.e. "63.16.0.3:443"), or a port (i.e. ":2222") |
| `NatPool` | `*string` | Optional | if `routed`==`false` (usually at Spoke), but some hosts needs to be reachable from Hub, a subnet is required to create and advertise the route to Hub |
| `NoReadvertiseToLanBgp` | `*bool` | Optional | toward LAN-side BGP peers<br>**Default**: `false` |
| `NoReadvertiseToLanOspf` | `*bool` | Optional | toward LAN-side OSPF peers<br>**Default**: `false` |
| `NoReadvertiseToOverlay` | `*bool` | Optional | toward overlay<br>how HUB should deal with routes it received from Spokes |
| `OtherVrfs` | `[]string` | Optional | by default, the routes are only readvertised toward the same vrf on spoke<br>to allow it to be leaked to other vrfs |
| `Routed` | `*bool` | Optional | whether this network is routable |
| `SourceNat` | [`*models.NetworkSourceNat`](../../doc/models/network-source-nat.md) | Optional | if `routed`==`false` (usually at Spoke), but some hosts needs to be reachable from Hub |
| `StaticNat` | [`map[string]models.NetworkStaticNatProperty`](../../doc/models/network-static-nat-property.md) | Optional | Property key may be an IP Address (i.e. "172.16.0.1"), and IP Address and Port (i.e. "172.16.0.1:8443") or a CIDR (i.e. "172.16.0.12/20") |
| `SummarizedSubnet` | `*string` | Optional | toward overlay<br>how HUB should deal with routes it received from Spokes |
| `SummarizedSubnetToLanBgp` | `*string` | Optional | toward LAN-side BGP peers |
| `SummarizedSubnetToLanOspf` | `*string` | Optional | toward LAN-side OSPF peers |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
      "port": 162,
      "wan_name": "wan_name0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "internal_ip": "internal_ip0",
      "name": "name4",
      "port": 162,
      "wan_name": "wan_name0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

