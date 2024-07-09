
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
| `Type` | [`*models.LicenseTypeEnum`](../../doc/models/license-type-enum.md) | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "created_time": 64.7,
  "end_time": 48,
  "id": "00000f00-0000-0000-0000-000000000000",
  "modified_time": 14.26,
  "order_id": "order_id4"
}
```

