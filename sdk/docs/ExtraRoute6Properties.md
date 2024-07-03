# ExtraRoute6Properties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discard** | Pointer to **bool** | this takes precedence | [optional] [default to false]
**Metric** | Pointer to **NullableInt32** |  | [optional] 
**NextQualified** | Pointer to [**map[string]ExtraRoute6PropertiesNextQualifiedProperties**](ExtraRoute6PropertiesNextQualifiedProperties.md) |  | [optional] 
**NoResolve** | Pointer to **bool** |  | [optional] [default to false]
**Preference** | Pointer to **NullableInt32** |  | [optional] 
**Via** | Pointer to **string** | next-hop IP Address | [optional] 

## Methods

### NewExtraRoute6Properties

`func NewExtraRoute6Properties() *ExtraRoute6Properties`

NewExtraRoute6Properties instantiates a new ExtraRoute6Properties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtraRoute6PropertiesWithDefaults

`func NewExtraRoute6PropertiesWithDefaults() *ExtraRoute6Properties`

NewExtraRoute6PropertiesWithDefaults instantiates a new ExtraRoute6Properties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscard

`func (o *ExtraRoute6Properties) GetDiscard() bool`

GetDiscard returns the Discard field if non-nil, zero value otherwise.

### GetDiscardOk

`func (o *ExtraRoute6Properties) GetDiscardOk() (*bool, bool)`

GetDiscardOk returns a tuple with the Discard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscard

`func (o *ExtraRoute6Properties) SetDiscard(v bool)`

SetDiscard sets Discard field to given value.

### HasDiscard

`func (o *ExtraRoute6Properties) HasDiscard() bool`

HasDiscard returns a boolean if a field has been set.

### GetMetric

`func (o *ExtraRoute6Properties) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ExtraRoute6Properties) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ExtraRoute6Properties) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *ExtraRoute6Properties) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### SetMetricNil

`func (o *ExtraRoute6Properties) SetMetricNil(b bool)`

 SetMetricNil sets the value for Metric to be an explicit nil

### UnsetMetric
`func (o *ExtraRoute6Properties) UnsetMetric()`

UnsetMetric ensures that no value is present for Metric, not even an explicit nil
### GetNextQualified

`func (o *ExtraRoute6Properties) GetNextQualified() map[string]ExtraRoute6PropertiesNextQualifiedProperties`

GetNextQualified returns the NextQualified field if non-nil, zero value otherwise.

### GetNextQualifiedOk

`func (o *ExtraRoute6Properties) GetNextQualifiedOk() (*map[string]ExtraRoute6PropertiesNextQualifiedProperties, bool)`

GetNextQualifiedOk returns a tuple with the NextQualified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextQualified

`func (o *ExtraRoute6Properties) SetNextQualified(v map[string]ExtraRoute6PropertiesNextQualifiedProperties)`

SetNextQualified sets NextQualified field to given value.

### HasNextQualified

`func (o *ExtraRoute6Properties) HasNextQualified() bool`

HasNextQualified returns a boolean if a field has been set.

### GetNoResolve

`func (o *ExtraRoute6Properties) GetNoResolve() bool`

GetNoResolve returns the NoResolve field if non-nil, zero value otherwise.

### GetNoResolveOk

`func (o *ExtraRoute6Properties) GetNoResolveOk() (*bool, bool)`

GetNoResolveOk returns a tuple with the NoResolve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoResolve

`func (o *ExtraRoute6Properties) SetNoResolve(v bool)`

SetNoResolve sets NoResolve field to given value.

### HasNoResolve

`func (o *ExtraRoute6Properties) HasNoResolve() bool`

HasNoResolve returns a boolean if a field has been set.

### GetPreference

`func (o *ExtraRoute6Properties) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *ExtraRoute6Properties) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *ExtraRoute6Properties) SetPreference(v int32)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *ExtraRoute6Properties) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### SetPreferenceNil

`func (o *ExtraRoute6Properties) SetPreferenceNil(b bool)`

 SetPreferenceNil sets the value for Preference to be an explicit nil

### UnsetPreference
`func (o *ExtraRoute6Properties) UnsetPreference()`

UnsetPreference ensures that no value is present for Preference, not even an explicit nil
### GetVia

`func (o *ExtraRoute6Properties) GetVia() string`

GetVia returns the Via field if non-nil, zero value otherwise.

### GetViaOk

`func (o *ExtraRoute6Properties) GetViaOk() (*string, bool)`

GetViaOk returns a tuple with the Via field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVia

`func (o *ExtraRoute6Properties) SetVia(v string)`

SetVia sets Via field to given value.

### HasVia

`func (o *ExtraRoute6Properties) HasVia() bool`

HasVia returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


