
# Response Claim License License Item

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseClaimLicenseLicenseItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Quantity` | `int` | Required | - |
| `Start` | `int` | Required | - |
| `Type` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 12,
  "quantity": 82,
  "start": 226,
  "type": "type4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

