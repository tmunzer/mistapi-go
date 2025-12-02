
# Org Ui Settings Tile

## Structure

`OrgUiSettingsTile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Description` | `*string` | Optional | Description of the tile |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `IsAutoTitle` | `*bool` | Optional | Whether the tile title is auto generated or not |
| `Name` | `*string` | Optional | Name of the tile |
| `NlQuery` | `*string` | Optional | Natural Language query for the tile |
| `Position` | [`*models.OrgUiSettingsTilePosition`](../../doc/models/org-ui-settings-tile-position.md) | Optional | - |

## Example (as JSON)

```json
{
  "description": "This tile shows the top 10 APs by bandwidth",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "Top 10 APs by Bandwidth",
  "nl_query": "List top 10 APs by bandwidth",
  "isAutoTitle": false
}
```

