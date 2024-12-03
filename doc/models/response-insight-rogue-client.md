
# Response Insight Rogue Client

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseInsightRogueClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `string` | Required | - |
| `Results` | [`[]models.InsightRogueClient`](../../doc/models/insight-rogue-client.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 226,
  "limit": 200,
  "next": "next2",
  "results": [
    {
      "annotation": "annotation0",
      "ap_mac": "ap_mac8",
      "avg_rssi": 170.84,
      "band": "band8",
      "bssid": "bssid0",
      "client_mac": "client_mac4",
      "num_aps": 140,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 184,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

