
# Response Zone Search

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseZoneSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*float64` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.ResponseZoneSearchItem`](../../doc/models/response-zone-search-item.md) | Optional | - |
| `Start` | `*float64` | Optional | - |
| `Total` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1541705289.76991,
  "limit": 1,
  "next": "/api/v1/sites/67970e46-4e12-11e6-9188-0242ac110007/zones/visits/search?limit=2&end=1541705247.000&scope_id=85fbba9e-4e12-11e6-9188-0242ac110007&user_type=asset&start=1541618889.77",
  "start": 1541618889.76989,
  "total": 5892,
  "results": [
    {
      "enter": 226,
      "scope": "scope6",
      "timestamp": 8,
      "user": "user6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "enter": 226,
      "scope": "scope6",
      "timestamp": 8,
      "user": "user6",
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

