# AutoPreemption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to [**DayOfWeek**](DayOfWeek.md) |  | [optional] 
**Enabled** | Pointer to **bool** | whether auto preemption should happen | [optional] [default to false]
**TimeOfDay** | Pointer to **string** | any / HH:MM (24-hour format) | [optional] [default to "any"]

## Methods

### NewAutoPreemption

`func NewAutoPreemption() *AutoPreemption`

NewAutoPreemption instantiates a new AutoPreemption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoPreemptionWithDefaults

`func NewAutoPreemptionWithDefaults() *AutoPreemption`

NewAutoPreemptionWithDefaults instantiates a new AutoPreemption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *AutoPreemption) GetDayOfWeek() DayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *AutoPreemption) GetDayOfWeekOk() (*DayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *AutoPreemption) SetDayOfWeek(v DayOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *AutoPreemption) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetEnabled

`func (o *AutoPreemption) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoPreemption) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoPreemption) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AutoPreemption) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTimeOfDay

`func (o *AutoPreemption) GetTimeOfDay() string`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *AutoPreemption) GetTimeOfDayOk() (*string, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *AutoPreemption) SetTimeOfDay(v string)`

SetTimeOfDay sets TimeOfDay field to given value.

### HasTimeOfDay

`func (o *AutoPreemption) HasTimeOfDay() bool`

HasTimeOfDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


