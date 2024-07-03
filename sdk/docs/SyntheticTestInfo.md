# SyntheticTestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | Pointer to **string** |  | [optional] 
**DeviceType** | Pointer to [**SyntheticTestInfoDeviceType**](SyntheticTestInfoDeviceType.md) |  | [optional] 
**Failed** | Pointer to **bool** |  | [optional] 
**Latency** | Pointer to **int32** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**PortId** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**RxMbps** | Pointer to **int32** |  | [optional] 
**StartTime** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **float32** |  | [optional] 
**TxMbps** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**SyntheticTestType**](SyntheticTestType.md) |  | [optional] 
**VlanId** | Pointer to **int32** |  | [optional] 

## Methods

### NewSyntheticTestInfo

`func NewSyntheticTestInfo() *SyntheticTestInfo`

NewSyntheticTestInfo instantiates a new SyntheticTestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticTestInfoWithDefaults

`func NewSyntheticTestInfoWithDefaults() *SyntheticTestInfo`

NewSyntheticTestInfoWithDefaults instantiates a new SyntheticTestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *SyntheticTestInfo) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *SyntheticTestInfo) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *SyntheticTestInfo) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *SyntheticTestInfo) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetDeviceType

`func (o *SyntheticTestInfo) GetDeviceType() SyntheticTestInfoDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *SyntheticTestInfo) GetDeviceTypeOk() (*SyntheticTestInfoDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *SyntheticTestInfo) SetDeviceType(v SyntheticTestInfoDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *SyntheticTestInfo) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetFailed

`func (o *SyntheticTestInfo) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *SyntheticTestInfo) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *SyntheticTestInfo) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *SyntheticTestInfo) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetLatency

`func (o *SyntheticTestInfo) GetLatency() int32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *SyntheticTestInfo) GetLatencyOk() (*int32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *SyntheticTestInfo) SetLatency(v int32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *SyntheticTestInfo) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetMac

`func (o *SyntheticTestInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SyntheticTestInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SyntheticTestInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SyntheticTestInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetPortId

`func (o *SyntheticTestInfo) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *SyntheticTestInfo) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *SyntheticTestInfo) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *SyntheticTestInfo) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetReason

`func (o *SyntheticTestInfo) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SyntheticTestInfo) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SyntheticTestInfo) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SyntheticTestInfo) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRxMbps

`func (o *SyntheticTestInfo) GetRxMbps() int32`

GetRxMbps returns the RxMbps field if non-nil, zero value otherwise.

### GetRxMbpsOk

`func (o *SyntheticTestInfo) GetRxMbpsOk() (*int32, bool)`

GetRxMbpsOk returns a tuple with the RxMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMbps

`func (o *SyntheticTestInfo) SetRxMbps(v int32)`

SetRxMbps sets RxMbps field to given value.

### HasRxMbps

`func (o *SyntheticTestInfo) HasRxMbps() bool`

HasRxMbps returns a boolean if a field has been set.

### GetStartTime

`func (o *SyntheticTestInfo) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SyntheticTestInfo) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SyntheticTestInfo) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SyntheticTestInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *SyntheticTestInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticTestInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticTestInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyntheticTestInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *SyntheticTestInfo) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SyntheticTestInfo) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SyntheticTestInfo) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SyntheticTestInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTxMbps

`func (o *SyntheticTestInfo) GetTxMbps() int32`

GetTxMbps returns the TxMbps field if non-nil, zero value otherwise.

### GetTxMbpsOk

`func (o *SyntheticTestInfo) GetTxMbpsOk() (*int32, bool)`

GetTxMbpsOk returns a tuple with the TxMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMbps

`func (o *SyntheticTestInfo) SetTxMbps(v int32)`

SetTxMbps sets TxMbps field to given value.

### HasTxMbps

`func (o *SyntheticTestInfo) HasTxMbps() bool`

HasTxMbps returns a boolean if a field has been set.

### GetType

`func (o *SyntheticTestInfo) GetType() SyntheticTestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticTestInfo) GetTypeOk() (*SyntheticTestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticTestInfo) SetType(v SyntheticTestType)`

SetType sets Type field to given value.

### HasType

`func (o *SyntheticTestInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlanId

`func (o *SyntheticTestInfo) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *SyntheticTestInfo) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *SyntheticTestInfo) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *SyntheticTestInfo) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


