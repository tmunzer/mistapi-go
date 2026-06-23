
# Org Ui Settings Tile Position

Grid position for a databoard tile

## Structure

`OrgUiSettingsTilePosition`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Col` | `*int` | Optional | Grid column index for the tile |
| `ColSpan` | `*int` | Optional | Grid column span for the tile |
| `Row` | `*int` | Optional | Grid row index for the tile |
| `RowSpan` | `*int` | Optional | Grid row span for the tile |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgUiSettingsTilePosition := models.OrgUiSettingsTilePosition{
        Col:                  models.ToPointer(1),
        ColSpan:              models.ToPointer(5),
        Row:                  models.ToPointer(1),
        RowSpan:              models.ToPointer(2),
    }

}
```

