# UiSettingsTile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartBand** | Pointer to **string** |  | [optional] 
**ChartColor** | Pointer to **string** |  | [optional] 
**ChartDirection** | Pointer to **string** |  | [optional] 
**ChartRankBy** | Pointer to **string** |  | [optional] 
**ChartType** | Pointer to **string** |  | [optional] 
**Colspan** | Pointer to **int32** |  | [optional] 
**Column** | Pointer to **int32** |  | [optional] 
**HideEmptyRows** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Metric** | Pointer to [**UiSettingsTileMetric**](UiSettingsTileMetric.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Row** | Pointer to **int32** |  | [optional] 
**Rowspan** | Pointer to **int32** |  | [optional] 
**ScopeId** | Pointer to **string** |  | [optional] 
**ScopeType** | Pointer to **string** |  | [optional] 
**SortedColumnIds** | Pointer to **[]string** |  | [optional] 
**TimeRange** | Pointer to [**UiSettingsTileTimeRange**](UiSettingsTileTimeRange.md) |  | [optional] 
**TrendType** | Pointer to **string** |  | [optional] 
**VizType** | Pointer to **string** |  | [optional] 

## Methods

### NewUiSettingsTile

`func NewUiSettingsTile() *UiSettingsTile`

NewUiSettingsTile instantiates a new UiSettingsTile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiSettingsTileWithDefaults

`func NewUiSettingsTileWithDefaults() *UiSettingsTile`

NewUiSettingsTileWithDefaults instantiates a new UiSettingsTile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartBand

`func (o *UiSettingsTile) GetChartBand() string`

GetChartBand returns the ChartBand field if non-nil, zero value otherwise.

### GetChartBandOk

`func (o *UiSettingsTile) GetChartBandOk() (*string, bool)`

GetChartBandOk returns a tuple with the ChartBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartBand

`func (o *UiSettingsTile) SetChartBand(v string)`

SetChartBand sets ChartBand field to given value.

### HasChartBand

`func (o *UiSettingsTile) HasChartBand() bool`

HasChartBand returns a boolean if a field has been set.

### GetChartColor

`func (o *UiSettingsTile) GetChartColor() string`

GetChartColor returns the ChartColor field if non-nil, zero value otherwise.

### GetChartColorOk

`func (o *UiSettingsTile) GetChartColorOk() (*string, bool)`

GetChartColorOk returns a tuple with the ChartColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartColor

`func (o *UiSettingsTile) SetChartColor(v string)`

SetChartColor sets ChartColor field to given value.

### HasChartColor

`func (o *UiSettingsTile) HasChartColor() bool`

HasChartColor returns a boolean if a field has been set.

### GetChartDirection

`func (o *UiSettingsTile) GetChartDirection() string`

GetChartDirection returns the ChartDirection field if non-nil, zero value otherwise.

### GetChartDirectionOk

`func (o *UiSettingsTile) GetChartDirectionOk() (*string, bool)`

GetChartDirectionOk returns a tuple with the ChartDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartDirection

`func (o *UiSettingsTile) SetChartDirection(v string)`

SetChartDirection sets ChartDirection field to given value.

### HasChartDirection

`func (o *UiSettingsTile) HasChartDirection() bool`

HasChartDirection returns a boolean if a field has been set.

### GetChartRankBy

`func (o *UiSettingsTile) GetChartRankBy() string`

GetChartRankBy returns the ChartRankBy field if non-nil, zero value otherwise.

### GetChartRankByOk

`func (o *UiSettingsTile) GetChartRankByOk() (*string, bool)`

GetChartRankByOk returns a tuple with the ChartRankBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartRankBy

`func (o *UiSettingsTile) SetChartRankBy(v string)`

SetChartRankBy sets ChartRankBy field to given value.

### HasChartRankBy

`func (o *UiSettingsTile) HasChartRankBy() bool`

HasChartRankBy returns a boolean if a field has been set.

### GetChartType

`func (o *UiSettingsTile) GetChartType() string`

GetChartType returns the ChartType field if non-nil, zero value otherwise.

### GetChartTypeOk

`func (o *UiSettingsTile) GetChartTypeOk() (*string, bool)`

GetChartTypeOk returns a tuple with the ChartType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartType

`func (o *UiSettingsTile) SetChartType(v string)`

SetChartType sets ChartType field to given value.

### HasChartType

`func (o *UiSettingsTile) HasChartType() bool`

HasChartType returns a boolean if a field has been set.

### GetColspan

`func (o *UiSettingsTile) GetColspan() int32`

GetColspan returns the Colspan field if non-nil, zero value otherwise.

### GetColspanOk

`func (o *UiSettingsTile) GetColspanOk() (*int32, bool)`

GetColspanOk returns a tuple with the Colspan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColspan

`func (o *UiSettingsTile) SetColspan(v int32)`

SetColspan sets Colspan field to given value.

### HasColspan

`func (o *UiSettingsTile) HasColspan() bool`

HasColspan returns a boolean if a field has been set.

### GetColumn

`func (o *UiSettingsTile) GetColumn() int32`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *UiSettingsTile) GetColumnOk() (*int32, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *UiSettingsTile) SetColumn(v int32)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *UiSettingsTile) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### GetHideEmptyRows

`func (o *UiSettingsTile) GetHideEmptyRows() bool`

GetHideEmptyRows returns the HideEmptyRows field if non-nil, zero value otherwise.

### GetHideEmptyRowsOk

`func (o *UiSettingsTile) GetHideEmptyRowsOk() (*bool, bool)`

GetHideEmptyRowsOk returns a tuple with the HideEmptyRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideEmptyRows

`func (o *UiSettingsTile) SetHideEmptyRows(v bool)`

SetHideEmptyRows sets HideEmptyRows field to given value.

### HasHideEmptyRows

`func (o *UiSettingsTile) HasHideEmptyRows() bool`

HasHideEmptyRows returns a boolean if a field has been set.

### GetId

`func (o *UiSettingsTile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UiSettingsTile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UiSettingsTile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UiSettingsTile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetric

`func (o *UiSettingsTile) GetMetric() UiSettingsTileMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *UiSettingsTile) GetMetricOk() (*UiSettingsTileMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *UiSettingsTile) SetMetric(v UiSettingsTileMetric)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *UiSettingsTile) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetName

`func (o *UiSettingsTile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UiSettingsTile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UiSettingsTile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UiSettingsTile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRow

`func (o *UiSettingsTile) GetRow() int32`

GetRow returns the Row field if non-nil, zero value otherwise.

### GetRowOk

`func (o *UiSettingsTile) GetRowOk() (*int32, bool)`

GetRowOk returns a tuple with the Row field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRow

`func (o *UiSettingsTile) SetRow(v int32)`

SetRow sets Row field to given value.

### HasRow

`func (o *UiSettingsTile) HasRow() bool`

HasRow returns a boolean if a field has been set.

### GetRowspan

`func (o *UiSettingsTile) GetRowspan() int32`

GetRowspan returns the Rowspan field if non-nil, zero value otherwise.

### GetRowspanOk

`func (o *UiSettingsTile) GetRowspanOk() (*int32, bool)`

GetRowspanOk returns a tuple with the Rowspan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowspan

`func (o *UiSettingsTile) SetRowspan(v int32)`

SetRowspan sets Rowspan field to given value.

### HasRowspan

`func (o *UiSettingsTile) HasRowspan() bool`

HasRowspan returns a boolean if a field has been set.

### GetScopeId

`func (o *UiSettingsTile) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *UiSettingsTile) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *UiSettingsTile) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *UiSettingsTile) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetScopeType

`func (o *UiSettingsTile) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *UiSettingsTile) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *UiSettingsTile) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *UiSettingsTile) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### GetSortedColumnIds

`func (o *UiSettingsTile) GetSortedColumnIds() []string`

GetSortedColumnIds returns the SortedColumnIds field if non-nil, zero value otherwise.

### GetSortedColumnIdsOk

`func (o *UiSettingsTile) GetSortedColumnIdsOk() (*[]string, bool)`

GetSortedColumnIdsOk returns a tuple with the SortedColumnIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortedColumnIds

`func (o *UiSettingsTile) SetSortedColumnIds(v []string)`

SetSortedColumnIds sets SortedColumnIds field to given value.

### HasSortedColumnIds

`func (o *UiSettingsTile) HasSortedColumnIds() bool`

HasSortedColumnIds returns a boolean if a field has been set.

### GetTimeRange

`func (o *UiSettingsTile) GetTimeRange() UiSettingsTileTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *UiSettingsTile) GetTimeRangeOk() (*UiSettingsTileTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *UiSettingsTile) SetTimeRange(v UiSettingsTileTimeRange)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *UiSettingsTile) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### GetTrendType

`func (o *UiSettingsTile) GetTrendType() string`

GetTrendType returns the TrendType field if non-nil, zero value otherwise.

### GetTrendTypeOk

`func (o *UiSettingsTile) GetTrendTypeOk() (*string, bool)`

GetTrendTypeOk returns a tuple with the TrendType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendType

`func (o *UiSettingsTile) SetTrendType(v string)`

SetTrendType sets TrendType field to given value.

### HasTrendType

`func (o *UiSettingsTile) HasTrendType() bool`

HasTrendType returns a boolean if a field has been set.

### GetVizType

`func (o *UiSettingsTile) GetVizType() string`

GetVizType returns the VizType field if non-nil, zero value otherwise.

### GetVizTypeOk

`func (o *UiSettingsTile) GetVizTypeOk() (*string, bool)`

GetVizTypeOk returns a tuple with the VizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizType

`func (o *UiSettingsTile) SetVizType(v string)`

SetVizType sets VizType field to given value.

### HasVizType

`func (o *UiSettingsTile) HasVizType() bool`

HasVizType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


