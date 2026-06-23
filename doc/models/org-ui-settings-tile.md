
# Org Ui Settings Tile

Tile shown on an organization UI databoard

## Structure

`OrgUiSettingsTile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Description` | `*string` | Optional | Text describing the databoard tile |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `IsAutoTitle` | `*bool` | Optional | Whether the tile title is auto generated or not |
| `Name` | `*string` | Optional | Display name of the databoard tile |
| `NlQuery` | `*string` | Optional | Natural Language query for the tile |
| `Position` | [`*models.OrgUiSettingsTilePosition`](../../doc/models/org-ui-settings-tile-position.md) | Optional | Grid position for a databoard tile |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    orgUiSettingsTile := models.OrgUiSettingsTile{
        Description:          models.ToPointer("This tile shows the top 10 APs by bandwidth"),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        IsAutoTitle:          models.ToPointer(false),
        Name:                 models.ToPointer("Top 10 APs by Bandwidth"),
        NlQuery:              models.ToPointer("List top 10 APs by bandwidth"),
    }

}
```

