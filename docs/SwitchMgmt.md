# SwitchMgmt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigRevert** | Pointer to **int32** |  | [optional] [default to 10]
**ProtectRe** | Pointer to [**ProtectRe**](ProtectRe.md) |  | [optional] 
**RootPassword** | Pointer to **string** |  | [optional] 
**Tacacs** | Pointer to [**Tacacs**](Tacacs.md) |  | [optional] 

## Methods

### NewSwitchMgmt

`func NewSwitchMgmt() *SwitchMgmt`

NewSwitchMgmt instantiates a new SwitchMgmt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchMgmtWithDefaults

`func NewSwitchMgmtWithDefaults() *SwitchMgmt`

NewSwitchMgmtWithDefaults instantiates a new SwitchMgmt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigRevert

`func (o *SwitchMgmt) GetConfigRevert() int32`

GetConfigRevert returns the ConfigRevert field if non-nil, zero value otherwise.

### GetConfigRevertOk

`func (o *SwitchMgmt) GetConfigRevertOk() (*int32, bool)`

GetConfigRevertOk returns a tuple with the ConfigRevert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigRevert

`func (o *SwitchMgmt) SetConfigRevert(v int32)`

SetConfigRevert sets ConfigRevert field to given value.

### HasConfigRevert

`func (o *SwitchMgmt) HasConfigRevert() bool`

HasConfigRevert returns a boolean if a field has been set.

### GetProtectRe

`func (o *SwitchMgmt) GetProtectRe() ProtectRe`

GetProtectRe returns the ProtectRe field if non-nil, zero value otherwise.

### GetProtectReOk

`func (o *SwitchMgmt) GetProtectReOk() (*ProtectRe, bool)`

GetProtectReOk returns a tuple with the ProtectRe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectRe

`func (o *SwitchMgmt) SetProtectRe(v ProtectRe)`

SetProtectRe sets ProtectRe field to given value.

### HasProtectRe

`func (o *SwitchMgmt) HasProtectRe() bool`

HasProtectRe returns a boolean if a field has been set.

### GetRootPassword

`func (o *SwitchMgmt) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *SwitchMgmt) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *SwitchMgmt) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *SwitchMgmt) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetTacacs

`func (o *SwitchMgmt) GetTacacs() Tacacs`

GetTacacs returns the Tacacs field if non-nil, zero value otherwise.

### GetTacacsOk

`func (o *SwitchMgmt) GetTacacsOk() (*Tacacs, bool)`

GetTacacsOk returns a tuple with the Tacacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacacs

`func (o *SwitchMgmt) SetTacacs(v Tacacs)`

SetTacacs sets Tacacs field to given value.

### HasTacacs

`func (o *SwitchMgmt) HasTacacs() bool`

HasTacacs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


