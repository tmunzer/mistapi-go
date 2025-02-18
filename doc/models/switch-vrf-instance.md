
# Switch Vrf Instance

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchVrfInstance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
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
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

