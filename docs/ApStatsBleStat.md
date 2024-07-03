# ApStatsBleStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeaconRate** | Pointer to **int32** |  | [optional] [default to 0]
**EddystoneUidEnabled** | Pointer to **bool** |  | [optional] [default to false]
**EddystoneUidFreqMsec** | Pointer to **int32** |  | [optional] [default to 0]
**EddystoneUidInstance** | Pointer to **string** |  | [optional] [default to ""]
**EddystoneUidNamespace** | Pointer to **string** |  | [optional] [default to ""]
**EddystoneUrlEnabled** | Pointer to **bool** |  | [optional] [default to false]
**EddystoneUrlFreqMsec** | Pointer to **int32** | Frequency (msec) of data emmit by Eddystone-UID beacon | [optional] [default to 0]
**EddystoneUrlUrl** | Pointer to **string** |  | [optional] [default to ""]
**IbeaconEnabled** | Pointer to **bool** |  | [optional] [default to false]
**IbeaconMajor** | Pointer to **int32** |  | [optional] [default to 0]
**IbeaconMinor** | Pointer to **int32** |  | [optional] [default to 0]
**IbeaconUuid** | Pointer to **string** |  | [optional] 
**Major** | Pointer to **int32** |  | [optional] [default to 0]
**Minors** | Pointer to **[]int32** |  | [optional] 
**Power** | Pointer to **int32** |  | [optional] [default to 9]
**RxBytes** | Pointer to **int32** |  | [optional] [default to 0]
**RxPkts** | Pointer to **int32** |  | [optional] [default to 0]
**TxBytes** | Pointer to **int32** |  | [optional] [default to 0]
**TxPkts** | Pointer to **int32** |  | [optional] [default to 0]
**TxResets** | Pointer to **int32** | resets due to tx hung | [optional] [default to 0]
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewApStatsBleStat

`func NewApStatsBleStat() *ApStatsBleStat`

NewApStatsBleStat instantiates a new ApStatsBleStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApStatsBleStatWithDefaults

`func NewApStatsBleStatWithDefaults() *ApStatsBleStat`

NewApStatsBleStatWithDefaults instantiates a new ApStatsBleStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeaconRate

`func (o *ApStatsBleStat) GetBeaconRate() int32`

GetBeaconRate returns the BeaconRate field if non-nil, zero value otherwise.

### GetBeaconRateOk

`func (o *ApStatsBleStat) GetBeaconRateOk() (*int32, bool)`

GetBeaconRateOk returns a tuple with the BeaconRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconRate

`func (o *ApStatsBleStat) SetBeaconRate(v int32)`

SetBeaconRate sets BeaconRate field to given value.

### HasBeaconRate

`func (o *ApStatsBleStat) HasBeaconRate() bool`

HasBeaconRate returns a boolean if a field has been set.

### GetEddystoneUidEnabled

`func (o *ApStatsBleStat) GetEddystoneUidEnabled() bool`

GetEddystoneUidEnabled returns the EddystoneUidEnabled field if non-nil, zero value otherwise.

### GetEddystoneUidEnabledOk

`func (o *ApStatsBleStat) GetEddystoneUidEnabledOk() (*bool, bool)`

GetEddystoneUidEnabledOk returns a tuple with the EddystoneUidEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidEnabled

`func (o *ApStatsBleStat) SetEddystoneUidEnabled(v bool)`

SetEddystoneUidEnabled sets EddystoneUidEnabled field to given value.

### HasEddystoneUidEnabled

`func (o *ApStatsBleStat) HasEddystoneUidEnabled() bool`

HasEddystoneUidEnabled returns a boolean if a field has been set.

### GetEddystoneUidFreqMsec

`func (o *ApStatsBleStat) GetEddystoneUidFreqMsec() int32`

GetEddystoneUidFreqMsec returns the EddystoneUidFreqMsec field if non-nil, zero value otherwise.

### GetEddystoneUidFreqMsecOk

`func (o *ApStatsBleStat) GetEddystoneUidFreqMsecOk() (*int32, bool)`

GetEddystoneUidFreqMsecOk returns a tuple with the EddystoneUidFreqMsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidFreqMsec

`func (o *ApStatsBleStat) SetEddystoneUidFreqMsec(v int32)`

SetEddystoneUidFreqMsec sets EddystoneUidFreqMsec field to given value.

### HasEddystoneUidFreqMsec

`func (o *ApStatsBleStat) HasEddystoneUidFreqMsec() bool`

HasEddystoneUidFreqMsec returns a boolean if a field has been set.

### GetEddystoneUidInstance

`func (o *ApStatsBleStat) GetEddystoneUidInstance() string`

GetEddystoneUidInstance returns the EddystoneUidInstance field if non-nil, zero value otherwise.

### GetEddystoneUidInstanceOk

`func (o *ApStatsBleStat) GetEddystoneUidInstanceOk() (*string, bool)`

GetEddystoneUidInstanceOk returns a tuple with the EddystoneUidInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidInstance

`func (o *ApStatsBleStat) SetEddystoneUidInstance(v string)`

SetEddystoneUidInstance sets EddystoneUidInstance field to given value.

### HasEddystoneUidInstance

`func (o *ApStatsBleStat) HasEddystoneUidInstance() bool`

HasEddystoneUidInstance returns a boolean if a field has been set.

### GetEddystoneUidNamespace

`func (o *ApStatsBleStat) GetEddystoneUidNamespace() string`

GetEddystoneUidNamespace returns the EddystoneUidNamespace field if non-nil, zero value otherwise.

### GetEddystoneUidNamespaceOk

`func (o *ApStatsBleStat) GetEddystoneUidNamespaceOk() (*string, bool)`

GetEddystoneUidNamespaceOk returns a tuple with the EddystoneUidNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUidNamespace

`func (o *ApStatsBleStat) SetEddystoneUidNamespace(v string)`

SetEddystoneUidNamespace sets EddystoneUidNamespace field to given value.

### HasEddystoneUidNamespace

`func (o *ApStatsBleStat) HasEddystoneUidNamespace() bool`

HasEddystoneUidNamespace returns a boolean if a field has been set.

### GetEddystoneUrlEnabled

`func (o *ApStatsBleStat) GetEddystoneUrlEnabled() bool`

GetEddystoneUrlEnabled returns the EddystoneUrlEnabled field if non-nil, zero value otherwise.

### GetEddystoneUrlEnabledOk

`func (o *ApStatsBleStat) GetEddystoneUrlEnabledOk() (*bool, bool)`

GetEddystoneUrlEnabledOk returns a tuple with the EddystoneUrlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrlEnabled

`func (o *ApStatsBleStat) SetEddystoneUrlEnabled(v bool)`

SetEddystoneUrlEnabled sets EddystoneUrlEnabled field to given value.

### HasEddystoneUrlEnabled

`func (o *ApStatsBleStat) HasEddystoneUrlEnabled() bool`

HasEddystoneUrlEnabled returns a boolean if a field has been set.

### GetEddystoneUrlFreqMsec

`func (o *ApStatsBleStat) GetEddystoneUrlFreqMsec() int32`

GetEddystoneUrlFreqMsec returns the EddystoneUrlFreqMsec field if non-nil, zero value otherwise.

### GetEddystoneUrlFreqMsecOk

`func (o *ApStatsBleStat) GetEddystoneUrlFreqMsecOk() (*int32, bool)`

GetEddystoneUrlFreqMsecOk returns a tuple with the EddystoneUrlFreqMsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrlFreqMsec

`func (o *ApStatsBleStat) SetEddystoneUrlFreqMsec(v int32)`

SetEddystoneUrlFreqMsec sets EddystoneUrlFreqMsec field to given value.

### HasEddystoneUrlFreqMsec

`func (o *ApStatsBleStat) HasEddystoneUrlFreqMsec() bool`

HasEddystoneUrlFreqMsec returns a boolean if a field has been set.

### GetEddystoneUrlUrl

`func (o *ApStatsBleStat) GetEddystoneUrlUrl() string`

GetEddystoneUrlUrl returns the EddystoneUrlUrl field if non-nil, zero value otherwise.

### GetEddystoneUrlUrlOk

`func (o *ApStatsBleStat) GetEddystoneUrlUrlOk() (*string, bool)`

GetEddystoneUrlUrlOk returns a tuple with the EddystoneUrlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddystoneUrlUrl

`func (o *ApStatsBleStat) SetEddystoneUrlUrl(v string)`

SetEddystoneUrlUrl sets EddystoneUrlUrl field to given value.

### HasEddystoneUrlUrl

`func (o *ApStatsBleStat) HasEddystoneUrlUrl() bool`

HasEddystoneUrlUrl returns a boolean if a field has been set.

### GetIbeaconEnabled

`func (o *ApStatsBleStat) GetIbeaconEnabled() bool`

GetIbeaconEnabled returns the IbeaconEnabled field if non-nil, zero value otherwise.

### GetIbeaconEnabledOk

`func (o *ApStatsBleStat) GetIbeaconEnabledOk() (*bool, bool)`

GetIbeaconEnabledOk returns a tuple with the IbeaconEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconEnabled

`func (o *ApStatsBleStat) SetIbeaconEnabled(v bool)`

SetIbeaconEnabled sets IbeaconEnabled field to given value.

### HasIbeaconEnabled

`func (o *ApStatsBleStat) HasIbeaconEnabled() bool`

HasIbeaconEnabled returns a boolean if a field has been set.

### GetIbeaconMajor

`func (o *ApStatsBleStat) GetIbeaconMajor() int32`

GetIbeaconMajor returns the IbeaconMajor field if non-nil, zero value otherwise.

### GetIbeaconMajorOk

`func (o *ApStatsBleStat) GetIbeaconMajorOk() (*int32, bool)`

GetIbeaconMajorOk returns a tuple with the IbeaconMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMajor

`func (o *ApStatsBleStat) SetIbeaconMajor(v int32)`

SetIbeaconMajor sets IbeaconMajor field to given value.

### HasIbeaconMajor

`func (o *ApStatsBleStat) HasIbeaconMajor() bool`

HasIbeaconMajor returns a boolean if a field has been set.

### GetIbeaconMinor

`func (o *ApStatsBleStat) GetIbeaconMinor() int32`

GetIbeaconMinor returns the IbeaconMinor field if non-nil, zero value otherwise.

### GetIbeaconMinorOk

`func (o *ApStatsBleStat) GetIbeaconMinorOk() (*int32, bool)`

GetIbeaconMinorOk returns a tuple with the IbeaconMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconMinor

`func (o *ApStatsBleStat) SetIbeaconMinor(v int32)`

SetIbeaconMinor sets IbeaconMinor field to given value.

### HasIbeaconMinor

`func (o *ApStatsBleStat) HasIbeaconMinor() bool`

HasIbeaconMinor returns a boolean if a field has been set.

### GetIbeaconUuid

`func (o *ApStatsBleStat) GetIbeaconUuid() string`

GetIbeaconUuid returns the IbeaconUuid field if non-nil, zero value otherwise.

### GetIbeaconUuidOk

`func (o *ApStatsBleStat) GetIbeaconUuidOk() (*string, bool)`

GetIbeaconUuidOk returns a tuple with the IbeaconUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeaconUuid

`func (o *ApStatsBleStat) SetIbeaconUuid(v string)`

SetIbeaconUuid sets IbeaconUuid field to given value.

### HasIbeaconUuid

`func (o *ApStatsBleStat) HasIbeaconUuid() bool`

HasIbeaconUuid returns a boolean if a field has been set.

### GetMajor

`func (o *ApStatsBleStat) GetMajor() int32`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *ApStatsBleStat) GetMajorOk() (*int32, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *ApStatsBleStat) SetMajor(v int32)`

SetMajor sets Major field to given value.

### HasMajor

`func (o *ApStatsBleStat) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetMinors

`func (o *ApStatsBleStat) GetMinors() []int32`

GetMinors returns the Minors field if non-nil, zero value otherwise.

### GetMinorsOk

`func (o *ApStatsBleStat) GetMinorsOk() (*[]int32, bool)`

GetMinorsOk returns a tuple with the Minors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinors

`func (o *ApStatsBleStat) SetMinors(v []int32)`

SetMinors sets Minors field to given value.

### HasMinors

`func (o *ApStatsBleStat) HasMinors() bool`

HasMinors returns a boolean if a field has been set.

### GetPower

`func (o *ApStatsBleStat) GetPower() int32`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *ApStatsBleStat) GetPowerOk() (*int32, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *ApStatsBleStat) SetPower(v int32)`

SetPower sets Power field to given value.

### HasPower

`func (o *ApStatsBleStat) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetRxBytes

`func (o *ApStatsBleStat) GetRxBytes() int32`

GetRxBytes returns the RxBytes field if non-nil, zero value otherwise.

### GetRxBytesOk

`func (o *ApStatsBleStat) GetRxBytesOk() (*int32, bool)`

GetRxBytesOk returns a tuple with the RxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBytes

`func (o *ApStatsBleStat) SetRxBytes(v int32)`

SetRxBytes sets RxBytes field to given value.

### HasRxBytes

`func (o *ApStatsBleStat) HasRxBytes() bool`

HasRxBytes returns a boolean if a field has been set.

### GetRxPkts

`func (o *ApStatsBleStat) GetRxPkts() int32`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *ApStatsBleStat) GetRxPktsOk() (*int32, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *ApStatsBleStat) SetRxPkts(v int32)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *ApStatsBleStat) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetTxBytes

`func (o *ApStatsBleStat) GetTxBytes() int32`

GetTxBytes returns the TxBytes field if non-nil, zero value otherwise.

### GetTxBytesOk

`func (o *ApStatsBleStat) GetTxBytesOk() (*int32, bool)`

GetTxBytesOk returns a tuple with the TxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBytes

`func (o *ApStatsBleStat) SetTxBytes(v int32)`

SetTxBytes sets TxBytes field to given value.

### HasTxBytes

`func (o *ApStatsBleStat) HasTxBytes() bool`

HasTxBytes returns a boolean if a field has been set.

### GetTxPkts

`func (o *ApStatsBleStat) GetTxPkts() int32`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *ApStatsBleStat) GetTxPktsOk() (*int32, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *ApStatsBleStat) SetTxPkts(v int32)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *ApStatsBleStat) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### GetTxResets

`func (o *ApStatsBleStat) GetTxResets() int32`

GetTxResets returns the TxResets field if non-nil, zero value otherwise.

### GetTxResetsOk

`func (o *ApStatsBleStat) GetTxResetsOk() (*int32, bool)`

GetTxResetsOk returns a tuple with the TxResets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxResets

`func (o *ApStatsBleStat) SetTxResets(v int32)`

SetTxResets sets TxResets field to given value.

### HasTxResets

`func (o *ApStatsBleStat) HasTxResets() bool`

HasTxResets returns a boolean if a field has been set.

### GetUuid

`func (o *ApStatsBleStat) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApStatsBleStat) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApStatsBleStat) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApStatsBleStat) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


