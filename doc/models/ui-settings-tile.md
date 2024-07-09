
# Ui Settings Tile

## Structure

`UiSettingsTile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChartBand` | `*string` | Optional | - |
| `ChartColor` | `*string` | Optional | - |
| `ChartDirection` | `*string` | Optional | - |
| `ChartRankBy` | `*string` | Optional | - |
| `ChartType` | `*string` | Optional | - |
| `Colspan` | `*int` | Optional | - |
| `Column` | `*int` | Optional | - |
| `HideEmptyRows` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `Metric` | [`*models.UiSettingsTileMetric`](../../doc/models/ui-settings-tile-metric.md) | Optional | - |
| `Name` | `*string` | Optional | - |
| `Row` | `*int` | Optional | - |
| `Rowspan` | `*int` | Optional | - |
| `ScopeId` | `*string` | Optional | - |
| `ScopeType` | `*string` | Optional | - |
| `SortedColumnIds` | `[]string` | Optional | - |
| `TimeRange` | [`*models.UiSettingsTileTimeRange`](../../doc/models/ui-settings-tile-time-range.md) | Optional | - |
| `TrendType` | `*string` | Optional | - |
| `VizType` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "chartBand": "2.4 ghz",
  "chartColor": "#00B4AD",
  "chartDirection": "tx + rx",
  "chartType": "timeSeries",
  "colspan": 5,
  "column": 1,
  "id": "7a9ab38c-cfc3-483d-b51a-0aec571fadc0",
  "name": "New Analysis",
  "row": 1,
  "rowspan": 2,
  "scopeId": "e0c767834b4c",
  "scopeType": "client",
  "trendType": "line",
  "vizType": "averageTimeSeriesChart",
  "chartRankBy": "chartRankBy0"
}
```

