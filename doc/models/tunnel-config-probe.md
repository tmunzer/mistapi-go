
# Tunnel Config Probe

Only if `provider`==`custom-ipsec`

## Structure

`TunnelConfigProbe`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interval` | `*int` | Optional | How often to trigger the probe |
| `Threshold` | `*int` | Optional | Number of consecutive misses before declaring the tunnel down |
| `Timeout` | `*int` | Optional | Time within which to complete the connectivity check |
| `Type` | [`*models.TunnelConfigProbeTypeEnum`](../../doc/models/tunnel-config-probe-type-enum.md) | Optional | enum: `http`, `icmp`<br><br>**Default**: `"icmp"` |

## Example (as JSON)

```json
{
  "type": "icmp",
  "interval": 6,
  "threshold": 54,
  "timeout": 118
}
```

