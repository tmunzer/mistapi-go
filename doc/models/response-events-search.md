
# Response Events Search

## Structure

`ResponseEventsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.EventsClient`](../../doc/models/events-client.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 36,
  "limit": 134,
  "results": [
    {
      "band": "6",
      "key_mgmt": "WPA2-PSK",
      "timestamp": 2.64,
      "ap": "ap8",
      "bssid": "bssid0",
      "channel": 30,
      "proto": "b"
    }
  ],
  "start": 250,
  "total": 28,
  "next": "next0"
}
```

