# SwitchVirtualChassisMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | fpc0, same as the mac of device_id | [optional] 
**VcRole** | Pointer to [**SwitchVirtualChassisMemberVcRole**](SwitchVirtualChassisMemberVcRole.md) |  | [optional] 

## Methods

### NewSwitchVirtualChassisMember

`func NewSwitchVirtualChassisMember() *SwitchVirtualChassisMember`

NewSwitchVirtualChassisMember instantiates a new SwitchVirtualChassisMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchVirtualChassisMemberWithDefaults

`func NewSwitchVirtualChassisMemberWithDefaults() *SwitchVirtualChassisMember`

NewSwitchVirtualChassisMemberWithDefaults instantiates a new SwitchVirtualChassisMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *SwitchVirtualChassisMember) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SwitchVirtualChassisMember) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SwitchVirtualChassisMember) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SwitchVirtualChassisMember) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetVcRole

`func (o *SwitchVirtualChassisMember) GetVcRole() SwitchVirtualChassisMemberVcRole`

GetVcRole returns the VcRole field if non-nil, zero value otherwise.

### GetVcRoleOk

`func (o *SwitchVirtualChassisMember) GetVcRoleOk() (*SwitchVirtualChassisMemberVcRole, bool)`

GetVcRoleOk returns a tuple with the VcRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcRole

`func (o *SwitchVirtualChassisMember) SetVcRole(v SwitchVirtualChassisMemberVcRole)`

SetVcRole sets VcRole field to given value.

### HasVcRole

`func (o *SwitchVirtualChassisMember) HasVcRole() bool`

HasVcRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


