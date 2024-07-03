# MxedgeStatsPortStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullDuplex** | Pointer to **bool** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**RxBytes** | Pointer to **float32** |  | [optional] 
**RxErrors** | Pointer to **int32** |  | [optional] 
**RxPkts** | Pointer to **int32** |  | [optional] 
**Speed** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TxBytes** | Pointer to **int32** |  | [optional] 
**TxErrors** | Pointer to **int32** |  | [optional] 
**TxPkts** | Pointer to **int32** |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 

## Methods

### NewMxedgeStatsPortStat

`func NewMxedgeStatsPortStat() *MxedgeStatsPortStat`

NewMxedgeStatsPortStat instantiates a new MxedgeStatsPortStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeStatsPortStatWithDefaults

`func NewMxedgeStatsPortStatWithDefaults() *MxedgeStatsPortStat`

NewMxedgeStatsPortStatWithDefaults instantiates a new MxedgeStatsPortStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullDuplex

`func (o *MxedgeStatsPortStat) GetFullDuplex() bool`

GetFullDuplex returns the FullDuplex field if non-nil, zero value otherwise.

### GetFullDuplexOk

`func (o *MxedgeStatsPortStat) GetFullDuplexOk() (*bool, bool)`

GetFullDuplexOk returns a tuple with the FullDuplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullDuplex

`func (o *MxedgeStatsPortStat) SetFullDuplex(v bool)`

SetFullDuplex sets FullDuplex field to given value.

### HasFullDuplex

`func (o *MxedgeStatsPortStat) HasFullDuplex() bool`

HasFullDuplex returns a boolean if a field has been set.

### GetMac

`func (o *MxedgeStatsPortStat) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MxedgeStatsPortStat) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *MxedgeStatsPortStat) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *MxedgeStatsPortStat) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetRxBytes

`func (o *MxedgeStatsPortStat) GetRxBytes() float32`

GetRxBytes returns the RxBytes field if non-nil, zero value otherwise.

### GetRxBytesOk

`func (o *MxedgeStatsPortStat) GetRxBytesOk() (*float32, bool)`

GetRxBytesOk returns a tuple with the RxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBytes

`func (o *MxedgeStatsPortStat) SetRxBytes(v float32)`

SetRxBytes sets RxBytes field to given value.

### HasRxBytes

`func (o *MxedgeStatsPortStat) HasRxBytes() bool`

HasRxBytes returns a boolean if a field has been set.

### GetRxErrors

`func (o *MxedgeStatsPortStat) GetRxErrors() int32`

GetRxErrors returns the RxErrors field if non-nil, zero value otherwise.

### GetRxErrorsOk

`func (o *MxedgeStatsPortStat) GetRxErrorsOk() (*int32, bool)`

GetRxErrorsOk returns a tuple with the RxErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrors

`func (o *MxedgeStatsPortStat) SetRxErrors(v int32)`

SetRxErrors sets RxErrors field to given value.

### HasRxErrors

`func (o *MxedgeStatsPortStat) HasRxErrors() bool`

HasRxErrors returns a boolean if a field has been set.

### GetRxPkts

`func (o *MxedgeStatsPortStat) GetRxPkts() int32`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *MxedgeStatsPortStat) GetRxPktsOk() (*int32, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *MxedgeStatsPortStat) SetRxPkts(v int32)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *MxedgeStatsPortStat) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### GetSpeed

`func (o *MxedgeStatsPortStat) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *MxedgeStatsPortStat) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *MxedgeStatsPortStat) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *MxedgeStatsPortStat) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetState

`func (o *MxedgeStatsPortStat) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MxedgeStatsPortStat) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MxedgeStatsPortStat) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MxedgeStatsPortStat) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTxBytes

`func (o *MxedgeStatsPortStat) GetTxBytes() int32`

GetTxBytes returns the TxBytes field if non-nil, zero value otherwise.

### GetTxBytesOk

`func (o *MxedgeStatsPortStat) GetTxBytesOk() (*int32, bool)`

GetTxBytesOk returns a tuple with the TxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBytes

`func (o *MxedgeStatsPortStat) SetTxBytes(v int32)`

SetTxBytes sets TxBytes field to given value.

### HasTxBytes

`func (o *MxedgeStatsPortStat) HasTxBytes() bool`

HasTxBytes returns a boolean if a field has been set.

### GetTxErrors

`func (o *MxedgeStatsPortStat) GetTxErrors() int32`

GetTxErrors returns the TxErrors field if non-nil, zero value otherwise.

### GetTxErrorsOk

`func (o *MxedgeStatsPortStat) GetTxErrorsOk() (*int32, bool)`

GetTxErrorsOk returns a tuple with the TxErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrors

`func (o *MxedgeStatsPortStat) SetTxErrors(v int32)`

SetTxErrors sets TxErrors field to given value.

### HasTxErrors

`func (o *MxedgeStatsPortStat) HasTxErrors() bool`

HasTxErrors returns a boolean if a field has been set.

### GetTxPkts

`func (o *MxedgeStatsPortStat) GetTxPkts() int32`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *MxedgeStatsPortStat) GetTxPktsOk() (*int32, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *MxedgeStatsPortStat) SetTxPkts(v int32)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *MxedgeStatsPortStat) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### GetUp

`func (o *MxedgeStatsPortStat) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *MxedgeStatsPortStat) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *MxedgeStatsPortStat) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *MxedgeStatsPortStat) HasUp() bool`

HasUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


