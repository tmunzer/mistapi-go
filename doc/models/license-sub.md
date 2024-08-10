
# License Sub

## Structure

`LicenseSub`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `EndTime` | `*int` | Optional | end date of the license term |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `OrderId` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Quantity` | `*int` | Optional | number of devices entitled for this license |
| `RemainingQuantity` | `*int` | Optional | Number of licenses left in this subscription |
| `StartTime` | `*int` | Optional | start date of the license term |
| `SubscriptionId` | `*string` | Optional | - |
| `Type` | [`*models.LicenseTypeEnum`](../../doc/models/license-type-enum.md) | Optional | enum: `SUB-AST`, `SUB-DATA`, `SUB-ENG`, `SUB-EX12`, `SUB-EX24`, `SUB-EX48`, `SUB-MAN`, `SUB-ME`, `SUB-PMA`, `SUB-SRX1`, `SUB-SRX2`, `SUB-SVNA`, `SUB-VNA`, `SUB-WAN1`, `SUB-WAN2`, `SUB-WVNA1`, `SUB-WVNA2` |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "created_time": 46.48,
  "end_time": 18,
  "id": "00001912-0000-0000-0000-000000000000",
  "modified_time": 32.48,
  "order_id": "order_id8"
}
```

