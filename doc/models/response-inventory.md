
# Response Inventory

## Structure

`ResponseInventory`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Added` | `[]string` | Optional | - |
| `Duplicated` | `[]string` | Optional | - |
| `Error` | `[]string` | Optional | - |
| `InventoryAdded` | [`[]models.ResponseInventoryInventoryAddedItems`](../../doc/models/response-inventory-inventory-added-items.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `InventoryDuplicated` | [`[]models.ResponseInventoryInventoryDuplicatedItems`](../../doc/models/response-inventory-inventory-duplicated-items.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

## Example (as JSON)

```json
{
  "added": [
    "added0",
    "added1",
    "added2"
  ],
  "duplicated": [
    "duplicated9",
    "duplicated0"
  ],
  "error": [
    "error3"
  ],
  "inventory_added": [
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

