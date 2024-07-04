# ResponseAutoPlacementInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **float32** | time when autoplacement completed or was manually stopped | [optional] 
**EstTimeLeft** | Pointer to **float32** | (Only when inprogress) estimate of the time to completion | [optional] 
**StartTime** | Pointer to **float32** | time when autoplacement process was last queued for this map | [optional] 
**Status** | Pointer to [**AutoPlacementInfoStatus**](AutoPlacementInfoStatus.md) |  | [optional] 

## Methods

### NewResponseAutoPlacementInfo

`func NewResponseAutoPlacementInfo() *ResponseAutoPlacementInfo`

NewResponseAutoPlacementInfo instantiates a new ResponseAutoPlacementInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAutoPlacementInfoWithDefaults

`func NewResponseAutoPlacementInfoWithDefaults() *ResponseAutoPlacementInfo`

NewResponseAutoPlacementInfoWithDefaults instantiates a new ResponseAutoPlacementInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *ResponseAutoPlacementInfo) GetEndTime() float32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ResponseAutoPlacementInfo) GetEndTimeOk() (*float32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ResponseAutoPlacementInfo) SetEndTime(v float32)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ResponseAutoPlacementInfo) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEstTimeLeft

`func (o *ResponseAutoPlacementInfo) GetEstTimeLeft() float32`

GetEstTimeLeft returns the EstTimeLeft field if non-nil, zero value otherwise.

### GetEstTimeLeftOk

`func (o *ResponseAutoPlacementInfo) GetEstTimeLeftOk() (*float32, bool)`

GetEstTimeLeftOk returns a tuple with the EstTimeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstTimeLeft

`func (o *ResponseAutoPlacementInfo) SetEstTimeLeft(v float32)`

SetEstTimeLeft sets EstTimeLeft field to given value.

### HasEstTimeLeft

`func (o *ResponseAutoPlacementInfo) HasEstTimeLeft() bool`

HasEstTimeLeft returns a boolean if a field has been set.

### GetStartTime

`func (o *ResponseAutoPlacementInfo) GetStartTime() float32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ResponseAutoPlacementInfo) GetStartTimeOk() (*float32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ResponseAutoPlacementInfo) SetStartTime(v float32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ResponseAutoPlacementInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseAutoPlacementInfo) GetStatus() AutoPlacementInfoStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseAutoPlacementInfo) GetStatusOk() (*AutoPlacementInfoStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseAutoPlacementInfo) SetStatus(v AutoPlacementInfoStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseAutoPlacementInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


