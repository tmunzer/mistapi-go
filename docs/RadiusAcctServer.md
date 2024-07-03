# RadiusAcctServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | ip / hostname of RADIUS server | 
**KeywrapEnabled** | Pointer to **bool** |  | [optional] 
**KeywrapFormat** | Pointer to [**RadiusKeywrapFormat**](RadiusKeywrapFormat.md) |  | [optional] 
**KeywrapKek** | Pointer to **string** |  | [optional] 
**KeywrapMack** | Pointer to **string** |  | [optional] 
**Port** | **int32** | Acct port of RADIUS server | [default to 1813]
**Secret** | **string** | secret of RADIUS server | 

## Methods

### NewRadiusAcctServer

`func NewRadiusAcctServer(host string, port int32, secret string, ) *RadiusAcctServer`

NewRadiusAcctServer instantiates a new RadiusAcctServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusAcctServerWithDefaults

`func NewRadiusAcctServerWithDefaults() *RadiusAcctServer`

NewRadiusAcctServerWithDefaults instantiates a new RadiusAcctServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *RadiusAcctServer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RadiusAcctServer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RadiusAcctServer) SetHost(v string)`

SetHost sets Host field to given value.


### GetKeywrapEnabled

`func (o *RadiusAcctServer) GetKeywrapEnabled() bool`

GetKeywrapEnabled returns the KeywrapEnabled field if non-nil, zero value otherwise.

### GetKeywrapEnabledOk

`func (o *RadiusAcctServer) GetKeywrapEnabledOk() (*bool, bool)`

GetKeywrapEnabledOk returns a tuple with the KeywrapEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywrapEnabled

`func (o *RadiusAcctServer) SetKeywrapEnabled(v bool)`

SetKeywrapEnabled sets KeywrapEnabled field to given value.

### HasKeywrapEnabled

`func (o *RadiusAcctServer) HasKeywrapEnabled() bool`

HasKeywrapEnabled returns a boolean if a field has been set.

### GetKeywrapFormat

`func (o *RadiusAcctServer) GetKeywrapFormat() RadiusKeywrapFormat`

GetKeywrapFormat returns the KeywrapFormat field if non-nil, zero value otherwise.

### GetKeywrapFormatOk

`func (o *RadiusAcctServer) GetKeywrapFormatOk() (*RadiusKeywrapFormat, bool)`

GetKeywrapFormatOk returns a tuple with the KeywrapFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywrapFormat

`func (o *RadiusAcctServer) SetKeywrapFormat(v RadiusKeywrapFormat)`

SetKeywrapFormat sets KeywrapFormat field to given value.

### HasKeywrapFormat

`func (o *RadiusAcctServer) HasKeywrapFormat() bool`

HasKeywrapFormat returns a boolean if a field has been set.

### GetKeywrapKek

`func (o *RadiusAcctServer) GetKeywrapKek() string`

GetKeywrapKek returns the KeywrapKek field if non-nil, zero value otherwise.

### GetKeywrapKekOk

`func (o *RadiusAcctServer) GetKeywrapKekOk() (*string, bool)`

GetKeywrapKekOk returns a tuple with the KeywrapKek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywrapKek

`func (o *RadiusAcctServer) SetKeywrapKek(v string)`

SetKeywrapKek sets KeywrapKek field to given value.

### HasKeywrapKek

`func (o *RadiusAcctServer) HasKeywrapKek() bool`

HasKeywrapKek returns a boolean if a field has been set.

### GetKeywrapMack

`func (o *RadiusAcctServer) GetKeywrapMack() string`

GetKeywrapMack returns the KeywrapMack field if non-nil, zero value otherwise.

### GetKeywrapMackOk

`func (o *RadiusAcctServer) GetKeywrapMackOk() (*string, bool)`

GetKeywrapMackOk returns a tuple with the KeywrapMack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywrapMack

`func (o *RadiusAcctServer) SetKeywrapMack(v string)`

SetKeywrapMack sets KeywrapMack field to given value.

### HasKeywrapMack

`func (o *RadiusAcctServer) HasKeywrapMack() bool`

HasKeywrapMack returns a boolean if a field has been set.

### GetPort

`func (o *RadiusAcctServer) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RadiusAcctServer) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RadiusAcctServer) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSecret

`func (o *RadiusAcctServer) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *RadiusAcctServer) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *RadiusAcctServer) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


