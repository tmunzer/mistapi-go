
# Response Mxedge Events Search

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseMxedgeEventsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Page` | `*int` | Optional | - |
| `Results` | [`[]models.MxedgeEvent`](../../doc/models/mxedge-event.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1694708579,
  "limit": 10,
  "page": 3,
  "start": 1694622179,
  "results": [
    {
      "component": "component4",
      "mxcluster_id": "mxcluster_id6",
      "mxedge_id": "mxedge_id2",
      "org_id": "00002492-0000-0000-0000-000000000000",
      "service": "service4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "component": "component4",
      "mxcluster_id": "mxcluster_id6",
      "mxedge_id": "mxedge_id2",
      "org_id": "00002492-0000-0000-0000-000000000000",
      "service": "service4",
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

