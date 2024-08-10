
# Response Events Fastroam

## Structure

`ResponseEventsFastroam`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | the link to query next set of results. value is null if no next page exists. |
| `Results` | [`[]models.EventFastroam`](../../doc/models/event-fastroam.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 182,
  "limit": 244,
  "results": [
    {
      "ap_mac": "ap_mac8",
      "client_mac": "client_mac4",
      "fromap": "fromap6",
      "latency": 250.14,
      "ssid": "ssid6",
      "subtype": "subtype8",
      "timestamp": 2.64,
      "type": "pingpong"
    }
  ],
  "start": 140,
  "next": "next6"
}
```

