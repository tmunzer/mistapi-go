
# License Sub

*This model accepts additional fields of type interface{}.*

## Structure

`LicenseSub`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `EndTime` | `*int` | Optional | End date of the license term |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `OrderId` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Quantity` | `*int` | Optional | Number of devices entitled for this license |
| `RemainingQuantity` | `*int` | Optional | Number of licenses left in this subscription |
| `StartTime` | `*int` | Optional | Start date of the license term |
| `SubscriptionId` | `*string` | Optional | - |
| `Type` | `*string` | Optional | Type of license. The list of supported license type can be retrieve with the [List License Type](/#operations/listLicenseTypes) API request. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "created_time": 46.48,
  "end_time": 18,
  "modified_time": 32.48,
  "order_id": "order_id8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

