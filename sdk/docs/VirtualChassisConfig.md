# VirtualChassisConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locating** | Pointer to **bool** |  | [optional] [readonly] 
**Members** | Pointer to [**[]VirtualChassisConfigMember**](VirtualChassisConfigMember.md) |  | [optional] 
**Preprovisioned** | Pointer to **bool** | To create the Virtual Chassis in Pre-Provisioned mode | [optional] [default to false]

## Methods

### NewVirtualChassisConfig

`func NewVirtualChassisConfig() *VirtualChassisConfig`

NewVirtualChassisConfig instantiates a new VirtualChassisConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualChassisConfigWithDefaults

`func NewVirtualChassisConfigWithDefaults() *VirtualChassisConfig`

NewVirtualChassisConfigWithDefaults instantiates a new VirtualChassisConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocating

`func (o *VirtualChassisConfig) GetLocating() bool`

GetLocating returns the Locating field if non-nil, zero value otherwise.

### GetLocatingOk

`func (o *VirtualChassisConfig) GetLocatingOk() (*bool, bool)`

GetLocatingOk returns a tuple with the Locating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocating

`func (o *VirtualChassisConfig) SetLocating(v bool)`

SetLocating sets Locating field to given value.

### HasLocating

`func (o *VirtualChassisConfig) HasLocating() bool`

HasLocating returns a boolean if a field has been set.

### GetMembers

`func (o *VirtualChassisConfig) GetMembers() []VirtualChassisConfigMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *VirtualChassisConfig) GetMembersOk() (*[]VirtualChassisConfigMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *VirtualChassisConfig) SetMembers(v []VirtualChassisConfigMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *VirtualChassisConfig) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetPreprovisioned

`func (o *VirtualChassisConfig) GetPreprovisioned() bool`

GetPreprovisioned returns the Preprovisioned field if non-nil, zero value otherwise.

### GetPreprovisionedOk

`func (o *VirtualChassisConfig) GetPreprovisionedOk() (*bool, bool)`

GetPreprovisionedOk returns a tuple with the Preprovisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprovisioned

`func (o *VirtualChassisConfig) SetPreprovisioned(v bool)`

SetPreprovisioned sets Preprovisioned field to given value.

### HasPreprovisioned

`func (o *VirtualChassisConfig) HasPreprovisioned() bool`

HasPreprovisioned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


