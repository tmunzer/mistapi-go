
# Tunnel Configs Auto Provision

## Structure

`TunnelConfigsAutoProvision`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enable` | `*bool` | Optional | - |
| `Latlng` | [`*models.LatLng`](../../doc/models/lat-lng.md) | Optional | - |
| `Primary` | [`*models.TunnelConfigsAutoProvisionNode`](../../doc/models/tunnel-configs-auto-provision-node.md) | Optional | - |
| `Region` | [`*models.TunnelConfigsAutoProvisionRegionEnum`](../../doc/models/tunnel-configs-auto-provision-region-enum.md) | Optional | enum: `APAC`, `Americas`, `EMEA`, `auto`<br>**Default**: `"auto"` |
| `Secondary` | [`*models.TunnelConfigsAutoProvisionNode`](../../doc/models/tunnel-configs-auto-provision-node.md) | Optional | - |

## Example (as JSON)

```json
{
  "region": "auto",
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

