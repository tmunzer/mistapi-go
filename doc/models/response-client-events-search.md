
# Response Client Events Search

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

## Example (as JSON)

```json
{
  "end": 10,
  "limit": 96,
  "results": [
    {
      "ap": "ap8",
      "band": "6",
      "bssid": "bssid0",
      "channel": 30,
      "proto": "a",
      "ssid": "ssid6",
      "text": "text4",
      "timestamp": 2.64
    }
  ],
  "start": 224,
  "total": 2,
  "next": "next4"
}
```

