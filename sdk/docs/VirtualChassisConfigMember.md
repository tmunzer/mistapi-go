# VirtualChassisConfigMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locating** | Pointer to **bool** |  | [optional] [readonly] 
**Mac** | Pointer to **string** | fpc0, same as the mac of device_id | [optional] 
**Member** | Pointer to **int32** | to create a preprovisionned virtual chassis | [optional] 
**VcPorts** | Pointer to **[]string** |  | [optional] 
**VcRole** | Pointer to [**VirtualChassisConfigMemberVcRole**](VirtualChassisConfigMemberVcRole.md) |  | [optional] 

## Methods

### NewVirtualChassisConfigMember

`func NewVirtualChassisConfigMember() *VirtualChassisConfigMember`

NewVirtualChassisConfigMember instantiates a new VirtualChassisConfigMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualChassisConfigMemberWithDefaults

`func NewVirtualChassisConfigMemberWithDefaults() *VirtualChassisConfigMember`

NewVirtualChassisConfigMemberWithDefaults instantiates a new VirtualChassisConfigMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocating

`func (o *VirtualChassisConfigMember) GetLocating() bool`

GetLocating returns the Locating field if non-nil, zero value otherwise.

### GetLocatingOk

`func (o *VirtualChassisConfigMember) GetLocatingOk() (*bool, bool)`

GetLocatingOk returns a tuple with the Locating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocating

`func (o *VirtualChassisConfigMember) SetLocating(v bool)`

SetLocating sets Locating field to given value.

### HasLocating

`func (o *VirtualChassisConfigMember) HasLocating() bool`

HasLocating returns a boolean if a field has been set.

### GetMac

`func (o *VirtualChassisConfigMember) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *VirtualChassisConfigMember) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *VirtualChassisConfigMember) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *VirtualChassisConfigMember) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMember

`func (o *VirtualChassisConfigMember) GetMember() int32`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *VirtualChassisConfigMember) GetMemberOk() (*int32, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *VirtualChassisConfigMember) SetMember(v int32)`

SetMember sets Member field to given value.

### HasMember

`func (o *VirtualChassisConfigMember) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetVcPorts

`func (o *VirtualChassisConfigMember) GetVcPorts() []string`

GetVcPorts returns the VcPorts field if non-nil, zero value otherwise.

### GetVcPortsOk

`func (o *VirtualChassisConfigMember) GetVcPortsOk() (*[]string, bool)`

GetVcPortsOk returns a tuple with the VcPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPorts

`func (o *VirtualChassisConfigMember) SetVcPorts(v []string)`

SetVcPorts sets VcPorts field to given value.

### HasVcPorts

`func (o *VirtualChassisConfigMember) HasVcPorts() bool`

HasVcPorts returns a boolean if a field has been set.

### GetVcRole

`func (o *VirtualChassisConfigMember) GetVcRole() VirtualChassisConfigMemberVcRole`

GetVcRole returns the VcRole field if non-nil, zero value otherwise.

### GetVcRoleOk

`func (o *VirtualChassisConfigMember) GetVcRoleOk() (*VirtualChassisConfigMemberVcRole, bool)`

GetVcRoleOk returns a tuple with the VcRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcRole

`func (o *VirtualChassisConfigMember) SetVcRole(v VirtualChassisConfigMemberVcRole)`

SetVcRole sets VcRole field to given value.

### HasVcRole

`func (o *VirtualChassisConfigMember) HasVcRole() bool`

HasVcRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


