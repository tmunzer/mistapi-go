
# Response Search Var

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSearchVar`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Results` | [`[]models.ResponseSearchVarItem`](../../doc/models/response-search-var-item.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 44,
  "limit": 126,
  "results": [
    {
      "created_time": 73.76,
      "modified_time": 5.2,
      "org_id": "00002492-0000-0000-0000-000000000000",
      "site_id": "00001420-0000-0000-0000-000000000000",
      "src": "src8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 2,
  "total": 220,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

