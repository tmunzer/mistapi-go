
# Ospf Config

Junos OSPF config

## Structure

`OspfConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Areas` | [`map[string]models.OspfConfigArea`](../../doc/models/ospf-config-area.md) | Optional | OSPF areas to run on this device and the corresponding per-area-specific configs. Property key is the area |
| `Enabled` | `*bool` | Optional | whether to rung OSPF on this device |
| `ReferenceBandwidth` | `*string` | Optional | Bandwidth for calculating metric defaults (9600..4000000000000)<br>**Default**: `"100M"` |

## Example (as JSON)

```json
{
  "reference_bandwidth": "100M",
  "areas": {
    "key0": {
      "no_summary": false
    }
  },
  "enabled": false
}
```

