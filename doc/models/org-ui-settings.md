
# Org Ui Settings

*This model accepts additional fields of type interface{}.*

## Structure

`OrgUiSettings`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Description` | `*string` | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `IsCustomDataboard` | `*bool` | Optional | Whether this is a custom databoard or not |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Name of the databoard |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Purpose` | [`*models.OrgUiSettingsPurposeEnum`](../../doc/models/org-ui-settings-purpose-enum.md) | Optional | enum: `marvisdashboard` |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Tiles` | [`[]models.OrgUiSettingsTile`](../../doc/models/org-ui-settings-tile.md) | Optional | List of tiles in the databoard |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "description": "This databoard shows AP stats",
  "for_site": false,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "AP Stats",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 193.48,
  "isCustomDataboard": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

