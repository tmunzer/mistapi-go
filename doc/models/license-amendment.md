
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
| `Type` | `*string` | Optional | Type of license. The list of supported license type can be retrieve with the [List License Type](../../doc/controllers/constants-definitions.md#list-license-types) API request. |
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

