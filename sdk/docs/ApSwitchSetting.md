# ApSwitchSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalVlanIds** | Pointer to **[]int32** | additional VLAN IDs, only valid in mesh base mode | [optional] 
**EnableVlan** | Pointer to **bool** |  | [optional] 
**PortVlanId** | Pointer to [**ApSwitchSettingPortVlanId**](ApSwitchSettingPortVlanId.md) |  | [optional] 
**VlanIds** | Pointer to **[]int32** | list of VLAN ids this | [optional] 

## Methods

### NewApSwitchSetting

`func NewApSwitchSetting() *ApSwitchSetting`

NewApSwitchSetting instantiates a new ApSwitchSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApSwitchSettingWithDefaults

`func NewApSwitchSettingWithDefaults() *ApSwitchSetting`

NewApSwitchSettingWithDefaults instantiates a new ApSwitchSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalVlanIds

`func (o *ApSwitchSetting) GetAdditionalVlanIds() []int32`

GetAdditionalVlanIds returns the AdditionalVlanIds field if non-nil, zero value otherwise.

### GetAdditionalVlanIdsOk

`func (o *ApSwitchSetting) GetAdditionalVlanIdsOk() (*[]int32, bool)`

GetAdditionalVlanIdsOk returns a tuple with the AdditionalVlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalVlanIds

`func (o *ApSwitchSetting) SetAdditionalVlanIds(v []int32)`

SetAdditionalVlanIds sets AdditionalVlanIds field to given value.

### HasAdditionalVlanIds

`func (o *ApSwitchSetting) HasAdditionalVlanIds() bool`

HasAdditionalVlanIds returns a boolean if a field has been set.

### GetEnableVlan

`func (o *ApSwitchSetting) GetEnableVlan() bool`

GetEnableVlan returns the EnableVlan field if non-nil, zero value otherwise.

### GetEnableVlanOk

`func (o *ApSwitchSetting) GetEnableVlanOk() (*bool, bool)`

GetEnableVlanOk returns a tuple with the EnableVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVlan

`func (o *ApSwitchSetting) SetEnableVlan(v bool)`

SetEnableVlan sets EnableVlan field to given value.

### HasEnableVlan

`func (o *ApSwitchSetting) HasEnableVlan() bool`

HasEnableVlan returns a boolean if a field has been set.

### GetPortVlanId

`func (o *ApSwitchSetting) GetPortVlanId() ApSwitchSettingPortVlanId`

GetPortVlanId returns the PortVlanId field if non-nil, zero value otherwise.

### GetPortVlanIdOk

`func (o *ApSwitchSetting) GetPortVlanIdOk() (*ApSwitchSettingPortVlanId, bool)`

GetPortVlanIdOk returns a tuple with the PortVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortVlanId

`func (o *ApSwitchSetting) SetPortVlanId(v ApSwitchSettingPortVlanId)`

SetPortVlanId sets PortVlanId field to given value.

### HasPortVlanId

`func (o *ApSwitchSetting) HasPortVlanId() bool`

HasPortVlanId returns a boolean if a field has been set.

### GetVlanIds

`func (o *ApSwitchSetting) GetVlanIds() []int32`

GetVlanIds returns the VlanIds field if non-nil, zero value otherwise.

### GetVlanIdsOk

`func (o *ApSwitchSetting) GetVlanIdsOk() (*[]int32, bool)`

GetVlanIdsOk returns a tuple with the VlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIds

`func (o *ApSwitchSetting) SetVlanIds(v []int32)`

SetVlanIds sets VlanIds field to given value.

### HasVlanIds

`func (o *ApSwitchSetting) HasVlanIds() bool`

HasVlanIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


