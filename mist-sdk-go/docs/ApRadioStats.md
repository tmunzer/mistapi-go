# ApRadioStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bandwidth** | Pointer to [**Dot11Bandwidth**](Dot11Bandwidth.md) |  | [optional] 
**Channel** | Pointer to **int32** | current channel the radio is running on | [optional] 
**DynamicChainingEnalbed** | Pointer to **bool** | Use dynamic chaining for downlink | [optional] 
**Mac** | Pointer to **string** | radio (base) mac, it can have 16 bssids (e.g. 5c5b350001a0-5c5b350001af) | [optional] 
**NumClients** | Pointer to **int32** |  | [optional] 
**Power** | Pointer to **int32** | transmit power (in dBm) | [optional] 
**RxBytes** | Pointer to **float32** |  | [optional] 
**RxPkts** | Pointer to **float32** |  | [optional] 
**TxBytes** | Pointer to **float32** |  | [optional] 
**TxPkts** | Pointer to **float32** |  | [optional] 
**UtilAll** | Pointer to **int32** | all utilization in percentage | [optional] 
**UtilNonWifi** | Pointer to **int32** | reception of “No Packets” utilization in percentage, received frames with invalid PLCPs and CRS glitches as noise | [optional] 
**UtilRxInBss** | Pointer to **int32** | reception of “In BSS” utilization in percentage, only frames that are received from AP/STAs within the BSS | [optional] 
**UtilRxOtherBss** | Pointer to **int32** | reception of “Other BSS” utilization in percentage, all frames received from AP/STAs that are outside the BSS | [optional] 
**UtilTx** | Pointer to **int32** | transmission utilization in percentage | [optional] 
**UtilUndecodableWifi** | Pointer to **int32** | reception of “UnDecodable Wifi“ utilization in percentage, only Preamble, PLCP header is decoded, Rest is undecodable in this radio | [optional] 
**UtilUnknownWifi** | Pointer to **int32** | reception of “No Category” utilization in percentage, all 802.11 frames that are corrupted at the receiver | [optional] 

## Methods

### NewApRadioStats

`func NewApRadioStats() *ApRadioStats`

NewApRadioStats instantiates a new ApRadioStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApRadioStatsWithDefaults

`func NewApRadioStatsWithDefaults() *ApRadioStats`

NewApRadioStatsWithDefaults instantiates a new ApRadioStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidth

`func (o *ApRadioStats) GetBandwidth() Dot11Bandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *ApRadioStats) GetBandwidthOk() (*Dot11Bandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *ApRadioStats) SetBandwidth(v Dot11Bandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *ApRadioStats) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetChannel

`func (o *ApRadioStats) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ApRadioStats) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ApRadioStats) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ApRadioStats) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDynamicChainingEnalbed

`func (o *ApRadioStats) GetDynamicChainingEnalbed() bool`

GetDynamicChainingEnalbed returns the DynamicChainingEnalbed field if non-nil, zero value otherwise.

### GetDynamicChainingEnalbedOk

`func (o *ApRadioStats) GetDynamicChainingEnalbedOk() (*bool, bool)`

GetDynamicChainingEnalbedOk returns a tuple with the DynamicChainingEnalbed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicChainingEnalbed

`func (o *ApRadioStats) SetDynamicChainingEnalbed(v bool)`

SetDynamicChainingEnalbed sets DynamicChainingEnalbed field to given value.

### HasDynamicChainingEnalbed

`func (o *ApRadioStats) HasDynamicChainingEnalbed() bool`

HasDynamicChainingEnalbed returns a boolean if a field has been set.

### GetMac

`func (o *ApRadioStats) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ApRadioStats) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ApRadioStats) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ApRadioStats) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNumClients

`func (o *ApRadioStats) GetNumClients() int32`

GetNumClients returns the NumClients field if non-nil, zero value otherwise.

### GetNumClientsOk

`func (o *ApRadioStats) GetNumClientsOk() (*int32, bool)`

GetNumClientsOk returns a tuple with the NumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumClients

`func (o *ApRadioStats) SetNumClients(v int32)`

SetNumClients sets NumClients field to given value.

### HasNumClients

`func (o *ApRadioStats) HasNumClients() bool`

HasNumClients returns a boolean if a field has been set.

### GetPower

`func (o *ApRadioStats) GetPower() int32`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *ApRadioStats) GetPowerOk() (*int32, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *ApRadioStats) SetPower(v int32)`

SetPower sets Power field to given value.

### HasPower

`func (o *ApRadioStats) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetRxBytes

`func (o *ApRadioStats) GetRxBytes() float32`

GetRxBytes returns the RxBytes field if non-nil, zero value otherwise.

### GetRxBytesOk

`func (o *ApRadioStats) GetRxBytesOk() (*float32, bool)`

GetRxBytesOk returns a tuple with the RxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBytes

`func (o *ApRadioStats) SetRxBytes(v float32)`

SetRxBytes sets RxBytes field to given value.

### HasRxBytes

`func (o *ApRadioStats) HasRxBytes() bool`

HasRxBytes returns a boolean if a field has been set.

### GetRxPkts

`func (o *ApRadioStats) GetRxPkts() float32`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *ApRadioStats) GetRxPktsOk() (*float32, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *ApRadioStats) SetRxPkts(v float32)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *ApRadioStats) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetTxBytes

`func (o *ApRadioStats) GetTxBytes() float32`

GetTxBytes returns the TxBytes field if non-nil, zero value otherwise.

### GetTxBytesOk

`func (o *ApRadioStats) GetTxBytesOk() (*float32, bool)`

GetTxBytesOk returns a tuple with the TxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBytes

`func (o *ApRadioStats) SetTxBytes(v float32)`

SetTxBytes sets TxBytes field to given value.

### HasTxBytes

`func (o *ApRadioStats) HasTxBytes() bool`

HasTxBytes returns a boolean if a field has been set.

### GetTxPkts

`func (o *ApRadioStats) GetTxPkts() float32`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *ApRadioStats) GetTxPktsOk() (*float32, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *ApRadioStats) SetTxPkts(v float32)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *ApRadioStats) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### GetUtilAll

`func (o *ApRadioStats) GetUtilAll() int32`

GetUtilAll returns the UtilAll field if non-nil, zero value otherwise.

### GetUtilAllOk

`func (o *ApRadioStats) GetUtilAllOk() (*int32, bool)`

GetUtilAllOk returns a tuple with the UtilAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilAll

`func (o *ApRadioStats) SetUtilAll(v int32)`

SetUtilAll sets UtilAll field to given value.

### HasUtilAll

`func (o *ApRadioStats) HasUtilAll() bool`

HasUtilAll returns a boolean if a field has been set.

### GetUtilNonWifi

`func (o *ApRadioStats) GetUtilNonWifi() int32`

GetUtilNonWifi returns the UtilNonWifi field if non-nil, zero value otherwise.

### GetUtilNonWifiOk

`func (o *ApRadioStats) GetUtilNonWifiOk() (*int32, bool)`

GetUtilNonWifiOk returns a tuple with the UtilNonWifi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilNonWifi

`func (o *ApRadioStats) SetUtilNonWifi(v int32)`

SetUtilNonWifi sets UtilNonWifi field to given value.

### HasUtilNonWifi

`func (o *ApRadioStats) HasUtilNonWifi() bool`

HasUtilNonWifi returns a boolean if a field has been set.

### GetUtilRxInBss

`func (o *ApRadioStats) GetUtilRxInBss() int32`

GetUtilRxInBss returns the UtilRxInBss field if non-nil, zero value otherwise.

### GetUtilRxInBssOk

`func (o *ApRadioStats) GetUtilRxInBssOk() (*int32, bool)`

GetUtilRxInBssOk returns a tuple with the UtilRxInBss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilRxInBss

`func (o *ApRadioStats) SetUtilRxInBss(v int32)`

SetUtilRxInBss sets UtilRxInBss field to given value.

### HasUtilRxInBss

`func (o *ApRadioStats) HasUtilRxInBss() bool`

HasUtilRxInBss returns a boolean if a field has been set.

### GetUtilRxOtherBss

`func (o *ApRadioStats) GetUtilRxOtherBss() int32`

GetUtilRxOtherBss returns the UtilRxOtherBss field if non-nil, zero value otherwise.

### GetUtilRxOtherBssOk

`func (o *ApRadioStats) GetUtilRxOtherBssOk() (*int32, bool)`

GetUtilRxOtherBssOk returns a tuple with the UtilRxOtherBss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilRxOtherBss

`func (o *ApRadioStats) SetUtilRxOtherBss(v int32)`

SetUtilRxOtherBss sets UtilRxOtherBss field to given value.

### HasUtilRxOtherBss

`func (o *ApRadioStats) HasUtilRxOtherBss() bool`

HasUtilRxOtherBss returns a boolean if a field has been set.

### GetUtilTx

`func (o *ApRadioStats) GetUtilTx() int32`

GetUtilTx returns the UtilTx field if non-nil, zero value otherwise.

### GetUtilTxOk

`func (o *ApRadioStats) GetUtilTxOk() (*int32, bool)`

GetUtilTxOk returns a tuple with the UtilTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilTx

`func (o *ApRadioStats) SetUtilTx(v int32)`

SetUtilTx sets UtilTx field to given value.

### HasUtilTx

`func (o *ApRadioStats) HasUtilTx() bool`

HasUtilTx returns a boolean if a field has been set.

### GetUtilUndecodableWifi

`func (o *ApRadioStats) GetUtilUndecodableWifi() int32`

GetUtilUndecodableWifi returns the UtilUndecodableWifi field if non-nil, zero value otherwise.

### GetUtilUndecodableWifiOk

`func (o *ApRadioStats) GetUtilUndecodableWifiOk() (*int32, bool)`

GetUtilUndecodableWifiOk returns a tuple with the UtilUndecodableWifi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilUndecodableWifi

`func (o *ApRadioStats) SetUtilUndecodableWifi(v int32)`

SetUtilUndecodableWifi sets UtilUndecodableWifi field to given value.

### HasUtilUndecodableWifi

`func (o *ApRadioStats) HasUtilUndecodableWifi() bool`

HasUtilUndecodableWifi returns a boolean if a field has been set.

### GetUtilUnknownWifi

`func (o *ApRadioStats) GetUtilUnknownWifi() int32`

GetUtilUnknownWifi returns the UtilUnknownWifi field if non-nil, zero value otherwise.

### GetUtilUnknownWifiOk

`func (o *ApRadioStats) GetUtilUnknownWifiOk() (*int32, bool)`

GetUtilUnknownWifiOk returns a tuple with the UtilUnknownWifi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilUnknownWifi

`func (o *ApRadioStats) SetUtilUnknownWifi(v int32)`

SetUtilUnknownWifi sets UtilUnknownWifi field to given value.

### HasUtilUnknownWifi

`func (o *ApRadioStats) HasUtilUnknownWifi() bool`

HasUtilUnknownWifi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


