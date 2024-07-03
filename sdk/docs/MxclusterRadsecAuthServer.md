# MxclusterRadsecAuthServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | ip / hostname of RADIUS server | [optional] 
**KeywrapEnabled** | Pointer to **bool** | if used for Mist APs, enable keywrap algorithm. Default is false | [optional] 
**KeywrapFormat** | Pointer to [**NullableMxclusterRadAuthServerKeywrapFormat**](MxclusterRadAuthServerKeywrapFormat.md) |  | [optional] [default to MXCLUSTERRADAUTHSERVERKEYWRAPFORMAT_ASCII]
**KeywrapKek** | Pointer to **string** | if used for Mist APs, encryption key | [optional] 
**KeywrapMack** | Pointer to **string** | if used for Mist APs, Message Authentication Code Key | [optional] 
**Port** | Pointer to **int32** | Auth port of RADIUS server | [optional] [default to 1812]
**Secret** | Pointer to **string** | secret of RADIUS server | [optional] 
**Ssids** | Pointer to **[]string** | list of ssids that will use this server if match_ssid is true and match is found | [optional] 

## Methods

### NewMxclusterRadsecAuthServer

`func NewMxclusterRadsecAuthServer() *MxclusterRadsecAuthServer`

NewMxclusterRadsecAuthServer instantiates a new MxclusterRadsecAuthServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxclusterRadsecAuthServerWithDefaults

`func NewMxclusterRadsecAuthServerWithDefaults() *MxclusterRadsecAuthServer`

NewMxclusterRadsecAuthServerWithDefaults instantiates a new MxclusterRadsecAuthServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *MxclusterRadsecAuthServer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MxclusterRadsecAuthServer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MxclusterRadsecAuthServer) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *MxclusterRadsecAuthServer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKeywrapEnabled

`func (o *MxclusterRadsecAuthServer) GetKeywrapEnabled() bool`

GetKeywrapEnabled returns the KeywrapEnabled field if non-nil, zero value otherwise.

### GetKeywrapEnabledOk

`func (o *MxclusterRadsecAuthServer) GetKeywrapEnabledOk() (*bool, bool)`

GetKeywrapEnabledOk returns a tuple with the KeywrapEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywrapEnabled

`func (o *MxclusterRadsecAuthServer) SetKeywrapEnabled(v bool)`

SetKeywrapEnabled sets KeywrapEnabled field to given value.

### HasKeywrapEnabled

`func (o *MxclusterRadsecAuthServer) HasKeywrapEnabled() bool`

HasKeywrapEnabled returns a boolean if a field has been set.

### GetKeywrapFormat

`func (o *MxclusterRadsecAuthServer) GetKeywrapFormat() MxclusterRadAuthServerKeywrapFormat`

GetKeywrapFormat returns the KeywrapFormat field if non-nil, zero value otherwise.

### GetKeywrapFormatOk

`func (o *MxclusterRadsecAuthServer) GetKeywrapFormatOk() (*MxclusterRadAuthServerKeywrapFormat, bool)`

GetKeywrapFormatOk returns a tuple with the KeywrapFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywrapFormat

`func (o *MxclusterRadsecAuthServer) SetKeywrapFormat(v MxclusterRadAuthServerKeywrapFormat)`

SetKeywrapFormat sets KeywrapFormat field to given value.

### HasKeywrapFormat

`func (o *MxclusterRadsecAuthServer) HasKeywrapFormat() bool`

HasKeywrapFormat returns a boolean if a field has been set.

### SetKeywrapFormatNil

`func (o *MxclusterRadsecAuthServer) SetKeywrapFormatNil(b bool)`

 SetKeywrapFormatNil sets the value for KeywrapFormat to be an explicit nil

### UnsetKeywrapFormat
`func (o *MxclusterRadsecAuthServer) UnsetKeywrapFormat()`

UnsetKeywrapFormat ensures that no value is present for KeywrapFormat, not even an explicit nil
### GetKeywrapKek

`func (o *MxclusterRadsecAuthServer) GetKeywrapKek() string`

GetKeywrapKek returns the KeywrapKek field if non-nil, zero value otherwise.

### GetKeywrapKekOk

`func (o *MxclusterRadsecAuthServer) GetKeywrapKekOk() (*string, bool)`

GetKeywrapKekOk returns a tuple with the KeywrapKek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywrapKek

`func (o *MxclusterRadsecAuthServer) SetKeywrapKek(v string)`

SetKeywrapKek sets KeywrapKek field to given value.

### HasKeywrapKek

`func (o *MxclusterRadsecAuthServer) HasKeywrapKek() bool`

HasKeywrapKek returns a boolean if a field has been set.

### GetKeywrapMack

`func (o *MxclusterRadsecAuthServer) GetKeywrapMack() string`

GetKeywrapMack returns the KeywrapMack field if non-nil, zero value otherwise.

### GetKeywrapMackOk

`func (o *MxclusterRadsecAuthServer) GetKeywrapMackOk() (*string, bool)`

GetKeywrapMackOk returns a tuple with the KeywrapMack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywrapMack

`func (o *MxclusterRadsecAuthServer) SetKeywrapMack(v string)`

SetKeywrapMack sets KeywrapMack field to given value.

### HasKeywrapMack

`func (o *MxclusterRadsecAuthServer) HasKeywrapMack() bool`

HasKeywrapMack returns a boolean if a field has been set.

### GetPort

`func (o *MxclusterRadsecAuthServer) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MxclusterRadsecAuthServer) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MxclusterRadsecAuthServer) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *MxclusterRadsecAuthServer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecret

`func (o *MxclusterRadsecAuthServer) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *MxclusterRadsecAuthServer) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *MxclusterRadsecAuthServer) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *MxclusterRadsecAuthServer) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSsids

`func (o *MxclusterRadsecAuthServer) GetSsids() []string`

GetSsids returns the Ssids field if non-nil, zero value otherwise.

### GetSsidsOk

`func (o *MxclusterRadsecAuthServer) GetSsidsOk() (*[]string, bool)`

GetSsidsOk returns a tuple with the Ssids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsids

`func (o *MxclusterRadsecAuthServer) SetSsids(v []string)`

SetSsids sets Ssids field to given value.

### HasSsids

`func (o *MxclusterRadsecAuthServer) HasSsids() bool`

HasSsids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


