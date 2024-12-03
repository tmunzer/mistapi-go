
# Response Login Success

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseLoginSuccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Email` | `*string` | Optional | - |
| `TwoFactorPassed` | `*bool` | Optional | - |
| `TwoFactorRequired` | `*bool` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "email": "email6",
  "two_factor_passed": false,
  "two_factor_required": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

