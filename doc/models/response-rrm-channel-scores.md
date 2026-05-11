
# Response Rrm Channel Scores

## Structure

`ResponseRrmChannelScores`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Results` | [`[]models.RrmChannelScore`](../../doc/models/rrm-channel-score.md) | Required | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "results": [
    {
      "channel": 30,
      "util_score": 76.78,
      "util_score_noise_floor": 140.42,
      "util_score_non_wifi": 74.88,
      "util_score_other": 108.4,
      "util_score_radar": 115.1,
      "util_score_undecodable_wifi": 176.2,
      "util_score_unknown_wifi": 27.24
    }
  ]
}
```

