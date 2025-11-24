
# Inventory Search

*This model accepts additional fields of type interface{}.*

## Structure

`InventorySearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.InventorySearchResult`](../../doc/models/inventory-search-result.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "limit": 1000,
  "total": 1,
  "end": 206,
  "next": "next6",
  "results": [
    {
      "mac": "mac0",
      "master": false,
      "members": [
        {
          "mac": "mac2",
          "model": "model6",
          "serial": "serial8",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "mac": "mac2",
          "model": "model6",
          "serial": "serial8",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
      "model": "model4",
      "name": "name6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 164,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

