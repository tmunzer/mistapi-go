
# Response Client Sessions Search

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseClientSessionsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.ResponseClientSessionsSearchItem`](../../doc/models/response-client-sessions-search-item.md) | Required | - |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 230,
  "limit": 196,
  "results": [
    {
      "ap": "ap8",
      "band": "band8",
      "client_manufacture": "client_manufacture0",
      "connect": 206.78,
      "disconnect": 129.48,
      "duration": 104.42,
      "mac": "mac0",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "ssid": "ssid6",
      "timestamp": 2.64,
      "wlan_id": "00000742-0000-0000-0000-000000000000",
      "tags": [
        "tags1",
        "tags2"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 188,
  "total": 222,
  "next": "next4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

