
# License Amendment

*This model accepts additional fields of type interface{}.*

## Structure

`LicenseAmendment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `EndTime` | `*int` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Quantity` | `*int` | Optional | - |
| `StartTime` | `*int` | Optional | - |
| `SubscriptionId` | `*string` | Optional | - |
| `Type` | [`*models.LicenseTypeEnum`](../../doc/models/license-type-enum.md) | Optional | enum: `SUB-AST`, `SUB-DATA`, `SUB-ENG`, `SUB-EX12`, `SUB-EX24`, `SUB-EX48`, `SUB-MAN`, `SUB-ME`, `SUB-PMA`, `SUB-SRX1`, `SUB-SRX2`, `SUB-SVNA`, `SUB-VNA`, `SUB-WAN1`, `SUB-WAN2`, `SUB-WVNA1`, `SUB-WVNA2` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "created_time": 222.36,
  "end_time": 198,
  "modified_time": 112.6,
  "quantity": 170,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

