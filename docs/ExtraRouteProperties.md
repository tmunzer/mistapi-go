# ExtraRouteProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discard** | Pointer to **bool** | this takes precedence | [optional] [default to false]
**Metric** | Pointer to **NullableInt32** |  | [optional] 
**NextQualified** | Pointer to [**map[string]ExtraRoutePropertiesNextQualifiedProperties**](ExtraRoutePropertiesNextQualifiedProperties.md) |  | [optional] 
**NoResolve** | Pointer to **bool** |  | [optional] [default to false]
**Preference** | Pointer to **NullableInt32** |  | [optional] 
**Via** | Pointer to **string** | next-hop IP Address | [optional] 

## Methods

### NewExtraRouteProperties

`func NewExtraRouteProperties() *ExtraRouteProperties`

NewExtraRouteProperties instantiates a new ExtraRouteProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtraRoutePropertiesWithDefaults

`func NewExtraRoutePropertiesWithDefaults() *ExtraRouteProperties`

NewExtraRoutePropertiesWithDefaults instantiates a new ExtraRouteProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscard

`func (o *ExtraRouteProperties) GetDiscard() bool`

GetDiscard returns the Discard field if non-nil, zero value otherwise.

### GetDiscardOk

`func (o *ExtraRouteProperties) GetDiscardOk() (*bool, bool)`

GetDiscardOk returns a tuple with the Discard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscard

`func (o *ExtraRouteProperties) SetDiscard(v bool)`

SetDiscard sets Discard field to given value.

### HasDiscard

`func (o *ExtraRouteProperties) HasDiscard() bool`

HasDiscard returns a boolean if a field has been set.

### GetMetric

`func (o *ExtraRouteProperties) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ExtraRouteProperties) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ExtraRouteProperties) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *ExtraRouteProperties) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### SetMetricNil

`func (o *ExtraRouteProperties) SetMetricNil(b bool)`

 SetMetricNil sets the value for Metric to be an explicit nil

### UnsetMetric
`func (o *ExtraRouteProperties) UnsetMetric()`

UnsetMetric ensures that no value is present for Metric, not even an explicit nil
### GetNextQualified

`func (o *ExtraRouteProperties) GetNextQualified() map[string]ExtraRoutePropertiesNextQualifiedProperties`

GetNextQualified returns the NextQualified field if non-nil, zero value otherwise.

### GetNextQualifiedOk

`func (o *ExtraRouteProperties) GetNextQualifiedOk() (*map[string]ExtraRoutePropertiesNextQualifiedProperties, bool)`

GetNextQualifiedOk returns a tuple with the NextQualified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextQualified

`func (o *ExtraRouteProperties) SetNextQualified(v map[string]ExtraRoutePropertiesNextQualifiedProperties)`

SetNextQualified sets NextQualified field to given value.

### HasNextQualified

`func (o *ExtraRouteProperties) HasNextQualified() bool`

HasNextQualified returns a boolean if a field has been set.

### GetNoResolve

`func (o *ExtraRouteProperties) GetNoResolve() bool`

GetNoResolve returns the NoResolve field if non-nil, zero value otherwise.

### GetNoResolveOk

`func (o *ExtraRouteProperties) GetNoResolveOk() (*bool, bool)`

GetNoResolveOk returns a tuple with the NoResolve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoResolve

`func (o *ExtraRouteProperties) SetNoResolve(v bool)`

SetNoResolve sets NoResolve field to given value.

### HasNoResolve

`func (o *ExtraRouteProperties) HasNoResolve() bool`

HasNoResolve returns a boolean if a field has been set.

### GetPreference

`func (o *ExtraRouteProperties) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *ExtraRouteProperties) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *ExtraRouteProperties) SetPreference(v int32)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *ExtraRouteProperties) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### SetPreferenceNil

`func (o *ExtraRouteProperties) SetPreferenceNil(b bool)`

 SetPreferenceNil sets the value for Preference to be an explicit nil

### UnsetPreference
`func (o *ExtraRouteProperties) UnsetPreference()`

UnsetPreference ensures that no value is present for Preference, not even an explicit nil
### GetVia

`func (o *ExtraRouteProperties) GetVia() string`

GetVia returns the Via field if non-nil, zero value otherwise.

### GetViaOk

`func (o *ExtraRouteProperties) GetViaOk() (*string, bool)`

GetViaOk returns a tuple with the Via field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVia

`func (o *ExtraRouteProperties) SetVia(v string)`

SetVia sets Via field to given value.

### HasVia

`func (o *ExtraRouteProperties) HasVia() bool`

HasVia returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


