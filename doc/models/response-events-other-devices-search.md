
# Response Events Other Devices Search

## Structure

`ResponseEventsOtherDevicesSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Results` | [`*models.EventOtherdevice`](../../doc/models/event-otherdevice.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 206,
  "limit": 220,
  "results": {
    "device_mac": "device_mac0",
    "mac": "mac0",
    "org_id": "00002492-0000-0000-0000-000000000000",
    "site_id": "00001420-0000-0000-0000-000000000000",
    "text": "text4"
  },
  "start": 164,
  "total": 58
}
```

