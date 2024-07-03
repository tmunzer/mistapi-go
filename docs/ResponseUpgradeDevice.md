# ResponseUpgradeDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**UpgradeInfoStatus**](UpgradeInfoStatus.md) |  | 
**Timestamp** | **float32** | timestamp | 

## Methods

### NewResponseUpgradeDevice

`func NewResponseUpgradeDevice(status UpgradeInfoStatus, timestamp float32, ) *ResponseUpgradeDevice`

NewResponseUpgradeDevice instantiates a new ResponseUpgradeDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseUpgradeDeviceWithDefaults

`func NewResponseUpgradeDeviceWithDefaults() *ResponseUpgradeDevice`

NewResponseUpgradeDeviceWithDefaults instantiates a new ResponseUpgradeDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResponseUpgradeDevice) GetStatus() UpgradeInfoStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseUpgradeDevice) GetStatusOk() (*UpgradeInfoStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseUpgradeDevice) SetStatus(v UpgradeInfoStatus)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *ResponseUpgradeDevice) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ResponseUpgradeDevice) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ResponseUpgradeDevice) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


