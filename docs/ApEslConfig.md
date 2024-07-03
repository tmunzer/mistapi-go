# ApEslConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cacert** | Pointer to **string** | Only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;imagotag&#x60; or &#x60;type&#x60;&#x3D;&#x3D;&#x60;native&#x60; | [optional] 
**Channel** | Pointer to **int32** | Only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;imagotag&#x60; or &#x60;type&#x60;&#x3D;&#x3D;&#x60;native&#x60; | [optional] 
**Enabled** | Pointer to **bool** | usb_config is ignored if esl_config enabled | [optional] [default to false]
**Host** | Pointer to **string** | Only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;imagotag&#x60; or &#x60;type&#x60;&#x3D;&#x3D;&#x60;native&#x60; | [optional] 
**Port** | Pointer to **int32** | Only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;imagotag&#x60; or &#x60;type&#x60;&#x3D;&#x3D;&#x60;native&#x60; | [optional] 
**Type** | Pointer to [**ApEslType**](ApEslType.md) |  | [optional] 
**VerifyCert** | Pointer to **bool** | Only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;imagotag&#x60; or &#x60;type&#x60;&#x3D;&#x3D;&#x60;native&#x60; | [optional] 
**VlanId** | Pointer to **int32** | Only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;solum&#x60; or &#x60;type&#x60;&#x3D;&#x3D;&#x60;hansho&#x60; | [optional] [default to 1]

## Methods

### NewApEslConfig

`func NewApEslConfig() *ApEslConfig`

NewApEslConfig instantiates a new ApEslConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApEslConfigWithDefaults

`func NewApEslConfigWithDefaults() *ApEslConfig`

NewApEslConfigWithDefaults instantiates a new ApEslConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacert

`func (o *ApEslConfig) GetCacert() string`

GetCacert returns the Cacert field if non-nil, zero value otherwise.

### GetCacertOk

`func (o *ApEslConfig) GetCacertOk() (*string, bool)`

GetCacertOk returns a tuple with the Cacert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacert

`func (o *ApEslConfig) SetCacert(v string)`

SetCacert sets Cacert field to given value.

### HasCacert

`func (o *ApEslConfig) HasCacert() bool`

HasCacert returns a boolean if a field has been set.

### GetChannel

`func (o *ApEslConfig) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ApEslConfig) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ApEslConfig) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ApEslConfig) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetEnabled

`func (o *ApEslConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApEslConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApEslConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApEslConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHost

`func (o *ApEslConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ApEslConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ApEslConfig) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ApEslConfig) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *ApEslConfig) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApEslConfig) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApEslConfig) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApEslConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetType

`func (o *ApEslConfig) GetType() ApEslType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApEslConfig) GetTypeOk() (*ApEslType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApEslConfig) SetType(v ApEslType)`

SetType sets Type field to given value.

### HasType

`func (o *ApEslConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVerifyCert

`func (o *ApEslConfig) GetVerifyCert() bool`

GetVerifyCert returns the VerifyCert field if non-nil, zero value otherwise.

### GetVerifyCertOk

`func (o *ApEslConfig) GetVerifyCertOk() (*bool, bool)`

GetVerifyCertOk returns a tuple with the VerifyCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCert

`func (o *ApEslConfig) SetVerifyCert(v bool)`

SetVerifyCert sets VerifyCert field to given value.

### HasVerifyCert

`func (o *ApEslConfig) HasVerifyCert() bool`

HasVerifyCert returns a boolean if a field has been set.

### GetVlanId

`func (o *ApEslConfig) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *ApEslConfig) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *ApEslConfig) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *ApEslConfig) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


