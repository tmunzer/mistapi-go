# RrmNeighbors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** |  | [optional] 
**Neighbors** | Pointer to [**[]RrmNeighborsNeighbor**](RrmNeighborsNeighbor.md) |  | [optional] 

## Methods

### NewRrmNeighbors

`func NewRrmNeighbors() *RrmNeighbors`

NewRrmNeighbors instantiates a new RrmNeighbors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRrmNeighborsWithDefaults

`func NewRrmNeighborsWithDefaults() *RrmNeighbors`

NewRrmNeighborsWithDefaults instantiates a new RrmNeighbors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *RrmNeighbors) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *RrmNeighbors) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *RrmNeighbors) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *RrmNeighbors) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNeighbors

`func (o *RrmNeighbors) GetNeighbors() []RrmNeighborsNeighbor`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *RrmNeighbors) GetNeighborsOk() (*[]RrmNeighborsNeighbor, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *RrmNeighbors) SetNeighbors(v []RrmNeighborsNeighbor)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *RrmNeighbors) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


