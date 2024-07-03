# WxlanTunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Dmvpn** | Pointer to [**WxlanTunnelDmvpn**](WxlanTunnelDmvpn.md) |  | [optional] 
**ForMgmt** | Pointer to **bool** | determined during creation time and cannot be toggled. A management tunnel cannot be used by wxlan rule or by wlan | [optional] [default to false]
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**HelloInterval** | Pointer to **int32** | in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries. | [optional] [default to 60]
**HelloRetries** | Pointer to **int32** |  | [optional] [default to 7]
**Hostname** | Pointer to **string** | optional, overwrite the hostname in SCCRQ control message, default is “” or null, %H and %M can be used, which will be replace with corresponding values: * %H: name of the ap if provided (and will be stripped so it can be used for hostname) and fallbacks to MAC * %M: MAC (e.g. 5c5b350e0060) | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Ipsec** | Pointer to [**WxlanTunnelIpsec**](WxlanTunnelIpsec.md) |  | [optional] 
**IsStatic** | Pointer to **bool** | whether it’s static/unmanaged (i.e. no control session). As the session configurations are not compatible, cannot be toggled. | [optional] [default to false]
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Mtu** | Pointer to **int32** | 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU | [optional] [default to 0]
**Name** | **string** | The name of the tunnel | 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Peers** | Pointer to **[]string** | list of remote peers’ IP or hostname | [optional] 
**RouterId** | Pointer to **string** | optional, overwrite the router-id in SCCRQ control message, default is “0” or null, can also be an IPv4 address | [optional] 
**Secret** | Pointer to **string** | secret, ‘’ if no auth is used | [optional] 
**Sessions** | Pointer to [**[]WxlanTunnelSession**](WxlanTunnelSession.md) | sessions to be established with the tunnel. Has to be &gt;&#x3D; 1 in order for this tunnel to be useful. For management tunnel, it can only have 1 | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**UdpPort** | Pointer to **int32** | udp port if &#x60;use_udp&#x60;&#x3D;&#x3D;&#x60;true&#x60; | [optional] 
**UseUdp** | Pointer to **bool** | whether to use UDP instead of IP (proto&#x3D;115, which is default of L2TPv3) | [optional] [default to false]

## Methods

### NewWxlanTunnel

`func NewWxlanTunnel(name string, ) *WxlanTunnel`

NewWxlanTunnel instantiates a new WxlanTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWxlanTunnelWithDefaults

`func NewWxlanTunnelWithDefaults() *WxlanTunnel`

NewWxlanTunnelWithDefaults instantiates a new WxlanTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *WxlanTunnel) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *WxlanTunnel) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *WxlanTunnel) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *WxlanTunnel) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDmvpn

`func (o *WxlanTunnel) GetDmvpn() WxlanTunnelDmvpn`

GetDmvpn returns the Dmvpn field if non-nil, zero value otherwise.

### GetDmvpnOk

`func (o *WxlanTunnel) GetDmvpnOk() (*WxlanTunnelDmvpn, bool)`

GetDmvpnOk returns a tuple with the Dmvpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmvpn

`func (o *WxlanTunnel) SetDmvpn(v WxlanTunnelDmvpn)`

SetDmvpn sets Dmvpn field to given value.

### HasDmvpn

`func (o *WxlanTunnel) HasDmvpn() bool`

HasDmvpn returns a boolean if a field has been set.

### GetForMgmt

`func (o *WxlanTunnel) GetForMgmt() bool`

GetForMgmt returns the ForMgmt field if non-nil, zero value otherwise.

### GetForMgmtOk

`func (o *WxlanTunnel) GetForMgmtOk() (*bool, bool)`

GetForMgmtOk returns a tuple with the ForMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForMgmt

`func (o *WxlanTunnel) SetForMgmt(v bool)`

SetForMgmt sets ForMgmt field to given value.

### HasForMgmt

`func (o *WxlanTunnel) HasForMgmt() bool`

HasForMgmt returns a boolean if a field has been set.

### GetForSite

`func (o *WxlanTunnel) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *WxlanTunnel) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *WxlanTunnel) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *WxlanTunnel) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetHelloInterval

`func (o *WxlanTunnel) GetHelloInterval() int32`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *WxlanTunnel) GetHelloIntervalOk() (*int32, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *WxlanTunnel) SetHelloInterval(v int32)`

SetHelloInterval sets HelloInterval field to given value.

### HasHelloInterval

`func (o *WxlanTunnel) HasHelloInterval() bool`

HasHelloInterval returns a boolean if a field has been set.

### GetHelloRetries

`func (o *WxlanTunnel) GetHelloRetries() int32`

GetHelloRetries returns the HelloRetries field if non-nil, zero value otherwise.

### GetHelloRetriesOk

`func (o *WxlanTunnel) GetHelloRetriesOk() (*int32, bool)`

GetHelloRetriesOk returns a tuple with the HelloRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloRetries

`func (o *WxlanTunnel) SetHelloRetries(v int32)`

SetHelloRetries sets HelloRetries field to given value.

### HasHelloRetries

`func (o *WxlanTunnel) HasHelloRetries() bool`

HasHelloRetries returns a boolean if a field has been set.

### GetHostname

`func (o *WxlanTunnel) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *WxlanTunnel) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *WxlanTunnel) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *WxlanTunnel) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *WxlanTunnel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WxlanTunnel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WxlanTunnel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WxlanTunnel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpsec

`func (o *WxlanTunnel) GetIpsec() WxlanTunnelIpsec`

GetIpsec returns the Ipsec field if non-nil, zero value otherwise.

### GetIpsecOk

`func (o *WxlanTunnel) GetIpsecOk() (*WxlanTunnelIpsec, bool)`

GetIpsecOk returns a tuple with the Ipsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsec

`func (o *WxlanTunnel) SetIpsec(v WxlanTunnelIpsec)`

SetIpsec sets Ipsec field to given value.

### HasIpsec

`func (o *WxlanTunnel) HasIpsec() bool`

HasIpsec returns a boolean if a field has been set.

### GetIsStatic

`func (o *WxlanTunnel) GetIsStatic() bool`

GetIsStatic returns the IsStatic field if non-nil, zero value otherwise.

### GetIsStaticOk

`func (o *WxlanTunnel) GetIsStaticOk() (*bool, bool)`

GetIsStaticOk returns a tuple with the IsStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStatic

`func (o *WxlanTunnel) SetIsStatic(v bool)`

SetIsStatic sets IsStatic field to given value.

### HasIsStatic

`func (o *WxlanTunnel) HasIsStatic() bool`

HasIsStatic returns a boolean if a field has been set.

### GetModifiedTime

`func (o *WxlanTunnel) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *WxlanTunnel) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *WxlanTunnel) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *WxlanTunnel) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetMtu

`func (o *WxlanTunnel) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *WxlanTunnel) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *WxlanTunnel) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *WxlanTunnel) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *WxlanTunnel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WxlanTunnel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WxlanTunnel) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *WxlanTunnel) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WxlanTunnel) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WxlanTunnel) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *WxlanTunnel) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPeers

`func (o *WxlanTunnel) GetPeers() []string`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *WxlanTunnel) GetPeersOk() (*[]string, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *WxlanTunnel) SetPeers(v []string)`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *WxlanTunnel) HasPeers() bool`

HasPeers returns a boolean if a field has been set.

### GetRouterId

`func (o *WxlanTunnel) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *WxlanTunnel) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *WxlanTunnel) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *WxlanTunnel) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetSecret

`func (o *WxlanTunnel) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WxlanTunnel) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WxlanTunnel) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *WxlanTunnel) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSessions

`func (o *WxlanTunnel) GetSessions() []WxlanTunnelSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *WxlanTunnel) GetSessionsOk() (*[]WxlanTunnelSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *WxlanTunnel) SetSessions(v []WxlanTunnelSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *WxlanTunnel) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetSiteId

`func (o *WxlanTunnel) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WxlanTunnel) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WxlanTunnel) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *WxlanTunnel) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetUdpPort

`func (o *WxlanTunnel) GetUdpPort() int32`

GetUdpPort returns the UdpPort field if non-nil, zero value otherwise.

### GetUdpPortOk

`func (o *WxlanTunnel) GetUdpPortOk() (*int32, bool)`

GetUdpPortOk returns a tuple with the UdpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpPort

`func (o *WxlanTunnel) SetUdpPort(v int32)`

SetUdpPort sets UdpPort field to given value.

### HasUdpPort

`func (o *WxlanTunnel) HasUdpPort() bool`

HasUdpPort returns a boolean if a field has been set.

### GetUseUdp

`func (o *WxlanTunnel) GetUseUdp() bool`

GetUseUdp returns the UseUdp field if non-nil, zero value otherwise.

### GetUseUdpOk

`func (o *WxlanTunnel) GetUseUdpOk() (*bool, bool)`

GetUseUdpOk returns a tuple with the UseUdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUdp

`func (o *WxlanTunnel) SetUseUdp(v bool)`

SetUseUdp sets UseUdp field to given value.

### HasUseUdp

`func (o *WxlanTunnel) HasUseUdp() bool`

HasUseUdp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


