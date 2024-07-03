# ResponseAutoOrientation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**AutoOrientationState**](AutoOrientationState.md) |  | [optional] 
**TimeQueued** | Pointer to **float32** | Time when auto orient process was last queued for this map | [optional] 

## Methods

### NewResponseAutoOrientation

`func NewResponseAutoOrientation() *ResponseAutoOrientation`

NewResponseAutoOrientation instantiates a new ResponseAutoOrientation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAutoOrientationWithDefaults

`func NewResponseAutoOrientationWithDefaults() *ResponseAutoOrientation`

NewResponseAutoOrientationWithDefaults instantiates a new ResponseAutoOrientation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseAutoOrientation) GetState() AutoOrientationState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseAutoOrientation) GetStateOk() (*AutoOrientationState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseAutoOrientation) SetState(v AutoOrientationState)`

SetState sets State field to given value.

### HasState

`func (o *ResponseAutoOrientation) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTimeQueued

`func (o *ResponseAutoOrientation) GetTimeQueued() float32`

GetTimeQueued returns the TimeQueued field if non-nil, zero value otherwise.

### GetTimeQueuedOk

`func (o *ResponseAutoOrientation) GetTimeQueuedOk() (*float32, bool)`

GetTimeQueuedOk returns a tuple with the TimeQueued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeQueued

`func (o *ResponseAutoOrientation) SetTimeQueued(v float32)`

SetTimeQueued sets TimeQueued field to given value.

### HasTimeQueued

`func (o *ResponseAutoOrientation) HasTimeQueued() bool`

HasTimeQueued returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


