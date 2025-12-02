
# Response Mxedge Events Search

## Structure

`ResponseMxedgeEventsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Next` | `*string` | Optional | - |
| `Page` | `*int` | Optional | - |
| `Results` | [`[]models.MxedgeEvent`](../../doc/models/mxedge-event.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Start` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1694708579,
  "limit": 10,
  "page": 3,
  "start": 1694622179,
  "next": "next4",
  "results": [
    {
      "audit_id": "00000d00-0000-0000-0000-000000000000",
      "component": "component4",
      "device_type": "device_type4",
      "from_version": "from_version8",
      "mac": "mac0"
    },
    {
      "audit_id": "00000d00-0000-0000-0000-000000000000",
      "component": "component4",
      "device_type": "device_type4",
      "from_version": "from_version8",
      "mac": "mac0"
    }
  ]
}
```

