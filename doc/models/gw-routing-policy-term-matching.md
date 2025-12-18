
# Gw Routing Policy Term Matching

zero or more criteria/filter can be specified to match the term, all criteria have to be met

## Structure

`GwRoutingPolicyTermMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AsPath` | [`[]models.BgpAs`](../../doc/models/containers/bgp-as.md) | Optional | BGP AS, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}` ) |
| `Community` | `[]string` | Optional | - |
| `Network` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Prefix` | `[]string` | Optional | zero or more criteria/filter can be specified to match the term, all criteria have to be met |
| `Protocol` | [`[]models.GwRoutingPolicyTermMatchingProtocolEnum`](../../doc/models/gw-routing-policy-term-matching-protocol-enum.md) | Optional | - |
| `RouteExists` | [`*models.GwRoutingPolicyTermMatchingRouteExists`](../../doc/models/gw-routing-policy-term-matching-route-exists.md) | Optional | - |
| `VpnNeighborMac` | `[]string` | Optional | overlay-facing criteria (used for bgp_config where via=vpn) |
| `VpnPath` | `[]string` | Optional | overlay-facing criteria (used for bgp_config where via=vpn). ordered- |
| `VpnPathSla` | [`*models.GwRoutingPolicyTermMatchingVpnPathSla`](../../doc/models/gw-routing-policy-term-matching-vpn-path-sla.md) | Optional | - |

## Example (as JSON)

```json
{
  "as_path": [
    "String3",
    "String2",
    "String1"
  ],
  "community": [
    "community4",
    "community5",
    "community6"
  ],
  "network": [
    "network7",
    "network8"
  ],
  "prefix": [
    "prefix5",
    "prefix6"
  ],
  "protocol": [
    "aggregate"
  ]
}
```

