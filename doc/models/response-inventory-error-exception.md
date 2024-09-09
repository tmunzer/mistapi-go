
# Response Inventory Error Exception

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
      "type": "type6"
    },
    {
      "mac": "mac0",
      "magic": "magic6",
      "model": "model4",
      "serial": "serial6",
      "type": "type6"
    },
    {
      "mac": "mac0",
      "magic": "magic6",
      "model": "model4",
      "serial": "serial6",
      "type": "type6"
    }
  ],
  "inventory_duplicated": [
    {
      "mac": "mac0",
      "magic": "magic6",
      "model": "model4",
      "serial": "serial6",
      "type": "type6"
    },
    {
      "mac": "mac0",
      "magic": "magic6",
      "model": "model4",
      "serial": "serial6",
      "type": "type6"
    }
  ]
}
```

