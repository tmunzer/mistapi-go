# SsoMxedgeProxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctServers** | Pointer to [**[]SsoMxedgeProxyAcctServer**](SsoMxedgeProxyAcctServer.md) |  | [optional] 
**AuthServers** | Pointer to [**[]SsoMxedgeProxyAuthServer**](SsoMxedgeProxyAuthServer.md) |  | [optional] 
**MxclusterId** | Pointer to **string** |  | [optional] 
**OperatorName** | Pointer to **string** | Operator name as Radius attribute while proxying | [optional] 
**ProxyHosts** | Pointer to **[]string** | public hostname/IPs | [optional] 
**Ssids** | Pointer to **[]string** | SSIDs that support eduroam | [optional] 

## Methods

### NewSsoMxedgeProxy

`func NewSsoMxedgeProxy() *SsoMxedgeProxy`

NewSsoMxedgeProxy instantiates a new SsoMxedgeProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoMxedgeProxyWithDefaults

`func NewSsoMxedgeProxyWithDefaults() *SsoMxedgeProxy`

NewSsoMxedgeProxyWithDefaults instantiates a new SsoMxedgeProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctServers

`func (o *SsoMxedgeProxy) GetAcctServers() []SsoMxedgeProxyAcctServer`

GetAcctServers returns the AcctServers field if non-nil, zero value otherwise.

### GetAcctServersOk

`func (o *SsoMxedgeProxy) GetAcctServersOk() (*[]SsoMxedgeProxyAcctServer, bool)`

GetAcctServersOk returns a tuple with the AcctServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctServers

`func (o *SsoMxedgeProxy) SetAcctServers(v []SsoMxedgeProxyAcctServer)`

SetAcctServers sets AcctServers field to given value.

### HasAcctServers

`func (o *SsoMxedgeProxy) HasAcctServers() bool`

HasAcctServers returns a boolean if a field has been set.

### GetAuthServers

`func (o *SsoMxedgeProxy) GetAuthServers() []SsoMxedgeProxyAuthServer`

GetAuthServers returns the AuthServers field if non-nil, zero value otherwise.

### GetAuthServersOk

`func (o *SsoMxedgeProxy) GetAuthServersOk() (*[]SsoMxedgeProxyAuthServer, bool)`

GetAuthServersOk returns a tuple with the AuthServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServers

`func (o *SsoMxedgeProxy) SetAuthServers(v []SsoMxedgeProxyAuthServer)`

SetAuthServers sets AuthServers field to given value.

### HasAuthServers

`func (o *SsoMxedgeProxy) HasAuthServers() bool`

HasAuthServers returns a boolean if a field has been set.

### GetMxclusterId

`func (o *SsoMxedgeProxy) GetMxclusterId() string`

GetMxclusterId returns the MxclusterId field if non-nil, zero value otherwise.

### GetMxclusterIdOk

`func (o *SsoMxedgeProxy) GetMxclusterIdOk() (*string, bool)`

GetMxclusterIdOk returns a tuple with the MxclusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxclusterId

`func (o *SsoMxedgeProxy) SetMxclusterId(v string)`

SetMxclusterId sets MxclusterId field to given value.

### HasMxclusterId

`func (o *SsoMxedgeProxy) HasMxclusterId() bool`

HasMxclusterId returns a boolean if a field has been set.

### GetOperatorName

`func (o *SsoMxedgeProxy) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *SsoMxedgeProxy) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *SsoMxedgeProxy) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.

### HasOperatorName

`func (o *SsoMxedgeProxy) HasOperatorName() bool`

HasOperatorName returns a boolean if a field has been set.

### GetProxyHosts

`func (o *SsoMxedgeProxy) GetProxyHosts() []string`

GetProxyHosts returns the ProxyHosts field if non-nil, zero value otherwise.

### GetProxyHostsOk

`func (o *SsoMxedgeProxy) GetProxyHostsOk() (*[]string, bool)`

GetProxyHostsOk returns a tuple with the ProxyHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHosts

`func (o *SsoMxedgeProxy) SetProxyHosts(v []string)`

SetProxyHosts sets ProxyHosts field to given value.

### HasProxyHosts

`func (o *SsoMxedgeProxy) HasProxyHosts() bool`

HasProxyHosts returns a boolean if a field has been set.

### GetSsids

`func (o *SsoMxedgeProxy) GetSsids() []string`

GetSsids returns the Ssids field if non-nil, zero value otherwise.

### GetSsidsOk

`func (o *SsoMxedgeProxy) GetSsidsOk() (*[]string, bool)`

GetSsidsOk returns a tuple with the Ssids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsids

`func (o *SsoMxedgeProxy) SetSsids(v []string)`

SetSsids sets Ssids field to given value.

### HasSsids

`func (o *SsoMxedgeProxy) HasSsids() bool`

HasSsids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


