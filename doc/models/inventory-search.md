
# Inventory Search

*This model accepts additional fields of type interface{}.*

## Structure

`InventorySearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Limit` | `*int` | Optional | - |
| `Page` | `*int` | Optional | - |
| `Results` | [`[]models.InventorySearchResult`](../../doc/models/inventory-search-result.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "limit": 1000,
  "page": 1,
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
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

