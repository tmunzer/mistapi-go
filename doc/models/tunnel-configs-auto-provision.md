
# Tunnel Configs Auto Provision

## Structure

`TunnelConfigsAutoProvision`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enable` | `*bool` | Optional | - |
| `Latlng` | [`*models.LatLng`](../../doc/models/lat-lng.md) | Optional | - |
| `Primary` | [`*models.TunnelConfigsAutoProvisionNode`](../../doc/models/tunnel-configs-auto-provision-node.md) | Optional | - |
| `Secondary` | [`*models.TunnelConfigsAutoProvisionNode`](../../doc/models/tunnel-configs-auto-provision-node.md) | Optional | - |

## Example (as JSON)

```json
{
  "enable": false,
  "latlng": {
    "lat": 144.64,
    "lng": 22.82
  },
  "primary": {
    "num_hosts": "num_hosts6",
    "wan_names": [
      "wan_names8"
    ]
  },
  "secondary": {
    "num_hosts": "num_hosts8",
    "wan_names": [
      "wan_names0"
    ]
  }
}
```

