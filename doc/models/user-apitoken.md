
# User Apitoken

User API Token

## Structure

`UserApitoken`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `Key` | `*string` | Optional | - |
| `LastUsed` | `models.Optional[int]` | Optional | - |
| `Name` | `*string` | Optional | name of the token |

## Example (as JSON)

```json
{
  "created_time": 1626875902.0,
  "id": "864f351a-1377-4ad9-83f8-72f3fe6199ba",
  "key": "1qkb...QQCL",
  "last_used": 1690115110,
  "name": "org_token_xyz"
}
```

