
# Response Events Orgs Search

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseEventsOrgsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.OrgEvent`](../../doc/models/org-event.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1688035193,
  "limit": 10,
  "start": 1687948793,
  "next": "next4",
  "results": [
    {
      "org_id": "00002492-0000-0000-0000-000000000000",
      "text": "text4",
      "timestamp": 2.64,
      "type": "type4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "org_id": "00002492-0000-0000-0000-000000000000",
      "text": "text4",
      "timestamp": 2.64,
      "type": "type4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

