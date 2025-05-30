
# Response Inventory Error Exception

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseInventoryErrorException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Added` | `[]string` | Optional | - |
| `Duplicated` | `[]string` | Optional | - |
| `Error` | `[]string` | Optional | - |
| `InventoryAdded` | [`[]models.ResponseInventoryInventoryAddedItems`](../../doc/models/response-inventory-inventory-added-items.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `InventoryDuplicated` | [`[]models.ResponseInventoryInventoryDuplicatedItems`](../../doc/models/response-inventory-inventory-duplicated-items.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Reason` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "added": [
    "added8"
  ],
  "duplicated": [
    "duplicated7",
    "duplicated8",
    "duplicated9"
  ],
  "error": [
    "error1",
    "error2"
  ],
  "inventory_added": [
    {
      "mac": "mac0",
      "magic": "magic6",
      "model": "model4",
      "serial": "serial6",
      "type": "type6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "mac": "mac0",
      "magic": "magic6",
      "model": "model4",
      "serial": "serial6",
      "type": "type6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "mac": "mac0",
      "magic": "magic6",
      "model": "model4",
      "serial": "serial6",
      "type": "type6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "inventory_duplicated": [
    {
      "mac": "mac0",
      "magic": "magic6",
      "model": "model4",
      "serial": "serial6",
      "type": "type6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "mac": "mac0",
      "magic": "magic6",
      "model": "model4",
      "serial": "serial6",
      "type": "type6",
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

