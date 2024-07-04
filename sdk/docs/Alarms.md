# Alarms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlarmIds** | **[]string** |  | 
**Note** | Pointer to **string** | Some text note describing the intent | [optional] 

## Methods

### NewAlarms

`func NewAlarms(alarmIds []string, ) *Alarms`

NewAlarms instantiates a new Alarms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmsWithDefaults

`func NewAlarmsWithDefaults() *Alarms`

NewAlarmsWithDefaults instantiates a new Alarms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmIds

`func (o *Alarms) GetAlarmIds() []string`

GetAlarmIds returns the AlarmIds field if non-nil, zero value otherwise.

### GetAlarmIdsOk

`func (o *Alarms) GetAlarmIdsOk() (*[]string, bool)`

GetAlarmIdsOk returns a tuple with the AlarmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmIds

`func (o *Alarms) SetAlarmIds(v []string)`

SetAlarmIds sets AlarmIds field to given value.


### GetNote

`func (o *Alarms) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Alarms) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Alarms) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Alarms) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


