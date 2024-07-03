# WlanCiscoCwa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedHostnames** | Pointer to **[]string** | list of hostnames without http(s):// (matched by substring) | [optional] 
**AllowedSubnets** | Pointer to **[]string** | list of CIDRs | [optional] 
**BlockedSubnets** | Pointer to **[]string** | list of blocked CIDRs | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewWlanCiscoCwa

`func NewWlanCiscoCwa() *WlanCiscoCwa`

NewWlanCiscoCwa instantiates a new WlanCiscoCwa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanCiscoCwaWithDefaults

`func NewWlanCiscoCwaWithDefaults() *WlanCiscoCwa`

NewWlanCiscoCwaWithDefaults instantiates a new WlanCiscoCwa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedHostnames

`func (o *WlanCiscoCwa) GetAllowedHostnames() []string`

GetAllowedHostnames returns the AllowedHostnames field if non-nil, zero value otherwise.

### GetAllowedHostnamesOk

`func (o *WlanCiscoCwa) GetAllowedHostnamesOk() (*[]string, bool)`

GetAllowedHostnamesOk returns a tuple with the AllowedHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHostnames

`func (o *WlanCiscoCwa) SetAllowedHostnames(v []string)`

SetAllowedHostnames sets AllowedHostnames field to given value.

### HasAllowedHostnames

`func (o *WlanCiscoCwa) HasAllowedHostnames() bool`

HasAllowedHostnames returns a boolean if a field has been set.

### GetAllowedSubnets

`func (o *WlanCiscoCwa) GetAllowedSubnets() []string`

GetAllowedSubnets returns the AllowedSubnets field if non-nil, zero value otherwise.

### GetAllowedSubnetsOk

`func (o *WlanCiscoCwa) GetAllowedSubnetsOk() (*[]string, bool)`

GetAllowedSubnetsOk returns a tuple with the AllowedSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSubnets

`func (o *WlanCiscoCwa) SetAllowedSubnets(v []string)`

SetAllowedSubnets sets AllowedSubnets field to given value.

### HasAllowedSubnets

`func (o *WlanCiscoCwa) HasAllowedSubnets() bool`

HasAllowedSubnets returns a boolean if a field has been set.

### GetBlockedSubnets

`func (o *WlanCiscoCwa) GetBlockedSubnets() []string`

GetBlockedSubnets returns the BlockedSubnets field if non-nil, zero value otherwise.

### GetBlockedSubnetsOk

`func (o *WlanCiscoCwa) GetBlockedSubnetsOk() (*[]string, bool)`

GetBlockedSubnetsOk returns a tuple with the BlockedSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedSubnets

`func (o *WlanCiscoCwa) SetBlockedSubnets(v []string)`

SetBlockedSubnets sets BlockedSubnets field to given value.

### HasBlockedSubnets

`func (o *WlanCiscoCwa) HasBlockedSubnets() bool`

HasBlockedSubnets returns a boolean if a field has been set.

### GetEnabled

`func (o *WlanCiscoCwa) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WlanCiscoCwa) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WlanCiscoCwa) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WlanCiscoCwa) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


