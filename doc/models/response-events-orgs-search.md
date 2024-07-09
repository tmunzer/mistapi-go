
# Response Events Orgs Search

## Structure

`ResponseEventsOrgsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Results` | [`[]models.OrgEvent`](../../doc/models/org-event.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1688035193,
  "limit": 10,
  "start": 1687948793,
  "results": [
    {
      "org_id": "00002492-0000-0000-0000-000000000000",
      "text": "text4",
      "timestamp": 2.64,
      "type": "type4"
    },
    {
      "org_id": "00002492-0000-0000-0000-000000000000",
      "text": "text4",
      "timestamp": 2.64,
      "type": "type4"
    }
  ],
  "total": 198
}
```

