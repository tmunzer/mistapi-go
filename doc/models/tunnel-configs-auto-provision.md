
# Tunnel Configs Auto Provision

*This model accepts additional fields of type interface{}.*

## Structure

`TunnelConfigsAutoProvision`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enable` | `*bool` | Optional | - |
| `Latlng` | [`*models.LatLng`](../../doc/models/lat-lng.md) | Optional | - |
| `Primary` | [`*models.TunnelConfigsAutoProvisionNode`](../../doc/models/tunnel-configs-auto-provision-node.md) | Optional | - |
| `Secondary` | [`*models.TunnelConfigsAutoProvisionNode`](../../doc/models/tunnel-configs-auto-provision-node.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enable": false,
  "latlng": {
    "lat": 144.64,
    "lng": 22.82,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "primary": {
    "num_hosts": "num_hosts6",
    "wan_names": [
      "wan_names8"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "secondary": {
    "num_hosts": "num_hosts8",
    "wan_names": [
      "wan_names0"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

