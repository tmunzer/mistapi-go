
# Vrf Instance

## Structure

`VrfInstance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExtraRoutes` | [`map[string]models.VrfExtraRoute`](../../doc/models/vrf-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8") |
| `Networks` | `[]string` | Optional | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "extra_routes": {
    "0.0.0.0/0": {
      "via": "192.168.31.1"
    }
  },
  "networks": [
    "guest"
  ]
}
```

