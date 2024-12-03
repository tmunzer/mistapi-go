
# Ui Settings

UI Settings

*This model accepts additional fields of type interface{}.*

## Structure

`UiSettings`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | when the object has been created, in epoch |
| `DefaultScopeId` | `*string` | Optional | - |
| `DefaultScopeType` | `*string` | Optional | - |
| `DefaultTimeRange` | [`*models.UiSettingsDefaultTimeRange`](../../doc/models/ui-settings-default-time-range.md) | Optional | - |
| `Description` | `string` | Required | - |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `IsCustomDataboard` | `*bool` | Optional | - |
| `IsScopeLinked` | `*bool` | Optional | - |
| `IsTimeRangeLinked` | `*bool` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | when the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Purpose` | `string` | Required | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Tiles` | [`[]models.UiSettingsTile`](../../doc/models/ui-settings-tile.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "defaultScopeId": "67970e46-4e12-11e6-9188-0242ad112847",
  "defaultScopeType": "site",
  "description": "Description of the databoard",
  "for_site": true,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "New Databoard",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "purpose": "databoard",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 52.82,
  "defaultTimeRange": {
    "end": 224,
    "endDate": "endDate8",
    "interval": "interval4",
    "name": "name6",
    "shortName": "shortName4",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

