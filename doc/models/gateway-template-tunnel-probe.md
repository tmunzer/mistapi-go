
# Gateway Template Tunnel Probe

Only if `provider`== `custom-ipsec`

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayTemplateTunnelProbe`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interval` | `*int` | Optional | how often to trigger the probe |
| `Threshold` | `*int` | Optional | number of consecutive misses before declaring the tunnel down |
| `Timeout` | `*int` | Optional | time within which to complete the connectivity check |
| `Type` | [`*models.GatewayTemplateProbeTypeEnum`](../../doc/models/gateway-template-probe-type-enum.md) | Optional | enum: `http`, `icmp`<br>**Default**: `"icmp"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "type": "icmp",
  "interval": 192,
  "threshold": 252,
  "timeout": 60,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

