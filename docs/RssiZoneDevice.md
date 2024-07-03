# RssiZoneDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** |  | 
**Rssi** | **int32** | RSSI threshold | 

## Methods

### NewRssiZoneDevice

`func NewRssiZoneDevice(deviceId string, rssi int32, ) *RssiZoneDevice`

NewRssiZoneDevice instantiates a new RssiZoneDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRssiZoneDeviceWithDefaults

`func NewRssiZoneDeviceWithDefaults() *RssiZoneDevice`

NewRssiZoneDeviceWithDefaults instantiates a new RssiZoneDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *RssiZoneDevice) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *RssiZoneDevice) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *RssiZoneDevice) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetRssi

`func (o *RssiZoneDevice) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *RssiZoneDevice) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *RssiZoneDevice) SetRssi(v int32)`

SetRssi sets Rssi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


