
# Response Zone Search

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

## Example (as JSON)

```json
{
  "end": 1541705289.769911,
  "limit": 1,
  "next": "/api/v1/sites/67970e46-4e12-11e6-9188-0242ac110007/zones/visits/search?limit=2&end=1541705247.000&scope_id=85fbba9e-4e12-11e6-9188-0242ac110007&user_type=asset&start=1541618889.77",
  "start": 1541618889.769886,
  "total": 5892,
  "results": [
    {
      "enter": 226,
      "scope": "scope6",
      "timestamp": 2.64,
      "user": "user6"
    },
    {
      "enter": 226,
      "scope": "scope6",
      "timestamp": 2.64,
      "user": "user6"
    }
  ]
}
```

