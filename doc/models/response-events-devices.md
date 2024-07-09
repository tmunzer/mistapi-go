
# Response Events Devices

## Structure

`ResponseEventsDevices`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `string` | Required | - |
| `Results` | [`[]models.EventsDeviceAp`](../../doc/models/events-device-ap.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 236,
  "limit": 190,
  "next": "next0",
  "results": [
    {
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "timestamp": 2.64,
      "ap": "ap8",
      "apfw": "apfw8",
      "count": 226,
      "device_type": "device_type4",
      "mac": "mac0"
    }
  ],
  "start": 194,
  "total": 28
}
```

