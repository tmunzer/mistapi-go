
# Aamw Profile

*This model accepts additional fields of type interface{}.*

## Structure

`AamwProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Categories` | [`[]models.AamwProfileCategory`](../../doc/models/aamw-profile-category.md) | Optional | - |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `FallbackAction` | [`*models.AamwProfileActionEnum`](../../doc/models/aamw-profile-action-enum.md) | Optional | enum: `block`, `permit`<br><br>**Default**: `"block"` |
| `FileAction` | [`*models.AamwProfileActionEnum`](../../doc/models/aamw-profile-action-enum.md) | Optional | enum: `block`, `permit`<br><br>**Default**: `"block"` |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `VerdictThreshold` | `*int` | Optional | **Default**: `8`<br><br>**Constraints**: `>= 1`, `<= 10` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "fallback_action": "block",
  "file_action": "block",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "aamw-custom",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "verdict_threshold": 8,
  "categories": [
    {
      "category": "executable",
      "hash_lookup_only": false,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "category": "executable",
      "hash_lookup_only": false,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "category": "executable",
      "hash_lookup_only": false,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "created_time": 134.84,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

