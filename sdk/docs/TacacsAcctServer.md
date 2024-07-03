# TacacsAcctServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] [default to 10]

## Methods

### NewTacacsAcctServer

`func NewTacacsAcctServer() *TacacsAcctServer`

NewTacacsAcctServer instantiates a new TacacsAcctServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTacacsAcctServerWithDefaults

`func NewTacacsAcctServerWithDefaults() *TacacsAcctServer`

NewTacacsAcctServerWithDefaults instantiates a new TacacsAcctServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *TacacsAcctServer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TacacsAcctServer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TacacsAcctServer) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TacacsAcctServer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *TacacsAcctServer) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TacacsAcctServer) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TacacsAcctServer) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TacacsAcctServer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecret

`func (o *TacacsAcctServer) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TacacsAcctServer) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TacacsAcctServer) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *TacacsAcctServer) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetTimeout

`func (o *TacacsAcctServer) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TacacsAcctServer) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TacacsAcctServer) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TacacsAcctServer) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


