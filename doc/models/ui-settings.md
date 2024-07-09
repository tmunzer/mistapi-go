
# Ui Settings

UI Settings

## Structure

`UiSettings`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `DefaultScopeId` | `*string` | Optional | - |
| `DefaultScopeType` | `*string` | Optional | - |
| `DefaultTimeRange` | [`*models.UiSettingsDefaultTimeRange`](../../doc/models/ui-settings-default-time-range.md) | Optional | - |
| `Description` | `string` | Required | - |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `IsCustomDataboard` | `*bool` | Optional | - |
| `IsScopeLinked` | `*bool` | Optional | - |
| `IsTimeRangeLinked` | `*bool` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Purpose` | `string` | Required | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Tiles` | [`[]models.UiSettingsTile`](../../doc/models/ui-settings-tile.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

## Example (as JSON)

```json
{
  "created_time": 1508823803.0,
  "defaultScopeId": "67970e46-4e12-11e6-9188-0242ad112847",
  "defaultScopeType": "site",
  "description": "Description of the databoard",
  "for_site": true,
  "id": "3bdcc7e8-c04d-4512-b4fc-093da9057eb0",
  "modified_time": 1508823803,
  "name": "New Databoard",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "purpose": "databoard",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "defaultTimeRange": {
    "end": 224,
    "endDate": "endDate8",
    "interval": "interval4",
    "name": "name6",
    "shortName": "shortName4"
  }
}
```

