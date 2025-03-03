
# Routing Policy Term Matching

zero or more criteria/filter can be specified to match the term, all criteria have to be met

*This model accepts additional fields of type interface{}.*

## Structure

`RoutingPolicyTermMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AsPath` | `[]string` | Optional | takes regular expression |
| `Community` | `[]string` | Optional | - |
| `Network` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Prefix` | `[]string` | Optional | zero or more criteria/filter can be specified to match the term, all criteria have to be met |
| `Protocol` | `[]string` | Optional | `direct`, `bgp`, `osp`, `static`, `aggregate`... |
| `RouteExists` | [`*models.RoutingPolicyTermMatchingRouteExists`](../../doc/models/routing-policy-term-matching-route-exists.md) | Optional | - |
| `VpnNeighborMac` | `[]string` | Optional | overlay-facing criteria (used for bgp_config where via=vpn) |
| `VpnPath` | `[]string` | Optional | overlay-facing criteria (used for bgp_config where via=vpn). ordered- |
| `VpnPathSla` | [`*models.RoutingPolicyTermMatchingVpnPathSla`](../../doc/models/routing-policy-term-matching-vpn-path-sla.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "as_path": [
    "as_path2",
    "as_path1"
  ],
  "community": [
    "community4",
    "community5"
  ],
  "network": [
    "network7"
  ],
  "prefix": [
    "prefix5"
  ],
  "protocol": [
    "protocol5",
    "protocol6",
    "protocol7"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

