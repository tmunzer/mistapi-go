# MxedgeDasCoaServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableEventTimestampCheck** | Pointer to **bool** | whether to disable Event-Timestamp Check | [optional] [default to false]
**Enabled** | Pointer to **bool** |  | [optional] 
**Host** | Pointer to **string** | this server configured to send CoA|DM to mist edges | [optional] 
**Port** | Pointer to **int32** | mist edges will allow this host on this port | [optional] [default to 3799]
**Secret** | Pointer to **string** |  | [optional] 

## Methods

### NewMxedgeDasCoaServer

`func NewMxedgeDasCoaServer() *MxedgeDasCoaServer`

NewMxedgeDasCoaServer instantiates a new MxedgeDasCoaServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeDasCoaServerWithDefaults

`func NewMxedgeDasCoaServerWithDefaults() *MxedgeDasCoaServer`

NewMxedgeDasCoaServerWithDefaults instantiates a new MxedgeDasCoaServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableEventTimestampCheck

`func (o *MxedgeDasCoaServer) GetDisableEventTimestampCheck() bool`

GetDisableEventTimestampCheck returns the DisableEventTimestampCheck field if non-nil, zero value otherwise.

### GetDisableEventTimestampCheckOk

`func (o *MxedgeDasCoaServer) GetDisableEventTimestampCheckOk() (*bool, bool)`

GetDisableEventTimestampCheckOk returns a tuple with the DisableEventTimestampCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEventTimestampCheck

`func (o *MxedgeDasCoaServer) SetDisableEventTimestampCheck(v bool)`

SetDisableEventTimestampCheck sets DisableEventTimestampCheck field to given value.

### HasDisableEventTimestampCheck

`func (o *MxedgeDasCoaServer) HasDisableEventTimestampCheck() bool`

HasDisableEventTimestampCheck returns a boolean if a field has been set.

### GetEnabled

`func (o *MxedgeDasCoaServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MxedgeDasCoaServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MxedgeDasCoaServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MxedgeDasCoaServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHost

`func (o *MxedgeDasCoaServer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MxedgeDasCoaServer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MxedgeDasCoaServer) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *MxedgeDasCoaServer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *MxedgeDasCoaServer) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MxedgeDasCoaServer) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MxedgeDasCoaServer) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *MxedgeDasCoaServer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecret

`func (o *MxedgeDasCoaServer) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *MxedgeDasCoaServer) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *MxedgeDasCoaServer) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *MxedgeDasCoaServer) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


