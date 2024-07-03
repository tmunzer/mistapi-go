# VirtualChassisMemberUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Required if &#x60;op&#x60;&#x3D;&#x3D;&#x60;add&#x60; or &#x60;op&#x60;&#x3D;&#x3D;&#x60;preprovision&#x60;. | [optional] 
**Member** | Pointer to **int32** | Required if &#x60;op&#x60;&#x3D;&#x3D;&#x60;remove&#x60; or &#x60;op&#x60;&#x3D;&#x3D;&#x60;preprovision&#x60;. Optional if &#x60;op&#x60;&#x3D;&#x3D;&#x60;add&#x60; | [optional] 
**VcPorts** | Pointer to **[]string** | Required if &#x60;op&#x60;&#x3D;&#x3D;&#x60;add&#x60; or &#x60;op&#x60;&#x3D;&#x3D;&#x60;preprovision&#x60; | [optional] 
**VcRole** | Pointer to [**VirtualChassisMemberUpdateVcRole**](VirtualChassisMemberUpdateVcRole.md) |  | [optional] 

## Methods

### NewVirtualChassisMemberUpdate

`func NewVirtualChassisMemberUpdate() *VirtualChassisMemberUpdate`

NewVirtualChassisMemberUpdate instantiates a new VirtualChassisMemberUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualChassisMemberUpdateWithDefaults

`func NewVirtualChassisMemberUpdateWithDefaults() *VirtualChassisMemberUpdate`

NewVirtualChassisMemberUpdateWithDefaults instantiates a new VirtualChassisMemberUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *VirtualChassisMemberUpdate) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *VirtualChassisMemberUpdate) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *VirtualChassisMemberUpdate) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *VirtualChassisMemberUpdate) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMember

`func (o *VirtualChassisMemberUpdate) GetMember() int32`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *VirtualChassisMemberUpdate) GetMemberOk() (*int32, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *VirtualChassisMemberUpdate) SetMember(v int32)`

SetMember sets Member field to given value.

### HasMember

`func (o *VirtualChassisMemberUpdate) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetVcPorts

`func (o *VirtualChassisMemberUpdate) GetVcPorts() []string`

GetVcPorts returns the VcPorts field if non-nil, zero value otherwise.

### GetVcPortsOk

`func (o *VirtualChassisMemberUpdate) GetVcPortsOk() (*[]string, bool)`

GetVcPortsOk returns a tuple with the VcPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPorts

`func (o *VirtualChassisMemberUpdate) SetVcPorts(v []string)`

SetVcPorts sets VcPorts field to given value.

### HasVcPorts

`func (o *VirtualChassisMemberUpdate) HasVcPorts() bool`

HasVcPorts returns a boolean if a field has been set.

### GetVcRole

`func (o *VirtualChassisMemberUpdate) GetVcRole() VirtualChassisMemberUpdateVcRole`

GetVcRole returns the VcRole field if non-nil, zero value otherwise.

### GetVcRoleOk

`func (o *VirtualChassisMemberUpdate) GetVcRoleOk() (*VirtualChassisMemberUpdateVcRole, bool)`

GetVcRoleOk returns a tuple with the VcRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcRole

`func (o *VirtualChassisMemberUpdate) SetVcRole(v VirtualChassisMemberUpdateVcRole)`

SetVcRole sets VcRole field to given value.

### HasVcRole

`func (o *VirtualChassisMemberUpdate) HasVcRole() bool`

HasVcRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


