
# Response Client Events Search

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseClientEventsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.EventsClient`](../../doc/models/events-client.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 20,
  "limit": 150,
  "results": [
    {
      "ap": "ap8",
      "band": "6",
      "bssid": "bssid0",
      "channel": 30,
      "proto": "a",
      "ssid": "ssid6",
      "text": "text4",
      "timestamp": 2.64,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 234,
  "total": 244,
  "next": "next4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

