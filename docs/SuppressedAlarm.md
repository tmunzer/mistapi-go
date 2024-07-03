# SuppressedAlarm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applies** | Pointer to [**SuppressedAlarmApplies**](SuppressedAlarmApplies.md) |  | [optional] 
**Duration** | Pointer to **float32** | duration, in seconds. Maximum duration is 86400 * 180 (180 days). 0 is to un-suppress alarms | [optional] [default to 3600]
**ScheduledTime** | Pointer to **int32** | poch_time in seconds, Default as now, accepted time range is from now to now + 7 days | [optional] 
**Scope** | Pointer to [**SuppressedAlarmScope**](SuppressedAlarmScope.md) |  | [optional] [default to SUPPRESSEDALARMSCOPE_SITE]

## Methods

### NewSuppressedAlarm

`func NewSuppressedAlarm() *SuppressedAlarm`

NewSuppressedAlarm instantiates a new SuppressedAlarm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppressedAlarmWithDefaults

`func NewSuppressedAlarmWithDefaults() *SuppressedAlarm`

NewSuppressedAlarmWithDefaults instantiates a new SuppressedAlarm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplies

`func (o *SuppressedAlarm) GetApplies() SuppressedAlarmApplies`

GetApplies returns the Applies field if non-nil, zero value otherwise.

### GetAppliesOk

`func (o *SuppressedAlarm) GetAppliesOk() (*SuppressedAlarmApplies, bool)`

GetAppliesOk returns a tuple with the Applies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplies

`func (o *SuppressedAlarm) SetApplies(v SuppressedAlarmApplies)`

SetApplies sets Applies field to given value.

### HasApplies

`func (o *SuppressedAlarm) HasApplies() bool`

HasApplies returns a boolean if a field has been set.

### GetDuration

`func (o *SuppressedAlarm) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SuppressedAlarm) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SuppressedAlarm) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SuppressedAlarm) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetScheduledTime

`func (o *SuppressedAlarm) GetScheduledTime() int32`

GetScheduledTime returns the ScheduledTime field if non-nil, zero value otherwise.

### GetScheduledTimeOk

`func (o *SuppressedAlarm) GetScheduledTimeOk() (*int32, bool)`

GetScheduledTimeOk returns a tuple with the ScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTime

`func (o *SuppressedAlarm) SetScheduledTime(v int32)`

SetScheduledTime sets ScheduledTime field to given value.

### HasScheduledTime

`func (o *SuppressedAlarm) HasScheduledTime() bool`

HasScheduledTime returns a boolean if a field has been set.

### GetScope

`func (o *SuppressedAlarm) GetScope() SuppressedAlarmScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SuppressedAlarm) GetScopeOk() (*SuppressedAlarmScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SuppressedAlarm) SetScope(v SuppressedAlarmScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SuppressedAlarm) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


