
# Response Search Var

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

## Example (as JSON)

```json
{
  "end": 18,
  "limit": 152,
  "results": [
    {
      "created_time": 73.76,
      "modified_time": 5.2,
      "org_id": "00002492-0000-0000-0000-000000000000",
      "site_id": "00001420-0000-0000-0000-000000000000",
      "src": "src8"
    }
  ],
  "start": 232,
  "total": 10
}
```

