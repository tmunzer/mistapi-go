package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// UiSettingsTile represents a UiSettingsTile struct.
type UiSettingsTile struct {
    ChartBand            *string                  `json:"chartBand,omitempty"`
    ChartColor           *string                  `json:"chartColor,omitempty"`
    ChartDirection       *string                  `json:"chartDirection,omitempty"`
    ChartRankBy          *string                  `json:"chartRankBy,omitempty"`
    ChartType            *string                  `json:"chartType,omitempty"`
    Colspan              *int                     `json:"colspan,omitempty"`
    Column               *int                     `json:"column,omitempty"`
    HideEmptyRows        *bool                    `json:"hideEmptyRows,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID               `json:"id,omitempty"`
    Metric               *UiSettingsTileMetric    `json:"metric,omitempty"`
    Name                 *string                  `json:"name,omitempty"`
    Row                  *int                     `json:"row,omitempty"`
    Rowspan              *int                     `json:"rowspan,omitempty"`
    ScopeId              *string                  `json:"scopeId,omitempty"`
    ScopeType            *string                  `json:"scopeType,omitempty"`
    SortedColumnIds      []string                 `json:"sortedColumnIds,omitempty"`
    TimeRange            *UiSettingsTileTimeRange `json:"timeRange,omitempty"`
    TrendType            *string                  `json:"trendType,omitempty"`
    VizType              *string                  `json:"vizType,omitempty"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for UiSettingsTile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UiSettingsTile) String() string {
    return fmt.Sprintf(
    	"UiSettingsTile[ChartBand=%v, ChartColor=%v, ChartDirection=%v, ChartRankBy=%v, ChartType=%v, Colspan=%v, Column=%v, HideEmptyRows=%v, Id=%v, Metric=%v, Name=%v, Row=%v, Rowspan=%v, ScopeId=%v, ScopeType=%v, SortedColumnIds=%v, TimeRange=%v, TrendType=%v, VizType=%v, AdditionalProperties=%v]",
    	u.ChartBand, u.ChartColor, u.ChartDirection, u.ChartRankBy, u.ChartType, u.Colspan, u.Column, u.HideEmptyRows, u.Id, u.Metric, u.Name, u.Row, u.Rowspan, u.ScopeId, u.ScopeType, u.SortedColumnIds, u.TimeRange, u.TrendType, u.VizType, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UiSettingsTile.
// It customizes the JSON marshaling process for UiSettingsTile objects.
func (u UiSettingsTile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "chartBand", "chartColor", "chartDirection", "chartRankBy", "chartType", "colspan", "column", "hideEmptyRows", "id", "metric", "name", "row", "rowspan", "scopeId", "scopeType", "sortedColumnIds", "timeRange", "trendType", "vizType"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UiSettingsTile object to a map representation for JSON marshaling.
func (u UiSettingsTile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.ChartBand != nil {
        structMap["chartBand"] = u.ChartBand
    }
    if u.ChartColor != nil {
        structMap["chartColor"] = u.ChartColor
    }
    if u.ChartDirection != nil {
        structMap["chartDirection"] = u.ChartDirection
    }
    if u.ChartRankBy != nil {
        structMap["chartRankBy"] = u.ChartRankBy
    }
    if u.ChartType != nil {
        structMap["chartType"] = u.ChartType
    }
    if u.Colspan != nil {
        structMap["colspan"] = u.Colspan
    }
    if u.Column != nil {
        structMap["column"] = u.Column
    }
    if u.HideEmptyRows != nil {
        structMap["hideEmptyRows"] = u.HideEmptyRows
    }
    if u.Id != nil {
        structMap["id"] = u.Id
    }
    if u.Metric != nil {
        structMap["metric"] = u.Metric.toMap()
    }
    if u.Name != nil {
        structMap["name"] = u.Name
    }
    if u.Row != nil {
        structMap["row"] = u.Row
    }
    if u.Rowspan != nil {
        structMap["rowspan"] = u.Rowspan
    }
    if u.ScopeId != nil {
        structMap["scopeId"] = u.ScopeId
    }
    if u.ScopeType != nil {
        structMap["scopeType"] = u.ScopeType
    }
    if u.SortedColumnIds != nil {
        structMap["sortedColumnIds"] = u.SortedColumnIds
    }
    if u.TimeRange != nil {
        structMap["timeRange"] = u.TimeRange.toMap()
    }
    if u.TrendType != nil {
        structMap["trendType"] = u.TrendType
    }
    if u.VizType != nil {
        structMap["vizType"] = u.VizType
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UiSettingsTile.
// It customizes the JSON unmarshaling process for UiSettingsTile objects.
func (u *UiSettingsTile) UnmarshalJSON(input []byte) error {
    var temp tempUiSettingsTile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "chartBand", "chartColor", "chartDirection", "chartRankBy", "chartType", "colspan", "column", "hideEmptyRows", "id", "metric", "name", "row", "rowspan", "scopeId", "scopeType", "sortedColumnIds", "timeRange", "trendType", "vizType")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.ChartBand = temp.ChartBand
    u.ChartColor = temp.ChartColor
    u.ChartDirection = temp.ChartDirection
    u.ChartRankBy = temp.ChartRankBy
    u.ChartType = temp.ChartType
    u.Colspan = temp.Colspan
    u.Column = temp.Column
    u.HideEmptyRows = temp.HideEmptyRows
    u.Id = temp.Id
    u.Metric = temp.Metric
    u.Name = temp.Name
    u.Row = temp.Row
    u.Rowspan = temp.Rowspan
    u.ScopeId = temp.ScopeId
    u.ScopeType = temp.ScopeType
    u.SortedColumnIds = temp.SortedColumnIds
    u.TimeRange = temp.TimeRange
    u.TrendType = temp.TrendType
    u.VizType = temp.VizType
    return nil
}

// tempUiSettingsTile is a temporary struct used for validating the fields of UiSettingsTile.
type tempUiSettingsTile  struct {
    ChartBand       *string                  `json:"chartBand,omitempty"`
    ChartColor      *string                  `json:"chartColor,omitempty"`
    ChartDirection  *string                  `json:"chartDirection,omitempty"`
    ChartRankBy     *string                  `json:"chartRankBy,omitempty"`
    ChartType       *string                  `json:"chartType,omitempty"`
    Colspan         *int                     `json:"colspan,omitempty"`
    Column          *int                     `json:"column,omitempty"`
    HideEmptyRows   *bool                    `json:"hideEmptyRows,omitempty"`
    Id              *uuid.UUID               `json:"id,omitempty"`
    Metric          *UiSettingsTileMetric    `json:"metric,omitempty"`
    Name            *string                  `json:"name,omitempty"`
    Row             *int                     `json:"row,omitempty"`
    Rowspan         *int                     `json:"rowspan,omitempty"`
    ScopeId         *string                  `json:"scopeId,omitempty"`
    ScopeType       *string                  `json:"scopeType,omitempty"`
    SortedColumnIds []string                 `json:"sortedColumnIds,omitempty"`
    TimeRange       *UiSettingsTileTimeRange `json:"timeRange,omitempty"`
    TrendType       *string                  `json:"trendType,omitempty"`
    VizType         *string                  `json:"vizType,omitempty"`
}
