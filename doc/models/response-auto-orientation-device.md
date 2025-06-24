
# Response Auto Orientation Device

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseAutoOrientationDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reason` | `*string` | Optional | Provides the reason for the status if the AP is invalid. |
| `Valid` | `*bool` | Optional | Indicates whether the auto orient request is valid for the device. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "reason": "reason2",
  "valid": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

