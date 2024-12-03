
# Account Jse Config

*This model accepts additional fields of type interface{}.*

## Structure

`AccountJseConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CloudName` | `*string` | Optional | - |
| `Password` | `string` | Required | - |
| `Username` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "cloud_name": "devcentral.juniperclouds.net",
  "password": "foryoureyesonly",
  "username": "john@abc.com",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

