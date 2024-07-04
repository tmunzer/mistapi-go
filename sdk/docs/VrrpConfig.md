# VrrpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Groups** | Pointer to [**map[string]VrrpConfigGroup**](VrrpConfigGroup.md) | Property key is the VRRP name | [optional] 

## Methods

### NewVrrpConfig

`func NewVrrpConfig() *VrrpConfig`

NewVrrpConfig instantiates a new VrrpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrrpConfigWithDefaults

`func NewVrrpConfigWithDefaults() *VrrpConfig`

NewVrrpConfigWithDefaults instantiates a new VrrpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *VrrpConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VrrpConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VrrpConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VrrpConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGroups

`func (o *VrrpConfig) GetGroups() map[string]VrrpConfigGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *VrrpConfig) GetGroupsOk() (*map[string]VrrpConfigGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *VrrpConfig) SetGroups(v map[string]VrrpConfigGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *VrrpConfig) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


