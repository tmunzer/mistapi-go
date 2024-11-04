
# User Apitoken

User API Token

## Structure

`UserApitoken`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | when the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `Key` | `*string` | Optional | - |
| `LastUsed` | `models.Optional[int]` | Optional | - |
| `Name` | `*string` | Optional | name of the token |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "key": "1qkb...QQCL",
  "last_used": 1690115110,
  "name": "org_token_xyz",
  "created_time": 116.44
}
```

