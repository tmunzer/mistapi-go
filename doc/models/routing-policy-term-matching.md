
# Routing Policy Term Matching

zero or more criteria/filter can be specified to match the term, all criteria have to be met

## Structure

`RoutingPolicyTermMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AsPath` | `[]string` | Optional | takes regular expression |
| `Community` | `[]string` | Optional | - |
| `Network` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Prefix` | `[]string` | Optional | zero or more criteria/filter can be specified to match the term, all criteria have to be met |
| `Protocol` | `[]string` | Optional | `direct`, `bgp`, `osp`, ... |
| `RouteExists` | [`*models.RoutingPolicyTermMatchingRouteExists`](../../doc/models/routing-policy-term-matching-route-exists.md) | Optional | - |
| `VpnNeighborMac` | `[]string` | Optional | overlay-facing criteria (used for bgp_config where via=vpn) |
| `VpnPath` | `[]string` | Optional | overlay-facing criteria (used for bgp_config where via=vpn)<br>ordered- |
| `VpnPathSla` | [`*models.RoutingPolicyTermMatchingVpnPathSla`](../../doc/models/routing-policy-term-matching-vpn-path-sla.md) | Optional | - |

## Example (as JSON)

```json
{
  "as_path": [
    "as_path4"
  ],
  "community": [
    "community2"
  ],
  "network": [
    "network1",
    "network2"
  ],
  "prefix": [
    "prefix3",
    "prefix4",
    "prefix5"
  ],
  "protocol": [
    "protocol3",
    "protocol4"
  ]
}
```

