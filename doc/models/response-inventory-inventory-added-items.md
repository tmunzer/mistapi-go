
# Response Inventory Inventory Added Items

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseInventoryInventoryAddedItems`

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
  "mac": "5c5b35000018",
  "magic": "6JG8EPTFV2A9Z2N",
  "model": "AP41",
  "serial": "FXLH2015150025",
  "type": "ap",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

