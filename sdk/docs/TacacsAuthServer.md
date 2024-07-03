# TacacsAuthServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] [default to 10]

## Methods

### NewTacacsAuthServer

`func NewTacacsAuthServer() *TacacsAuthServer`

NewTacacsAuthServer instantiates a new TacacsAuthServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTacacsAuthServerWithDefaults

`func NewTacacsAuthServerWithDefaults() *TacacsAuthServer`

NewTacacsAuthServerWithDefaults instantiates a new TacacsAuthServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *TacacsAuthServer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TacacsAuthServer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TacacsAuthServer) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TacacsAuthServer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *TacacsAuthServer) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TacacsAuthServer) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TacacsAuthServer) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TacacsAuthServer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecret

`func (o *TacacsAuthServer) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TacacsAuthServer) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TacacsAuthServer) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *TacacsAuthServer) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetTimeout

`func (o *TacacsAuthServer) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TacacsAuthServer) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TacacsAuthServer) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TacacsAuthServer) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


