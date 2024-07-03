# ResponseTunnelSearchItem

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
**Uptime** | Pointer to **int32** | duration from first (or last) SA was established | [optional] 
**AuthAlgo** | Pointer to **string** | authentication algorithm | [optional] 
**EncryptAlgo** | Pointer to **string** | encryption algorithm | [optional] 
**IkeVersion** | Pointer to **string** | ike version | [optional] 
**Ip** | Pointer to **string** | ip address | [optional] 
**LastEvent** | Pointer to **string** | reason of why the tunnel is down | [optional] 
**Mac** | Pointer to **string** | router mac address | [optional] 
**Node** | Pointer to **string** | node0/node1 | [optional] 
**PeerHost** | Pointer to **string** | peer host | [optional] 
**PeerIp** | Pointer to **string** | peer ip address | [optional] 
**Priority** | Pointer to [**WanTunnelStatsPriority**](WanTunnelStatsPriority.md) |  | [optional] 
**Protocol** | Pointer to [**WanTunnelProtocol**](WanTunnelProtocol.md) |  | [optional] 
**RxBytes** | Pointer to **int32** |  | [optional] 
**RxPkts** | Pointer to **int32** |  | [optional] 
**TunnelName** | Pointer to **string** | Mist Tunnel Name | [optional] 
**TxBytes** | Pointer to **int32** |  | [optional] 
**TxPkts** | Pointer to **int32** |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 
**WanName** | Pointer to **string** | wan interface name | [optional] 

## Methods

### NewResponseTunnelSearchItem

`func NewResponseTunnelSearchItem() *ResponseTunnelSearchItem`

NewResponseTunnelSearchItem instantiates a new ResponseTunnelSearchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTunnelSearchItemWithDefaults

`func NewResponseTunnelSearchItemWithDefaults() *ResponseTunnelSearchItem`

NewResponseTunnelSearchItemWithDefaults instantiates a new ResponseTunnelSearchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *ResponseTunnelSearchItem) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *ResponseTunnelSearchItem) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *ResponseTunnelSearchItem) SetAp(v string)`

SetAp sets Ap field to given value.

### HasAp

`func (o *ResponseTunnelSearchItem) HasAp() bool`

HasAp returns a boolean if a field has been set.

### GetForSite

`func (o *ResponseTunnelSearchItem) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *ResponseTunnelSearchItem) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *ResponseTunnelSearchItem) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *ResponseTunnelSearchItem) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetLastSeen

`func (o *ResponseTunnelSearchItem) GetLastSeen() int32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ResponseTunnelSearchItem) GetLastSeenOk() (*int32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ResponseTunnelSearchItem) SetLastSeen(v int32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ResponseTunnelSearchItem) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMxclusterId

`func (o *ResponseTunnelSearchItem) GetMxclusterId() string`

GetMxclusterId returns the MxclusterId field if non-nil, zero value otherwise.

### GetMxclusterIdOk

`func (o *ResponseTunnelSearchItem) GetMxclusterIdOk() (*string, bool)`

GetMxclusterIdOk returns a tuple with the MxclusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxclusterId

`func (o *ResponseTunnelSearchItem) SetMxclusterId(v string)`

SetMxclusterId sets MxclusterId field to given value.

### HasMxclusterId

`func (o *ResponseTunnelSearchItem) HasMxclusterId() bool`

HasMxclusterId returns a boolean if a field has been set.

### GetMxedgeId

`func (o *ResponseTunnelSearchItem) GetMxedgeId() string`

GetMxedgeId returns the MxedgeId field if non-nil, zero value otherwise.

### GetMxedgeIdOk

`func (o *ResponseTunnelSearchItem) GetMxedgeIdOk() (*string, bool)`

GetMxedgeIdOk returns a tuple with the MxedgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxedgeId

`func (o *ResponseTunnelSearchItem) SetMxedgeId(v string)`

SetMxedgeId sets MxedgeId field to given value.

### HasMxedgeId

`func (o *ResponseTunnelSearchItem) HasMxedgeId() bool`

HasMxedgeId returns a boolean if a field has been set.

### GetMxtunnelId

`func (o *ResponseTunnelSearchItem) GetMxtunnelId() string`

GetMxtunnelId returns the MxtunnelId field if non-nil, zero value otherwise.

### GetMxtunnelIdOk

`func (o *ResponseTunnelSearchItem) GetMxtunnelIdOk() (*string, bool)`

GetMxtunnelIdOk returns a tuple with the MxtunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxtunnelId

`func (o *ResponseTunnelSearchItem) SetMxtunnelId(v string)`

SetMxtunnelId sets MxtunnelId field to given value.

### HasMxtunnelId

`func (o *ResponseTunnelSearchItem) HasMxtunnelId() bool`

HasMxtunnelId returns a boolean if a field has been set.

### GetOrgId

`func (o *ResponseTunnelSearchItem) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ResponseTunnelSearchItem) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ResponseTunnelSearchItem) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ResponseTunnelSearchItem) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPeerMxedgeId

`func (o *ResponseTunnelSearchItem) GetPeerMxedgeId() string`

GetPeerMxedgeId returns the PeerMxedgeId field if non-nil, zero value otherwise.

### GetPeerMxedgeIdOk

`func (o *ResponseTunnelSearchItem) GetPeerMxedgeIdOk() (*string, bool)`

GetPeerMxedgeIdOk returns a tuple with the PeerMxedgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerMxedgeId

`func (o *ResponseTunnelSearchItem) SetPeerMxedgeId(v string)`

SetPeerMxedgeId sets PeerMxedgeId field to given value.

### HasPeerMxedgeId

`func (o *ResponseTunnelSearchItem) HasPeerMxedgeId() bool`

HasPeerMxedgeId returns a boolean if a field has been set.

### GetRemoteIp

`func (o *ResponseTunnelSearchItem) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *ResponseTunnelSearchItem) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *ResponseTunnelSearchItem) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *ResponseTunnelSearchItem) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetRemotePort

`func (o *ResponseTunnelSearchItem) GetRemotePort() int32`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *ResponseTunnelSearchItem) GetRemotePortOk() (*int32, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *ResponseTunnelSearchItem) SetRemotePort(v int32)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *ResponseTunnelSearchItem) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetRxControlPkts

`func (o *ResponseTunnelSearchItem) GetRxControlPkts() int32`

GetRxControlPkts returns the RxControlPkts field if non-nil, zero value otherwise.

### GetRxControlPktsOk

`func (o *ResponseTunnelSearchItem) GetRxControlPktsOk() (*int32, bool)`

GetRxControlPktsOk returns a tuple with the RxControlPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxControlPkts

`func (o *ResponseTunnelSearchItem) SetRxControlPkts(v int32)`

SetRxControlPkts sets RxControlPkts field to given value.

### HasRxControlPkts

`func (o *ResponseTunnelSearchItem) HasRxControlPkts() bool`

HasRxControlPkts returns a boolean if a field has been set.

### GetSessions

`func (o *ResponseTunnelSearchItem) GetSessions() []MxtunnelStatsSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *ResponseTunnelSearchItem) GetSessionsOk() (*[]MxtunnelStatsSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *ResponseTunnelSearchItem) SetSessions(v []MxtunnelStatsSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *ResponseTunnelSearchItem) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetSiteId

`func (o *ResponseTunnelSearchItem) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ResponseTunnelSearchItem) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ResponseTunnelSearchItem) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ResponseTunnelSearchItem) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetState

`func (o *ResponseTunnelSearchItem) GetState() MxtunnelStatsState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseTunnelSearchItem) GetStateOk() (*MxtunnelStatsState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseTunnelSearchItem) SetState(v MxtunnelStatsState)`

SetState sets State field to given value.

### HasState

`func (o *ResponseTunnelSearchItem) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTxControlPkts

`func (o *ResponseTunnelSearchItem) GetTxControlPkts() int32`

GetTxControlPkts returns the TxControlPkts field if non-nil, zero value otherwise.

### GetTxControlPktsOk

`func (o *ResponseTunnelSearchItem) GetTxControlPktsOk() (*int32, bool)`

GetTxControlPktsOk returns a tuple with the TxControlPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxControlPkts

`func (o *ResponseTunnelSearchItem) SetTxControlPkts(v int32)`

SetTxControlPkts sets TxControlPkts field to given value.

### HasTxControlPkts

`func (o *ResponseTunnelSearchItem) HasTxControlPkts() bool`

HasTxControlPkts returns a boolean if a field has been set.

### GetUptime

`func (o *ResponseTunnelSearchItem) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ResponseTunnelSearchItem) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ResponseTunnelSearchItem) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *ResponseTunnelSearchItem) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetAuthAlgo

`func (o *ResponseTunnelSearchItem) GetAuthAlgo() string`

GetAuthAlgo returns the AuthAlgo field if non-nil, zero value otherwise.

### GetAuthAlgoOk

`func (o *ResponseTunnelSearchItem) GetAuthAlgoOk() (*string, bool)`

GetAuthAlgoOk returns a tuple with the AuthAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthAlgo

`func (o *ResponseTunnelSearchItem) SetAuthAlgo(v string)`

SetAuthAlgo sets AuthAlgo field to given value.

### HasAuthAlgo

`func (o *ResponseTunnelSearchItem) HasAuthAlgo() bool`

HasAuthAlgo returns a boolean if a field has been set.

### GetEncryptAlgo

`func (o *ResponseTunnelSearchItem) GetEncryptAlgo() string`

GetEncryptAlgo returns the EncryptAlgo field if non-nil, zero value otherwise.

### GetEncryptAlgoOk

`func (o *ResponseTunnelSearchItem) GetEncryptAlgoOk() (*string, bool)`

GetEncryptAlgoOk returns a tuple with the EncryptAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAlgo

`func (o *ResponseTunnelSearchItem) SetEncryptAlgo(v string)`

SetEncryptAlgo sets EncryptAlgo field to given value.

### HasEncryptAlgo

`func (o *ResponseTunnelSearchItem) HasEncryptAlgo() bool`

HasEncryptAlgo returns a boolean if a field has been set.

### GetIkeVersion

`func (o *ResponseTunnelSearchItem) GetIkeVersion() string`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *ResponseTunnelSearchItem) GetIkeVersionOk() (*string, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *ResponseTunnelSearchItem) SetIkeVersion(v string)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *ResponseTunnelSearchItem) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetIp

`func (o *ResponseTunnelSearchItem) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ResponseTunnelSearchItem) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ResponseTunnelSearchItem) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ResponseTunnelSearchItem) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLastEvent

`func (o *ResponseTunnelSearchItem) GetLastEvent() string`

GetLastEvent returns the LastEvent field if non-nil, zero value otherwise.

### GetLastEventOk

`func (o *ResponseTunnelSearchItem) GetLastEventOk() (*string, bool)`

GetLastEventOk returns a tuple with the LastEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEvent

`func (o *ResponseTunnelSearchItem) SetLastEvent(v string)`

SetLastEvent sets LastEvent field to given value.

### HasLastEvent

`func (o *ResponseTunnelSearchItem) HasLastEvent() bool`

HasLastEvent returns a boolean if a field has been set.

### GetMac

`func (o *ResponseTunnelSearchItem) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ResponseTunnelSearchItem) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ResponseTunnelSearchItem) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ResponseTunnelSearchItem) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNode

`func (o *ResponseTunnelSearchItem) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *ResponseTunnelSearchItem) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *ResponseTunnelSearchItem) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *ResponseTunnelSearchItem) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetPeerHost

`func (o *ResponseTunnelSearchItem) GetPeerHost() string`

GetPeerHost returns the PeerHost field if non-nil, zero value otherwise.

### GetPeerHostOk

`func (o *ResponseTunnelSearchItem) GetPeerHostOk() (*string, bool)`

GetPeerHostOk returns a tuple with the PeerHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerHost

`func (o *ResponseTunnelSearchItem) SetPeerHost(v string)`

SetPeerHost sets PeerHost field to given value.

### HasPeerHost

`func (o *ResponseTunnelSearchItem) HasPeerHost() bool`

HasPeerHost returns a boolean if a field has been set.

### GetPeerIp

`func (o *ResponseTunnelSearchItem) GetPeerIp() string`

GetPeerIp returns the PeerIp field if non-nil, zero value otherwise.

### GetPeerIpOk

`func (o *ResponseTunnelSearchItem) GetPeerIpOk() (*string, bool)`

GetPeerIpOk returns a tuple with the PeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIp

`func (o *ResponseTunnelSearchItem) SetPeerIp(v string)`

SetPeerIp sets PeerIp field to given value.

### HasPeerIp

`func (o *ResponseTunnelSearchItem) HasPeerIp() bool`

HasPeerIp returns a boolean if a field has been set.

### GetPriority

`func (o *ResponseTunnelSearchItem) GetPriority() WanTunnelStatsPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ResponseTunnelSearchItem) GetPriorityOk() (*WanTunnelStatsPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ResponseTunnelSearchItem) SetPriority(v WanTunnelStatsPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ResponseTunnelSearchItem) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProtocol

`func (o *ResponseTunnelSearchItem) GetProtocol() WanTunnelProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ResponseTunnelSearchItem) GetProtocolOk() (*WanTunnelProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ResponseTunnelSearchItem) SetProtocol(v WanTunnelProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ResponseTunnelSearchItem) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRxBytes

`func (o *ResponseTunnelSearchItem) GetRxBytes() int32`

GetRxBytes returns the RxBytes field if non-nil, zero value otherwise.

### GetRxBytesOk

`func (o *ResponseTunnelSearchItem) GetRxBytesOk() (*int32, bool)`

GetRxBytesOk returns a tuple with the RxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBytes

`func (o *ResponseTunnelSearchItem) SetRxBytes(v int32)`

SetRxBytes sets RxBytes field to given value.

### HasRxBytes

`func (o *ResponseTunnelSearchItem) HasRxBytes() bool`

HasRxBytes returns a boolean if a field has been set.

### GetRxPkts

`func (o *ResponseTunnelSearchItem) GetRxPkts() int32`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *ResponseTunnelSearchItem) GetRxPktsOk() (*int32, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *ResponseTunnelSearchItem) SetRxPkts(v int32)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *ResponseTunnelSearchItem) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetTunnelName

`func (o *ResponseTunnelSearchItem) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *ResponseTunnelSearchItem) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *ResponseTunnelSearchItem) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.

### HasTunnelName

`func (o *ResponseTunnelSearchItem) HasTunnelName() bool`

HasTunnelName returns a boolean if a field has been set.

### GetTxBytes

`func (o *ResponseTunnelSearchItem) GetTxBytes() int32`

GetTxBytes returns the TxBytes field if non-nil, zero value otherwise.

### GetTxBytesOk

`func (o *ResponseTunnelSearchItem) GetTxBytesOk() (*int32, bool)`

GetTxBytesOk returns a tuple with the TxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBytes

`func (o *ResponseTunnelSearchItem) SetTxBytes(v int32)`

SetTxBytes sets TxBytes field to given value.

### HasTxBytes

`func (o *ResponseTunnelSearchItem) HasTxBytes() bool`

HasTxBytes returns a boolean if a field has been set.

### GetTxPkts

`func (o *ResponseTunnelSearchItem) GetTxPkts() int32`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *ResponseTunnelSearchItem) GetTxPktsOk() (*int32, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *ResponseTunnelSearchItem) SetTxPkts(v int32)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *ResponseTunnelSearchItem) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### GetUp

`func (o *ResponseTunnelSearchItem) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *ResponseTunnelSearchItem) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *ResponseTunnelSearchItem) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *ResponseTunnelSearchItem) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetWanName

`func (o *ResponseTunnelSearchItem) GetWanName() string`

GetWanName returns the WanName field if non-nil, zero value otherwise.

### GetWanNameOk

`func (o *ResponseTunnelSearchItem) GetWanNameOk() (*string, bool)`

GetWanNameOk returns a tuple with the WanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanName

`func (o *ResponseTunnelSearchItem) SetWanName(v string)`

SetWanName sets WanName field to given value.

### HasWanName

`func (o *ResponseTunnelSearchItem) HasWanName() bool`

HasWanName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


