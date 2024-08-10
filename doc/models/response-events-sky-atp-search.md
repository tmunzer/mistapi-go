
# Response Events Sky Atp Search

## Structure

`ResponseEventsSkyAtpSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.EventsSkyatp`](../../doc/models/events-skyatp.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 218,
  "limit": 208,
  "results": [
    {
      "device_mac": "device_mac0",
      "ip": "ip0",
      "mac": "mac0",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "threat_level": 152,
      "timestamp": 2.64,
      "type": "type4",
      "for_site": false
    }
  ],
  "start": 176,
  "total": 46,
  "next": "next4"
}
```

