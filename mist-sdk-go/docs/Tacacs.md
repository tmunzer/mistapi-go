# Tacacs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctServers** | Pointer to [**[]TacacsAcctServer**](TacacsAcctServer.md) |  | [optional] 
**DefaultRole** | Pointer to [**TacacsDefaultRole**](TacacsDefaultRole.md) |  | [optional] [default to TACACSDEFAULTROLE_NONE]
**Enabled** | Pointer to **bool** |  | [optional] 
**Network** | Pointer to **string** | which network the TACACS server resides | [optional] 
**TacplusServers** | Pointer to [**[]TacacsAuthServer**](TacacsAuthServer.md) |  | [optional] 

## Methods

### NewTacacs

`func NewTacacs() *Tacacs`

NewTacacs instantiates a new Tacacs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTacacsWithDefaults

`func NewTacacsWithDefaults() *Tacacs`

NewTacacsWithDefaults instantiates a new Tacacs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctServers

`func (o *Tacacs) GetAcctServers() []TacacsAcctServer`

GetAcctServers returns the AcctServers field if non-nil, zero value otherwise.

### GetAcctServersOk

`func (o *Tacacs) GetAcctServersOk() (*[]TacacsAcctServer, bool)`

GetAcctServersOk returns a tuple with the AcctServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctServers

`func (o *Tacacs) SetAcctServers(v []TacacsAcctServer)`

SetAcctServers sets AcctServers field to given value.

### HasAcctServers

`func (o *Tacacs) HasAcctServers() bool`

HasAcctServers returns a boolean if a field has been set.

### GetDefaultRole

`func (o *Tacacs) GetDefaultRole() TacacsDefaultRole`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *Tacacs) GetDefaultRoleOk() (*TacacsDefaultRole, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *Tacacs) SetDefaultRole(v TacacsDefaultRole)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *Tacacs) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetEnabled

`func (o *Tacacs) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Tacacs) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Tacacs) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Tacacs) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNetwork

`func (o *Tacacs) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Tacacs) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Tacacs) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Tacacs) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetTacplusServers

`func (o *Tacacs) GetTacplusServers() []TacacsAuthServer`

GetTacplusServers returns the TacplusServers field if non-nil, zero value otherwise.

### GetTacplusServersOk

`func (o *Tacacs) GetTacplusServersOk() (*[]TacacsAuthServer, bool)`

GetTacplusServersOk returns a tuple with the TacplusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacplusServers

`func (o *Tacacs) SetTacplusServers(v []TacacsAuthServer)`

SetTacplusServers sets TacplusServers field to given value.

### HasTacplusServers

`func (o *Tacacs) HasTacplusServers() bool`

HasTacplusServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


