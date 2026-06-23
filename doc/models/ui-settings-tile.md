
# Ui Settings Tile

Tile shown on a site UI databoard

## Structure

`UiSettingsTile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChartBand` | `*string` | Optional | Band or category filter applied to the tile chart |
| `ChartColor` | `*string` | Optional | Hex color used to render the tile chart |
| `ChartDirection` | `*string` | Optional | Traffic direction displayed by the tile chart |
| `ChartRankBy` | `*string` | Optional | Ranking field used by the tile chart |
| `ChartType` | `*string` | Optional | Visualization chart type for this tile |
| `Colspan` | `*int` | Optional | Number of grid columns spanned by this tile |
| `Column` | `*int` | Optional | Grid column where this tile starts |
| `HideEmptyRows` | `*bool` | Optional | Whether empty rows are hidden in this tile |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Metric` | [`*models.UiSettingsTileMetric`](../../doc/models/ui-settings-tile-metric.md) | Optional | Metric selected for a site UI databoard tile |
| `Name` | `*string` | Optional | Display name of this databoard tile |
| `Row` | `*int` | Optional | Grid row where this tile starts |
| `Rowspan` | `*int` | Optional | Number of grid rows spanned by this tile |
| `ScopeId` | `*string` | Optional | Identifier of the scope used by this tile |
| `ScopeType` | `*string` | Optional | Type of scope used by this tile |
| `SortedColumnIds` | `[]string` | Optional | Ordered tile table column identifiers used for sorting |
| `TimeRange` | [`*models.UiSettingsTileTimeRange`](../../doc/models/ui-settings-tile-time-range.md) | Optional | Time range override for a site UI databoard tile |
| `TrendType` | `*string` | Optional | Rendering style for trend data in this tile |
| `VizType` | `*string` | Optional | Display visualization type used by this tile |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    uiSettingsTile := models.UiSettingsTile{
        ChartBand:            models.ToPointer("2.4 ghz"),
        ChartColor:           models.ToPointer("#00B4AD"),
        ChartDirection:       models.ToPointer("tx + rx"),
        ChartRankBy:          models.ToPointer("chartRankBy2"),
        ChartType:            models.ToPointer("timeSeries"),
        Colspan:              models.ToPointer(5),
        Column:               models.ToPointer(1),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Name:                 models.ToPointer("New Analysis"),
        Row:                  models.ToPointer(1),
        Rowspan:              models.ToPointer(2),
        ScopeId:              models.ToPointer("e0c767834b4c"),
        ScopeType:            models.ToPointer("client"),
        TrendType:            models.ToPointer("line"),
        VizType:              models.ToPointer("averageTimeSeriesChart"),
    }

}
```

