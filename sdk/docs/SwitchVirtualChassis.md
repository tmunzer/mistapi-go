# SwitchVirtualChassis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]SwitchVirtualChassisMember**](SwitchVirtualChassisMember.md) | list of Virtual Chassis members | [optional] 
**Preprovisioned** | Pointer to **bool** | to configure whether the VC is preprovisioned or nonprovisioned | [optional] [default to false]

## Methods

### NewSwitchVirtualChassis

`func NewSwitchVirtualChassis() *SwitchVirtualChassis`

NewSwitchVirtualChassis instantiates a new SwitchVirtualChassis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchVirtualChassisWithDefaults

`func NewSwitchVirtualChassisWithDefaults() *SwitchVirtualChassis`

NewSwitchVirtualChassisWithDefaults instantiates a new SwitchVirtualChassis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *SwitchVirtualChassis) GetMembers() []SwitchVirtualChassisMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *SwitchVirtualChassis) GetMembersOk() (*[]SwitchVirtualChassisMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *SwitchVirtualChassis) SetMembers(v []SwitchVirtualChassisMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *SwitchVirtualChassis) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetPreprovisioned

`func (o *SwitchVirtualChassis) GetPreprovisioned() bool`

GetPreprovisioned returns the Preprovisioned field if non-nil, zero value otherwise.

### GetPreprovisionedOk

`func (o *SwitchVirtualChassis) GetPreprovisionedOk() (*bool, bool)`

GetPreprovisionedOk returns a tuple with the Preprovisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprovisioned

`func (o *SwitchVirtualChassis) SetPreprovisioned(v bool)`

SetPreprovisioned sets Preprovisioned field to given value.

### HasPreprovisioned

`func (o *SwitchVirtualChassis) HasPreprovisioned() bool`

HasPreprovisioned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


