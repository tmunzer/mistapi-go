
# Inventory Search

## Structure

`InventorySearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Limit` | `*int` | Optional | - |
| `Page` | `*int` | Optional | - |
| `Results` | [`[]models.InventorySearchResult`](../../doc/models/inventory-search-result.md) | Optional | - |

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
          "serial": "serial8"
        },
        {
          "mac": "mac2",
          "model": "model6",
          "serial": "serial8"
        }
      ],
      "model": "model4",
      "name": "name6"
    }
  ]
}
```

