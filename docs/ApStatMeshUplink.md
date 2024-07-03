# ApStatMeshUplink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Band** | Pointer to **string** |  | [optional] 
**Channel** | Pointer to **int32** |  | [optional] 
**IdleTime** | Pointer to **int32** |  | [optional] 
**LastSeen** | Pointer to **int32** |  | [optional] 
**Proto** | Pointer to **string** |  | [optional] 
**Rssi** | Pointer to **int32** |  | [optional] 
**RxBps** | Pointer to **int32** |  | [optional] 
**RxBytes** | Pointer to **int32** |  | [optional] 
**RxPackets** | Pointer to **int32** |  | [optional] 
**RxRate** | Pointer to **int32** |  | [optional] 
**RxRetries** | Pointer to **int32** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Snr** | Pointer to **int32** |  | [optional] 
**TxBps** | Pointer to **int32** |  | [optional] 
**TxBytes** | Pointer to **int32** |  | [optional] 
**TxPackets** | Pointer to **int32** |  | [optional] 
**TxRate** | Pointer to **int32** |  | [optional] 
**TxRetries** | Pointer to **int32** |  | [optional] 
**UplinkApId** | Pointer to **string** |  | [optional] 

## Methods

### NewApStatMeshUplink

`func NewApStatMeshUplink() *ApStatMeshUplink`

NewApStatMeshUplink instantiates a new ApStatMeshUplink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApStatMeshUplinkWithDefaults

`func NewApStatMeshUplinkWithDefaults() *ApStatMeshUplink`

NewApStatMeshUplinkWithDefaults instantiates a new ApStatMeshUplink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBand

`func (o *ApStatMeshUplink) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *ApStatMeshUplink) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *ApStatMeshUplink) SetBand(v string)`

SetBand sets Band field to given value.

### HasBand

`func (o *ApStatMeshUplink) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetChannel

`func (o *ApStatMeshUplink) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ApStatMeshUplink) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ApStatMeshUplink) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ApStatMeshUplink) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetIdleTime

`func (o *ApStatMeshUplink) GetIdleTime() int32`

GetIdleTime returns the IdleTime field if non-nil, zero value otherwise.

### GetIdleTimeOk

`func (o *ApStatMeshUplink) GetIdleTimeOk() (*int32, bool)`

GetIdleTimeOk returns a tuple with the IdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTime

`func (o *ApStatMeshUplink) SetIdleTime(v int32)`

SetIdleTime sets IdleTime field to given value.

### HasIdleTime

`func (o *ApStatMeshUplink) HasIdleTime() bool`

HasIdleTime returns a boolean if a field has been set.

### GetLastSeen

`func (o *ApStatMeshUplink) GetLastSeen() int32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ApStatMeshUplink) GetLastSeenOk() (*int32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ApStatMeshUplink) SetLastSeen(v int32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ApStatMeshUplink) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetProto

`func (o *ApStatMeshUplink) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *ApStatMeshUplink) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *ApStatMeshUplink) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *ApStatMeshUplink) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetRssi

`func (o *ApStatMeshUplink) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *ApStatMeshUplink) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *ApStatMeshUplink) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *ApStatMeshUplink) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetRxBps

`func (o *ApStatMeshUplink) GetRxBps() int32`

GetRxBps returns the RxBps field if non-nil, zero value otherwise.

### GetRxBpsOk

`func (o *ApStatMeshUplink) GetRxBpsOk() (*int32, bool)`

GetRxBpsOk returns a tuple with the RxBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBps

`func (o *ApStatMeshUplink) SetRxBps(v int32)`

SetRxBps sets RxBps field to given value.

### HasRxBps

`func (o *ApStatMeshUplink) HasRxBps() bool`

HasRxBps returns a boolean if a field has been set.

### GetRxBytes

`func (o *ApStatMeshUplink) GetRxBytes() int32`

GetRxBytes returns the RxBytes field if non-nil, zero value otherwise.

### GetRxBytesOk

`func (o *ApStatMeshUplink) GetRxBytesOk() (*int32, bool)`

GetRxBytesOk returns a tuple with the RxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBytes

`func (o *ApStatMeshUplink) SetRxBytes(v int32)`

SetRxBytes sets RxBytes field to given value.

### HasRxBytes

`func (o *ApStatMeshUplink) HasRxBytes() bool`

HasRxBytes returns a boolean if a field has been set.

### GetRxPackets

`func (o *ApStatMeshUplink) GetRxPackets() int32`

GetRxPackets returns the RxPackets field if non-nil, zero value otherwise.

### GetRxPacketsOk

`func (o *ApStatMeshUplink) GetRxPacketsOk() (*int32, bool)`

GetRxPacketsOk returns a tuple with the RxPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPackets

`func (o *ApStatMeshUplink) SetRxPackets(v int32)`

SetRxPackets sets RxPackets field to given value.

### HasRxPackets

`func (o *ApStatMeshUplink) HasRxPackets() bool`

HasRxPackets returns a boolean if a field has been set.

### GetRxRate

`func (o *ApStatMeshUplink) GetRxRate() int32`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *ApStatMeshUplink) GetRxRateOk() (*int32, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *ApStatMeshUplink) SetRxRate(v int32)`

SetRxRate sets RxRate field to given value.

### HasRxRate

`func (o *ApStatMeshUplink) HasRxRate() bool`

HasRxRate returns a boolean if a field has been set.

### GetRxRetries

`func (o *ApStatMeshUplink) GetRxRetries() int32`

GetRxRetries returns the RxRetries field if non-nil, zero value otherwise.

### GetRxRetriesOk

`func (o *ApStatMeshUplink) GetRxRetriesOk() (*int32, bool)`

GetRxRetriesOk returns a tuple with the RxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRetries

`func (o *ApStatMeshUplink) SetRxRetries(v int32)`

SetRxRetries sets RxRetries field to given value.

### HasRxRetries

`func (o *ApStatMeshUplink) HasRxRetries() bool`

HasRxRetries returns a boolean if a field has been set.

### GetSiteId

`func (o *ApStatMeshUplink) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ApStatMeshUplink) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ApStatMeshUplink) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ApStatMeshUplink) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSnr

`func (o *ApStatMeshUplink) GetSnr() int32`

GetSnr returns the Snr field if non-nil, zero value otherwise.

### GetSnrOk

`func (o *ApStatMeshUplink) GetSnrOk() (*int32, bool)`

GetSnrOk returns a tuple with the Snr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnr

`func (o *ApStatMeshUplink) SetSnr(v int32)`

SetSnr sets Snr field to given value.

### HasSnr

`func (o *ApStatMeshUplink) HasSnr() bool`

HasSnr returns a boolean if a field has been set.

### GetTxBps

`func (o *ApStatMeshUplink) GetTxBps() int32`

GetTxBps returns the TxBps field if non-nil, zero value otherwise.

### GetTxBpsOk

`func (o *ApStatMeshUplink) GetTxBpsOk() (*int32, bool)`

GetTxBpsOk returns a tuple with the TxBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBps

`func (o *ApStatMeshUplink) SetTxBps(v int32)`

SetTxBps sets TxBps field to given value.

### HasTxBps

`func (o *ApStatMeshUplink) HasTxBps() bool`

HasTxBps returns a boolean if a field has been set.

### GetTxBytes

`func (o *ApStatMeshUplink) GetTxBytes() int32`

GetTxBytes returns the TxBytes field if non-nil, zero value otherwise.

### GetTxBytesOk

`func (o *ApStatMeshUplink) GetTxBytesOk() (*int32, bool)`

GetTxBytesOk returns a tuple with the TxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBytes

`func (o *ApStatMeshUplink) SetTxBytes(v int32)`

SetTxBytes sets TxBytes field to given value.

### HasTxBytes

`func (o *ApStatMeshUplink) HasTxBytes() bool`

HasTxBytes returns a boolean if a field has been set.

### GetTxPackets

`func (o *ApStatMeshUplink) GetTxPackets() int32`

GetTxPackets returns the TxPackets field if non-nil, zero value otherwise.

### GetTxPacketsOk

`func (o *ApStatMeshUplink) GetTxPacketsOk() (*int32, bool)`

GetTxPacketsOk returns a tuple with the TxPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPackets

`func (o *ApStatMeshUplink) SetTxPackets(v int32)`

SetTxPackets sets TxPackets field to given value.

### HasTxPackets

`func (o *ApStatMeshUplink) HasTxPackets() bool`

HasTxPackets returns a boolean if a field has been set.

### GetTxRate

`func (o *ApStatMeshUplink) GetTxRate() int32`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *ApStatMeshUplink) GetTxRateOk() (*int32, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *ApStatMeshUplink) SetTxRate(v int32)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *ApStatMeshUplink) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetTxRetries

`func (o *ApStatMeshUplink) GetTxRetries() int32`

GetTxRetries returns the TxRetries field if non-nil, zero value otherwise.

### GetTxRetriesOk

`func (o *ApStatMeshUplink) GetTxRetriesOk() (*int32, bool)`

GetTxRetriesOk returns a tuple with the TxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRetries

`func (o *ApStatMeshUplink) SetTxRetries(v int32)`

SetTxRetries sets TxRetries field to given value.

### HasTxRetries

`func (o *ApStatMeshUplink) HasTxRetries() bool`

HasTxRetries returns a boolean if a field has been set.

### GetUplinkApId

`func (o *ApStatMeshUplink) GetUplinkApId() string`

GetUplinkApId returns the UplinkApId field if non-nil, zero value otherwise.

### GetUplinkApIdOk

`func (o *ApStatMeshUplink) GetUplinkApIdOk() (*string, bool)`

GetUplinkApIdOk returns a tuple with the UplinkApId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkApId

`func (o *ApStatMeshUplink) SetUplinkApId(v string)`

SetUplinkApId sets UplinkApId field to given value.

### HasUplinkApId

`func (o *ApStatMeshUplink) HasUplinkApId() bool`

HasUplinkApId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


