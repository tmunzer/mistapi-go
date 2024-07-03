# ApUsb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cacert** | Pointer to **NullableString** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;imagotag&#x60; | [optional] 
**Channel** | Pointer to **int32** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;imagotag&#x60; channel selection, not needed by default, required for manual channel override only | [optional] 
**Enabled** | Pointer to **bool** | whether to enable any usb config | [optional] 
**Host** | Pointer to **string** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;imagotag&#x60; | [optional] 
**Port** | Pointer to **int32** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;imagotag&#x60; | [optional] [default to 0]
**Type** | Pointer to [**ApUsbType**](ApUsbType.md) |  | [optional] 
**VerifyCert** | Pointer to **bool** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;imagotag&#x60;, whether to turn on SSL verification | [optional] 
**VlanId** | Pointer to **int32** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;solum&#x60; or &#x60;type&#x60;&#x3D;&#x3D;&#x60;hanshow&#x60; | [optional] [default to 1]

## Methods

### NewApUsb

`func NewApUsb() *ApUsb`

NewApUsb instantiates a new ApUsb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApUsbWithDefaults

`func NewApUsbWithDefaults() *ApUsb`

NewApUsbWithDefaults instantiates a new ApUsb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacert

`func (o *ApUsb) GetCacert() string`

GetCacert returns the Cacert field if non-nil, zero value otherwise.

### GetCacertOk

`func (o *ApUsb) GetCacertOk() (*string, bool)`

GetCacertOk returns a tuple with the Cacert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacert

`func (o *ApUsb) SetCacert(v string)`

SetCacert sets Cacert field to given value.

### HasCacert

`func (o *ApUsb) HasCacert() bool`

HasCacert returns a boolean if a field has been set.

### SetCacertNil

`func (o *ApUsb) SetCacertNil(b bool)`

 SetCacertNil sets the value for Cacert to be an explicit nil

### UnsetCacert
`func (o *ApUsb) UnsetCacert()`

UnsetCacert ensures that no value is present for Cacert, not even an explicit nil
### GetChannel

`func (o *ApUsb) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ApUsb) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ApUsb) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ApUsb) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetEnabled

`func (o *ApUsb) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApUsb) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApUsb) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApUsb) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHost

`func (o *ApUsb) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ApUsb) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ApUsb) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ApUsb) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *ApUsb) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApUsb) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApUsb) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApUsb) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetType

`func (o *ApUsb) GetType() ApUsbType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApUsb) GetTypeOk() (*ApUsbType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApUsb) SetType(v ApUsbType)`

SetType sets Type field to given value.

### HasType

`func (o *ApUsb) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVerifyCert

`func (o *ApUsb) GetVerifyCert() bool`

GetVerifyCert returns the VerifyCert field if non-nil, zero value otherwise.

### GetVerifyCertOk

`func (o *ApUsb) GetVerifyCertOk() (*bool, bool)`

GetVerifyCertOk returns a tuple with the VerifyCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCert

`func (o *ApUsb) SetVerifyCert(v bool)`

SetVerifyCert sets VerifyCert field to given value.

### HasVerifyCert

`func (o *ApUsb) HasVerifyCert() bool`

HasVerifyCert returns a boolean if a field has been set.

### GetVlanId

`func (o *ApUsb) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *ApUsb) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *ApUsb) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *ApUsb) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


