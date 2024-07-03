# WanTunnelStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthAlgo** | Pointer to **string** | authentication algorithm | [optional] 
**EncryptAlgo** | Pointer to **string** | encryption algorithm | [optional] 
**IkeVersion** | Pointer to **string** | ike version | [optional] 
**Ip** | Pointer to **string** | ip address | [optional] 
**LastEvent** | Pointer to **string** | reason of why the tunnel is down | [optional] 
**Mac** | Pointer to **string** | router mac address | [optional] 
**Node** | Pointer to **string** | node0/node1 | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PeerHost** | Pointer to **string** | peer host | [optional] 
**PeerIp** | Pointer to **string** | peer ip address | [optional] 
**Priority** | Pointer to [**WanTunnelStatsPriority**](WanTunnelStatsPriority.md) |  | [optional] 
**Protocol** | Pointer to [**WanTunnelProtocol**](WanTunnelProtocol.md) |  | [optional] 
**RxBytes** | Pointer to **int32** |  | [optional] 
**RxPkts** | Pointer to **int32** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**TunnelName** | Pointer to **string** | Mist Tunnel Name | [optional] 
**TxBytes** | Pointer to **int32** |  | [optional] 
**TxPkts** | Pointer to **int32** |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 
**Uptime** | Pointer to **int32** | duration from first (or last) SA was established | [optional] 
**WanName** | Pointer to **string** | wan interface name | [optional] 

## Methods

### NewWanTunnelStats

`func NewWanTunnelStats() *WanTunnelStats`

NewWanTunnelStats instantiates a new WanTunnelStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanTunnelStatsWithDefaults

`func NewWanTunnelStatsWithDefaults() *WanTunnelStats`

NewWanTunnelStatsWithDefaults instantiates a new WanTunnelStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthAlgo

`func (o *WanTunnelStats) GetAuthAlgo() string`

GetAuthAlgo returns the AuthAlgo field if non-nil, zero value otherwise.

### GetAuthAlgoOk

`func (o *WanTunnelStats) GetAuthAlgoOk() (*string, bool)`

GetAuthAlgoOk returns a tuple with the AuthAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthAlgo

`func (o *WanTunnelStats) SetAuthAlgo(v string)`

SetAuthAlgo sets AuthAlgo field to given value.

### HasAuthAlgo

`func (o *WanTunnelStats) HasAuthAlgo() bool`

HasAuthAlgo returns a boolean if a field has been set.

### GetEncryptAlgo

`func (o *WanTunnelStats) GetEncryptAlgo() string`

GetEncryptAlgo returns the EncryptAlgo field if non-nil, zero value otherwise.

### GetEncryptAlgoOk

`func (o *WanTunnelStats) GetEncryptAlgoOk() (*string, bool)`

GetEncryptAlgoOk returns a tuple with the EncryptAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAlgo

`func (o *WanTunnelStats) SetEncryptAlgo(v string)`

SetEncryptAlgo sets EncryptAlgo field to given value.

### HasEncryptAlgo

`func (o *WanTunnelStats) HasEncryptAlgo() bool`

HasEncryptAlgo returns a boolean if a field has been set.

### GetIkeVersion

`func (o *WanTunnelStats) GetIkeVersion() string`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *WanTunnelStats) GetIkeVersionOk() (*string, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *WanTunnelStats) SetIkeVersion(v string)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *WanTunnelStats) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetIp

`func (o *WanTunnelStats) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *WanTunnelStats) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *WanTunnelStats) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *WanTunnelStats) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLastEvent

`func (o *WanTunnelStats) GetLastEvent() string`

GetLastEvent returns the LastEvent field if non-nil, zero value otherwise.

### GetLastEventOk

`func (o *WanTunnelStats) GetLastEventOk() (*string, bool)`

GetLastEventOk returns a tuple with the LastEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEvent

`func (o *WanTunnelStats) SetLastEvent(v string)`

SetLastEvent sets LastEvent field to given value.

### HasLastEvent

`func (o *WanTunnelStats) HasLastEvent() bool`

HasLastEvent returns a boolean if a field has been set.

### GetMac

`func (o *WanTunnelStats) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WanTunnelStats) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WanTunnelStats) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *WanTunnelStats) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNode

`func (o *WanTunnelStats) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *WanTunnelStats) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *WanTunnelStats) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *WanTunnelStats) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetOrgId

`func (o *WanTunnelStats) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WanTunnelStats) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WanTunnelStats) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *WanTunnelStats) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPeerHost

`func (o *WanTunnelStats) GetPeerHost() string`

GetPeerHost returns the PeerHost field if non-nil, zero value otherwise.

### GetPeerHostOk

`func (o *WanTunnelStats) GetPeerHostOk() (*string, bool)`

GetPeerHostOk returns a tuple with the PeerHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerHost

`func (o *WanTunnelStats) SetPeerHost(v string)`

SetPeerHost sets PeerHost field to given value.

### HasPeerHost

`func (o *WanTunnelStats) HasPeerHost() bool`

HasPeerHost returns a boolean if a field has been set.

### GetPeerIp

`func (o *WanTunnelStats) GetPeerIp() string`

GetPeerIp returns the PeerIp field if non-nil, zero value otherwise.

### GetPeerIpOk

`func (o *WanTunnelStats) GetPeerIpOk() (*string, bool)`

GetPeerIpOk returns a tuple with the PeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIp

`func (o *WanTunnelStats) SetPeerIp(v string)`

SetPeerIp sets PeerIp field to given value.

### HasPeerIp

`func (o *WanTunnelStats) HasPeerIp() bool`

HasPeerIp returns a boolean if a field has been set.

### GetPriority

`func (o *WanTunnelStats) GetPriority() WanTunnelStatsPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WanTunnelStats) GetPriorityOk() (*WanTunnelStatsPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WanTunnelStats) SetPriority(v WanTunnelStatsPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *WanTunnelStats) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProtocol

`func (o *WanTunnelStats) GetProtocol() WanTunnelProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *WanTunnelStats) GetProtocolOk() (*WanTunnelProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *WanTunnelStats) SetProtocol(v WanTunnelProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *WanTunnelStats) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRxBytes

`func (o *WanTunnelStats) GetRxBytes() int32`

GetRxBytes returns the RxBytes field if non-nil, zero value otherwise.

### GetRxBytesOk

`func (o *WanTunnelStats) GetRxBytesOk() (*int32, bool)`

GetRxBytesOk returns a tuple with the RxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBytes

`func (o *WanTunnelStats) SetRxBytes(v int32)`

SetRxBytes sets RxBytes field to given value.

### HasRxBytes

`func (o *WanTunnelStats) HasRxBytes() bool`

HasRxBytes returns a boolean if a field has been set.

### GetRxPkts

`func (o *WanTunnelStats) GetRxPkts() int32`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *WanTunnelStats) GetRxPktsOk() (*int32, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *WanTunnelStats) SetRxPkts(v int32)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *WanTunnelStats) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetSiteId

`func (o *WanTunnelStats) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WanTunnelStats) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WanTunnelStats) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *WanTunnelStats) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTunnelName

`func (o *WanTunnelStats) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *WanTunnelStats) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *WanTunnelStats) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.

### HasTunnelName

`func (o *WanTunnelStats) HasTunnelName() bool`

HasTunnelName returns a boolean if a field has been set.

### GetTxBytes

`func (o *WanTunnelStats) GetTxBytes() int32`

GetTxBytes returns the TxBytes field if non-nil, zero value otherwise.

### GetTxBytesOk

`func (o *WanTunnelStats) GetTxBytesOk() (*int32, bool)`

GetTxBytesOk returns a tuple with the TxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBytes

`func (o *WanTunnelStats) SetTxBytes(v int32)`

SetTxBytes sets TxBytes field to given value.

### HasTxBytes

`func (o *WanTunnelStats) HasTxBytes() bool`

HasTxBytes returns a boolean if a field has been set.

### GetTxPkts

`func (o *WanTunnelStats) GetTxPkts() int32`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *WanTunnelStats) GetTxPktsOk() (*int32, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *WanTunnelStats) SetTxPkts(v int32)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *WanTunnelStats) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### GetUp

`func (o *WanTunnelStats) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *WanTunnelStats) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *WanTunnelStats) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *WanTunnelStats) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUptime

`func (o *WanTunnelStats) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *WanTunnelStats) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *WanTunnelStats) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *WanTunnelStats) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetWanName

`func (o *WanTunnelStats) GetWanName() string`

GetWanName returns the WanName field if non-nil, zero value otherwise.

### GetWanNameOk

`func (o *WanTunnelStats) GetWanNameOk() (*string, bool)`

GetWanNameOk returns a tuple with the WanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanName

`func (o *WanTunnelStats) SetWanName(v string)`

SetWanName sets WanName field to given value.

### HasWanName

`func (o *WanTunnelStats) HasWanName() bool`

HasWanName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


