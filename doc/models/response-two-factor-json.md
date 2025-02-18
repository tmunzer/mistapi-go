
# Response Two Factor Json

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseTwoFactorJson`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TwoFactorSecret` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "two_factor_secret": "NRMTSTRWNBVECY3GJVYEY3DDJFRGSNCZGJUDO4RVN5FDM3DUMJSA",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

