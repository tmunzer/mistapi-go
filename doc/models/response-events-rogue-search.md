
# Response Events Rogue Search

## Structure

`ResponseEventsRogueSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.EventsRogue`](../../doc/models/events-rogue.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 160,
  "limit": 10,
  "results": [
    {
      "ap": "ap8",
      "bssid": "bssid0",
      "channel": 30,
      "rssi": 68,
      "ssid": "ssid6",
      "timestamp": 2.64
    }
  ],
  "start": 118,
  "total": 104,
  "next": "next8"
}
```

