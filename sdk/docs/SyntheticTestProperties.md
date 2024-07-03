# SyntheticTestProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomTestUrls** | Pointer to **[]string** |  | [optional] 
**Disabled** | Pointer to **bool** | for some vlans where we don&#39;t want this to run | [optional] [default to false]
**VlanIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewSyntheticTestProperties

`func NewSyntheticTestProperties() *SyntheticTestProperties`

NewSyntheticTestProperties instantiates a new SyntheticTestProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticTestPropertiesWithDefaults

`func NewSyntheticTestPropertiesWithDefaults() *SyntheticTestProperties`

NewSyntheticTestPropertiesWithDefaults instantiates a new SyntheticTestProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomTestUrls

`func (o *SyntheticTestProperties) GetCustomTestUrls() []string`

GetCustomTestUrls returns the CustomTestUrls field if non-nil, zero value otherwise.

### GetCustomTestUrlsOk

`func (o *SyntheticTestProperties) GetCustomTestUrlsOk() (*[]string, bool)`

GetCustomTestUrlsOk returns a tuple with the CustomTestUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTestUrls

`func (o *SyntheticTestProperties) SetCustomTestUrls(v []string)`

SetCustomTestUrls sets CustomTestUrls field to given value.

### HasCustomTestUrls

`func (o *SyntheticTestProperties) HasCustomTestUrls() bool`

HasCustomTestUrls returns a boolean if a field has been set.

### GetDisabled

`func (o *SyntheticTestProperties) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SyntheticTestProperties) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SyntheticTestProperties) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SyntheticTestProperties) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetVlanIds

`func (o *SyntheticTestProperties) GetVlanIds() []int32`

GetVlanIds returns the VlanIds field if non-nil, zero value otherwise.

### GetVlanIdsOk

`func (o *SyntheticTestProperties) GetVlanIdsOk() (*[]int32, bool)`

GetVlanIdsOk returns a tuple with the VlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIds

`func (o *SyntheticTestProperties) SetVlanIds(v []int32)`

SetVlanIds sets VlanIds field to given value.

### HasVlanIds

`func (o *SyntheticTestProperties) HasVlanIds() bool`

HasVlanIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


