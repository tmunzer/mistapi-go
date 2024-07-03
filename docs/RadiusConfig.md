# RadiusConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctInterimInterval** | Pointer to **int32** | how frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the radius server, 600 and up is recommended when enabled | [optional] [default to 0]
**AcctServers** | Pointer to [**[]RadiusAcctServer**](RadiusAcctServer.md) |  | [optional] 
**AuthServers** | Pointer to [**[]RadiusAuthServer**](RadiusAuthServer.md) |  | [optional] 
**AuthServersRetries** | Pointer to **int32** | radius auth session retries | [optional] [default to 3]
**AuthServersTimeout** | Pointer to **int32** | radius auth session timeout | [optional] [default to 5]
**CoaEnabled** | Pointer to **bool** |  | [optional] [default to false]
**CoaPort** | Pointer to **int32** |  | [optional] [default to 3799]
**Network** | Pointer to **string** | use &#x60;network&#x60;or &#x60;source_ip&#x60; which network the RADIUS server resides, if there&#39;s static IP for this network, we&#39;d use it as source-ip | [optional] 
**SourceIp** | Pointer to **string** | use &#x60;network&#x60;or &#x60;source_ip&#x60; | [optional] 

## Methods

### NewRadiusConfig

`func NewRadiusConfig() *RadiusConfig`

NewRadiusConfig instantiates a new RadiusConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusConfigWithDefaults

`func NewRadiusConfigWithDefaults() *RadiusConfig`

NewRadiusConfigWithDefaults instantiates a new RadiusConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctInterimInterval

`func (o *RadiusConfig) GetAcctInterimInterval() int32`

GetAcctInterimInterval returns the AcctInterimInterval field if non-nil, zero value otherwise.

### GetAcctInterimIntervalOk

`func (o *RadiusConfig) GetAcctInterimIntervalOk() (*int32, bool)`

GetAcctInterimIntervalOk returns a tuple with the AcctInterimInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctInterimInterval

`func (o *RadiusConfig) SetAcctInterimInterval(v int32)`

SetAcctInterimInterval sets AcctInterimInterval field to given value.

### HasAcctInterimInterval

`func (o *RadiusConfig) HasAcctInterimInterval() bool`

HasAcctInterimInterval returns a boolean if a field has been set.

### GetAcctServers

`func (o *RadiusConfig) GetAcctServers() []RadiusAcctServer`

GetAcctServers returns the AcctServers field if non-nil, zero value otherwise.

### GetAcctServersOk

`func (o *RadiusConfig) GetAcctServersOk() (*[]RadiusAcctServer, bool)`

GetAcctServersOk returns a tuple with the AcctServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctServers

`func (o *RadiusConfig) SetAcctServers(v []RadiusAcctServer)`

SetAcctServers sets AcctServers field to given value.

### HasAcctServers

`func (o *RadiusConfig) HasAcctServers() bool`

HasAcctServers returns a boolean if a field has been set.

### GetAuthServers

`func (o *RadiusConfig) GetAuthServers() []RadiusAuthServer`

GetAuthServers returns the AuthServers field if non-nil, zero value otherwise.

### GetAuthServersOk

`func (o *RadiusConfig) GetAuthServersOk() (*[]RadiusAuthServer, bool)`

GetAuthServersOk returns a tuple with the AuthServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServers

`func (o *RadiusConfig) SetAuthServers(v []RadiusAuthServer)`

SetAuthServers sets AuthServers field to given value.

### HasAuthServers

`func (o *RadiusConfig) HasAuthServers() bool`

HasAuthServers returns a boolean if a field has been set.

### GetAuthServersRetries

`func (o *RadiusConfig) GetAuthServersRetries() int32`

GetAuthServersRetries returns the AuthServersRetries field if non-nil, zero value otherwise.

### GetAuthServersRetriesOk

`func (o *RadiusConfig) GetAuthServersRetriesOk() (*int32, bool)`

GetAuthServersRetriesOk returns a tuple with the AuthServersRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServersRetries

`func (o *RadiusConfig) SetAuthServersRetries(v int32)`

SetAuthServersRetries sets AuthServersRetries field to given value.

### HasAuthServersRetries

`func (o *RadiusConfig) HasAuthServersRetries() bool`

HasAuthServersRetries returns a boolean if a field has been set.

### GetAuthServersTimeout

`func (o *RadiusConfig) GetAuthServersTimeout() int32`

GetAuthServersTimeout returns the AuthServersTimeout field if non-nil, zero value otherwise.

### GetAuthServersTimeoutOk

`func (o *RadiusConfig) GetAuthServersTimeoutOk() (*int32, bool)`

GetAuthServersTimeoutOk returns a tuple with the AuthServersTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServersTimeout

`func (o *RadiusConfig) SetAuthServersTimeout(v int32)`

SetAuthServersTimeout sets AuthServersTimeout field to given value.

### HasAuthServersTimeout

`func (o *RadiusConfig) HasAuthServersTimeout() bool`

HasAuthServersTimeout returns a boolean if a field has been set.

### GetCoaEnabled

`func (o *RadiusConfig) GetCoaEnabled() bool`

GetCoaEnabled returns the CoaEnabled field if non-nil, zero value otherwise.

### GetCoaEnabledOk

`func (o *RadiusConfig) GetCoaEnabledOk() (*bool, bool)`

GetCoaEnabledOk returns a tuple with the CoaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoaEnabled

`func (o *RadiusConfig) SetCoaEnabled(v bool)`

SetCoaEnabled sets CoaEnabled field to given value.

### HasCoaEnabled

`func (o *RadiusConfig) HasCoaEnabled() bool`

HasCoaEnabled returns a boolean if a field has been set.

### GetCoaPort

`func (o *RadiusConfig) GetCoaPort() int32`

GetCoaPort returns the CoaPort field if non-nil, zero value otherwise.

### GetCoaPortOk

`func (o *RadiusConfig) GetCoaPortOk() (*int32, bool)`

GetCoaPortOk returns a tuple with the CoaPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoaPort

`func (o *RadiusConfig) SetCoaPort(v int32)`

SetCoaPort sets CoaPort field to given value.

### HasCoaPort

`func (o *RadiusConfig) HasCoaPort() bool`

HasCoaPort returns a boolean if a field has been set.

### GetNetwork

`func (o *RadiusConfig) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RadiusConfig) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RadiusConfig) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *RadiusConfig) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSourceIp

`func (o *RadiusConfig) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *RadiusConfig) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *RadiusConfig) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *RadiusConfig) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


