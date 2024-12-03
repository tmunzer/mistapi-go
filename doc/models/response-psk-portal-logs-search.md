
# Response Psk Portal Logs Search

*This model accepts additional fields of type interface{}.*

## Structure

`ResponsePskPortalLogsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Results` | [`[]models.ResponsePskPortalLogsSearchItem`](../../doc/models/response-psk-portal-logs-search-item.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1428954000,
  "limit": 100,
  "start": 1428939600,
  "total": 135,
  "results": [
    {
      "id": "000023ba-0000-0000-0000-000000000000",
      "message": "message6",
      "name_id": "name_id8",
      "org_id": "00002492-0000-0000-0000-000000000000",
      "psk_id": "00000b40-0000-0000-0000-000000000000",
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

