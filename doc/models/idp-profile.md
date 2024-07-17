
# Idp Profile

## Structure

`IdpProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BaseProfile` | [`*models.IdpProfileBaseProfileEnum`](../../doc/models/idp-profile-base-profile-enum.md) | Optional | - |
| `CreatedTime` | `*float64` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Overwrites` | [`[]models.IdpProfileOverwrite`](../../doc/models/idp-profile-overwrite.md) | Optional | - |

## Example (as JSON)

```json
{
  "base_profile": "strict",
  "created_time": 1508823803.0,
  "id": "874ca978-d736-4d4b-bc90-a49a29eec133",
  "name": "relaxed",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "modified_time": 155.08
}
```

