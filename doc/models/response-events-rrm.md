
# Response Events Rrm

## Structure

`ResponseEventsRrm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | the link to query next set of results. value is null if no next page exists. |
| `Results` | [`[]models.RrmEvent`](../../doc/models/rrm-event.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 234,
  "limit": 192,
  "results": [
    {
      "ap_id": "00001b9c-0000-0000-0000-000000000000",
      "band": "6",
      "bandwidth": 20,
      "channel": 30,
      "event": "radar-detected",
      "power": 226,
      "pre_bandwidth": 80,
      "pre_channel": 164,
      "pre_power": 92.2,
      "pre_usage": "pre_usage6",
      "timestamp": 2.64,
      "usage": "usage4"
    }
  ],
  "start": 192,
  "next": "next4"
}
```

