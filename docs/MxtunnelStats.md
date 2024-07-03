# MxtunnelStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | Pointer to **string** |  | [optional] [readonly] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**LastSeen** | Pointer to **int32** |  | [optional] [readonly] 
**MxclusterId** | Pointer to **string** |  | [optional] [readonly] 
**MxedgeId** | Pointer to **string** |  | [optional] [readonly] 
**MxtunnelId** | Pointer to **string** |  | [optional] [readonly] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PeerMxedgeId** | Pointer to **string** | MxEdge ID of the peer(mist edge to mist edge tunnel) | [optional] [readonly] 
**RemoteIp** | Pointer to **string** |  | [optional] [readonly] 
**RemotePort** | Pointer to **int32** |  | [optional] [readonly] 
**RxControlPkts** | Pointer to **int32** |  | [optional] [readonly] 
**Sessions** | Pointer to [**[]MxtunnelStatsSession**](MxtunnelStatsSession.md) | list of sessions | [optional] [readonly] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**State** | Pointer to [**MxtunnelStatsState**](MxtunnelStatsState.md) |  | [optional] 
**TxControlPkts** | Pointer to **int32** |  | [optional] [readonly] 
**Uptime** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewMxtunnelStats

`func NewMxtunnelStats() *MxtunnelStats`

NewMxtunnelStats instantiates a new MxtunnelStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxtunnelStatsWithDefaults

`func NewMxtunnelStatsWithDefaults() *MxtunnelStats`

NewMxtunnelStatsWithDefaults instantiates a new MxtunnelStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *MxtunnelStats) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *MxtunnelStats) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *MxtunnelStats) SetAp(v string)`

SetAp sets Ap field to given value.

### HasAp

`func (o *MxtunnelStats) HasAp() bool`

HasAp returns a boolean if a field has been set.

### GetForSite

`func (o *MxtunnelStats) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *MxtunnelStats) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *MxtunnelStats) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *MxtunnelStats) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetLastSeen

`func (o *MxtunnelStats) GetLastSeen() int32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *MxtunnelStats) GetLastSeenOk() (*int32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *MxtunnelStats) SetLastSeen(v int32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *MxtunnelStats) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMxclusterId

`func (o *MxtunnelStats) GetMxclusterId() string`

GetMxclusterId returns the MxclusterId field if non-nil, zero value otherwise.

### GetMxclusterIdOk

`func (o *MxtunnelStats) GetMxclusterIdOk() (*string, bool)`

GetMxclusterIdOk returns a tuple with the MxclusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxclusterId

`func (o *MxtunnelStats) SetMxclusterId(v string)`

SetMxclusterId sets MxclusterId field to given value.

### HasMxclusterId

`func (o *MxtunnelStats) HasMxclusterId() bool`

HasMxclusterId returns a boolean if a field has been set.

### GetMxedgeId

`func (o *MxtunnelStats) GetMxedgeId() string`

GetMxedgeId returns the MxedgeId field if non-nil, zero value otherwise.

### GetMxedgeIdOk

`func (o *MxtunnelStats) GetMxedgeIdOk() (*string, bool)`

GetMxedgeIdOk returns a tuple with the MxedgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxedgeId

`func (o *MxtunnelStats) SetMxedgeId(v string)`

SetMxedgeId sets MxedgeId field to given value.

### HasMxedgeId

`func (o *MxtunnelStats) HasMxedgeId() bool`

HasMxedgeId returns a boolean if a field has been set.

### GetMxtunnelId

`func (o *MxtunnelStats) GetMxtunnelId() string`

GetMxtunnelId returns the MxtunnelId field if non-nil, zero value otherwise.

### GetMxtunnelIdOk

`func (o *MxtunnelStats) GetMxtunnelIdOk() (*string, bool)`

GetMxtunnelIdOk returns a tuple with the MxtunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxtunnelId

`func (o *MxtunnelStats) SetMxtunnelId(v string)`

SetMxtunnelId sets MxtunnelId field to given value.

### HasMxtunnelId

`func (o *MxtunnelStats) HasMxtunnelId() bool`

HasMxtunnelId returns a boolean if a field has been set.

### GetOrgId

`func (o *MxtunnelStats) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *MxtunnelStats) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *MxtunnelStats) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *MxtunnelStats) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPeerMxedgeId

`func (o *MxtunnelStats) GetPeerMxedgeId() string`

GetPeerMxedgeId returns the PeerMxedgeId field if non-nil, zero value otherwise.

### GetPeerMxedgeIdOk

`func (o *MxtunnelStats) GetPeerMxedgeIdOk() (*string, bool)`

GetPeerMxedgeIdOk returns a tuple with the PeerMxedgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerMxedgeId

`func (o *MxtunnelStats) SetPeerMxedgeId(v string)`

SetPeerMxedgeId sets PeerMxedgeId field to given value.

### HasPeerMxedgeId

`func (o *MxtunnelStats) HasPeerMxedgeId() bool`

HasPeerMxedgeId returns a boolean if a field has been set.

### GetRemoteIp

`func (o *MxtunnelStats) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *MxtunnelStats) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *MxtunnelStats) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *MxtunnelStats) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetRemotePort

`func (o *MxtunnelStats) GetRemotePort() int32`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *MxtunnelStats) GetRemotePortOk() (*int32, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *MxtunnelStats) SetRemotePort(v int32)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *MxtunnelStats) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetRxControlPkts

`func (o *MxtunnelStats) GetRxControlPkts() int32`

GetRxControlPkts returns the RxControlPkts field if non-nil, zero value otherwise.

### GetRxControlPktsOk

`func (o *MxtunnelStats) GetRxControlPktsOk() (*int32, bool)`

GetRxControlPktsOk returns a tuple with the RxControlPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxControlPkts

`func (o *MxtunnelStats) SetRxControlPkts(v int32)`

SetRxControlPkts sets RxControlPkts field to given value.

### HasRxControlPkts

`func (o *MxtunnelStats) HasRxControlPkts() bool`

HasRxControlPkts returns a boolean if a field has been set.

### GetSessions

`func (o *MxtunnelStats) GetSessions() []MxtunnelStatsSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *MxtunnelStats) GetSessionsOk() (*[]MxtunnelStatsSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *MxtunnelStats) SetSessions(v []MxtunnelStatsSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *MxtunnelStats) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetSiteId

`func (o *MxtunnelStats) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *MxtunnelStats) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *MxtunnelStats) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *MxtunnelStats) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetState

`func (o *MxtunnelStats) GetState() MxtunnelStatsState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MxtunnelStats) GetStateOk() (*MxtunnelStatsState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MxtunnelStats) SetState(v MxtunnelStatsState)`

SetState sets State field to given value.

### HasState

`func (o *MxtunnelStats) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTxControlPkts

`func (o *MxtunnelStats) GetTxControlPkts() int32`

GetTxControlPkts returns the TxControlPkts field if non-nil, zero value otherwise.

### GetTxControlPktsOk

`func (o *MxtunnelStats) GetTxControlPktsOk() (*int32, bool)`

GetTxControlPktsOk returns a tuple with the TxControlPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxControlPkts

`func (o *MxtunnelStats) SetTxControlPkts(v int32)`

SetTxControlPkts sets TxControlPkts field to given value.

### HasTxControlPkts

`func (o *MxtunnelStats) HasTxControlPkts() bool`

HasTxControlPkts returns a boolean if a field has been set.

### GetUptime

`func (o *MxtunnelStats) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *MxtunnelStats) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *MxtunnelStats) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *MxtunnelStats) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


