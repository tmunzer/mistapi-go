
# Jse Inventory Item

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
| `CustomerPo` | `*string` | Optional | po number associated with this SKU |
| `Distributor` | `*string` | Optional | distributor name |
| `EolTime` | `*int` | Optional | end of life time |
| `EosTime` | `*int` | Optional | end of support time |
| `InstalledAddress` | `*string` | Optional | address where the device is installed. It is a combination of address , region , country , zip |
| `Model` | `*string` | Optional | model of the install base inventory |
| `OrderId` | `*string` | Optional | order ID associated with this SKU |
| `Reseller` | `*string` | Optional | reseller name |
| `Serial` | `*string` | Optional | serial Number of the inventory |
| `ShippedTime` | `*float64` | Optional | Shipped date |
| `Sku` | `*string` | Optional | serviceable device stock |
| `Type` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Optional | - |
| `WarrantyEndTime` | `*int` | Optional | - |
| `WarrantyStartTime` | `*int` | Optional | - |
| `WarrantyType` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "contract_end_time": 250,
  "contract_id": "contract_id0",
  "contract_sku": "contract_sku0",
  "contract_start_time": 68,
  "contract_type": "contract_type0"
}
```

