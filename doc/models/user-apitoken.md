
# User Apitoken

User API Token

*This model accepts additional fields of type interface{}.*

## Structure

`UserApitoken`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Key` | `*string` | Optional | - |
| `LastUsed` | `models.Optional[int]` | Optional | - |
| `Name` | `*string` | Optional | Name of the token |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "key": "1qkb...QQCL",
  "last_used": 1690115110,
  "name": "org_token_xyz",
  "created_time": 116.44,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

