
# Synthetictest Config Custom Probe

## Structure

`SynthetictestConfigCustomProbe`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aggressiveness` | [`*models.SynthetictestConfigAggressivenessEnum`](../../doc/models/synthetictest-config-aggressiveness-enum.md) | Optional | enum: `auto`, `high`, `low`<br><br>**Default**: `"auto"` |
| `Target` | `*string` | Optional | Can be URL (e.g. http://x.com, https://x.com:8080/path/to/resource), IP address, or IP:port combination |
| `Threshold` | `*int` | Optional | In milliseconds |
| `Type` | [`*models.SynthetictestConfigCustomProbeTypeEnum`](../../doc/models/synthetictest-config-custom-probe-type-enum.md) | Optional | enum: `application`, `curl`, `icmp`, `reachability`, `tcp`<br><br>**Default**: `"icmp"` |

## Example (as JSON)

```json
{
  "aggressiveness": "auto",
  "target": "10.3.5.3:8080",
  "threshold": 100,
  "type": "icmp"
}
```

