
# Response Events Search

## Structure

`ResponseEventsSearch`

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
  "end": 24,
  "limit": 146,
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
  "start": 238,
  "total": 16,
  "next": "next0"
}
```

