# VirtualChassisPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**[]ConfigVcPortMember**](ConfigVcPortMember.md) |  | 
**Op** | [**VirtualChassisPortOperation**](VirtualChassisPortOperation.md) |  | 

## Methods

### NewVirtualChassisPort

`func NewVirtualChassisPort(members []ConfigVcPortMember, op VirtualChassisPortOperation, ) *VirtualChassisPort`

NewVirtualChassisPort instantiates a new VirtualChassisPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualChassisPortWithDefaults

`func NewVirtualChassisPortWithDefaults() *VirtualChassisPort`

NewVirtualChassisPortWithDefaults instantiates a new VirtualChassisPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *VirtualChassisPort) GetMembers() []ConfigVcPortMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *VirtualChassisPort) GetMembersOk() (*[]ConfigVcPortMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *VirtualChassisPort) SetMembers(v []ConfigVcPortMember)`

SetMembers sets Members field to given value.


### GetOp

`func (o *VirtualChassisPort) GetOp() VirtualChassisPortOperation`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *VirtualChassisPort) GetOpOk() (*VirtualChassisPortOperation, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *VirtualChassisPort) SetOp(v VirtualChassisPortOperation)`

SetOp sets Op field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


