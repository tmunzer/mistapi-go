
# Js Inventory Search

*This model accepts additional fields of type interface{}.*

## Structure

`JsInventorySearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Offset to end at |
| `Limit` | `*int` | Optional | Number of results to return |
| `Results` | [`[]models.JsInventoryItem`](../../doc/models/js-inventory-item.md) | Optional | - |
| `Start` | `*int` | Optional | Offset to start from |
| `Total` | `*int` | Optional | Total number of results |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 82,
  "limit": 88,
  "results": [
    {
      "device_name": "device_name6",
      "eol_time": 174,
      "eos_time": 220,
      "master": false,
      "model": "model4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 40,
  "total": 74,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

