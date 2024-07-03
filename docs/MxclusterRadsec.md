# MxclusterRadsec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctServers** | Pointer to [**[]MxclusterRadsecAcctServer**](MxclusterRadsecAcctServer.md) | list of RADIUS accounting servers, optional, order matters where the first one is treated as primary | [optional] 
**AuthServers** | Pointer to [**[]MxclusterRadsecAuthServer**](MxclusterRadsecAuthServer.md) | list of RADIUS authentication servers, order matters where the first one is treated as primary | [optional] 
**Enabled** | Pointer to **bool** | whether to enable service on Mist Edge i.e. RADIUS proxy over TLS | [optional] 
**MatchSsid** | Pointer to **bool** | whether to match ssid in request message to select from a subset of RADIUS servers | [optional] 
**ProxyHosts** | Pointer to **[]string** | hostnames or IPs for Mist AP to use as the TLS Server (i.e. they are reachable from AP) in addition to &#x60;tunterm_hosts&#x60; | [optional] 
**ServerSelection** | Pointer to [**MxclusterRadsecServerSelection**](MxclusterRadsecServerSelection.md) |  | [optional] [default to MXCLUSTERRADSECSERVERSELECTION_ORDERED]
**Source** | Pointer to [**MxclusterRadsecSource**](MxclusterRadsecSource.md) |  | [optional] [default to MXCLUSTERRADSECSOURCE_ANY]

## Methods

### NewMxclusterRadsec

`func NewMxclusterRadsec() *MxclusterRadsec`

NewMxclusterRadsec instantiates a new MxclusterRadsec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxclusterRadsecWithDefaults

`func NewMxclusterRadsecWithDefaults() *MxclusterRadsec`

NewMxclusterRadsecWithDefaults instantiates a new MxclusterRadsec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctServers

`func (o *MxclusterRadsec) GetAcctServers() []MxclusterRadsecAcctServer`

GetAcctServers returns the AcctServers field if non-nil, zero value otherwise.

### GetAcctServersOk

`func (o *MxclusterRadsec) GetAcctServersOk() (*[]MxclusterRadsecAcctServer, bool)`

GetAcctServersOk returns a tuple with the AcctServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctServers

`func (o *MxclusterRadsec) SetAcctServers(v []MxclusterRadsecAcctServer)`

SetAcctServers sets AcctServers field to given value.

### HasAcctServers

`func (o *MxclusterRadsec) HasAcctServers() bool`

HasAcctServers returns a boolean if a field has been set.

### GetAuthServers

`func (o *MxclusterRadsec) GetAuthServers() []MxclusterRadsecAuthServer`

GetAuthServers returns the AuthServers field if non-nil, zero value otherwise.

### GetAuthServersOk

`func (o *MxclusterRadsec) GetAuthServersOk() (*[]MxclusterRadsecAuthServer, bool)`

GetAuthServersOk returns a tuple with the AuthServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServers

`func (o *MxclusterRadsec) SetAuthServers(v []MxclusterRadsecAuthServer)`

SetAuthServers sets AuthServers field to given value.

### HasAuthServers

`func (o *MxclusterRadsec) HasAuthServers() bool`

HasAuthServers returns a boolean if a field has been set.

### GetEnabled

`func (o *MxclusterRadsec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MxclusterRadsec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MxclusterRadsec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MxclusterRadsec) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMatchSsid

`func (o *MxclusterRadsec) GetMatchSsid() bool`

GetMatchSsid returns the MatchSsid field if non-nil, zero value otherwise.

### GetMatchSsidOk

`func (o *MxclusterRadsec) GetMatchSsidOk() (*bool, bool)`

GetMatchSsidOk returns a tuple with the MatchSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchSsid

`func (o *MxclusterRadsec) SetMatchSsid(v bool)`

SetMatchSsid sets MatchSsid field to given value.

### HasMatchSsid

`func (o *MxclusterRadsec) HasMatchSsid() bool`

HasMatchSsid returns a boolean if a field has been set.

### GetProxyHosts

`func (o *MxclusterRadsec) GetProxyHosts() []string`

GetProxyHosts returns the ProxyHosts field if non-nil, zero value otherwise.

### GetProxyHostsOk

`func (o *MxclusterRadsec) GetProxyHostsOk() (*[]string, bool)`

GetProxyHostsOk returns a tuple with the ProxyHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHosts

`func (o *MxclusterRadsec) SetProxyHosts(v []string)`

SetProxyHosts sets ProxyHosts field to given value.

### HasProxyHosts

`func (o *MxclusterRadsec) HasProxyHosts() bool`

HasProxyHosts returns a boolean if a field has been set.

### GetServerSelection

`func (o *MxclusterRadsec) GetServerSelection() MxclusterRadsecServerSelection`

GetServerSelection returns the ServerSelection field if non-nil, zero value otherwise.

### GetServerSelectionOk

`func (o *MxclusterRadsec) GetServerSelectionOk() (*MxclusterRadsecServerSelection, bool)`

GetServerSelectionOk returns a tuple with the ServerSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSelection

`func (o *MxclusterRadsec) SetServerSelection(v MxclusterRadsecServerSelection)`

SetServerSelection sets ServerSelection field to given value.

### HasServerSelection

`func (o *MxclusterRadsec) HasServerSelection() bool`

HasServerSelection returns a boolean if a field has been set.

### GetSource

`func (o *MxclusterRadsec) GetSource() MxclusterRadsecSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MxclusterRadsec) GetSourceOk() (*MxclusterRadsecSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MxclusterRadsec) SetSource(v MxclusterRadsecSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *MxclusterRadsec) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


