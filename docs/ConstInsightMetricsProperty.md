# ConstInsightMetricsProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ctype** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Example** | Pointer to [**[]ConstInsightMetricsPropertyExample**](ConstInsightMetricsPropertyExample.md) |  | [optional] 
**Intervals** | Pointer to [**map[string]ConstInsightMetricsPropertyInterval**](ConstInsightMetricsPropertyInterval.md) | Property key is the interval (e.g. 10m, 1h, ...) | [optional] 
**Keys** | Pointer to **map[string]interface{}** |  | [optional] 
**Params** | Pointer to [**map[string]ConstInsightMetricsPropertyParam**](ConstInsightMetricsPropertyParam.md) | Property key is the parameter name | [optional] 
**ReportDuration** | Pointer to [**map[string]ConstInsightMetricsPropertyReportDuration**](ConstInsightMetricsPropertyReportDuration.md) | Property key is the duration (e.g. 1d, 1w, ...) | [optional] 
**ReportScopes** | Pointer to **[]string** |  | [optional] 
**Scopes** | Pointer to [**[]ConstInsightMetricsPropertyScope**](ConstInsightMetricsPropertyScope.md) |  | [optional] 
**SleBaselined** | Pointer to **bool** |  | [optional] 
**SleClassifiers** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Values** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConstInsightMetricsProperty

`func NewConstInsightMetricsProperty() *ConstInsightMetricsProperty`

NewConstInsightMetricsProperty instantiates a new ConstInsightMetricsProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstInsightMetricsPropertyWithDefaults

`func NewConstInsightMetricsPropertyWithDefaults() *ConstInsightMetricsProperty`

NewConstInsightMetricsPropertyWithDefaults instantiates a new ConstInsightMetricsProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCtype

`func (o *ConstInsightMetricsProperty) GetCtype() []string`

GetCtype returns the Ctype field if non-nil, zero value otherwise.

### GetCtypeOk

`func (o *ConstInsightMetricsProperty) GetCtypeOk() (*[]string, bool)`

GetCtypeOk returns a tuple with the Ctype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtype

`func (o *ConstInsightMetricsProperty) SetCtype(v []string)`

SetCtype sets Ctype field to given value.

### HasCtype

`func (o *ConstInsightMetricsProperty) HasCtype() bool`

HasCtype returns a boolean if a field has been set.

### GetDescription

`func (o *ConstInsightMetricsProperty) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConstInsightMetricsProperty) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConstInsightMetricsProperty) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConstInsightMetricsProperty) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExample

`func (o *ConstInsightMetricsProperty) GetExample() []ConstInsightMetricsPropertyExample`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *ConstInsightMetricsProperty) GetExampleOk() (*[]ConstInsightMetricsPropertyExample, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *ConstInsightMetricsProperty) SetExample(v []ConstInsightMetricsPropertyExample)`

SetExample sets Example field to given value.

### HasExample

`func (o *ConstInsightMetricsProperty) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetIntervals

`func (o *ConstInsightMetricsProperty) GetIntervals() map[string]ConstInsightMetricsPropertyInterval`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *ConstInsightMetricsProperty) GetIntervalsOk() (*map[string]ConstInsightMetricsPropertyInterval, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *ConstInsightMetricsProperty) SetIntervals(v map[string]ConstInsightMetricsPropertyInterval)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *ConstInsightMetricsProperty) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### GetKeys

`func (o *ConstInsightMetricsProperty) GetKeys() map[string]interface{}`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ConstInsightMetricsProperty) GetKeysOk() (*map[string]interface{}, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ConstInsightMetricsProperty) SetKeys(v map[string]interface{})`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ConstInsightMetricsProperty) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetParams

`func (o *ConstInsightMetricsProperty) GetParams() map[string]ConstInsightMetricsPropertyParam`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ConstInsightMetricsProperty) GetParamsOk() (*map[string]ConstInsightMetricsPropertyParam, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ConstInsightMetricsProperty) SetParams(v map[string]ConstInsightMetricsPropertyParam)`

SetParams sets Params field to given value.

### HasParams

`func (o *ConstInsightMetricsProperty) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetReportDuration

`func (o *ConstInsightMetricsProperty) GetReportDuration() map[string]ConstInsightMetricsPropertyReportDuration`

GetReportDuration returns the ReportDuration field if non-nil, zero value otherwise.

### GetReportDurationOk

`func (o *ConstInsightMetricsProperty) GetReportDurationOk() (*map[string]ConstInsightMetricsPropertyReportDuration, bool)`

GetReportDurationOk returns a tuple with the ReportDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDuration

`func (o *ConstInsightMetricsProperty) SetReportDuration(v map[string]ConstInsightMetricsPropertyReportDuration)`

SetReportDuration sets ReportDuration field to given value.

### HasReportDuration

`func (o *ConstInsightMetricsProperty) HasReportDuration() bool`

HasReportDuration returns a boolean if a field has been set.

### GetReportScopes

`func (o *ConstInsightMetricsProperty) GetReportScopes() []string`

GetReportScopes returns the ReportScopes field if non-nil, zero value otherwise.

### GetReportScopesOk

`func (o *ConstInsightMetricsProperty) GetReportScopesOk() (*[]string, bool)`

GetReportScopesOk returns a tuple with the ReportScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportScopes

`func (o *ConstInsightMetricsProperty) SetReportScopes(v []string)`

SetReportScopes sets ReportScopes field to given value.

### HasReportScopes

`func (o *ConstInsightMetricsProperty) HasReportScopes() bool`

HasReportScopes returns a boolean if a field has been set.

### GetScopes

`func (o *ConstInsightMetricsProperty) GetScopes() []ConstInsightMetricsPropertyScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ConstInsightMetricsProperty) GetScopesOk() (*[]ConstInsightMetricsPropertyScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ConstInsightMetricsProperty) SetScopes(v []ConstInsightMetricsPropertyScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ConstInsightMetricsProperty) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSleBaselined

`func (o *ConstInsightMetricsProperty) GetSleBaselined() bool`

GetSleBaselined returns the SleBaselined field if non-nil, zero value otherwise.

### GetSleBaselinedOk

`func (o *ConstInsightMetricsProperty) GetSleBaselinedOk() (*bool, bool)`

GetSleBaselinedOk returns a tuple with the SleBaselined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleBaselined

`func (o *ConstInsightMetricsProperty) SetSleBaselined(v bool)`

SetSleBaselined sets SleBaselined field to given value.

### HasSleBaselined

`func (o *ConstInsightMetricsProperty) HasSleBaselined() bool`

HasSleBaselined returns a boolean if a field has been set.

### GetSleClassifiers

`func (o *ConstInsightMetricsProperty) GetSleClassifiers() []string`

GetSleClassifiers returns the SleClassifiers field if non-nil, zero value otherwise.

### GetSleClassifiersOk

`func (o *ConstInsightMetricsProperty) GetSleClassifiersOk() (*[]string, bool)`

GetSleClassifiersOk returns a tuple with the SleClassifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleClassifiers

`func (o *ConstInsightMetricsProperty) SetSleClassifiers(v []string)`

SetSleClassifiers sets SleClassifiers field to given value.

### HasSleClassifiers

`func (o *ConstInsightMetricsProperty) HasSleClassifiers() bool`

HasSleClassifiers returns a boolean if a field has been set.

### GetType

`func (o *ConstInsightMetricsProperty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConstInsightMetricsProperty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConstInsightMetricsProperty) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConstInsightMetricsProperty) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *ConstInsightMetricsProperty) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ConstInsightMetricsProperty) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ConstInsightMetricsProperty) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ConstInsightMetricsProperty) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValues

`func (o *ConstInsightMetricsProperty) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ConstInsightMetricsProperty) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ConstInsightMetricsProperty) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *ConstInsightMetricsProperty) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


