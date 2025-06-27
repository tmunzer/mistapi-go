
# Synthetictest Config Custom Probe

*This model accepts additional fields of type interface{}.*

## Structure

`SynthetictestConfigCustomProbe`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aggressiveness` | [`*models.SynthetictestConfigAggressivenessEnum`](../../doc/models/synthetictest-config-aggressiveness-enum.md) | Optional | enum: `auto`, `high`, `low`<br><br>**Default**: `"auto"` |
| `Host` | `*string` | Optional | If `type`==`icmp` or `type`==`tcp`, Host to be used for the custom probe |
| `Port` | `*int` | Optional | If `type`==`tcp`, Port to be used for the custom probe |
| `Threshold` | `*int` | Optional | In milliseconds |
| `Type` | [`*models.SynthetictestConfigCustomProbeTypeEnum`](../../doc/models/synthetictest-config-custom-probe-type-enum.md) | Optional | enum: `curl`, `icmp`, `tcp`<br><br>**Default**: `"icmp"` |
| `Url` | `*string` | Optional | If `type`==`curl`, URL to be used for the custom probe, can be url or IP |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "aggressiveness": "auto",
  "host": "internal.host",
  "port": 443,
  "threshold": 100,
  "type": "icmp",
  "url": "https://10.3.5.1:8080/about",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

