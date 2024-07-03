# MxclusterRadsecAcctServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | ip / hostname of RADIUS server | [optional] 
**Port** | Pointer to **int32** | Acct port of RADIUS server | [optional] [default to 1813]
**Secret** | Pointer to **string** | secret of RADIUS server | [optional] 
**Ssids** | Pointer to **[]string** | list of ssids that will use this server if match_ssid is true and match is found | [optional] 

## Methods

### NewMxclusterRadsecAcctServer

`func NewMxclusterRadsecAcctServer() *MxclusterRadsecAcctServer`

NewMxclusterRadsecAcctServer instantiates a new MxclusterRadsecAcctServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxclusterRadsecAcctServerWithDefaults

`func NewMxclusterRadsecAcctServerWithDefaults() *MxclusterRadsecAcctServer`

NewMxclusterRadsecAcctServerWithDefaults instantiates a new MxclusterRadsecAcctServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *MxclusterRadsecAcctServer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MxclusterRadsecAcctServer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MxclusterRadsecAcctServer) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *MxclusterRadsecAcctServer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *MxclusterRadsecAcctServer) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MxclusterRadsecAcctServer) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MxclusterRadsecAcctServer) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *MxclusterRadsecAcctServer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecret

`func (o *MxclusterRadsecAcctServer) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *MxclusterRadsecAcctServer) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *MxclusterRadsecAcctServer) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *MxclusterRadsecAcctServer) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSsids

`func (o *MxclusterRadsecAcctServer) GetSsids() []string`

GetSsids returns the Ssids field if non-nil, zero value otherwise.

### GetSsidsOk

`func (o *MxclusterRadsecAcctServer) GetSsidsOk() (*[]string, bool)`

GetSsidsOk returns a tuple with the Ssids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsids

`func (o *MxclusterRadsecAcctServer) SetSsids(v []string)`

SetSsids sets Ssids field to given value.

### HasSsids

`func (o *MxclusterRadsecAcctServer) HasSsids() bool`

HasSsids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


