# StatsWiredClient

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

### NewStatsWiredClient

`func NewStatsWiredClient(mac string, ) *StatsWiredClient`

NewStatsWiredClient instantiates a new StatsWiredClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsWiredClientWithDefaults

`func NewStatsWiredClientWithDefaults() *StatsWiredClient`

NewStatsWiredClientWithDefaults instantiates a new StatsWiredClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StatsWiredClient) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StatsWiredClient) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StatsWiredClient) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StatsWiredClient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTtl

`func (o *StatsWiredClient) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *StatsWiredClient) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *StatsWiredClient) SetTtl(v float32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *StatsWiredClient) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetAuthState

`func (o *StatsWiredClient) GetAuthState() string`

GetAuthState returns the AuthState field if non-nil, zero value otherwise.

### GetAuthStateOk

`func (o *StatsWiredClient) GetAuthStateOk() (*string, bool)`

GetAuthStateOk returns a tuple with the AuthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthState

`func (o *StatsWiredClient) SetAuthState(v string)`

SetAuthState sets AuthState field to given value.

### HasAuthState

`func (o *StatsWiredClient) HasAuthState() bool`

HasAuthState returns a boolean if a field has been set.

### GetDeviceId

`func (o *StatsWiredClient) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *StatsWiredClient) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *StatsWiredClient) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *StatsWiredClient) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetEthPort

`func (o *StatsWiredClient) GetEthPort() string`

GetEthPort returns the EthPort field if non-nil, zero value otherwise.

### GetEthPortOk

`func (o *StatsWiredClient) GetEthPortOk() (*string, bool)`

GetEthPortOk returns a tuple with the EthPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthPort

`func (o *StatsWiredClient) SetEthPort(v string)`

SetEthPort sets EthPort field to given value.

### HasEthPort

`func (o *StatsWiredClient) HasEthPort() bool`

HasEthPort returns a boolean if a field has been set.

### GetLastSeen

`func (o *StatsWiredClient) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *StatsWiredClient) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *StatsWiredClient) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *StatsWiredClient) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMac

`func (o *StatsWiredClient) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *StatsWiredClient) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *StatsWiredClient) SetMac(v string)`

SetMac sets Mac field to given value.


### GetRxBytes

`func (o *StatsWiredClient) GetRxBytes() float32`

GetRxBytes returns the RxBytes field if non-nil, zero value otherwise.

### GetRxBytesOk

`func (o *StatsWiredClient) GetRxBytesOk() (*float32, bool)`

GetRxBytesOk returns a tuple with the RxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBytes

`func (o *StatsWiredClient) SetRxBytes(v float32)`

SetRxBytes sets RxBytes field to given value.

### HasRxBytes

`func (o *StatsWiredClient) HasRxBytes() bool`

HasRxBytes returns a boolean if a field has been set.

### GetRxPkts

`func (o *StatsWiredClient) GetRxPkts() float32`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *StatsWiredClient) GetRxPktsOk() (*float32, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *StatsWiredClient) SetRxPkts(v float32)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *StatsWiredClient) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetSiteId

`func (o *StatsWiredClient) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *StatsWiredClient) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *StatsWiredClient) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *StatsWiredClient) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTxBytes

`func (o *StatsWiredClient) GetTxBytes() float32`

GetTxBytes returns the TxBytes field if non-nil, zero value otherwise.

### GetTxBytesOk

`func (o *StatsWiredClient) GetTxBytesOk() (*float32, bool)`

GetTxBytesOk returns a tuple with the TxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBytes

`func (o *StatsWiredClient) SetTxBytes(v float32)`

SetTxBytes sets TxBytes field to given value.

### HasTxBytes

`func (o *StatsWiredClient) HasTxBytes() bool`

HasTxBytes returns a boolean if a field has been set.

### GetTxPkts

`func (o *StatsWiredClient) GetTxPkts() float32`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *StatsWiredClient) GetTxPktsOk() (*float32, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *StatsWiredClient) SetTxPkts(v float32)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *StatsWiredClient) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### GetUptime

`func (o *StatsWiredClient) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *StatsWiredClient) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *StatsWiredClient) SetUptime(v float32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *StatsWiredClient) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVlanId

`func (o *StatsWiredClient) GetVlanId() float32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *StatsWiredClient) GetVlanIdOk() (*float32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *StatsWiredClient) SetVlanId(v float32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *StatsWiredClient) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


