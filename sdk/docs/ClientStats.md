# ClientStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **float32** | TTL of the validity of the stat | [optional] 
**AuthState** | Pointer to **string** | client authorization status | [optional] 
**DeviceId** | Pointer to **string** | Device ID the client is connected to | [optional] 
**EthPort** | Pointer to **string** | port on AP where the wired client is connected | [optional] 
**LastSeen** | Pointer to **float32** | time when last Tx/Rx observed | [optional] 
**Mac** | **string** | client mac | 
**RxBytes** | Pointer to **float32** | amount of traffic sent to client since client connects | [optional] 
**RxPkts** | Pointer to **float32** | amount of traffic sent to client since client connects | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**TxBytes** | Pointer to **float32** |  amount of traffic received from client since client connects | [optional] 
**TxPkts** | Pointer to **float32** | amount of traffic received from client since client connects | [optional] 
**Uptime** | Pointer to **float32** | how long, in seconds, has the client been connected | [optional] 
**VlanId** | Pointer to **float32** | vlan id, could be empty | [optional] 

## Methods

### NewClientStats

`func NewClientStats(mac string, ) *ClientStats`

NewClientStats instantiates a new ClientStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientStatsWithDefaults

`func NewClientStatsWithDefaults() *ClientStats`

NewClientStatsWithDefaults instantiates a new ClientStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientStats) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientStats) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientStats) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientStats) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTtl

`func (o *ClientStats) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ClientStats) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ClientStats) SetTtl(v float32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ClientStats) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetAuthState

`func (o *ClientStats) GetAuthState() string`

GetAuthState returns the AuthState field if non-nil, zero value otherwise.

### GetAuthStateOk

`func (o *ClientStats) GetAuthStateOk() (*string, bool)`

GetAuthStateOk returns a tuple with the AuthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthState

`func (o *ClientStats) SetAuthState(v string)`

SetAuthState sets AuthState field to given value.

### HasAuthState

`func (o *ClientStats) HasAuthState() bool`

HasAuthState returns a boolean if a field has been set.

### GetDeviceId

`func (o *ClientStats) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ClientStats) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ClientStats) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ClientStats) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetEthPort

`func (o *ClientStats) GetEthPort() string`

GetEthPort returns the EthPort field if non-nil, zero value otherwise.

### GetEthPortOk

`func (o *ClientStats) GetEthPortOk() (*string, bool)`

GetEthPortOk returns a tuple with the EthPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthPort

`func (o *ClientStats) SetEthPort(v string)`

SetEthPort sets EthPort field to given value.

### HasEthPort

`func (o *ClientStats) HasEthPort() bool`

HasEthPort returns a boolean if a field has been set.

### GetLastSeen

`func (o *ClientStats) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ClientStats) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ClientStats) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ClientStats) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMac

`func (o *ClientStats) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ClientStats) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ClientStats) SetMac(v string)`

SetMac sets Mac field to given value.


### GetRxBytes

`func (o *ClientStats) GetRxBytes() float32`

GetRxBytes returns the RxBytes field if non-nil, zero value otherwise.

### GetRxBytesOk

`func (o *ClientStats) GetRxBytesOk() (*float32, bool)`

GetRxBytesOk returns a tuple with the RxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBytes

`func (o *ClientStats) SetRxBytes(v float32)`

SetRxBytes sets RxBytes field to given value.

### HasRxBytes

`func (o *ClientStats) HasRxBytes() bool`

HasRxBytes returns a boolean if a field has been set.

### GetRxPkts

`func (o *ClientStats) GetRxPkts() float32`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *ClientStats) GetRxPktsOk() (*float32, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *ClientStats) SetRxPkts(v float32)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *ClientStats) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetSiteId

`func (o *ClientStats) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ClientStats) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ClientStats) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ClientStats) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTxBytes

`func (o *ClientStats) GetTxBytes() float32`

GetTxBytes returns the TxBytes field if non-nil, zero value otherwise.

### GetTxBytesOk

`func (o *ClientStats) GetTxBytesOk() (*float32, bool)`

GetTxBytesOk returns a tuple with the TxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBytes

`func (o *ClientStats) SetTxBytes(v float32)`

SetTxBytes sets TxBytes field to given value.

### HasTxBytes

`func (o *ClientStats) HasTxBytes() bool`

HasTxBytes returns a boolean if a field has been set.

### GetTxPkts

`func (o *ClientStats) GetTxPkts() float32`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *ClientStats) GetTxPktsOk() (*float32, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *ClientStats) SetTxPkts(v float32)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *ClientStats) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### GetUptime

`func (o *ClientStats) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ClientStats) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ClientStats) SetUptime(v float32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *ClientStats) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVlanId

`func (o *ClientStats) GetVlanId() float32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *ClientStats) GetVlanIdOk() (*float32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *ClientStats) SetVlanId(v float32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *ClientStats) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


