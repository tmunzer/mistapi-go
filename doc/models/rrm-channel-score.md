
# Rrm Channel Score

## Structure

`RrmChannelScore`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `int` | Required | Channel number |
| `UtilScore` | `float64` | Required | Utilization score for the channel, 0-1, lower means cleaner RF |
| `UtilScoreNoiseFloor` | `float64` | Required | Score contribution from noise, 0-1, lower means cleaner RF |
| `UtilScoreNonWifi` | `float64` | Required | Score contribution from non-wifi utilization, 0-1, lower means cleaner RF |
| `UtilScoreOther` | `float64` | Required | Score contribution from RxOtherBss utilization (wifi packets destined for other radios), 0-1, lower means cleaner RF |
| `UtilScoreRadar` | `float64` | Required | Score contribution from radar detections, 0-1, lower means cleaner RF |
| `UtilScoreUndecodableWifi` | `float64` | Required | Score contribution from undecodable wifi utilization (wifi packets which can't be decoded), 0-1, lower means cleaner RF |
| `UtilScoreUnknownWifi` | `float64` | Required | Score contribution from unknown wifi utilization (wifi packets of unknown type), 0-1, lower means cleaner RF |

## Example (as JSON)

```json
{
  "channel": 18,
  "util_score": 43.62,
  "util_score_noise_floor": 107.26,
  "util_score_non_wifi": 108.04,
  "util_score_other": 75.24,
  "util_score_radar": 81.94,
  "util_score_undecodable_wifi": 143.04,
  "util_score_unknown_wifi": 250.08
}
```

