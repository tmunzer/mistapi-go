# Anomaly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | **[]string** |  | [readonly] 
**Since** | Pointer to **float32** |  | [optional] [readonly] 
**SleBaseline** | **float32** |  | [readonly] 
**SleDeviation** | **float32** |  | [readonly] 
**Timestamp** | **float32** |  | [readonly] 

## Methods

### NewAnomaly

`func NewAnomaly(events []string, sleBaseline float32, sleDeviation float32, timestamp float32, ) *Anomaly`

NewAnomaly instantiates a new Anomaly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyWithDefaults

`func NewAnomalyWithDefaults() *Anomaly`

NewAnomalyWithDefaults instantiates a new Anomaly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *Anomaly) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Anomaly) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Anomaly) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetSince

`func (o *Anomaly) GetSince() float32`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *Anomaly) GetSinceOk() (*float32, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *Anomaly) SetSince(v float32)`

SetSince sets Since field to given value.

### HasSince

`func (o *Anomaly) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetSleBaseline

`func (o *Anomaly) GetSleBaseline() float32`

GetSleBaseline returns the SleBaseline field if non-nil, zero value otherwise.

### GetSleBaselineOk

`func (o *Anomaly) GetSleBaselineOk() (*float32, bool)`

GetSleBaselineOk returns a tuple with the SleBaseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleBaseline

`func (o *Anomaly) SetSleBaseline(v float32)`

SetSleBaseline sets SleBaseline field to given value.


### GetSleDeviation

`func (o *Anomaly) GetSleDeviation() float32`

GetSleDeviation returns the SleDeviation field if non-nil, zero value otherwise.

### GetSleDeviationOk

`func (o *Anomaly) GetSleDeviationOk() (*float32, bool)`

GetSleDeviationOk returns a tuple with the SleDeviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleDeviation

`func (o *Anomaly) SetSleDeviation(v float32)`

SetSleDeviation sets SleDeviation field to given value.


### GetTimestamp

`func (o *Anomaly) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Anomaly) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Anomaly) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


