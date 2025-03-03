
# Switch Vrf Instance

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchVrfInstance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AggregateRoutes` | [`map[string]models.AggregateRoute`](../../doc/models/aggregate-route.md) | Optional | Property key is the destination subnet (e.g. "172.16.3.0/24") |
| `EvpnAutoLookbackSubnet` | `*string` | Optional | - |
| `ExtraRoutes` | [`map[string]models.VrfExtraRoute`](../../doc/models/vrf-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8") |
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
    "key0": {
      "discard": false,
      "metric": 28,
      "preference": 148,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "discard": false,
      "metric": 28,
      "preference": 148,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "evpn_auto_lookback_subnet": "evpn_auto_lookback_subnet0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

