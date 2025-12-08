
# Response Device Events Search

## Structure

`ResponseDeviceEventsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.DeviceEvent`](../../doc/models/device-event.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 162,
  "limit": 8,
  "results": [
    {
      "audit_id": "53f10664-3ce8-4c27-b382-0ef66432349f",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "timestamp": 2.64,
      "type": "type4",
      "ap": "ap8",
      "ap_name": "ap_name6",
      "apfw": "apfw8",
      "bandwidth": 198
    }
  ],
  "start": 120,
  "total": 154,
  "next": "next6"
}
```

