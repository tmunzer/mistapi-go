# SyntheticTestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** |  | [optional] [default to false]
**Vlans** | Pointer to [**[]SyntheticTestProperties**](SyntheticTestProperties.md) |  | [optional] 

## Methods

### NewSyntheticTestConfig

`func NewSyntheticTestConfig() *SyntheticTestConfig`

NewSyntheticTestConfig instantiates a new SyntheticTestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticTestConfigWithDefaults

`func NewSyntheticTestConfigWithDefaults() *SyntheticTestConfig`

NewSyntheticTestConfigWithDefaults instantiates a new SyntheticTestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *SyntheticTestConfig) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SyntheticTestConfig) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SyntheticTestConfig) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SyntheticTestConfig) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetVlans

`func (o *SyntheticTestConfig) GetVlans() []SyntheticTestProperties`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *SyntheticTestConfig) GetVlansOk() (*[]SyntheticTestProperties, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *SyntheticTestConfig) SetVlans(v []SyntheticTestProperties)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *SyntheticTestConfig) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


