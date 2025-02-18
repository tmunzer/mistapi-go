
# Synthetictest Radius Server

*This model accepts additional fields of type interface{}.*

## Structure

`SynthetictestRadiusServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Password` | `string` | Required | Specify the password associated with the username |
| `Profile` | `*string` | Optional | Specify the access profile associated with the subscriber<br>**Default**: `"dot1x"` |
| `User` | `string` | Required | Specify the subscriber username to test |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "password": "password6",
  "profile": "dot1x",
  "user": "user2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

