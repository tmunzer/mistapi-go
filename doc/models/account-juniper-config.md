
# Account Juniper Config

*This model accepts additional fields of type interface{}.*

## Structure

`AccountJuniperConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Password` | `string` | Required | customer account password |
| `Username` | `string` | Required | customer account user name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "password": "password",
  "username": "john@nmo.com",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

