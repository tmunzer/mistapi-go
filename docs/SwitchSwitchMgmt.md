# SwitchSwitchMgmt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigRevertTimer** | Pointer to **int32** | rollback timer for commit confirmed | [optional] [default to 10]
**ProtectRe** | Pointer to [**ProtectRe**](ProtectRe.md) |  | [optional] 

## Methods

### NewSwitchSwitchMgmt

`func NewSwitchSwitchMgmt() *SwitchSwitchMgmt`

NewSwitchSwitchMgmt instantiates a new SwitchSwitchMgmt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchSwitchMgmtWithDefaults

`func NewSwitchSwitchMgmtWithDefaults() *SwitchSwitchMgmt`

NewSwitchSwitchMgmtWithDefaults instantiates a new SwitchSwitchMgmt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigRevertTimer

`func (o *SwitchSwitchMgmt) GetConfigRevertTimer() int32`

GetConfigRevertTimer returns the ConfigRevertTimer field if non-nil, zero value otherwise.

### GetConfigRevertTimerOk

`func (o *SwitchSwitchMgmt) GetConfigRevertTimerOk() (*int32, bool)`

GetConfigRevertTimerOk returns a tuple with the ConfigRevertTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigRevertTimer

`func (o *SwitchSwitchMgmt) SetConfigRevertTimer(v int32)`

SetConfigRevertTimer sets ConfigRevertTimer field to given value.

### HasConfigRevertTimer

`func (o *SwitchSwitchMgmt) HasConfigRevertTimer() bool`

HasConfigRevertTimer returns a boolean if a field has been set.

### GetProtectRe

`func (o *SwitchSwitchMgmt) GetProtectRe() ProtectRe`

GetProtectRe returns the ProtectRe field if non-nil, zero value otherwise.

### GetProtectReOk

`func (o *SwitchSwitchMgmt) GetProtectReOk() (*ProtectRe, bool)`

GetProtectReOk returns a tuple with the ProtectRe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectRe

`func (o *SwitchSwitchMgmt) SetProtectRe(v ProtectRe)`

SetProtectRe sets ProtectRe field to given value.

### HasProtectRe

`func (o *SwitchSwitchMgmt) HasProtectRe() bool`

HasProtectRe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


