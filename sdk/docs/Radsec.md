# Radsec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoaEnabled** | Pointer to **bool** |  | [optional] [default to false]
**Enabled** | Pointer to **bool** |  | [optional] 
**IdleTimeout** | Pointer to **int32** |  | [optional] 
**MxclusterIds** | Pointer to **[]string** | To use Org mxedges when this WLAN does not use mxtunnel, specify their mxcluster_ids. Org mxedge(s) identified by mxcluster_ids | [optional] 
**ProxyHosts** | Pointer to **[]string** | default is site.mxedge.radsec.proxy_hosts which must be a superset of all wlans[*].radsec.proxy_hosts when radsec.proxy_hosts are not used, tunnel peers (org or site mxedges) are used irrespective of use_site_mxedge | [optional] 
**ServerName** | Pointer to **string** | name of the server to verify (against the cacerts in Org Setting). Only if not Mist Edge. | [optional] 
**Servers** | Pointer to [**[]RadsecServer**](RadsecServer.md) | List of Radsec Servers. Only if not Mist Edge. | [optional] 
**UseMxedge** | Pointer to **bool** | use mxedge(s) as radsecproxy | [optional] 
**UseSiteMxedge** | Pointer to **bool** | To use Site mxedges when this WLAN does not use mxtunnel | [optional] [default to false]

## Methods

### NewRadsec

`func NewRadsec() *Radsec`

NewRadsec instantiates a new Radsec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadsecWithDefaults

`func NewRadsecWithDefaults() *Radsec`

NewRadsecWithDefaults instantiates a new Radsec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoaEnabled

`func (o *Radsec) GetCoaEnabled() bool`

GetCoaEnabled returns the CoaEnabled field if non-nil, zero value otherwise.

### GetCoaEnabledOk

`func (o *Radsec) GetCoaEnabledOk() (*bool, bool)`

GetCoaEnabledOk returns a tuple with the CoaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoaEnabled

`func (o *Radsec) SetCoaEnabled(v bool)`

SetCoaEnabled sets CoaEnabled field to given value.

### HasCoaEnabled

`func (o *Radsec) HasCoaEnabled() bool`

HasCoaEnabled returns a boolean if a field has been set.

### GetEnabled

`func (o *Radsec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Radsec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Radsec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Radsec) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIdleTimeout

`func (o *Radsec) GetIdleTimeout() int32`

GetIdleTimeout returns the IdleTimeout field if non-nil, zero value otherwise.

### GetIdleTimeoutOk

`func (o *Radsec) GetIdleTimeoutOk() (*int32, bool)`

GetIdleTimeoutOk returns a tuple with the IdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeout

`func (o *Radsec) SetIdleTimeout(v int32)`

SetIdleTimeout sets IdleTimeout field to given value.

### HasIdleTimeout

`func (o *Radsec) HasIdleTimeout() bool`

HasIdleTimeout returns a boolean if a field has been set.

### GetMxclusterIds

`func (o *Radsec) GetMxclusterIds() []string`

GetMxclusterIds returns the MxclusterIds field if non-nil, zero value otherwise.

### GetMxclusterIdsOk

`func (o *Radsec) GetMxclusterIdsOk() (*[]string, bool)`

GetMxclusterIdsOk returns a tuple with the MxclusterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxclusterIds

`func (o *Radsec) SetMxclusterIds(v []string)`

SetMxclusterIds sets MxclusterIds field to given value.

### HasMxclusterIds

`func (o *Radsec) HasMxclusterIds() bool`

HasMxclusterIds returns a boolean if a field has been set.

### GetProxyHosts

`func (o *Radsec) GetProxyHosts() []string`

GetProxyHosts returns the ProxyHosts field if non-nil, zero value otherwise.

### GetProxyHostsOk

`func (o *Radsec) GetProxyHostsOk() (*[]string, bool)`

GetProxyHostsOk returns a tuple with the ProxyHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHosts

`func (o *Radsec) SetProxyHosts(v []string)`

SetProxyHosts sets ProxyHosts field to given value.

### HasProxyHosts

`func (o *Radsec) HasProxyHosts() bool`

HasProxyHosts returns a boolean if a field has been set.

### GetServerName

`func (o *Radsec) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *Radsec) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *Radsec) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *Radsec) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServers

`func (o *Radsec) GetServers() []RadsecServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *Radsec) GetServersOk() (*[]RadsecServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *Radsec) SetServers(v []RadsecServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *Radsec) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetUseMxedge

`func (o *Radsec) GetUseMxedge() bool`

GetUseMxedge returns the UseMxedge field if non-nil, zero value otherwise.

### GetUseMxedgeOk

`func (o *Radsec) GetUseMxedgeOk() (*bool, bool)`

GetUseMxedgeOk returns a tuple with the UseMxedge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMxedge

`func (o *Radsec) SetUseMxedge(v bool)`

SetUseMxedge sets UseMxedge field to given value.

### HasUseMxedge

`func (o *Radsec) HasUseMxedge() bool`

HasUseMxedge returns a boolean if a field has been set.

### GetUseSiteMxedge

`func (o *Radsec) GetUseSiteMxedge() bool`

GetUseSiteMxedge returns the UseSiteMxedge field if non-nil, zero value otherwise.

### GetUseSiteMxedgeOk

`func (o *Radsec) GetUseSiteMxedgeOk() (*bool, bool)`

GetUseSiteMxedgeOk returns a tuple with the UseSiteMxedge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSiteMxedge

`func (o *Radsec) SetUseSiteMxedge(v bool)`

SetUseSiteMxedge sets UseSiteMxedge field to given value.

### HasUseSiteMxedge

`func (o *Radsec) HasUseSiteMxedge() bool`

HasUseSiteMxedge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


