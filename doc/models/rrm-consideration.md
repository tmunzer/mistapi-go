
# Rrm Consideration

## Structure

`RrmConsideration`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `int` | Required | - |
| `Noise` | `float64` | Required | - |
| `OtherRssi` | `*float64` | Optional | the avg RSSI heard from other APs (that does NOT belongs to the same site) |
| `OtherSsid` | `*string` | Optional | SSID from other AP that we heard from with the max RSSI |
| `UtilScore` | `float64` | Required | utilization score, 0-1, lower means less utilization (cleaner RF) |
| `UtilScoreNonWifi` | `float64` | Required | non-wifi utilization score, 0-1, lower means less utilization (cleaner RF) |
| `UtilScoreOther` | `float64` | Required | other utilization score, 0-1, lower means less utilization (cleaner RF) |

## Example (as JSON)

```json
{
  "channel": 38,
  "noise": 50.3,
  "other_rssi": 128.0,
  "other_ssid": "other_ssid0",
  "util_score": 171.42,
  "util_score_non_wifi": 19.76,
  "util_score_other": 203.04
}
```

