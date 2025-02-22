
# Rrm Consideration

*This model accepts additional fields of type interface{}.*

## Structure

`RrmConsideration`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `int` | Required | - |
| `Noise` | `float64` | Required | - |
| `OtherRssi` | `*float64` | Optional | Avg RSSI heard from other APs (that does NOT belongs to the same site) |
| `OtherSsid` | `*string` | Optional | SSID from other AP that we heard from with the max RSSI |
| `UtilScore` | `float64` | Required | utilization score, 0-1, lower means less utilization (cleaner RF) |
| `UtilScoreNonWifi` | `float64` | Required | non-Wi-Fi utilization score, 0-1, lower means less utilization (cleaner RF) |
| `UtilScoreOther` | `float64` | Required | other utilization score, 0-1, lower means less utilization (cleaner RF) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "channel": 144,
  "noise": 137.64,
  "other_rssi": 196.06,
  "other_ssid": "other_ssid6",
  "util_score": 16.52,
  "util_score_non_wifi": 168.18,
  "util_score_other": 15.1,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

