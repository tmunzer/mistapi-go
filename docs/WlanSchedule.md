# WlanSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**Hours** | Pointer to [**Hours**](Hours.md) |  | [optional] 

## Methods

### NewWlanSchedule

`func NewWlanSchedule() *WlanSchedule`

NewWlanSchedule instantiates a new WlanSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanScheduleWithDefaults

`func NewWlanScheduleWithDefaults() *WlanSchedule`

NewWlanScheduleWithDefaults instantiates a new WlanSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *WlanSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WlanSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WlanSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WlanSchedule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHours

`func (o *WlanSchedule) GetHours() Hours`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *WlanSchedule) GetHoursOk() (*Hours, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *WlanSchedule) SetHours(v Hours)`

SetHours sets Hours field to given value.

### HasHours

`func (o *WlanSchedule) HasHours() bool`

HasHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


