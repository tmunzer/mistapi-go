
# Response Self Oauth Link Success

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSelfOauthLinkSuccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | `string` | Required | - |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organnization |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "action": "action8",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

