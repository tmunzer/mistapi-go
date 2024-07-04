# ResponseDeviceSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusId** | Pointer to **string** | the internal status id | [optional] 
**Staus** | Pointer to [**ResponseDeviceSnapshotStatus**](ResponseDeviceSnapshotStatus.md) |  | [optional] 
**Timestamp** | Pointer to **float32** |  | [optional] 

## Methods

### NewResponseDeviceSnapshot

`func NewResponseDeviceSnapshot() *ResponseDeviceSnapshot`

NewResponseDeviceSnapshot instantiates a new ResponseDeviceSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeviceSnapshotWithDefaults

`func NewResponseDeviceSnapshotWithDefaults() *ResponseDeviceSnapshot`

NewResponseDeviceSnapshotWithDefaults instantiates a new ResponseDeviceSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusId

`func (o *ResponseDeviceSnapshot) GetStatusId() string`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *ResponseDeviceSnapshot) GetStatusIdOk() (*string, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *ResponseDeviceSnapshot) SetStatusId(v string)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *ResponseDeviceSnapshot) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetStaus

`func (o *ResponseDeviceSnapshot) GetStaus() ResponseDeviceSnapshotStatus`

GetStaus returns the Staus field if non-nil, zero value otherwise.

### GetStausOk

`func (o *ResponseDeviceSnapshot) GetStausOk() (*ResponseDeviceSnapshotStatus, bool)`

GetStausOk returns a tuple with the Staus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaus

`func (o *ResponseDeviceSnapshot) SetStaus(v ResponseDeviceSnapshotStatus)`

SetStaus sets Staus field to given value.

### HasStaus

`func (o *ResponseDeviceSnapshot) HasStaus() bool`

HasStaus returns a boolean if a field has been set.

### GetTimestamp

`func (o *ResponseDeviceSnapshot) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ResponseDeviceSnapshot) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ResponseDeviceSnapshot) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ResponseDeviceSnapshot) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


