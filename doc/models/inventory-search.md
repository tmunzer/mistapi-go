
# Inventory Search

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
  ],
  "start": 164
}
```

