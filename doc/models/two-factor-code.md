
# Two Factor Code

*This model accepts additional fields of type interface{}.*

## Structure

`TwoFactorCode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TwoFactor` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "two_factor": "123456",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

