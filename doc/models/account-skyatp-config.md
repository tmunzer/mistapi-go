
# Account Skyatp Config

*This model accepts additional fields of type interface{}.*

## Structure

`AccountSkyatpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Password` | `string` | Required | - |
| `Realm` | `string` | Required | - |
| `Username` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "password": "foryoureyesonly",
  "realm": "mist-team",
  "username": "john@abc.com",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

