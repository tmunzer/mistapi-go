# MxclusterNac

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctServerPort** | Pointer to **int32** |  | [optional] [default to 1813]
**AuthServerPort** | Pointer to **int32** |  | [optional] [default to 1812]
**ClientIps** | Pointer to [**map[string]map[string]MxclusterNacClientIp**](map.md) | Property key is the RADIUS Client IP/Subnet. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**Secret** | Pointer to **string** |  | [optional] 

## Methods

### NewMxclusterNac

`func NewMxclusterNac() *MxclusterNac`

NewMxclusterNac instantiates a new MxclusterNac object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxclusterNacWithDefaults

`func NewMxclusterNacWithDefaults() *MxclusterNac`

NewMxclusterNacWithDefaults instantiates a new MxclusterNac object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctServerPort

`func (o *MxclusterNac) GetAcctServerPort() int32`

GetAcctServerPort returns the AcctServerPort field if non-nil, zero value otherwise.

### GetAcctServerPortOk

`func (o *MxclusterNac) GetAcctServerPortOk() (*int32, bool)`

GetAcctServerPortOk returns a tuple with the AcctServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctServerPort

`func (o *MxclusterNac) SetAcctServerPort(v int32)`

SetAcctServerPort sets AcctServerPort field to given value.

### HasAcctServerPort

`func (o *MxclusterNac) HasAcctServerPort() bool`

HasAcctServerPort returns a boolean if a field has been set.

### GetAuthServerPort

`func (o *MxclusterNac) GetAuthServerPort() int32`

GetAuthServerPort returns the AuthServerPort field if non-nil, zero value otherwise.

### GetAuthServerPortOk

`func (o *MxclusterNac) GetAuthServerPortOk() (*int32, bool)`

GetAuthServerPortOk returns a tuple with the AuthServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServerPort

`func (o *MxclusterNac) SetAuthServerPort(v int32)`

SetAuthServerPort sets AuthServerPort field to given value.

### HasAuthServerPort

`func (o *MxclusterNac) HasAuthServerPort() bool`

HasAuthServerPort returns a boolean if a field has been set.

### GetClientIps

`func (o *MxclusterNac) GetClientIps() map[string]map[string]MxclusterNacClientIp`

GetClientIps returns the ClientIps field if non-nil, zero value otherwise.

### GetClientIpsOk

`func (o *MxclusterNac) GetClientIpsOk() (*map[string]map[string]MxclusterNacClientIp, bool)`

GetClientIpsOk returns a tuple with the ClientIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIps

`func (o *MxclusterNac) SetClientIps(v map[string]map[string]MxclusterNacClientIp)`

SetClientIps sets ClientIps field to given value.

### HasClientIps

`func (o *MxclusterNac) HasClientIps() bool`

HasClientIps returns a boolean if a field has been set.

### GetEnabled

`func (o *MxclusterNac) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MxclusterNac) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MxclusterNac) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MxclusterNac) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSecret

`func (o *MxclusterNac) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *MxclusterNac) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *MxclusterNac) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *MxclusterNac) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


