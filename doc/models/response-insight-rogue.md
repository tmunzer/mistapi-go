
# Response Insight Rogue

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseInsightRogue`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | Link to next set of results. If more results aren’t present, next is null. |
| `Results` | [`[]models.InsightRogueAp`](../../doc/models/insight-rogue-ap.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 108,
  "limit": 62,
  "results": [
    {
      "ap_mac": "ap_mac8",
      "avg_rssi": 170.84,
      "bssid": "bssid0",
      "channel": "channel8",
      "delta_x": 25.74,
      "delta_y": 53.12,
      "num_aps": 140,
      "seen_on_lan": false,
      "ssid": "ssid6",
      "times_heard": 110,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 66,
  "next": "next6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

