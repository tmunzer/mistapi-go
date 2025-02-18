
# Response Inventory Inventory Duplicated Items

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseInventoryInventoryDuplicatedItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | - |
| `Magic` | `string` | Required | - |
| `Model` | `string` | Required | - |
| `Serial` | `string` | Required | - |
| `Type` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "5c5b35000012",
  "magic": "DVH4VSNMSZPDXBR",
  "model": "AP41",
  "serial": "FXLH2015150027",
  "type": "ap",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

