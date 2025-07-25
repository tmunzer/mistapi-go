
# Switch Vrf Instance

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchVrfInstance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AggregateRoutes` | [`map[string]models.AggregateRoute`](../../doc/models/aggregate-route.md) | Optional | Property key is the destination subnet (e.g. "172.16.3.0/24") |
| `AggregateRoutes6` | [`map[string]models.AggregateRoute`](../../doc/models/aggregate-route.md) | Optional | Property key is the destination subnet (e.g. "2a02:1234:420a:10c9::/64") |
| `EvpnAutoLoopbackSubnet` | `*string` | Optional | - |
| `EvpnAutoLoopbackSubnet6` | `*string` | Optional | - |
| `ExtraRoutes` | [`map[string]models.VrfExtraRoute`](../../doc/models/vrf-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8") |
| `ExtraRoutes6` | [`map[string]models.VrfExtraRoute`](../../doc/models/vrf-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64") |
| `Networks` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "extra_routes": {
    "0.0.0.0/0": {
      "via": "192.168.31.1",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "networks": [
    "guest"
  ],
  "aggregate_routes": {
    "key0": null,
    "key1": {}
  },
  "aggregate_routes6": {
    "key0": null,
    "key1": {},
    "key2": {}
  },
  "evpn_auto_loopback_subnet": "evpn_auto_loopback_subnet2",
  "evpn_auto_loopback_subnet6": "evpn_auto_loopback_subnet66",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

