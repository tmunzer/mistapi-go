# MapWayfinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Micello** | Pointer to [**MapWayfindingMicello**](MapWayfindingMicello.md) |  | [optional] 
**SnapToPath** | Pointer to **bool** |  | [optional] 

## Methods

### NewMapWayfinding

`func NewMapWayfinding() *MapWayfinding`

NewMapWayfinding instantiates a new MapWayfinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapWayfindingWithDefaults

`func NewMapWayfindingWithDefaults() *MapWayfinding`

NewMapWayfindingWithDefaults instantiates a new MapWayfinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMicello

`func (o *MapWayfinding) GetMicello() MapWayfindingMicello`

GetMicello returns the Micello field if non-nil, zero value otherwise.

### GetMicelloOk

`func (o *MapWayfinding) GetMicelloOk() (*MapWayfindingMicello, bool)`

GetMicelloOk returns a tuple with the Micello field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicello

`func (o *MapWayfinding) SetMicello(v MapWayfindingMicello)`

SetMicello sets Micello field to given value.

### HasMicello

`func (o *MapWayfinding) HasMicello() bool`

HasMicello returns a boolean if a field has been set.

### GetSnapToPath

`func (o *MapWayfinding) GetSnapToPath() bool`

GetSnapToPath returns the SnapToPath field if non-nil, zero value otherwise.

### GetSnapToPathOk

`func (o *MapWayfinding) GetSnapToPathOk() (*bool, bool)`

GetSnapToPathOk returns a tuple with the SnapToPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapToPath

`func (o *MapWayfinding) SetSnapToPath(v bool)`

SetSnapToPath sets SnapToPath field to given value.

### HasSnapToPath

`func (o *MapWayfinding) HasSnapToPath() bool`

HasSnapToPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


