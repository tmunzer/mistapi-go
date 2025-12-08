
# Response Events Path Search

## Structure

`ResponseEventsPathSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.ServicePathEvent`](../../doc/models/service-path-event.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1697096379,
  "limit": 10,
  "start": 1697009979,
  "total": 2,
  "next": "next4",
  "results": [
    {
      "mac": "mac0",
      "model": "model4",
      "org_id": "00002492-0000-0000-0000-000000000000",
      "policy": "policy8",
      "port_id": "port_id6"
    },
    {
      "mac": "mac0",
      "model": "model4",
      "org_id": "00002492-0000-0000-0000-000000000000",
      "policy": "policy8",
      "port_id": "port_id6"
    },
    {
      "mac": "mac0",
      "model": "model4",
      "org_id": "00002492-0000-0000-0000-000000000000",
      "policy": "policy8",
      "port_id": "port_id6"
    }
  ]
}
```

