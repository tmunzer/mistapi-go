
# Idp Profile

*This model accepts additional fields of type interface{}.*

## Structure

`IdpProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BaseProfile` | [`*models.IdpProfileBaseProfileEnum`](../../doc/models/idp-profile-base-profile-enum.md) | Optional | enum: `critical`, `standard`, `strict` |
| `CreatedTime` | `*float64` | Optional | when the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `ModifiedTime` | `*float64` | Optional | when the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Overwrites` | [`[]models.IdpProfileOverwrite`](../../doc/models/idp-profile-overwrite.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "base_profile": "strict",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "relaxed",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "created_time": 104.4,
  "modified_time": 230.56,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

