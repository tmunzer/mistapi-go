
# Response Rrm Consideration

## Structure

`ResponseRrmConsideration`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Results` | [`[]models.RrmConsideration`](../../doc/models/rrm-consideration.md) | Required | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "results": [
    {
      "channel": 30,
      "noise": 44.34,
      "other_rssi": 33.36,
      "other_ssid": "other_ssid6",
      "rssi": 128.68,
      "util_score": 76.78,
      "util_score_non_wifi": 74.88,
      "util_score_other": 108.4
    }
  ]
}
```

