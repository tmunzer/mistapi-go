
# Jse Inventory Item

*This model accepts additional fields of type interface{}.*

## Structure

`JseInventoryItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ContractEndTime` | `*int` | Optional | - |
| `ContractId` | `*string` | Optional | - |
| `ContractSku` | `*string` | Optional | - |
| `ContractStartTime` | `*int` | Optional | - |
| `ContractType` | `*string` | Optional | Contract type (Maintenance / Subscription / Premium / Gov AdvCare / Gov TAC / High Sec / AdvCare / Gov Premium) |
| `CustomerPo` | `*string` | Optional | PO number associated with this SKU |
| `Distributor` | `*string` | Optional | Distributor name |
| `EolTime` | `*int` | Optional | End of life time |
| `EosTime` | `*int` | Optional | End of support time |
| `InstalledAddress` | `*string` | Optional | Address where the device is installed. It is a combination of address , region , country , zip |
| `Model` | `*string` | Optional | Model of the install base inventory |
| `OrderId` | `*string` | Optional | Order ID associated with this SKU |
| `Reseller` | `*string` | Optional | Reseller name |
| `Serial` | `*string` | Optional | Serial Number of the inventory |
| `ShippedTime` | `*float64` | Optional | Shipped date |
| `Sku` | `*string` | Optional | Serviceable device stock |
| `Type` | [`*models.JseInventoryItemTypeEnum`](../../doc/models/jse-inventory-item-type-enum.md) | Optional | enum: `ap`, `gateway`, `switch` |
| `WarrantyEndTime` | `*int` | Optional | - |
| `WarrantyStartTime` | `*int` | Optional | - |
| `WarrantyType` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "contract_end_time": 120,
  "contract_id": "contract_id4",
  "contract_sku": "contract_sku4",
  "contract_start_time": 182,
  "contract_type": "contract_type4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

