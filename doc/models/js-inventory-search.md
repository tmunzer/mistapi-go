
# Js Inventory Search

## Structure

`JsInventorySearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Offset to end at |
| `Limit` | `*int` | Optional | Number of results to return |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.JsInventoryItem`](../../doc/models/js-inventory-item.md) | Optional | - |
| `Start` | `*int` | Optional | Offset to start from |
| `Total` | `*int` | Optional | Total number of results |

## Example (as JSON)

```json
{
  "end": 82,
  "limit": 88,
  "next": "next0",
  "results": [
    {
      "claimed": false,
      "device_name": "device_name6",
      "eol_time": 174,
      "eos_time": 220,
      "has_support": false
    }
  ],
  "start": 40
}
```

