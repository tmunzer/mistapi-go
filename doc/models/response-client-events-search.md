
# Response Client Events Search

## Structure

`ResponseClientEventsSearch`

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
  "end": 20,
  "limit": 150,
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
  "start": 234,
  "total": 244,
  "next": "next4"
}
```

