
# Login

*This model accepts additional fields of type interface{}.*

## Structure

`Login`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Email` | `string` | Required | - |
| `Password` | `string` | Required | - |
| `TwoFactor` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "email": "test@mistsys.com",
  "password": "foryoureyesonly",
  "two_factor": "123456",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

