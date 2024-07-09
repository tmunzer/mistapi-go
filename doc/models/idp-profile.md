
# Idp Profile

## Structure

`IdpProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BaseProfile` | [`*models.IdpProfileBaseProfileEnum`](../../doc/models/idp-profile-base-profile-enum.md) | Optional | - |
| `CreatedTime` | `*int` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Overwrites` | [`[]models.IdpProfileOverwrite`](../../doc/models/idp-profile-overwrite.md) | Optional | - |

## Example (as JSON)

```json
{
  "base_profile": "strict",
  "id": "874ca978-d736-4d4b-bc90-a49a29eec133",
  "name": "relaxed",
  "created_time": 68,
  "modified_time": 148
}
```

